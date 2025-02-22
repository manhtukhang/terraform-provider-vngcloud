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

type L7Rule struct {
	AdminStateUp       bool      `json:"adminStateUp,omitempty"`
	CompareType        string    `json:"compareType,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	Invert             bool      `json:"invert,omitempty"`
	OperatingStatus    string    `json:"operatingStatus,omitempty"`
	PolicyId           string    `json:"policyId,omitempty"`
	ProjectId          string    `json:"projectId,omitempty"`
	ProvisioningStatus string    `json:"provisioningStatus,omitempty"`
	RuleKey            string    `json:"ruleKey,omitempty"`
	RuleType           string    `json:"ruleType,omitempty"`
	RuleValue          string    `json:"ruleValue,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	Uuid               string    `json:"uuid,omitempty"`
}
