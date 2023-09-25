package models

import "fmt"

type QueryAttributes struct {
	TypeName   string `json:"typeName"`
	Attributes []QueryAttribute
}

type QueryAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (q *QueryAttribute) Encode() string {
	return fmt.Sprintf("attr:%s=%s", q.Key, q.Value)
}

func (q *QueryAttributes) Encode() string {
	encodedAttributes := ""
	for _, attr := range q.Attributes {
		encodedAttributes += attr.Encode() + "&"
	}
	return encodedAttributes
}

type QueryParams struct {
	Params map[string]string
}

func (qp *QueryParams) IgnoreRelationships(value string) {
	qp.Params["ignoreRelationships"] = value
}

func (qp *QueryParams) MinExtInfo(value string) {
	qp.Params["minExtInfo"] = value
}

func (qp *QueryParams) BusinessAttributeUpdateBehavior(value businessAttributeUpdateBehavior) {
	qp.Params["businessAttributeUpdateBehavior"] = string(value)
}

func (qp *QueryParams) IsOverwrite(value string) {
	qp.Params["isOverwrite"] = value
}

type businessAttributeUpdateBehavior string

func BusinessAttributeUpdateBehavior(value string) businessAttributeUpdateBehavior {
	if value == "merge" || value == "replace" || value == "ignore" {
		return businessAttributeUpdateBehavior(value)
	}

	return ""
}

/*const (
	BusinessAttributeUpdateBehaviorMerge   = "merge"
	BusinessAttributeUpdateBehaviorReplace = "replace"
	BusinessAttributeUpdateBehaviorIgnore  = "ignore"
)*/
