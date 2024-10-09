/*
Copyright 2023 The Kubernetes Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrefetchConfigSpec defines the desired state of PrefetchConfig
type PrefetchRequestSpec struct {

	// Datapath specifies the location of datasets to be
	// used when workload starts
	Datapath string `json:"datapath,omitempty"`

	// lppf: Parameters for prefetch state
	// Parameters contains all other parameters classes may require.
	//
	// +optional
	// +kubebuilder:validation:MaxProperties=100
	Parameters map[string]Parameter `json:"parameters,omitempty"`
}

// AdmissionCheckStatus defines the observed state of AdmissionCheck
type PrefetchRequestStatus struct {
	// observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Specify prefetch status: New, Prefetching, Incomplete, Completed, Failed
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AdmissionCheck is the Schema for the admissionchecks API
type PrefetchRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrefetchRequestSpec   `json:"spec,omitempty"`
	Status PrefetchRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AdmissionCheckList contains a list of AdmissionCheck
type PrefetchRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrefetchRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrefetchRequest{}, &PrefetchRequestList{})
}

// ////////////////////////////////////////////////////////////////////

// // // Parameter is limited to 255 characters.
// // // +kubebuilder:validation:MaxLength=255
// // type Parameter string

// // +genclient
// // +genclient:nonNamespaced
// // +kubebuilder:object:root=true
// // +kubebuilder:storageversion
// // +kubebuilder:resource:scope=Cluster

// // Schema for the prefetchconfig API
// type PrefetchConfig struct {
// 	metav1.TypeMeta   `json:",inline"`
// 	metav1.ObjectMeta `json:"metadata,omitempty"`

// 	Spec PrefetchConfigSpec `json:"spec,omitempty"`
// }

// // +kubebuilder:object:root=true

// // PrefetchConfigList contains a list of PrefetchConfig
// type PrefetchConfigList struct {
// 	metav1.TypeMeta `json:",inline"`
// 	metav1.ListMeta `json:"metadata,omitempty"`
// 	Items           []PrefetchConfig `json:"items"`
// }

// func init() {
// 	SchemeBuilder.Register(&PrefetchConfig{}, &PrefetchConfigList{})
// }
