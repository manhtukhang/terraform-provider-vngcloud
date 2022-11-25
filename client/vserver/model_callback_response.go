/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type CallbackResponse struct {
	Code     int32  `json:"code,omitempty"`
	ErrorMsg string `json:"errorMsg,omitempty"`
	Success  bool   `json:"success,omitempty"`
	Total    int32  `json:"total,omitempty"`
}
