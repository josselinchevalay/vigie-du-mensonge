package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Structs des composants
type FrontendSpec struct {
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
}

type BackendSpec struct {
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
}

type DatabaseSpec struct {
	Storage           string `json:"storage"`
	User              string `json:"username"`
	Image             string `json:"image"`
	PasswordSecretRef string `json:"passwordSecretRef"`
}

type DataImportSpec struct {
	Enabled bool   `json:"enabled"`
	Image   string `json:"image"`
}

type IngressSpec struct {
	Enabled      bool   `json:"enabled"`
	Domain       string `json:"domain"`
	TlsSecretRef string `json:"tlsSecretRef"`
}

// Spec de VigieApp
type VigieAppSpec struct {
	Frontend   FrontendSpec   `json:"frontend"`
	Backend    BackendSpec    `json:"backend"`
	Database   DatabaseSpec   `json:"database"`
	DataImport DataImportSpec `json:"dataImport"`
	Ingress    IngressSpec    `json:"ingress"`
}

// Status de VigieApp
type VigieAppStatus struct {
	Ready        bool   `json:"ready"`
	BackendPort  int32  `json:"backendPort"`
	FrontendPort int32  `json:"frontendPort"`
	Message      string `json:"message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VigieApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VigieAppSpec   `json:"spec"`
	Status VigieAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type VigieAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VigieApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VigieApp{}, &VigieAppList{})
}
