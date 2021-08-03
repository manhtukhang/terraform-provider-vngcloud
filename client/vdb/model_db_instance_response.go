/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vdb

type DbInstanceResponse struct {
	Code         int32           `json:"code,omitempty"`
	Data         *DbInstanceInfo `json:"data,omitempty"`
	DbInstanceId string          `json:"dbInstanceId,omitempty"`
	ErrorMsg     string          `json:"errorMsg,omitempty"`
	ProjectId    string          `json:"projectId,omitempty"`
	Success      bool            `json:"success,omitempty"`
	Total        int32           `json:"total,omitempty"`
}
