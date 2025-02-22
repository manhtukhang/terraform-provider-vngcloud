/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type PagingL7Policy struct {
	ListData  []L7Policy `json:"listData,omitempty"`
	Page      int32      `json:"page,omitempty"`
	PageSize  int32      `json:"pageSize,omitempty"`
	TotalItem int64      `json:"totalItem,omitempty"`
	TotalPage int32      `json:"totalPage,omitempty"`
}
