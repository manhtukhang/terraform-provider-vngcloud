/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type Listener struct {
	AdminStateUp    bool   `json:"adminStateUp,omitempty"`
	AllowedCidrs    string `json:"allowedCidrs,omitempty"`
	ConnectionLimit int32  `json:"connectionLimit,omitempty"`
	// CreatedAt time.Time `json:"createdAt,omitempty"`
	DefaultPoolId          string   `json:"defaultPoolId,omitempty"`
	DefaultTlsContainerRef string   `json:"defaultTlsContainerRef,omitempty"`
	Description            string   `json:"description,omitempty"`
	Headers                []string `json:"headers,omitempty"`
	LoadBalancerId         string   `json:"loadBalancerId,omitempty"`
	Name                   string   `json:"name,omitempty"`
	OperatingStatus        string   `json:"operatingStatus,omitempty"`
	ProjectId              string   `json:"projectId,omitempty"`
	Protocol               string   `json:"protocol,omitempty"`
	ProtocolPort           int32    `json:"protocolPort,omitempty"`
	ProvisioningStatus     string   `json:"provisioningStatus,omitempty"`
	Status                 string   `json:"status,omitempty"`
	TimeoutClient          int32    `json:"timeoutClient,omitempty"`
	TimeoutConnection      int32    `json:"timeoutConnection,omitempty"`
	TimeoutMember          int32    `json:"timeoutMember,omitempty"`
	// UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}
