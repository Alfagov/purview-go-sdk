package entity

import (
	"github.com/Alfagov/pwCatalog/src/client"
	"github.com/Alfagov/pwCatalog/src/models"
)

type Entity interface {
	AddClassification(
		classification *models.AtlasClassification,
		entityGuids []string,
	) error
	AddClassifications(
		guid string,
		classifications []*models.AtlasClassification,
	) error
	AddClassificationsByUniqueAttribute(
		classifications []*models.AtlasClassification,
		attributes *models.QueryAttributes,
	) error
	AddLabel(guid string, label []string) error
	AddLabelsByUniqueAttribute(
		labels []string, attributes *models.QueryAttributes,
	) error
	AddOrUpdateBusinessMetadata(
		guid string, businessMetadata map[string]interface{},
		options *models.QueryParams,
	) error
	AddOrUpdateBusinessMetadataAttributes(
		businessMetadataAttributes map[string]interface{},
		guid string, bmName string,
	) error
	CreateOrUpdate(
		data *models.EntityCreateOrUpdateRequest,
		options *models.QueryParams,
	) (*models.EntityMutationResponse, error)
	CreateOrUpdateEntities(
		data *models.EntityCreateOrUpdateBulkRequest,
		options *models.QueryParams,
	) (*models.EntityMutationResponse, error)
	GetByGuid(
		guid string,
		options *models.QueryParams,
	) (*models.AtlasEntityWithExtInfo, error)
	GetByUniqueAttributes(
		uniqueAttributes *models.QueryAttributes,
		params *models.QueryParams,
	) (*models.AtlasEntityWithExtInfo, error)
	DeleteBusinessMetadata(
		guid string,
		options *models.QueryParams,
		businessMetadata map[string]interface{},
	) error
	DeleteBusinessMetadataAttributes(
		guid string, bmName string,
		businessMetadata map[string]interface{},
	) error
	DeleteByGuid(guid string) error
	DeleteByGuids(guids []string) error
	DeleteByUniqueAttribute(uniqueAttributes *models.QueryAttributes) error
	DeleteClassification(guid string, classificationName string) error
	DeleteClassificationByUniqueAttribute(
		name string, uniqueAttributes *models.QueryAttributes,
	) error
	DeleteLabels(guid string, labels []string) error
	DeleteLabelsByUniqueAttribute(
		uniqueAttributes *models.QueryAttributes,
		labels []string,
	) error
	GetClassification(guid string, name string) (
		*models.AtlasClassification,
		error,
	)
	GetClassifications(guid string) (*models.AtlasClassifications, error)
	GetEntitiesByUniqueAttributes(
		uniqueAttributes *models.QueryAttributes,
		queryOptions *models.QueryParams,
	) (
		*models.AtlasEntityWithExtInfo, error,
	)
	GetHeader(guid string) (*models.AtlasEntityHeader, error)
	GetSampleBusinessMetadataTemplate() (
		[]byte, error,
	)
	ImportBusinessMetadata() (*models.BulkImportResponse, error)
	ListByGuids(guids []string, options *models.QueryParams) (
		*models.
			AtlasEntityWithExtInfo, error,
	)
	PartialUpdateEntityAttributeByGuid(
		guid string, name string, value interface{},
	) (*models.EntityMutationResponse, error)
	PartialUpdateEntityByUniqueAttributes(
		uniqueAttributes *models.QueryAttributes,
		entityData *models.AtlasEntityWithExtInfo,
	) (*models.EntityMutationResponse, error)
	SetClassifications(
		guid string, classifications map[string]models.AtlasEntityHeader,
	) ([]string, error)
	SetLabels(guid string, labels []string) error
	SetLabelsByUniqueAttribute(
		uniqueAttributes *models.QueryAttributes, labels []string,
	) error
	UpdateClassifications(
		guid string, classifications []*models.AtlasClassification,
	) error
	UpdateClassificationsByUniqueAttribute(
		classifications []*models.AtlasClassification,
		uniqueAttributes *models.QueryAttributes,
	) error
}

type entityImpl struct {
	client.HttpClient
}

func NewEntity(httpClient client.HttpClient) Entity {
	return &entityImpl{httpClient}
}
