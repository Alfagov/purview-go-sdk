package paths

import (
	"github.com/Alfagov/pwCatalog/src/models"
	"net/url"
)

var (
	entityBasePath                     = &url.URL{Path: "/catalog/api/atlas/v2/entity"}
	uniqueAttributeBasePath            = &url.URL{Path: "uniqueAttribute/type"}
	guidEntityBasePath                 = entityBasePath.JoinPath("/guid")
	BulkEntitiesPath                   = entityBasePath.JoinPath("/bulk")
	BusinessMetadataBasePath           = entityBasePath.JoinPath("/businessmetadata")
	SampleBusinessMetadataTemplatePath = BusinessMetadataBasePath.JoinPath("/import/template")
	ImportBusinessMetadataPath         = BusinessMetadataBasePath.JoinPath("/import")
	SetClassificationsPath             = BulkEntitiesPath.JoinPath("/setClassifications")
	AddClassificationPath              = BulkEntitiesPath.JoinPath("/classification")
)

func ClassificationsPath(guid string) *url.URL {
	return guidEntityBasePath.JoinPath(guid, "classifications")
}

func ClassesByUniqueAttributePath(
	attributes *models.QueryAttributes,
) *url.URL {
	path := uniqueAttributeBasePath.JoinPath(
		attributes.TypeName,
		"classifications",
	)
	path.RawQuery = attributes.Encode()

	return path
}

func LabelPath(guid string) *url.URL {
	return guidEntityBasePath.JoinPath(guid, "labels")
}

func LabelsByUniqueAttributePath(attributes *models.QueryAttributes) *url.URL {
	path := uniqueAttributeBasePath.JoinPath(attributes.TypeName, "labels")
	path.RawQuery = attributes.Encode()

	return path
}

func BusinessMetadataPath(
	guid string, options *models.QueryParams,
) *url.URL {
	path := guidEntityBasePath.JoinPath(guid, "businessmetadata")
	queryOptions := options.Encode(map[string]bool{"isOverwrite": true})
	path.RawQuery = queryOptions

	return path
}

func BusinessMetadataAttributePath(
	guid string,
	bmName string,
) *url.URL {
	return guidEntityBasePath.JoinPath(guid, "businessmetadata", bmName)
}

func CreateOrUpdateEntityPath(options *models.QueryParams) *url.URL {
	queryOptions := options.Encode(map[string]bool{"businessAttributeUpdateBehavior": true})
	path := &url.URL{Path: entityBasePath.Path}
	path.RawQuery = queryOptions

	return path
}

func CreateOrUpdateEntitiesPath(options *models.QueryParams) *url.URL {
	path := &url.URL{Path: BulkEntitiesPath.Path}
	path.RawQuery = options.Encode(map[string]bool{"businessAttributeUpdateBehavior": true})

	return path
}

func EntityByGuidPath(guid string, options *models.QueryParams) *url.URL {
	path := guidEntityBasePath.JoinPath(guid)
	path.RawQuery = options.Encode(map[string]bool{"ignoreRelationships": true, "minExtInfo": true})

	return path
}

func DeleteEntitiesByGuids(guids []string) *url.URL {
	path := &url.URL{Path: BulkEntitiesPath.Path}
	path.RawQuery = createGuidsParamValues(guids)

	return path
}

func GetEntitiesByUniqueAttributesPath(
	attributes *models.QueryAttributes, options *models.QueryParams,
) *url.URL {
	path := &url.URL{Path: BulkEntitiesPath.Path}
	path.RawQuery = attributes.EncodeASC()
	path.RawQuery += "&" + options.Encode(map[string]bool{"ignoreRelationships": true, "minExtInfo": true})

	return path
}

func EntityByUniqueAttributePath(
	attributes *models.QueryAttributes,
	options *models.QueryParams,
) *url.URL {
	path := uniqueAttributeBasePath.JoinPath(attributes.TypeName)
	path.RawQuery = attributes.Encode()
	path.RawQuery += "&" + options.Encode(map[string]bool{"ignoreRelationships": true, "minExtInfo": true})

	return path
}

func ClassificationPath(guid string, classificationName string) *url.URL {
	return guidEntityBasePath.JoinPath(
		guid,
		"classification",
		classificationName,
	)
}

func DeleteClassificationByUniqueAttributePath(
	attributes *models.QueryAttributes,
	classificationName string,
) *url.URL {
	path := uniqueAttributeBasePath.JoinPath(
		attributes.TypeName,
		"classification",
		classificationName,
	)

	path.RawQuery = attributes.Encode()

	return path
}

func HeaderPath(guid string) *url.URL {
	return guidEntityBasePath.JoinPath(guid, "header")
}

func ListEntitiesByGuidsPath(
	guids []string, options *models.QueryParams,
) *url.URL {
	path := &url.URL{Path: BulkEntitiesPath.Path}
	path.RawQuery = createGuidsParamValues(guids)
	path.RawQuery += "&" + options.Encode(map[string]bool{"excludeRelationshipTypes": true, "ignoreRelationships": true, "minExtInfo": true})

	return path
}

func PartialUpdateEntityAttributeByGuidPath(
	guid string, attributeName string,
) *url.URL {
	path := guidEntityBasePath.JoinPath(guid)
	path.RawQuery = "name=" + attributeName

	return path
}
