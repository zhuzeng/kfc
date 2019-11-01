/*
Copyright The Kubeform Authors.

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type EmrInstanceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrInstanceGroupSpec   `json:"spec,omitempty"`
	Status            EmrInstanceGroupStatus `json:"status,omitempty"`
}

type EmrInstanceGroupSpecEbsConfig struct {
	// +optional
	Iops int    `json:"iops,omitempty" tf:"iops,omitempty"`
	Size int    `json:"size" tf:"size"`
	Type string `json:"type" tf:"type"`
	// +optional
	VolumesPerInstance int `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type EmrInstanceGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterID string `json:"clusterID" tf:"cluster_id"`
	// +optional
	EbsConfig []EmrInstanceGroupSpecEbsConfig `json:"ebsConfig,omitempty" tf:"ebs_config,omitempty"`
	// +optional
	EbsOptimized bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized,omitempty"`
	// +optional
	InstanceCount int    `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`
	InstanceType  string `json:"instanceType" tf:"instance_type"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	RunningInstanceCount int `json:"runningInstanceCount,omitempty" tf:"running_instance_count,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
}

type EmrInstanceGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EmrInstanceGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EmrInstanceGroupList is a list of EmrInstanceGroups
type EmrInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EmrInstanceGroup CRD objects
	Items []EmrInstanceGroup `json:"items,omitempty"`
}
