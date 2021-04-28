// +build !providerless

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/stretchr/testify/assert"
)

func TestElbProtocolsAreEqual(t *testing.T) {
	grid := []struct {
		L        *string
		R        *string
		Expected bool
	}{
		{
			L:        aws.String("http"),
			R:        aws.String("http"),
			Expected: true,
		},
		{
			L:        aws.String("HTTP"),
			R:        aws.String("http"),
			Expected: true,
		},
		{
			L:        aws.String("HTTP"),
			R:        aws.String("TCP"),
			Expected: false,
		},
		{
			L:        aws.String(""),
			R:        aws.String("TCP"),
			Expected: false,
		},
		{
			L:        aws.String(""),
			R:        aws.String(""),
			Expected: true,
		},
		{
			L:        nil,
			R:        aws.String(""),
			Expected: false,
		},
		{
			L:        aws.String(""),
			R:        nil,
			Expected: false,
		},
		{
			L:        nil,
			R:        nil,
			Expected: true,
		},
	}
	for _, g := range grid {
		actual := elbProtocolsAreEqual(g.L, g.R)
		if actual != g.Expected {
			t.Errorf("unexpected result from protocolsEquals(%v, %v)", g.L, g.R)
		}
	}
}

func TestAWSARNEquals(t *testing.T) {
	grid := []struct {
		L        *string
		R        *string
		Expected bool
	}{
		{
			L:        aws.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"),
			R:        aws.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"),
			Expected: true,
		},
		{
			L:        aws.String("ARN:AWS:ACM:US-EAST-1:123456789012:CERTIFICATE/12345678-1234-1234-1234-123456789012"),
			R:        aws.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"),
			Expected: true,
		},
		{
			L:        aws.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012"),
			R:        aws.String(""),
			Expected: false,
		},
		{
			L:        aws.String(""),
			R:        aws.String(""),
			Expected: true,
		},
		{
			L:        nil,
			R:        aws.String(""),
			Expected: false,
		},
		{
			L:        aws.String(""),
			R:        nil,
			Expected: false,
		},
		{
			L:        nil,
			R:        nil,
			Expected: true,
		},
	}
	for _, g := range grid {
		actual := awsArnEquals(g.L, g.R)
		if actual != g.Expected {
			t.Errorf("unexpected result from awsArnEquals(%v, %v)", g.L, g.R)
		}
	}
}

func TestIsNLB(t *testing.T) {
	tests := []struct {
		name string

		annotations map[string]string
		want        bool
	}{
		{
			"NLB annotation provided",
			map[string]string{"service.beta.kubernetes.io/aws-load-balancer-type": "nlb"},
			true,
		},
		{
			"NLB annotation has invalid value",
			map[string]string{"service.beta.kubernetes.io/aws-load-balancer-type": "elb"},
			false,
		},
		{
			"NLB annotation absent",
			map[string]string{},
			false,
		},
	}

	for _, test := range tests {
		t.Logf("Running test case %s", test.name)
		got := isNLB(test.annotations)

		if got != test.want {
			t.Errorf("Incorrect value for isNLB() case %s. Got %t, expected %t.", test.name, got, test.want)
		}
	}
}

func TestIsLBExternal(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		want        bool
	}{
		{
			name:        "No annotation",
			annotations: map[string]string{},
			want:        false,
		},
		{
			name:        "Type NLB",
			annotations: map[string]string{"service.beta.kubernetes.io/aws-load-balancer-type": "nlb"},
			want:        false,
		},
		{
			name:        "Type NLB-IP",
			annotations: map[string]string{"service.beta.kubernetes.io/aws-load-balancer-type": "nlb-ip"},
			want:        true,
		},
		{
			name:        "Type External",
			annotations: map[string]string{"service.beta.kubernetes.io/aws-load-balancer-type": "external"},
			want:        true,
		},
	}
	for _, test := range tests {
		t.Logf("Running test case %s", test.name)
		got := isLBExternal(test.annotations)

		if got != test.want {
			t.Errorf("Incorrect value for isLBExternal() case %s. Got %t, expected %t.", test.name, got, test.want)
		}
	}
}

