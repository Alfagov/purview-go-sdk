package models

import "encoding/json"

type AutoCompleteRequest struct {
	Filter   map[string]interface{} `json:"filter"`
	KeyWords string                 `json:"keywords"`
	Limit    int                    `json:"limit"`
}

type AutoCompleteResult struct {
	Value AutoCompleteResultValue `json:"value"`
}

type AutoCompleteResultValue struct {
	QueryPlusText string `json:"queryPlusText"`
	Text          string `json:"text"`
}

type BrowseRequest struct {
	EntityType string `json:"entityType"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Path       string `json:"path"`
}

type BrowseResult struct {
	SearchCount int                 `json:"@search.count"`
	Value       []BrowseResultValue `json:"value"`
}

type BrowseResultValue struct {
	EntityType    string              `json:"entityType"`
	Id            string              `json:"id"`
	IsLeaf        bool                `json:"isLeaf"`
	Name          string              `json:"name"`
	Owner         []BrowseResultOwner `json:"owner"`
	Path          string              `json:"path"`
	QualifiedName string              `json:"qualifiedName"`
}

type BrowseResultOwner struct {
	ContactType string `json:"contactType"`
	DisplayName string `json:"displayName"`
	Id          string `json:"id"`
	Mail        string `json:"mail"`
}

type QueryRequest struct {
	Facets          []SearchFacetItem      `json:"facets"`
	Filter          map[string]interface{} `json:"filter"`
	KeyWords        string                 `json:"keywords"`
	Limit           int                    `json:"limit"`
	Offset          int                    `json:"offset"`
	OrderBy         []map[string]string    `json:"orderBy"`
	TaxonomySetting TaxonomySetting        `json:"taxonomySetting"`
}

type SearchFacetItem struct {
	Count int         `json:"count"`
	Facet string      `json:"facet"`
	Sort  interface{} `json:"sort"`
}

type TaxonomySetting struct {
	AssetTypes []string        `json:"assetTypes"`
	Facet      SearchFacetItem `json:"facet"`
}

type SearchResult struct {
	SearchCount  int                    `json:"@search.count"`
	SearchFacets SearchFacetResultValue `json:"@search.facets"`
	Value        []SearchResultValue    `json:"value"`
}

type SearchFacetResultValue struct {
	AssetType      []SearchFacetItemValue `json:"assetType"`
	Classification []SearchFacetItemValue `json:"classification"`
	ContactId      []SearchFacetItemValue `json:"contactId"`
	ContactType    []SearchFacetItemValue `json:"contactType"`
	EntityType     []SearchFacetItemValue `json:"entityType"`
	GlossaryType   []SearchFacetItemValue `json:"glossaryType"`
	Label          []SearchFacetItemValue `json:"label"`
	Term           []SearchFacetItemValue `json:"term"`
	TermStatus     []SearchFacetItemValue `json:"termStatus"`
	TermTemplate   []SearchFacetItemValue `json:"termTemplate"`
}

type SearchFacetItemValue struct {
	Count int    `json:"count"`
	Value string `json:"value"`
}

type SearchResultValue struct {
	Highlights      []SearchHighlights         `json:"@search.highlights"`
	Score           int                        `json:"@search.score"`
	AssetType       []string                   `json:"assetType"`
	Classification  []string                   `json:"classification"`
	Contact         []ContactSearchResultValue `json:"contact"`
	CreateTime      int                        `json:"createTime"`
	Description     string                     `json:"description"`
	Endorsement     string                     `json:"endorsement"`
	EntityType      string                     `json:"entityType"`
	Glossary        string                     `json:"glossary"`
	GlossaryType    string                     `json:"glossaryType"`
	Id              string                     `json:"id"`
	Label           []string                   `json:"label"`
	LongDescription string                     `json:"longDescription"`
	Name            string                     `json:"name"`
	ObjectType      string                     `json:"objectType"`
	Owner           string                     `json:"owner"`
	QualifiedName   string                     `json:"qualifiedName"`
	Term            []TermSearchResultValue    `json:"term"`
	TermStatus      string                     `json:"termStatus"`
	TermTemplate    []string                   `json:"termTemplate"`
	UpdateTime      int                        `json:"updateTime"`
}

type SearchHighlights struct {
	Description   []string `json:"description"`
	EntityType    []string `json:"entityType"`
	Id            []string `json:"id"`
	Name          []string `json:"name"`
	QualifiedName []string `json:"qualifiedName"`
}

type ContactSearchResultValue struct {
	ContactType string `json:"contactType"`
	Id          string `json:"id"`
	Info        string `json:"info"`
}

type TermSearchResultValue struct {
	GlossaryName string `json:"glossaryName"`
	Guid         string `json:"guid"`
	Name         string `json:"name"`
}

type SuggestRequest struct {
	Filter   map[string]interface{} `json:"filter"`
	KeyWords string                 `json:"keywords"`
	Limit    int                    `json:"limit"`
}

type SuggestResult struct {
	Value []SuggestResultValue `json:"value"`
}

type SuggestResultValue struct {
	Score           int                        `json:"@search.score"`
	Text            string                     `json:"@search.text"`
	AssetType       []string                   `json:"assetType"`
	Classification  []string                   `json:"classification"`
	Contact         []ContactSearchResultValue `json:"contact"`
	CreateTime      int                        `json:"createTime"`
	Description     string                     `json:"description"`
	Endorsement     string                     `json:"endorsement"`
	EntityType      string                     `json:"entityType"`
	Glossary        string                     `json:"glossary"`
	GlossaryType    string                     `json:"glossaryType"`
	Id              string                     `json:"id"`
	Label           []string                   `json:"label"`
	LongDescription string                     `json:"longDescription"`
	Name            string                     `json:"name"`
	ObjectType      string                     `json:"objectType"`
	Owner           string                     `json:"owner"`
	QualifiedName   string                     `json:"qualifiedName"`
	Term            []TermSearchResultValue    `json:"term"`
	TermStatus      string                     `json:"termStatus"`
	TermTemplate    []string                   `json:"termTemplate"`
	UpdateTime      int                        `json:"updateTime"`
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
