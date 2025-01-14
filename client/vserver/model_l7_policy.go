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

type L7Policy struct {
	Action             string    `json:"action,omitempty"`
	AdminStateUp       bool      `json:"adminStateUp,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	Description        string    `json:"description,omitempty"`
	KeepQueryString    bool      `json:"keepQueryString,omitempty"`
	L7Rule             *L7Rule   `json:"l7Rule,omitempty"`
	ListenerId         string    `json:"listenerId,omitempty"`
	Name               string    `json:"name,omitempty"`
	OperatingStatus    string    `json:"operatingStatus,omitempty"`
	Position           int64     `json:"position,omitempty"`
	ProjectId          string    `json:"projectId,omitempty"`
	ProvisioningStatus string    `json:"provisioningStatus,omitempty"`
	RedirectHttpCode   int32     `json:"redirectHttpCode,omitempty"`
	RedirectPoolId     string    `json:"redirectPoolId,omitempty"`
	RedirectUrl        string    `json:"redirectUrl,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	Uuid               string    `json:"uuid,omitempty"`
}
