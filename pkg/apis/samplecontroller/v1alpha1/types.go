package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Foo struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	metav1.TypeMeta   `"json:,inline"`
	Spec              FooSpec   `json:"spec"`
	Status            FooStatus `json:"status"`
}

// FooSpec is the spec for a Foo resource
type FooSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// ForrStatus is the status for Foo resource
type FooStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FooList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Items []Foo `json:"items"`
}
