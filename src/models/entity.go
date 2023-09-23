package models

import "encoding/json"

type AtlasClassification struct {
	Attributes                       map[string]interface{} `json:"attributes"`
	EntityGuid                       string                 `json:"entityGuid"`
	EntityStatus                     Status                 `json:"entityStatus"`
	LastModifiedTS                   string                 `json:"lastModifiedTS"`
	RemovePropagationsOnEntityDelete bool                   `json:"removePropagationsOnEntityDelete"`
	Source                           string                 `json:"source"`
	TypeName                         string                 `json:"typeName"`
	ValidityPeriods                  []TimeBoundary         `json:"validityPeriods"`
}

type AtlasClassifications struct {
	List       []AtlasClassification `json:"list"`
	PageSize   int                   `json:"pageSize"`
	SortBy     string                `json:"sortBy"`
	SortType   string                `json:"sortType"`
	StartIndex int                   `json:"startIndex"`
	TotalCount int                   `json:"totalCount"`
}

type TimeBoundary struct {
	EndTime   string `json:"endTime"`
	StartTime string `json:"startTime"`
	TimeZone  string `json:"timeZone"`
}

type AtlasEntityHeader struct {
	Attributes          map[string]interface{}      `json:"attributes"`
	ClassificationNames []string                    `json:"classificationNames"`
	Classifications     []AtlasClassification       `json:"classifications"`
	DisplayText         string                      `json:"displayText"`
	Guid                string                      `json:"guid"`
	IsIncomplete        bool                        `json:"isIncomplete"`
	Labels              []string                    `json:"labels"`
	LastModifiedTS      string                      `json:"lastModifiedTS"`
	MeaningNames        []string                    `json:"meaningNames"`
	Meanings            []AtlasTermAssignmentHeader `json:"meanings"`
	Status              Status                      `json:"status"`
	TypeName            string                      `json:"typeName"`
}

type AtlasEntityWithExtInfo struct {
	Entity           Entity                 `json:"entity"`
	ReferredEntities map[string]interface{} `json:"referredEntities"`
}

type EntityMutationResponse struct {
	GuidAssignments       map[string]string      `json:"guidAssignments"`
	MutatedEntities       map[string]interface{} `json:"mutatedEntities"`
	PartialUpdateEntities []AtlasEntityHeader    `json:"partialUpdateEntities"`
}

type AtlasTermAssignmentHeader struct {
	Confidence   int                   `json:"confidence"`
	CreatedBy    string                `json:"createdBy"`
	Description  string                `json:"description"`
	DisplayText  string                `json:"displayText"`
	Expression   string                `json:"expression"`
	RelationGuid string                `json:"relationGuid"`
	Source       string                `json:"source"`
	Status       AtlasAssignmentStatus `json:"status"`
	Steward      string                `json:"steward"`
	TermGuid     string                `json:"termGuid"`
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
	FailedImportInfoList  []ImportInfo `json:"failedImportInfoList"`
	SuccessImportInfoList []ImportInfo `json:"successImportInfo"`
}

type ImportInfo struct {
	ChildObjectName  string `json:"childObjectName"`
	ImportStatus     string `json:"importStatus"`
	ParentObjectName string `json:"parentObjectName"`
	Remarks          string `json:"remarks"`
}

type EntityCreateOrUpdateRequest struct {
	Entity           Entity                 `json:"entity"`
	ReferredEntities map[string]interface{} `json:"referredEntities"`
}

type EntityCreateOrUpdateBulkRequest struct {
	Entity           []Entity               `json:"entity"`
	ReferredEntities map[string]interface{} `json:"referredEntities"`
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
