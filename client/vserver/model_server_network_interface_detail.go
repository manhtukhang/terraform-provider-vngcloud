/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ServerNetworkInterfaceDetail struct {
	ExternalInterfaces []InterfaceNetworkInterface `json:"externalInterfaces,omitempty"`
	InternalInterfaces []InterfaceNetworkInterface `json:"internalInterfaces,omitempty"`
}