/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type LoadBalancer struct {
	Address      string `json:"address,omitempty"`
	AdminStateUp bool   `json:"adminStateUp,omitempty"`
	// CreatedAt time.Time `json:"createdAt,omitempty"`
	Description        string `json:"description,omitempty"`
	LoadBalancerSchema string `json:"loadBalancerSchema,omitempty"`
	Name               string `json:"name,omitempty"`
	OperatingStatus    string `json:"operatingStatus,omitempty"`
	PackageId          string `json:"packageId,omitempty"`
	PrivateSubnetCidr  string `json:"privateSubnetCidr,omitempty"`
	PrivateSubnetId    string `json:"privateSubnetId,omitempty"`
	ProjectId          string `json:"projectId,omitempty"`
	ProvisioningStatus string `json:"provisioningStatus,omitempty"`
	Status             string `json:"status,omitempty"`
	SubnetId           string `json:"subnetId,omitempty"`
	Type_              string `json:"type,omitempty"`
	// UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
