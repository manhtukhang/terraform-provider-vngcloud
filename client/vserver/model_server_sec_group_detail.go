/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type ServerSecGroupDetail struct {
	Inbounds  []SecGroupRuleDetail `json:"inbounds,omitempty"`
	Outbounds []SecGroupRuleDetail `json:"outbounds,omitempty"`
}
