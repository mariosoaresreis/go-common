package datamodels

type AggregateType int

const (
	AggregateTypeDepartment        AggregateType = iota // Department
	AggregateTypeEmployee                               // Employee
	AggregateTypeJob                                    // Job
	AggregateTypeJobItem                                // JobItem
	AggregateTypeJobGQLView                             // JobGQLView
	AggregateTypeJobsSettings                           // JobSettings
	AggregateTypeLocation                               // Location
	AggregateTypeLocationType                           // LocationType
	AggregateTypeOrganization                           // Organization
	AggregateTypeProperty                               // Property
	AggregateTypeRole                                   // Role
	AggregateTypeRoomType                               // RoomType
	AggregateTypeUserInfo                               // UserInfo
	AggregateTypeTag                                    // Tag
	AggregateTypeChecklistTemplate                      // ChecklistTemplate
	AggregateTypeVirtualLocation                        // VirtualLocation
	AggregateTypeChecklistTask                          // ChecklistTask
	AggregateTypeChecklist                              // Checklist
	AggregateTypeProject                                // Project
	AggregateTypeRepeatingJob                           // RepeatingJob
	AggregateTypeAsset                                  // Asset
	AggregateTypeAssetType                              // AssetType
	AggregateTypeEmploymentType                         // EmploymentType
	AggregateTypeHousekeepingType                       // HousekeepingType
	AggregateTypeTimeline                               // Timeline
)
