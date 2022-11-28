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

type ServerGroupDetail struct {
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	Description   string    `json:"description,omitempty"`
	Name          string    `json:"name,omitempty"`
	PolicyId      string    `json:"policyId,omitempty"`
	PolicyName    string    `json:"policyName,omitempty"`
	ServerGroupId int32     `json:"serverGroupId,omitempty"`
	Servers       []Server  `json:"servers,omitempty"`
	Uuid          string    `json:"uuid,omitempty"`
}