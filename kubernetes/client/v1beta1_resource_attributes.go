/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface
type V1beta1ResourceAttributes struct {

	// Group is the API Group of the Resource.  \"*\" means all.
	Group string `json:"group,omitempty"`

	// Name is the name of the resource being requested for a \"get\" or deleted for a \"delete\". \"\" (empty) means all.
	Name string `json:"name,omitempty"`

	// Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces \"\" (empty) is defaulted for LocalSubjectAccessReviews \"\" (empty) is empty for cluster-scoped resources \"\" (empty) means \"all\" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
	Namespace string `json:"namespace,omitempty"`

	// Resource is one of the existing resource types.  \"*\" means all.
	Resource string `json:"resource,omitempty"`

	// Subresource is one of the existing resource types.  \"\" means none.
	Subresource string `json:"subresource,omitempty"`

	// Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  \"*\" means all.
	Verb string `json:"verb,omitempty"`

	// Version is the API Version of the Resource.  \"*\" means all.
	Version string `json:"version,omitempty"`
}
