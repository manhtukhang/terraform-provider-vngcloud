/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type UpdateQuota struct {
	ProjectId int64 `json:"projectId,omitempty"`
	Quota map[string]interface{} `json:"quota,omitempty"`
}