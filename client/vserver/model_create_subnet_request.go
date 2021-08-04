/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateSubnetRequest struct {
	// CIDR of Subnet, must be contained in its network.
	Cidr  string       `json:"cidr"`
	Extra *interface{} `json:"extra,omitempty"`
	// Name of the Subnet
	Name string `json:"name"`
	// Id of Network.
	NetworkId string `json:"networkId"`
}
