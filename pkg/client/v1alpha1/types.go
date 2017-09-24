package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Artillery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArtillerySpec    `json:"spec"`
	Status            *ArtilleryStatus `json:"status,omitempty"`
}

type ArtillerySpec struct {
	BaseImage  string                 `json:"baseImage,omitempty"`
	TestScript map[string]interface{} `json:"testScript"`
}

type ArtilleryStatus struct {
}

type ArtilleryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []*Artillery `json:"items"`
}
