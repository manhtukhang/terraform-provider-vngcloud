/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CreateListenerRequest struct {
	// Allowed cidr.
	AllowedCidrs string `json:"allowedCidrs"`
	// List of certificate authority
	CertificateAuthorities []string `json:"certificateAuthorities,omitempty"`
	// Default certificate authority that the listener will use
	DefaultCertificateAuthority string `json:"defaultCertificateAuthority,omitempty"`
	// Id of the pool that this listener will forward to.
	DefaultPoolId string `json:"defaultPoolId,omitempty"`
	// Name of the listener. Only letters (a-z, A-Z, 0-9, '_', '.') are allowed and your input data must be between 6 and 20 characters.
	ListenerName string `json:"listenerName"`
	// Protocol of the listener.
	ListenerProtocol string `json:"listenerProtocol"`
	// Port of the listener.
	ListenerProtocolPort int32 `json:"listenerProtocolPort"`
	// Id of the load balancer.
	LoadBalancerId string `json:"loadBalancerId,omitempty"`
	// Idle timeout of client. The value can be in range from 1 to 3600 seconds
	TimeoutClient int32 `json:"timeoutClient"`
	// Idle timeout of connection. The value can be in range from 1 to 3600 seconds
	TimeoutConnection int32 `json:"timeoutConnection"`
	// Idle timeout of member. The value can be in range from 1 to 3600 seconds
	TimeoutMember int32 `json:"timeoutMember"`
}
