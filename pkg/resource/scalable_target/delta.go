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

package scalable_target

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
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
	customSetDefaults(a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.MaxCapacity, b.ko.Spec.MaxCapacity) {
		delta.Add("Spec.MaxCapacity", a.ko.Spec.MaxCapacity, b.ko.Spec.MaxCapacity)
	} else if a.ko.Spec.MaxCapacity != nil && b.ko.Spec.MaxCapacity != nil {
		if *a.ko.Spec.MaxCapacity != *b.ko.Spec.MaxCapacity {
			delta.Add("Spec.MaxCapacity", a.ko.Spec.MaxCapacity, b.ko.Spec.MaxCapacity)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MinCapacity, b.ko.Spec.MinCapacity) {
		delta.Add("Spec.MinCapacity", a.ko.Spec.MinCapacity, b.ko.Spec.MinCapacity)
	} else if a.ko.Spec.MinCapacity != nil && b.ko.Spec.MinCapacity != nil {
		if *a.ko.Spec.MinCapacity != *b.ko.Spec.MinCapacity {
			delta.Add("Spec.MinCapacity", a.ko.Spec.MinCapacity, b.ko.Spec.MinCapacity)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResourceID, b.ko.Spec.ResourceID) {
		delta.Add("Spec.ResourceID", a.ko.Spec.ResourceID, b.ko.Spec.ResourceID)
	} else if a.ko.Spec.ResourceID != nil && b.ko.Spec.ResourceID != nil {
		if *a.ko.Spec.ResourceID != *b.ko.Spec.ResourceID {
			delta.Add("Spec.ResourceID", a.ko.Spec.ResourceID, b.ko.Spec.ResourceID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RoleARN, b.ko.Spec.RoleARN) {
		delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
	} else if a.ko.Spec.RoleARN != nil && b.ko.Spec.RoleARN != nil {
		if *a.ko.Spec.RoleARN != *b.ko.Spec.RoleARN {
			delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ScalableDimension, b.ko.Spec.ScalableDimension) {
		delta.Add("Spec.ScalableDimension", a.ko.Spec.ScalableDimension, b.ko.Spec.ScalableDimension)
	} else if a.ko.Spec.ScalableDimension != nil && b.ko.Spec.ScalableDimension != nil {
		if *a.ko.Spec.ScalableDimension != *b.ko.Spec.ScalableDimension {
			delta.Add("Spec.ScalableDimension", a.ko.Spec.ScalableDimension, b.ko.Spec.ScalableDimension)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ServiceNamespace, b.ko.Spec.ServiceNamespace) {
		delta.Add("Spec.ServiceNamespace", a.ko.Spec.ServiceNamespace, b.ko.Spec.ServiceNamespace)
	} else if a.ko.Spec.ServiceNamespace != nil && b.ko.Spec.ServiceNamespace != nil {
		if *a.ko.Spec.ServiceNamespace != *b.ko.Spec.ServiceNamespace {
			delta.Add("Spec.ServiceNamespace", a.ko.Spec.ServiceNamespace, b.ko.Spec.ServiceNamespace)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SuspendedState, b.ko.Spec.SuspendedState) {
		delta.Add("Spec.SuspendedState", a.ko.Spec.SuspendedState, b.ko.Spec.SuspendedState)
	} else if a.ko.Spec.SuspendedState != nil && b.ko.Spec.SuspendedState != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.SuspendedState.DynamicScalingInSuspended, b.ko.Spec.SuspendedState.DynamicScalingInSuspended) {
			delta.Add("Spec.SuspendedState.DynamicScalingInSuspended", a.ko.Spec.SuspendedState.DynamicScalingInSuspended, b.ko.Spec.SuspendedState.DynamicScalingInSuspended)
		} else if a.ko.Spec.SuspendedState.DynamicScalingInSuspended != nil && b.ko.Spec.SuspendedState.DynamicScalingInSuspended != nil {
			if *a.ko.Spec.SuspendedState.DynamicScalingInSuspended != *b.ko.Spec.SuspendedState.DynamicScalingInSuspended {
				delta.Add("Spec.SuspendedState.DynamicScalingInSuspended", a.ko.Spec.SuspendedState.DynamicScalingInSuspended, b.ko.Spec.SuspendedState.DynamicScalingInSuspended)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.SuspendedState.DynamicScalingOutSuspended, b.ko.Spec.SuspendedState.DynamicScalingOutSuspended) {
			delta.Add("Spec.SuspendedState.DynamicScalingOutSuspended", a.ko.Spec.SuspendedState.DynamicScalingOutSuspended, b.ko.Spec.SuspendedState.DynamicScalingOutSuspended)
		} else if a.ko.Spec.SuspendedState.DynamicScalingOutSuspended != nil && b.ko.Spec.SuspendedState.DynamicScalingOutSuspended != nil {
			if *a.ko.Spec.SuspendedState.DynamicScalingOutSuspended != *b.ko.Spec.SuspendedState.DynamicScalingOutSuspended {
				delta.Add("Spec.SuspendedState.DynamicScalingOutSuspended", a.ko.Spec.SuspendedState.DynamicScalingOutSuspended, b.ko.Spec.SuspendedState.DynamicScalingOutSuspended)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.SuspendedState.ScheduledScalingSuspended, b.ko.Spec.SuspendedState.ScheduledScalingSuspended) {
			delta.Add("Spec.SuspendedState.ScheduledScalingSuspended", a.ko.Spec.SuspendedState.ScheduledScalingSuspended, b.ko.Spec.SuspendedState.ScheduledScalingSuspended)
		} else if a.ko.Spec.SuspendedState.ScheduledScalingSuspended != nil && b.ko.Spec.SuspendedState.ScheduledScalingSuspended != nil {
			if *a.ko.Spec.SuspendedState.ScheduledScalingSuspended != *b.ko.Spec.SuspendedState.ScheduledScalingSuspended {
				delta.Add("Spec.SuspendedState.ScheduledScalingSuspended", a.ko.Spec.SuspendedState.ScheduledScalingSuspended, b.ko.Spec.SuspendedState.ScheduledScalingSuspended)
			}
		}
	}

	return delta
}
