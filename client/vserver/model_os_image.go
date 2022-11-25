/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type OsImage struct {
	FlavorZoneIds []string      `json:"flavorZoneIds,omitempty"`
	Id            string        `json:"id,omitempty"`
	ImageType     string        `json:"imageType,omitempty"`
	ImageVersion  string        `json:"imageVersion,omitempty"`
	Licence       bool          `json:"licence,omitempty"`
	PackageLimit  *PackageLimit `json:"packageLimit,omitempty"`
}
