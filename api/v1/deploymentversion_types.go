/*
Copyright 2021.

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

package v1

import (
	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeploymentVersionSpec defines the desired state of DeploymentVersion
type DeploymentVersionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// +optional
	TestProp string `json:"testProp,omitempty"`
	// +optional
	DeploymentSpec apps.DeploymentSpec `json:"deploymentSpec,omitempty"`
}

// DeploymentVersionStatus defines the observed state of DeploymentVersion
type DeploymentVersionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DeploymentVersion is the Schema for the deploymentversions API
type DeploymentVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentVersionSpec   `json:"spec,omitempty"`
	Status DeploymentVersionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeploymentVersionList contains a list of DeploymentVersion
type DeploymentVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeploymentVersion{}, &DeploymentVersionList{})
}
