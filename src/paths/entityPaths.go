package paths

import (
	"fmt"
	"github.com/Alfagov/purview-go-sdk/src/models"
	"strings"
)

const (
	entityBasePath                     = "/catalog/api/atlas/v2/entity"
	uniqueAttributeBasePath            = "uniqueAttribute/type"
	SampleBusinessMetadataTemplatePath = entityBasePath + "/businessmetadata/import/template"
	ImportBusinessMetadataPath         = entityBasePath + "/businessmetadata/import"
	AddClassificationPath              = entityBasePath + "/bulk/classification"
)

func formatPath(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func addGuidToPath(guid string) string {
	return formatPath("%s/guid/%s", entityBasePath, guid)
}

func addTypeNameToPath(typeName string) string {
	return formatPath(
		"%s/%s/%s", entityBasePath, uniqueAttributeBasePath, typeName,
	)
}

func constructPathWithAttributes(
	path string, attributes *models.QueryAttributes,
) string {
	return formatPath("%s?%s", path, attributes.Encode())
}

func createQueryOptions(
	options *models.QueryParams, allowedOptions map[string]bool,
) string {
	if options == nil {
		return ""
	}

	pathOptions := []string{}
	for key, value := range options.Params {
		if allowedOptions[key] {
			pathOptions = append(pathOptions, fmt.Sprintf("%s=%s", key, value))
		}
	}

	return strings.Join(pathOptions, "&")
}

func createGuidsParam(guids []string) string {
	s := make([]string, len(guids))
	for i, guid := range guids {
		s[i] = "guid=" + guid
	}
	return strings.Join(s, "&")
}

func createPathWithGUIDAndSuffix(guid string, suffix string) string {
	return formatPath("%s/%s", addGuidToPath(guid), suffix)
}

func createPathByUniqueAttrWithSuffix(
	attributes *models.QueryAttributes, suffix string,
) string {
	path := addTypeNameToPath(attributes.TypeName)
	return constructPathWithAttributes(
		formatPath("%s/%s", path, suffix), attributes,
	)
}

func ClassificationsPath(guid string) string {
	return createPathWithGUIDAndSuffix(guid, "classifications")
}

func AddClassificationsByUniqueAttributePath(
	attributes *models.QueryAttributes,
) string {
	return createPathByUniqueAttrWithSuffix(attributes, "classifications")
}

func LabelPath(guid string) string {
	return createPathWithGUIDAndSuffix(guid, "labels")
}

func LabelsByUniqueAttributePath(attributes *models.QueryAttributes) string {
	return createPathByUniqueAttrWithSuffix(attributes, "labels")
}

func BusinessMetadataPath(
	guid string, options *models.QueryParams,
) string {
	path := createPathWithGUIDAndSuffix(guid, "businessmetadata")
	queryOptions := createQueryOptions(
		options, map[string]bool{"isOverwrite": true},
	)
	return formatPath("%s?%s", path, queryOptions)
}

func BusinessMetadataAttributePath(
	guid string,
	bmName string,
) string {
	path := createPathWithGUIDAndSuffix(guid, "businessmetadata")
	path = fmt.Sprintf("%s/%s", path, bmName)
	return path
}

func CreateOrUpdateEntityPath(options *models.QueryParams) string {
	pathOptions := createQueryOptions(
		options, map[string]bool{
			"businessAttributeUpdateBehavior": true,
		},
	)
	return formatPath("%s?%s", entityBasePath, pathOptions)
}

func CreateOrUpdateEntitiesPath(options *models.QueryParams) string {
	path := entityBasePath + "/bulk"
	pathOptions := createQueryOptions(
		options, map[string]bool{
			"businessAttributeUpdateBehavior": true,
		},
	)
	return formatPath("%s?%s", path, pathOptions)
}

func EntityByGuidPath(guid string) string {
	return addGuidToPath(guid)
}

func DeleteEntitiesByGuids(guids []string) string {
	path := entityBasePath + "/bulk"
	pathGuids := createGuidsParam(guids)
	path = fmt.Sprintf("%s?%s", path, pathGuids)
	return path
}

func EntityByUniqueAttributePath(attributes *models.QueryAttributes) string {
	path := addTypeNameToPath(attributes.TypeName)
	path = fmt.Sprintf("%s?%s", path, attributes.Encode())
	return path
}

func ClassificationPath(guid string, classificationName string) string {
	return createPathWithGUIDAndSuffix(
		guid, formatPath("classification/%s", classificationName),
	)
}

func DeleteClassificationByUniqueAttributePath(
	attributes *models.QueryAttributes,
	classificationName string,
) string {
	return createPathByUniqueAttrWithSuffix(
		attributes,
		formatPath("classification/%s", classificationName),
	)
}

func HeaderPath(guid string) string {
	return createPathWithGUIDAndSuffix(guid, "header")
}

func ListEntitiesByGuidsPath(
	guids []string, options *models.QueryParams,
) string {
	path := entityBasePath + "/bulk"
	pathGuids := createGuidsParam(guids)
	pathOptions := createQueryOptions(
		options, map[string]bool{
			"excludeRelationshipTypes": true,
			"ignoreRelationships":      true,
			"minExtInfo":               true,
		},
	)
	path = fmt.Sprintf("%s?%s%s", path, pathOptions, pathGuids)
	return path
}
