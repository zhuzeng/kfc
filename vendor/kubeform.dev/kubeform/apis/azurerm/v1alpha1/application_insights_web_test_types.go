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

type ApplicationInsightsWebTest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationInsightsWebTestSpec   `json:"spec,omitempty"`
	Status            ApplicationInsightsWebTestStatus `json:"status,omitempty"`
}

type ApplicationInsightsWebTestSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApplicationInsightsID string `json:"applicationInsightsID" tf:"application_insights_id"`
	Configuration         string `json:"configuration" tf:"configuration"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Frequency int `json:"frequency,omitempty" tf:"frequency,omitempty"`
	// +kubebuilder:validation:MinItems=1
	GeoLocations      []string `json:"geoLocations" tf:"geo_locations"`
	Kind              string   `json:"kind" tf:"kind"`
	Location          string   `json:"location" tf:"location"`
	Name              string   `json:"name" tf:"name"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RetryEnabled bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`
	// +optional
	SyntheticMonitorID string `json:"syntheticMonitorID,omitempty" tf:"synthetic_monitor_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Timeout int `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type ApplicationInsightsWebTestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApplicationInsightsWebTestSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApplicationInsightsWebTestList is a list of ApplicationInsightsWebTests
type ApplicationInsightsWebTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationInsightsWebTest CRD objects
	Items []ApplicationInsightsWebTest `json:"items,omitempty"`
}
