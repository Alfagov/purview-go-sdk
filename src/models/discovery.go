package models

import "encoding/json"

type AutoCompleteRequest struct {
	Filter   map[string]interface{} `json:"filter,omitempty"`
	KeyWords string                 `json:"keywords,omitempty"`
	Limit    int                    `json:"limit,omitempty"`
}

type AutoCompleteResult struct {
	Value AutoCompleteResultValue `json:"value,omitempty"`
}

type AutoCompleteResultValue struct {
	QueryPlusText string `json:"queryPlusText,omitempty"`
	Text          string `json:"text,omitempty"`
}

type BrowseRequest struct {
	EntityType string `json:"entityType,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	Path       string `json:"path,omitempty"`
}

type BrowseResult struct {
	SearchCount int                 `json:"@search.count,omitempty"`
	Value       []BrowseResultValue `json:"value,omitempty"`
}

type BrowseResultValue struct {
	EntityType    string              `json:"entityType,omitempty"`
	Id            string              `json:"id,omitempty"`
	IsLeaf        bool                `json:"isLeaf,omitempty"`
	Name          string              `json:"name,omitempty"`
	Owner         []BrowseResultOwner `json:"owner,omitempty"`
	Path          string              `json:"path,omitempty"`
	QualifiedName string              `json:"qualifiedName,omitempty"`
}

type BrowseResultOwner struct {
	ContactType string `json:"contactType,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Id          string `json:"id,omitempty"`
	Mail        string `json:"mail,omitempty"`
}

type QueryRequest struct {
	Facets          []SearchFacetItem      `json:"facets,omitempty"`
	Filter          map[string]interface{} `json:"filter,omitempty"`
	KeyWords        string                 `json:"keywords,omitempty"`
	Limit           int                    `json:"limit,omitempty"`
	Offset          int                    `json:"offset,omitempty"`
	OrderBy         []map[string]string    `json:"orderBy,omitempty"`
	TaxonomySetting TaxonomySetting        `json:"taxonomySetting,omitempty"`
}

type SearchFacetItem struct {
	Count int         `json:"count,omitempty"`
	Facet string      `json:"facet,omitempty"`
	Sort  interface{} `json:"sort,omitempty"`
}

type TaxonomySetting struct {
	AssetTypes []string        `json:"assetTypes,omitempty"`
	Facet      SearchFacetItem `json:"facet,omitempty"`
}

type SearchResult struct {
	SearchCount  int                    `json:"@search.count,omitempty"`
	SearchFacets SearchFacetResultValue `json:"@search.facets,omitempty"`
	Value        []SearchResultValue    `json:"value,omitempty"`
}

type SearchFacetResultValue struct {
	AssetType      []SearchFacetItemValue `json:"assetType,omitempty"`
	Classification []SearchFacetItemValue `json:"classification,omitempty"`
	ContactId      []SearchFacetItemValue `json:"contactId,omitempty"`
	ContactType    []SearchFacetItemValue `json:"contactType,omitempty"`
	EntityType     []SearchFacetItemValue `json:"entityType,omitempty"`
	GlossaryType   []SearchFacetItemValue `json:"glossaryType,omitempty"`
	Label          []SearchFacetItemValue `json:"label,omitempty"`
	Term           []SearchFacetItemValue `json:"term,omitempty"`
	TermStatus     []SearchFacetItemValue `json:"termStatus,omitempty"`
	TermTemplate   []SearchFacetItemValue `json:"termTemplate,omitempty"`
}

type SearchFacetItemValue struct {
	Count int    `json:"count,omitempty"`
	Value string `json:"value,omitempty"`
}

