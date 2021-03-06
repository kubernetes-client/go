/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// CustomResourceSubresources defines the status and scale subresources for CustomResources.
type V1beta1CustomResourceSubresources struct {

	// Scale denotes the scale subresource for CustomResources
	Scale *V1beta1CustomResourceSubresourceScale `json:"scale,omitempty"`

	// Status denotes the status subresource for CustomResources
	Status *interface{} `json:"status,omitempty"`
}
