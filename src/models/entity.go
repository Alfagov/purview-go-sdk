package models

import "encoding/json"

type AtlasClassification struct {
	Attributes                       map[string]interface{} `json:"attributes,omitempty"`
	EntityGuid                       string                 `json:"entityGuid,omitempty"`
	EntityStatus                     Status                 `json:"entityStatus,omitempty"`
	LastModifiedTS                   string                 `json:"lastModifiedTS,omitempty"`
	RemovePropagationsOnEntityDelete bool                   `json:"removePropagationsOnEntityDelete,omitempty"`
	Source                           string                 `json:"source,omitempty"`
	TypeName                         string                 `json:"typeName,omitempty"`
	ValidityPeriods                  []TimeBoundary         `json:"validityPeriods,omitempty"`
}

type AtlasClassifications struct {
	List       []AtlasClassification `json:"list,omitempty"`
	PageSize   int                   `json:"pageSize,omitempty"`
	SortBy     string                `json:"sortBy,omitempty"`
	SortType   string                `json:"sortType,omitempty"`
	StartIndex int                   `json:"startIndex,omitempty"`
	TotalCount int                   `json:"totalCount,omitempty"`
}

type TimeBoundary struct {
	EndTime   string `json:"endTime,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	TimeZone  string `json:"timeZone,omitempty"`
}

type AtlasEntityHeader struct {
	Attributes          map[string]interface{}      `json:"attributes,omitempty"`
	ClassificationNames []string                    `json:"classificationNames,omitempty"`
	Classifications     []AtlasClassification       `json:"classifications,omitempty"`
	DisplayText         string                      `json:"displayText,omitempty"`
	Guid                string                      `json:"guid,omitempty"`
	IsIncomplete        bool                        `json:"isIncomplete,omitempty"`
	Labels              []string                    `json:"labels,omitempty"`
	LastModifiedTS      string                      `json:"lastModifiedTS,omitempty"`
	MeaningNames        []string                    `json:"meaningNames,omitempty"`
	Meanings            []AtlasTermAssignmentHeader `json:"meanings,omitempty"`
	Status              Status                      `json:"status,omitempty"`
	TypeName            string                      `json:"typeName,omitempty"`
}

type AtlasEntityWithExtInfo struct {
	Entity           Entity                 `json:"entity,omitempty"`
	ReferredEntities map[string]interface{} `json:"referredEntities,omitempty"`
}

type EntityMutationResponse struct {
	GuidAssignments       map[string]string      `json:"guidAssignments,omitempty"`
	MutatedEntities       map[string]interface{} `json:"mutatedEntities,omitempty"`
	PartialUpdateEntities []AtlasEntityHeader    `json:"partialUpdateEntities,omitempty"`
}

type AtlasTermAssignmentHeader struct {
	Confidence   int                   `json:"confidence,omitempty"`
	CreatedBy    string                `json:"createdBy,omitempty"`
	Description  string                `json:"description,omitempty"`
	DisplayText  string                `json:"displayText,omitempty"`
	Expression   string                `json:"expression,omitempty"`
	RelationGuid string                `json:"relationGuid,omitempty"`
	Source       string                `json:"source,omitempty"`
	Status       AtlasAssignmentStatus `json:"status,omitempty"`
	Steward      string                `json:"steward,omitempty"`
	TermGuid     string                `json:"termGuid,omitempty"`
}

type AtlasAssignmentStatus struct {
	string
}

type Status struct {
	string
}

type Entity struct {
	Attributes             map[string]interface{}      `json:"attributes,omitempty"`
	BusinessAttributes     map[string]interface{}      `json:"businessAttributes,omitempty"`
	Classifications        []AtlasClassification       `json:"classifications,omitempty"`
	CollectionId           string                      `json:"collectionId,omitempty"`
	Contacts               map[string]interface{}      `json:"contacts,omitempty"`
	CreateTime             int                         `json:"createTime ,omitempty"`
	CreatedBy              string                      `json:"createdBy,omitempty"`
	CustomAttributes       map[string]interface{}      `json:"customAttributes,omitempty"`
	Guid                   string                      `json:"guid,omitempty"`
	HomeId                 string                      `json:"homeId ,omitempty"`
	IsIncomplete           bool                        `json:"isIncomplete,omitempty"`
	Labels                 []string                    `json:"labels,omitempty"`
	LastModifiedTS         string                      `json:"lastModifiedTS,omitempty"`
	Meanings               []AtlasTermAssignmentHeader `json:"meanings,omitempty"`
	ProvenanceType         int                         `json:"provenanceType,omitempty"`
	Proxy                  bool                        `json:"proxy,omitempty"`
	RelationshipAttributes map[string]interface{}      `json:"relationshipAttributes optional"`
	Source                 string                      `json:"source,omitempty"`
	Status                 Status                      `json:"status ,omitempty"`
	TypeName               string                      `json:"typeName ,omitempty"`
	UpdateTime             int                         `json:"updateTime,omitempty"`
	UpdatedBy              string                      `json:"updatedBy,omitempty"`
	Version                int                         `json:"version,omitempty"`
}

type BulkImportResponse struct {
	FailedImportInfoList  []ImportInfo `json:"failedImportInfoList,omitempty"`
	SuccessImportInfoList []ImportInfo `json:"successImportInfo,omitempty"`
}

type ImportInfo struct {
	ChildObjectName  string `json:"childObjectName,omitempty"`
	ImportStatus     string `json:"importStatus,omitempty"`
	ParentObjectName string `json:"parentObjectName,omitempty"`
	Remarks          string `json:"remarks,omitempty"`
}

type EntityCreateOrUpdateRequest struct {
	Entity           Entity                 `json:"entity,omitempty"`
	ReferredEntities map[string]interface{} `json:"referredEntities,omitempty"`
}

type EntityCreateOrUpdateBulkRequest struct {
	Entity           []Entity               `json:"entity,omitempty"`
	ReferredEntities map[string]interface{} `json:"referredEntities,omitempty"`
}

func UnmarshalEntity(data []byte) (*Entity, error) {
	var r Entity
	err := json.Unmarshal(data, &r)
	return &r, err
}

func MarshalEntity(entity *Entity) ([]byte, error) {
	return json.Marshal(entity)
}

func (s AtlasEntityWithExtInfo) UnmarshalJson(data []byte) (
	*AtlasEntityWithExtInfo,
	error,
) {
	err := json.Unmarshal(data, &s)
	return &s, err
}

func (s EntityMutationResponse) UnmarshalJson(data []byte) (
	*EntityMutationResponse,
	error,
) {
	err := json.Unmarshal(data, &s)
	return &s, err
}

func (s AtlasClassification) UnmarshalJson(data []byte) (
	*AtlasClassification,
	error,
) {
	err := json.Unmarshal(data, &s)
	return &s, err
}

func (s AtlasClassifications) UnmarshalJson(data []byte) (
	*AtlasClassifications,
	error,
) {
	err := json.Unmarshal(data, &s)
	return &s, err
}

func (s AtlasEntityHeader) UnmarshalJson(data []byte) (
	*AtlasEntityHeader,
	error,
) {
	err := json.Unmarshal(data, &s)
	return &s, err
}

func (s BulkImportResponse) UnmarshalJson(data []byte) (
	*BulkImportResponse,
	error,
) {
	err := json.Unmarshal(data, &s)
	return &s, err
}

func (s *EntityCreateOrUpdateRequest) MarshalJson() ([]byte, error) {
	return json.Marshal(s)
}

func (s *EntityCreateOrUpdateBulkRequest) MarshalJson() ([]byte, error) {
	return json.Marshal(s)
}
