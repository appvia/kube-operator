package v1

import (
	corev1 "github.com/appvia/hub-apis/pkg/apis/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NamespaceClaimSpec defines the desired state of NamespaceClaim
// +k8s:openapi-gen=true
type NamespaceClaimSpec struct {
	// Use is the owner of the namespace i.e. the cluster
	// +kubebuilder:validation:Required
	Use corev1.Ownership `json:"use"`
	// Name is the name of the namespace to create
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// AnnotationsLabels is a series of annotations on the namespace
	AnnotationsLabels map[string]string `json:"annotationsLabels"`
	// NamespaceLabels is a series of labels for the namespace
	NamespaceLabels map[string]string `json:"namespaceLabels"`
	// @TODO we need to add limits, quotas etc
}

// NamespaceClaimStatus defines the observed state of NamespaceClaim
// +k8s:openapi-gen=true
type NamespaceClaimStatus struct {
	// Status is the status of the namespace
	Status corev1.Status `json:"status"`
	// Conditions is a series of things that caused the failure if any
	// +listType
	Conditions []corev1.Condition `json:"conditions"`
	// Phase is used to hold the current phase of the resource
	Phase string `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceClaim is the Schema for the namespaceclaims API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=namespaceclaims,scope=Namespaced
type NamespaceClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceClaimSpec   `json:"spec,omitempty"`
	Status NamespaceClaimStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceClaimList contains a list of NamespaceClaim
type NamespaceClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceClaim `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NamespaceClaim{}, &NamespaceClaimList{})
}
