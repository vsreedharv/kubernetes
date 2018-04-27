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

package v1beta1

import (
	"testing"

	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/kubernetes/pkg/apis/core/helper"
)

func role(rules []rbacv1beta1.PolicyRule, labels map[string]string, annotations map[string]string) *rbacv1beta1.ClusterRole {
	return &rbacv1beta1.ClusterRole{
		Rules:      rules,
		ObjectMeta: metav1.ObjectMeta{Labels: labels, Annotations: annotations},
	}
}

func rules(resources ...string) []rbacv1beta1.PolicyRule {
	r := []rbacv1beta1.PolicyRule{}
	for _, resource := range resources {
		r = append(r, rbacv1beta1.PolicyRule{APIGroups: []string{""}, Verbs: []string{"get"}, Resources: []string{resource}})
	}
	return r
}

type ss map[string]string

func TestComputeReconciledRoleRules(t *testing.T) {
	tests := map[string]struct {
		expectedRole           *rbacv1beta1.ClusterRole
		actualRole             *rbacv1beta1.ClusterRole
		removeExtraPermissions bool

		expectedReconciledRole       *rbacv1beta1.ClusterRole
		expectedReconciliationNeeded bool
	}{
		"empty": {
			expectedRole:           role(rules(), nil, nil),
			actualRole:             role(rules(), nil, nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"match without union": {
			expectedRole:           role(rules("a"), nil, nil),
			actualRole:             role(rules("a"), nil, nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"match with union": {
			expectedRole:           role(rules("a"), nil, nil),
			actualRole:             role(rules("a"), nil, nil),
			removeExtraPermissions: false,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"different rules without union": {
			expectedRole:           role(rules("a"), nil, nil),
			actualRole:             role(rules("b"), nil, nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), nil, nil),
			expectedReconciliationNeeded: true,
		},
		"different rules with union": {
			expectedRole:           role(rules("a"), nil, nil),
			actualRole:             role(rules("b"), nil, nil),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("b", "a"), nil, nil),
			expectedReconciliationNeeded: true,
		},
		"match labels without union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("a"), ss{"1": "a"}, nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"match labels with union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("a"), ss{"1": "a"}, nil),
			removeExtraPermissions: false,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"different labels without union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("a"), ss{"2": "b"}, nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), ss{"1": "a", "2": "b"}, nil),
			expectedReconciliationNeeded: true,
		},
		"different labels with union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("a"), ss{"2": "b"}, nil),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("a"), ss{"1": "a", "2": "b"}, nil),
			expectedReconciliationNeeded: true,
		},
		"different labels and rules without union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("b"), ss{"2": "b"}, nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), ss{"1": "a", "2": "b"}, nil),
			expectedReconciliationNeeded: true,
		},
		"different labels and rules with union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("b"), ss{"2": "b"}, nil),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("b", "a"), ss{"1": "a", "2": "b"}, nil),
			expectedReconciliationNeeded: true,
		},
		"conflicting labels and rules without union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("b"), ss{"1": "b"}, nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), ss{"1": "b"}, nil),
			expectedReconciliationNeeded: true,
		},
		"conflicting labels and rules with union": {
			expectedRole:           role(rules("a"), ss{"1": "a"}, nil),
			actualRole:             role(rules("b"), ss{"1": "b"}, nil),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("b", "a"), ss{"1": "b"}, nil),
			expectedReconciliationNeeded: true,
		},
		"match annotations without union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("a"), nil, ss{"1": "a"}),
			removeExtraPermissions: true,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"match annotations with union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("a"), nil, ss{"1": "a"}),
			removeExtraPermissions: false,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"different annotations without union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("a"), nil, ss{"2": "b"}),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), nil, ss{"1": "a", "2": "b"}),
			expectedReconciliationNeeded: true,
		},
		"different annotations with union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("a"), nil, ss{"2": "b"}),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("a"), nil, ss{"1": "a", "2": "b"}),
			expectedReconciliationNeeded: true,
		},
		"different annotations and rules without union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("b"), nil, ss{"2": "b"}),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), nil, ss{"1": "a", "2": "b"}),
			expectedReconciliationNeeded: true,
		},
		"different annotations and rules with union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("b"), nil, ss{"2": "b"}),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("b", "a"), nil, ss{"1": "a", "2": "b"}),
			expectedReconciliationNeeded: true,
		},
		"conflicting annotations and rules without union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("b"), nil, ss{"1": "b"}),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), nil, ss{"1": "b"}),
			expectedReconciliationNeeded: true,
		},
		"conflicting annotations and rules with union": {
			expectedRole:           role(rules("a"), nil, ss{"1": "a"}),
			actualRole:             role(rules("b"), nil, ss{"1": "b"}),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("b", "a"), nil, ss{"1": "b"}),
			expectedReconciliationNeeded: true,
		},
		"conflicting labels/annotations and rules without union": {
			expectedRole:           role(rules("a"), ss{"3": "d"}, ss{"1": "a"}),
			actualRole:             role(rules("b"), ss{"4": "e"}, ss{"1": "b"}),
			removeExtraPermissions: true,

			expectedReconciledRole:       role(rules("a"), ss{"3": "d", "4": "e"}, ss{"1": "b"}),
			expectedReconciliationNeeded: true,
		},
		"conflicting labels/annotations and rules with union": {
			expectedRole:           role(rules("a"), ss{"3": "d"}, ss{"1": "a"}),
			actualRole:             role(rules("b"), ss{"4": "e"}, ss{"1": "b"}),
			removeExtraPermissions: false,

			expectedReconciledRole:       role(rules("b", "a"), ss{"3": "d", "4": "e"}, ss{"1": "b"}),
			expectedReconciliationNeeded: true,
		},
		"complex labels/annotations and rules without union": {
			expectedRole:           role(rules("pods", "nodes", "secrets"), ss{"env": "prod", "color": "blue"}, ss{"description": "fancy", "system": "true"}),
			actualRole:             role(rules("nodes", "images", "projects"), ss{"color": "red", "team": "pm"}, ss{"system": "false", "owner": "admin", "vip": "yes"}),
			removeExtraPermissions: true,

			expectedReconciledRole: role(
				rules("pods", "nodes", "secrets"),
				ss{"env": "prod", "color": "red", "team": "pm"},
				ss{"description": "fancy", "system": "false", "owner": "admin", "vip": "yes"}),
			expectedReconciliationNeeded: true,
		},
		"complex labels/annotations and rules with union": {
			expectedRole:           role(rules("pods", "nodes", "secrets"), ss{"env": "prod", "color": "blue", "manager": "randy"}, ss{"description": "fancy", "system": "true", "up": "true"}),
			actualRole:             role(rules("nodes", "images", "projects"), ss{"color": "red", "team": "pm"}, ss{"system": "false", "owner": "admin", "vip": "yes", "rate": "down"}),
			removeExtraPermissions: false,

			expectedReconciledRole: role(
				rules("nodes", "images", "projects", "pods", "secrets"),
				ss{"env": "prod", "manager": "randy", "color": "red", "team": "pm"},
				ss{"description": "fancy", "system": "false", "owner": "admin", "vip": "yes", "rate": "down", "up": "true"}),
			expectedReconciliationNeeded: true,
		},
	}

	for k, tc := range tests {
		actualRole := ClusterRoleRuleOwner{ClusterRole: tc.actualRole}
		expectedRole := ClusterRoleRuleOwner{ClusterRole: tc.expectedRole}
		result, err := computeReconciledRole(actualRole, expectedRole, tc.removeExtraPermissions)
		if err != nil {
			t.Errorf("%s: %v", k, err)
			continue
		}
		reconciliationNeeded := result.Operation != ReconcileNone
		if reconciliationNeeded != tc.expectedReconciliationNeeded {
			t.Errorf("%s: Expected\n\t%v\ngot\n\t%v", k, tc.expectedReconciliationNeeded, reconciliationNeeded)
			continue
		}
		if reconciliationNeeded && !helper.Semantic.DeepEqual(result.Role.(ClusterRoleRuleOwner).ClusterRole, tc.expectedReconciledRole) {
			t.Errorf("%s: Expected\n\t%#v\ngot\n\t%#v", k, tc.expectedReconciledRole, result.Role)
		}
	}
}

