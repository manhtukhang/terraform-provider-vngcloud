/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ReorderPoliciesRequest struct {
	// Listener's id of the policies
	ListenerId string `json:"listenerId,omitempty"`
	// List of policies to reorder
	Policies []ReorderPolicyRequest `json:"policies,omitempty"`
}