type SearchResultValue struct {
	Highlights      []SearchHighlights         `json:"@search.highlights,omitempty"`
	Score           int                        `json:"@search.score,omitempty"`
	AssetType       []string                   `json:"assetType,omitempty"`
	Classification  []string                   `json:"classification,omitempty"`
	Contact         []ContactSearchResultValue `json:"contact,omitempty"`
	CreateTime      int                        `json:"createTime,omitempty"`
	Description     string                     `json:"description,omitempty"`
	Endorsement     string                     `json:"endorsement,omitempty"`
	EntityType      string                     `json:"entityType,omitempty"`
	Glossary        string                     `json:"glossary,omitempty"`
	GlossaryType    string                     `json:"glossaryType,omitempty"`
	Id              string                     `json:"id,omitempty"`
	Label           []string                   `json:"label,omitempty"`
	LongDescription string                     `json:"longDescription,omitempty"`
	Name            string                     `json:"name,omitempty"`
	ObjectType      string                     `json:"objectType,omitempty"`
	Owner           string                     `json:"owner,omitempty"`
	QualifiedName   string                     `json:"qualifiedName,omitempty"`
	Term            []TermSearchResultValue    `json:"term,omitempty"`
	TermStatus      string                     `json:"termStatus,omitempty"`
	TermTemplate    []string                   `json:"termTemplate,omitempty"`
	UpdateTime      int                        `json:"updateTime,omitempty"`
}

type SearchHighlights struct {
	Description   []string `json:"description,omitempty"`
	EntityType    []string `json:"entityType,omitempty"`
	Id            []string `json:"id,omitempty"`
	Name          []string `json:"name,omitempty"`
	QualifiedName []string `json:"qualifiedName,omitempty"`
}

type ContactSearchResultValue struct {
	ContactType string `json:"contactType,omitempty"`
	Id          string `json:"id,omitempty"`
	Info        string `json:"info,omitempty"`
}

type TermSearchResultValue struct {
	GlossaryName string `json:"glossaryName,omitempty"`
	Guid         string `json:"guid,omitempty"`
	Name         string `json:"name,omitempty"`
}

type SuggestRequest struct {
	Filter   map[string]interface{} `json:"filter,omitempty"`
	KeyWords string                 `json:"keywords,omitempty"`
	Limit    int                    `json:"limit,omitempty"`
}

type SuggestResult struct {
	Value []SuggestResultValue `json:"value"`
}

type SuggestResultValue struct {
	Score           int                        `json:"@search.score,omitempty"`
	Text            string                     `json:"@search.text,omitempty"`
	AssetType       []string                   `json:"assetType,omitempty"`
	Classification  []string                   `json:"classification,omitempty"`
	Contact         []ContactSearchResultValue `json:"contact,omitempty"`
	CreateTime      int                        `json:"createTime,omitempty"`
	Description     string                     `json:"description,omitempty"`
	Endorsement     string                     `json:"endorsement,omitempty"`
	EntityType      string                     `json:"entityType,omitempty"`
	Glossary        string                     `json:"glossary,omitempty"`
	GlossaryType    string                     `json:"glossaryType,omitempty"`
	Id              string                     `json:"id,omitempty"`
	Label           []string                   `json:"label,omitempty"`
	LongDescription string                     `json:"longDescription,omitempty"`
	Name            string                     `json:"name,omitempty"`
	ObjectType      string                     `json:"objectType,omitempty"`
	Owner           string                     `json:"owner,omitempty"`
	QualifiedName   string                     `json:"qualifiedName,omitempty"`
	Term            []TermSearchResultValue    `json:"term,omitempty"`
	TermStatus      string                     `json:"termStatus,omitempty"`
	TermTemplate    []string                   `json:"termTemplate,omitempty"`
	UpdateTime      int                        `json:"updateTime,omitempty"`
}

func (a AutoCompleteResult) UnmarshalJson(data []byte) (
	*AutoCompleteResult,
	error,
) {
	err := json.Unmarshal(data, &a)
	return &a, err
}

func (a BrowseResult) UnmarshalJson(data []byte) (
	*BrowseResult,
	error,
) {
	err := json.Unmarshal(data, &a)
	return &a, err
}

func (a SearchResult) UnmarshalJson(data []byte) (
	*SearchResult,
	error,
) {
	err := json.Unmarshal(data, &a)
	return &a, err
}

func (a SuggestResult) UnmarshalJson(data []byte) (
	*SuggestResult,
	error,
) {
	err := json.Unmarshal(data, &a)
	return &a, err
}
