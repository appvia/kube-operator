package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NamespaceClaimPolicySpec defines the desired state of NamespaceClaimPolicy
// +k8s:openapi-gen=true
type NamespaceClaimPolicySpec struct {
	// DefaultLabels is series of default labels to place on the namespace
	DefaultLabels map[string]string `json:"defaultLabels,omitempty"`
	// DefaultAnnotations is a series of annotations which are places on all
	// namespaces created by a claim
	DefaultAnnotations map[string]string `json:"defaultAnnotations,omitempty"`
	// @TODO we need to add limits, requests and quotas here
}

// NamespaceClaimPolicyStatus defines the observed state of NamespaceClaimPolicy
// +k8s:openapi-gen=true
type NamespaceClaimPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceClaimPolicy is the Schema for the namespaceclaimpolicies API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=namespaceclaimpolicies,scope=Namespaced
type NamespaceClaimPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceClaimPolicySpec   `json:"spec,omitempty"`
	Status NamespaceClaimPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NamespaceClaimPolicyList contains a list of NamespaceClaimPolicy
type NamespaceClaimPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceClaimPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NamespaceClaimPolicy{}, &NamespaceClaimPolicyList{})
}
