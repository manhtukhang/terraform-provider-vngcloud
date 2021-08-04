/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vdb

import (
	"time"
)

type CreateBackupStorageRequest struct {
	BackupPackageId   string      `json:"backupPackageId,omitempty"`
	BackupPackageName string      `json:"backupPackageName,omitempty"`
	EndDate           time.Time   `json:"endDate,omitempty"`
	EngineGroup       int32       `json:"engineGroup,omitempty"`
	Extra             interface{} `json:"extra,omitempty"`
	MonthlyCost       float64     `json:"monthlyCost,omitempty"`
	Name              string      `json:"name,omitempty"`
	Period            int32       `json:"period,omitempty"`
	ProjectId         string      `json:"projectId,omitempty"`
	Quota             int32       `json:"quota,omitempty"`
	StartDate         time.Time   `json:"startDate,omitempty"`
}