func TestSyncElbListeners(t *testing.T) {
	tests := []struct {
		name                 string
		loadBalancerName     string
		listeners            []*elb.Listener
		listenerDescriptions []*elb.ListenerDescription
		toCreate             []*elb.Listener
		toDelete             []*int64
	}{
		{
			name:             "no edge cases",
			loadBalancerName: "lb_one",
			listeners: []*elb.Listener{
				{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP"), SSLCertificateId: aws.String("abc-123")},
				{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP"), SSLCertificateId: aws.String("def-456")},
				{InstancePort: aws.Int64(8443), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(8443), Protocol: aws.String("TCP"), SSLCertificateId: aws.String("def-456")},
			},
			listenerDescriptions: []*elb.ListenerDescription{
				{Listener: &elb.Listener{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")}},
				{Listener: &elb.Listener{InstancePort: aws.Int64(8443), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(8443), Protocol: aws.String("TCP"), SSLCertificateId: aws.String("def-456")}},
			},
			toDelete: []*int64{
				aws.Int64(80),
			},
			toCreate: []*elb.Listener{
				{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP"), SSLCertificateId: aws.String("abc-123")},
				{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP"), SSLCertificateId: aws.String("def-456")},
			},
		},
		{
			name:             "no listeners to delete",
			loadBalancerName: "lb_two",
			listeners: []*elb.Listener{
				{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP"), SSLCertificateId: aws.String("abc-123")},
				{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP"), SSLCertificateId: aws.String("def-456")},
			},
			listenerDescriptions: []*elb.ListenerDescription{
				{Listener: &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP"), SSLCertificateId: aws.String("abc-123")}},
			},
			toCreate: []*elb.Listener{
				{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP"), SSLCertificateId: aws.String("def-456")},
			},
			toDelete: []*int64{},
		},
		{
			name:             "no listeners to create",
			loadBalancerName: "lb_three",
			listeners: []*elb.Listener{
				{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP"), SSLCertificateId: aws.String("abc-123")},
			},
			listenerDescriptions: []*elb.ListenerDescription{
				{Listener: &elb.Listener{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")}},
				{Listener: &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP"), SSLCertificateId: aws.String("abc-123")}},
			},
			toDelete: []*int64{
				aws.Int64(80),
			},
			toCreate: []*elb.Listener{},
		},
		{
			name:             "nil actual listener",
			loadBalancerName: "lb_four",
			listeners: []*elb.Listener{
				{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP")},
			},
			listenerDescriptions: []*elb.ListenerDescription{
				{Listener: &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP"), SSLCertificateId: aws.String("abc-123")}},
				{Listener: nil},
			},
			toDelete: []*int64{
				aws.Int64(443),
			},
			toCreate: []*elb.Listener{
				{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("HTTP")},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			additions, removals := syncElbListeners(test.loadBalancerName, test.listeners, test.listenerDescriptions)
			assert.Equal(t, additions, test.toCreate)
			assert.Equal(t, removals, test.toDelete)
		})
	}
}

func TestElbListenersAreEqual(t *testing.T) {
	tests := []struct {
		name             string
		expected, actual *elb.Listener
		equal            bool
	}{
		{
			name:     "should be equal",
			expected: &elb.Listener{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			actual:   &elb.Listener{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			equal:    true,
		},
		{
			name:     "instance port should be different",
			expected: &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			actual:   &elb.Listener{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			equal:    false,
		},
		{
			name:     "instance protocol should be different",
			expected: &elb.Listener{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("HTTP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			actual:   &elb.Listener{InstancePort: aws.Int64(80), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			equal:    false,
		},
		{
			name:     "load balancer port should be different",
			expected: &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(443), Protocol: aws.String("TCP")},
			actual:   &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			equal:    false,
		},
		{
			name:     "protocol should be different",
			expected: &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("TCP")},
			actual:   &elb.Listener{InstancePort: aws.Int64(443), InstanceProtocol: aws.String("TCP"), LoadBalancerPort: aws.Int64(80), Protocol: aws.String("HTTP")},
			equal:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.equal, elbListenersAreEqual(test.expected, test.actual))
		})
	}
}

func TestCloud_chunkTargetDescriptions(t *testing.T) {
	type args struct {
		targets   []*elbv2.TargetDescription
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want [][]*elbv2.TargetDescription
	}{
		{
			name: "can be evenly chunked",
			args: args{
				targets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
				chunkSize: 2,
			},
			want: [][]*elbv2.TargetDescription{
				{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
				},
				{
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
			},
		},
		{
			name: "cannot be evenly chunked",
			args: args{
				targets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
				chunkSize: 3,
			},
			want: [][]*elbv2.TargetDescription{
				{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
				},
				{

					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
			},
		},
		{
			name: "chunkSize equal to total count",
			args: args{
				targets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
				chunkSize: 4,
			},
			want: [][]*elbv2.TargetDescription{
				{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
			},
		},
		{
			name: "chunkSize greater than total count",
			args: args{
				targets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
				chunkSize: 10,
			},
			want: [][]*elbv2.TargetDescription{
				{
					{
						Id:   aws.String("i-abcdefg1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg3"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdefg4"),
						Port: aws.Int64(8080),
					},
				},
			},
		},
		{
			name: "chunk nil slice",
			args: args{
				targets:   nil,
				chunkSize: 2,
			},
			want: nil,
		},
		{
			name: "chunk empty slice",
			args: args{
				targets:   []*elbv2.TargetDescription{},
				chunkSize: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cloud{}
			got := c.chunkTargetDescriptions(tt.args.targets, tt.args.chunkSize)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCloud_diffTargetGroupTargets(t *testing.T) {
	type args struct {
		expectedTargets []*elbv2.TargetDescription
		actualTargets   []*elbv2.TargetDescription
	}
	tests := []struct {
		name                    string
		args                    args
		wantTargetsToRegister   []*elbv2.TargetDescription
		wantTargetsToDeregister []*elbv2.TargetDescription
	}{
		{
			name: "all targets to register",
			args: args{
				expectedTargets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdef1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef2"),
						Port: aws.Int64(8080),
					},
				},
				actualTargets: nil,
			},
			wantTargetsToRegister: []*elbv2.TargetDescription{
				{
					Id:   aws.String("i-abcdef1"),
					Port: aws.Int64(8080),
				},
				{
					Id:   aws.String("i-abcdef2"),
					Port: aws.Int64(8080),
				},
			},
			wantTargetsToDeregister: nil,
		},
		{
			name: "all targets to deregister",
			args: args{
				expectedTargets: nil,
				actualTargets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdef1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef2"),
						Port: aws.Int64(8080),
					},
				},
			},
			wantTargetsToRegister: nil,
			wantTargetsToDeregister: []*elbv2.TargetDescription{
				{
					Id:   aws.String("i-abcdef1"),
					Port: aws.Int64(8080),
				},
				{
					Id:   aws.String("i-abcdef2"),
					Port: aws.Int64(8080),
				},
			},
		},
		{
			name: "some targets to register and deregister",
			args: args{
				expectedTargets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdef1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef4"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef5"),
						Port: aws.Int64(8080),
					},
				},
				actualTargets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdef1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef3"),
						Port: aws.Int64(8080),
					},
				},
			},
			wantTargetsToRegister: []*elbv2.TargetDescription{
				{
					Id:   aws.String("i-abcdef4"),
					Port: aws.Int64(8080),
				},
				{
					Id:   aws.String("i-abcdef5"),
					Port: aws.Int64(8080),
				},
			},
			wantTargetsToDeregister: []*elbv2.TargetDescription{
				{
					Id:   aws.String("i-abcdef2"),
					Port: aws.Int64(8080),
				},
				{
					Id:   aws.String("i-abcdef3"),
					Port: aws.Int64(8080),
				},
			},
		},
		{
			name: "both expected and actual targets are empty",
			args: args{
				expectedTargets: nil,
				actualTargets:   nil,
			},
			wantTargetsToRegister:   nil,
			wantTargetsToDeregister: nil,
		},
		{
			name: "expected and actual targets equals",
			args: args{
				expectedTargets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdef1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef3"),
						Port: aws.Int64(8080),
					},
				},
				actualTargets: []*elbv2.TargetDescription{
					{
						Id:   aws.String("i-abcdef1"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef2"),
						Port: aws.Int64(8080),
					},
					{
						Id:   aws.String("i-abcdef3"),
						Port: aws.Int64(8080),
					},
				},
			},
			wantTargetsToRegister:   nil,
			wantTargetsToDeregister: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cloud{}
			gotTargetsToRegister, gotTargetsToDeregister := c.diffTargetGroupTargets(tt.args.expectedTargets, tt.args.actualTargets)
			assert.Equal(t, tt.wantTargetsToRegister, gotTargetsToRegister)
			assert.Equal(t, tt.wantTargetsToDeregister, gotTargetsToDeregister)
		})
	}
}

func TestCloud_computeTargetGroupExpectedTargets(t *testing.T) {
	type args struct {
		instanceIDs []string
		port        int64
	}
	tests := []struct {
		name string
		args args
		want []*elbv2.TargetDescription
	}{
		{
			name: "no instance",
			args: args{
				instanceIDs: nil,
				port:        8080,
			},
			want: []*elbv2.TargetDescription{},
		},
		{
			name: "one instance",
			args: args{
				instanceIDs: []string{"i-abcdef1"},
				port:        8080,
			},
			want: []*elbv2.TargetDescription{
				{
					Id:   aws.String("i-abcdef1"),
					Port: aws.Int64(8080),
				},
			},
		},
		{
			name: "multiple instances",
			args: args{
				instanceIDs: []string{"i-abcdef1", "i-abcdef2", "i-abcdef3"},
				port:        8080,
			},
			want: []*elbv2.TargetDescription{
				{
					Id:   aws.String("i-abcdef1"),
					Port: aws.Int64(8080),
				},
				{
					Id:   aws.String("i-abcdef2"),
					Port: aws.Int64(8080),
				},
				{
					Id:   aws.String("i-abcdef3"),
					Port: aws.Int64(8080),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cloud{}
			got := c.computeTargetGroupExpectedTargets(tt.args.instanceIDs, tt.args.port)
			assert.Equal(t, tt.want, got)
		})
	}
}
