package models

import (
	"fmt"
	"net/url"
	"strconv"
)

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

func (q *QueryAttribute) EncodeIdx(idx int) string {
	return fmt.Sprintf(
		"attr_%s:%s=%s",
		strconv.Itoa(idx),
		q.Key,
		q.Value,
	)
}

func (q *QueryAttributes) Encode() string {
	values := url.Values{}
	for _, attr := range q.Attributes {
		values.Add("attr:"+attr.Key, attr.Value)
	}
	return values.Encode()
}

func (q *QueryAttributes) ToUrlValues() url.Values {
	values := url.Values{}
	for _, attr := range q.Attributes {
		values.Add(attr.Key, attr.Value)
	}
	return values
}

func (q *QueryAttributes) EncodeASC() string {
	values := url.Values{}
	for i, attr := range q.Attributes {
		values.Add(
			fmt.Sprintf("attr_%s:%s", strconv.Itoa(i), attr.Key),
			attr.Value,
		)
	}
	return values.Encode()
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

func (qp *QueryParams) Encode(allowedParams map[string]bool) string {
	if allowedParams == nil || qp.Params == nil {
		return ""
	}

	values := url.Values{}
	for key, value := range qp.Params {
		if allowedParams[key] {
			values.Add(key, value)
		}
	}

	return values.Encode()
}

type businessAttributeUpdateBehavior string

func BusinessAttributeUpdateBehavior(value string) businessAttributeUpdateBehavior {
	if value == "merge" || value == "replace" || value == "ignore" {
		return businessAttributeUpdateBehavior(value)
	}

	return ""
}
