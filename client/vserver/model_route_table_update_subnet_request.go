/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type RouteTableUpdateSubnetRequest struct {
	// Id of project
	ProjectId string `json:"projectId,omitempty"`
	// The updated subnet uuid list of route-table
	Subnets []string `json:"subnets,omitempty"`
	// Id of user
	UserId int32 `json:"userId,omitempty"`
}