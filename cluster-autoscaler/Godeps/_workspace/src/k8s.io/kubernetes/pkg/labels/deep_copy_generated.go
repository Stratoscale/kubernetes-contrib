// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package labels

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
	sets "k8s.io/kubernetes/pkg/util/sets"
)

func DeepCopy_labels_Requirement(in Requirement, out *Requirement, c *conversion.Cloner) error {
	out.key = in.key
	out.operator = in.operator
	if in.strValues != nil {
		in, out := in.strValues, &out.strValues
		*out = make(sets.String)
		for key, val := range in {
			newVal := new(sets.Empty)
			if err := sets.DeepCopy_sets_Empty(val, newVal, c); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.strValues = nil
	}
	return nil
}