func aggregatedRole(aggregationRule *rbacv1beta1.AggregationRule) *rbacv1beta1.ClusterRole {
	return &rbacv1beta1.ClusterRole{
		AggregationRule: aggregationRule,
	}
}

func aggregationrule(selectors []map[string]string) *rbacv1beta1.AggregationRule {
	ret := &rbacv1beta1.AggregationRule{}
	for _, selector := range selectors {
		ret.ClusterRoleSelectors = append(ret.ClusterRoleSelectors,
			metav1.LabelSelector{MatchLabels: selector})
	}
	return ret
}

func TestComputeReconciledRoleAggregationRules(t *testing.T) {
	tests := map[string]struct {
		expectedRole           *rbacv1beta1.ClusterRole
		actualRole             *rbacv1beta1.ClusterRole
		removeExtraPermissions bool

		expectedReconciledRole       *rbacv1beta1.ClusterRole
		expectedReconciliationNeeded bool
	}{
		"empty": {
			expectedRole:           aggregatedRole(&rbacv1beta1.AggregationRule{}),
			actualRole:             aggregatedRole(nil),
			removeExtraPermissions: true,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"empty-2": {
			expectedRole:           aggregatedRole(&rbacv1beta1.AggregationRule{}),
			actualRole:             aggregatedRole(&rbacv1beta1.AggregationRule{}),
			removeExtraPermissions: true,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"match without union": {
			expectedRole:           aggregatedRole(aggregationrule([]map[string]string{{"foo": "bar"}})),
			actualRole:             aggregatedRole(aggregationrule([]map[string]string{{"foo": "bar"}})),
			removeExtraPermissions: true,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"match with union": {
			expectedRole:           aggregatedRole(aggregationrule([]map[string]string{{"foo": "bar"}})),
			actualRole:             aggregatedRole(aggregationrule([]map[string]string{{"foo": "bar"}})),
			removeExtraPermissions: false,

			expectedReconciledRole:       nil,
			expectedReconciliationNeeded: false,
		},
		"different rules without union": {
			expectedRole:           aggregatedRole(aggregationrule([]map[string]string{{"foo": "bar"}})),
			actualRole:             aggregatedRole(aggregationrule([]map[string]string{{"alpha": "bravo"}})),
			removeExtraPermissions: true,

			expectedReconciledRole:       aggregatedRole(aggregationrule([]map[string]string{{"foo": "bar"}})),
			expectedReconciliationNeeded: true,
		},
		"different rules with union": {
			expectedRole:           aggregatedRole(aggregationrule([]map[string]string{{"foo": "bar"}})),
			actualRole:             aggregatedRole(aggregationrule([]map[string]string{{"alpha": "bravo"}})),
			removeExtraPermissions: false,

			expectedReconciledRole:       aggregatedRole(aggregationrule([]map[string]string{{"alpha": "bravo"}, {"foo": "bar"}})),
			expectedReconciliationNeeded: true,
		},
	}

	for k, tc := range tests {
		actualRole := ClusterRoleRuleOwner{ClusterRole: tc.actualRole}
		expectedRole := ClusterRoleRuleOwner{ClusterRole: tc.expectedRole}
		result, err := computeReconciledRole(actualRole, expectedRole, tc.removeExtraPermissions)
		if err != nil {
			t.Errorf("%s: %v", k, err)
			continue
		}
		reconciliationNeeded := result.Operation != ReconcileNone
		if reconciliationNeeded != tc.expectedReconciliationNeeded {
			t.Errorf("%s: Expected\n\t%v\ngot\n\t%v", k, tc.expectedReconciliationNeeded, reconciliationNeeded)
			continue
		}
		if reconciliationNeeded && !helper.Semantic.DeepEqual(result.Role.(ClusterRoleRuleOwner).ClusterRole, tc.expectedReconciledRole) {
			t.Errorf("%s: %v", k, diff.ObjectDiff(tc.expectedReconciledRole, result.Role.(ClusterRoleRuleOwner).ClusterRole))
		}
	}
}
