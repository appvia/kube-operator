package v1

import (
	configv1 "github.com/appvia/hub-apis/pkg/apis/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NamespaceClaimSpec defines the desired state of NamespaceClaim
// +k8s:openapi-gen=true
type NamespaceClaimSpec struct {
	// Name is the name of the namespace to create
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// Cluster is the owner of the namespace i.e. the cluster
	// +kubebuilder:validation:Required
	Cluster configv1.Ownership `json:"cluster"`
	// Team is the team associated to the namespace (technically this can be inferred
	// from the namespace, but i don't like that idea
	// +kubebuilder:validation:Required
	Team configv1.Ownership `json:"team"`
	// Workspace is the workspace ownership reference
	// +kubebuilder:validation:Required
	Workspace configv1.Ownership `json:"workspace"`
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
	Status metav1.StatusReason `json:"status"`
	// Conditions is a series of things that caused the failure if any
	Conditions []metav1.Status `json:"conditions"`
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
