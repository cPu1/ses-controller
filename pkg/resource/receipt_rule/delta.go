// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package receipt_rule

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.After, b.ko.Spec.After) {
		delta.Add("Spec.After", a.ko.Spec.After, b.ko.Spec.After)
	} else if a.ko.Spec.After != nil && b.ko.Spec.After != nil {
		if *a.ko.Spec.After != *b.ko.Spec.After {
			delta.Add("Spec.After", a.ko.Spec.After, b.ko.Spec.After)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Rule, b.ko.Spec.Rule) {
		delta.Add("Spec.Rule", a.ko.Spec.Rule, b.ko.Spec.Rule)
	} else if a.ko.Spec.Rule != nil && b.ko.Spec.Rule != nil {
		if len(a.ko.Spec.Rule.Actions) != len(b.ko.Spec.Rule.Actions) {
			delta.Add("Spec.Rule.Actions", a.ko.Spec.Rule.Actions, b.ko.Spec.Rule.Actions)
		} else if len(a.ko.Spec.Rule.Actions) > 0 {
			if !reflect.DeepEqual(a.ko.Spec.Rule.Actions, b.ko.Spec.Rule.Actions) {
				delta.Add("Spec.Rule.Actions", a.ko.Spec.Rule.Actions, b.ko.Spec.Rule.Actions)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Rule.Enabled, b.ko.Spec.Rule.Enabled) {
			delta.Add("Spec.Rule.Enabled", a.ko.Spec.Rule.Enabled, b.ko.Spec.Rule.Enabled)
		} else if a.ko.Spec.Rule.Enabled != nil && b.ko.Spec.Rule.Enabled != nil {
			if *a.ko.Spec.Rule.Enabled != *b.ko.Spec.Rule.Enabled {
				delta.Add("Spec.Rule.Enabled", a.ko.Spec.Rule.Enabled, b.ko.Spec.Rule.Enabled)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Rule.Name, b.ko.Spec.Rule.Name) {
			delta.Add("Spec.Rule.Name", a.ko.Spec.Rule.Name, b.ko.Spec.Rule.Name)
		} else if a.ko.Spec.Rule.Name != nil && b.ko.Spec.Rule.Name != nil {
			if *a.ko.Spec.Rule.Name != *b.ko.Spec.Rule.Name {
				delta.Add("Spec.Rule.Name", a.ko.Spec.Rule.Name, b.ko.Spec.Rule.Name)
			}
		}
		if len(a.ko.Spec.Rule.Recipients) != len(b.ko.Spec.Rule.Recipients) {
			delta.Add("Spec.Rule.Recipients", a.ko.Spec.Rule.Recipients, b.ko.Spec.Rule.Recipients)
		} else if len(a.ko.Spec.Rule.Recipients) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.Rule.Recipients, b.ko.Spec.Rule.Recipients) {
				delta.Add("Spec.Rule.Recipients", a.ko.Spec.Rule.Recipients, b.ko.Spec.Rule.Recipients)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Rule.ScanEnabled, b.ko.Spec.Rule.ScanEnabled) {
			delta.Add("Spec.Rule.ScanEnabled", a.ko.Spec.Rule.ScanEnabled, b.ko.Spec.Rule.ScanEnabled)
		} else if a.ko.Spec.Rule.ScanEnabled != nil && b.ko.Spec.Rule.ScanEnabled != nil {
			if *a.ko.Spec.Rule.ScanEnabled != *b.ko.Spec.Rule.ScanEnabled {
				delta.Add("Spec.Rule.ScanEnabled", a.ko.Spec.Rule.ScanEnabled, b.ko.Spec.Rule.ScanEnabled)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Rule.TLSPolicy, b.ko.Spec.Rule.TLSPolicy) {
			delta.Add("Spec.Rule.TLSPolicy", a.ko.Spec.Rule.TLSPolicy, b.ko.Spec.Rule.TLSPolicy)
		} else if a.ko.Spec.Rule.TLSPolicy != nil && b.ko.Spec.Rule.TLSPolicy != nil {
			if *a.ko.Spec.Rule.TLSPolicy != *b.ko.Spec.Rule.TLSPolicy {
				delta.Add("Spec.Rule.TLSPolicy", a.ko.Spec.Rule.TLSPolicy, b.ko.Spec.Rule.TLSPolicy)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RuleSetName, b.ko.Spec.RuleSetName) {
		delta.Add("Spec.RuleSetName", a.ko.Spec.RuleSetName, b.ko.Spec.RuleSetName)
	} else if a.ko.Spec.RuleSetName != nil && b.ko.Spec.RuleSetName != nil {
		if *a.ko.Spec.RuleSetName != *b.ko.Spec.RuleSetName {
			delta.Add("Spec.RuleSetName", a.ko.Spec.RuleSetName, b.ko.Spec.RuleSetName)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.RuleSetRef, b.ko.Spec.RuleSetRef) {
		delta.Add("Spec.RuleSetRef", a.ko.Spec.RuleSetRef, b.ko.Spec.RuleSetRef)
	}

	return delta
}