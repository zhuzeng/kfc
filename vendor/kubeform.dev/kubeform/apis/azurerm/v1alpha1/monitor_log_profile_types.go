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

type MonitorLogProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorLogProfileSpec   `json:"spec,omitempty"`
	Status            MonitorLogProfileStatus `json:"status,omitempty"`
}

type MonitorLogProfileSpecRetentionPolicy struct {
	// +optional
	Days    int  `json:"days,omitempty" tf:"days,omitempty"`
	Enabled bool `json:"enabled" tf:"enabled"`
}

type MonitorLogProfileSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	Categories []string `json:"categories" tf:"categories"`
	// +kubebuilder:validation:MinItems=1
	Locations []string `json:"locations" tf:"locations"`
	Name      string   `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	RetentionPolicy []MonitorLogProfileSpecRetentionPolicy `json:"retentionPolicy" tf:"retention_policy"`
	// +optional
	ServicebusRuleID string `json:"servicebusRuleID,omitempty" tf:"servicebus_rule_id,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
}

type MonitorLogProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorLogProfileSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorLogProfileList is a list of MonitorLogProfiles
type MonitorLogProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorLogProfile CRD objects
	Items []MonitorLogProfile `json:"items,omitempty"`
}
