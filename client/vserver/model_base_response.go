/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type BaseResponse struct {
	ErrorCode int32       `json:"errorCode,omitempty"`
	ErrorMsg  string      `json:"errorMsg,omitempty"`
	Extra     interface{} `json:"extra,omitempty"`
	Success   bool        `json:"success,omitempty"`
}
