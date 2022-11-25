/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

import (
	"time"
)

type InterfaceNetworkInterfaceEntity struct {
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	DeletedAt     time.Time `json:"deletedAt,omitempty"`
	FixedIp       string    `json:"fixedIp,omitempty"`
	FloatingIp    string    `json:"floatingIp,omitempty"`
	Id            int32     `json:"id,omitempty"`
	InterfaceId   int32     `json:"interfaceId,omitempty"`
	InterfaceType string    `json:"interfaceType,omitempty"`
	Mac           string    `json:"mac,omitempty"`
	NetworkUuid   string    `json:"networkUuid,omitempty"`
	PortUuid      string    `json:"portUuid,omitempty"`
	Product       string    `json:"product,omitempty"`
	ServerUuid    string    `json:"serverUuid,omitempty"`
	Status        string    `json:"status,omitempty"`
	SubnetUuid    string    `json:"subnetUuid,omitempty"`
	Type_         string    `json:"type,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	Uuid          string    `json:"uuid,omitempty"`
}
