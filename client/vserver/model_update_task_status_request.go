/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type UpdateTaskStatusRequest struct {
	Data   *VpcInfo `json:"data,omitempty"`
	Status string   `json:"status,omitempty"`
	TaskID string   `json:"taskID,omitempty"`
}