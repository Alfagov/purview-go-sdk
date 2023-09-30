package Collection

import (
	"github.com/Alfagov/pwCatalog/src/client"
	"github.com/Alfagov/pwCatalog/src/models"
	"github.com/Alfagov/pwCatalog/src/paths"
	"github.com/Alfagov/pwCatalog/src/response"
	"github.com/Alfagov/pwCatalog/src/utils"
	"net/http"
)

type Collection interface {
	CreateOrUpdate(
		collName string,
		apiVersion string,
		data *models.Entity,
	) (*models.EntityMutationResponse, error)
	CreateOrUpdateBulk(
		collName string,
		apiVersion string,
		data *models.EntityCreateOrUpdateBulkRequest,
	) (*models.EntityMutationResponse, error)
	MoveEntitiesToCollection(
		collName string,
		apiVersion string,
		guids []string,
	) (*models.EntityMutationResponse, error)
}

type collectionImpl struct {
	client.HttpClient
}

func NewCollection(httpClient client.HttpClient) Collection {
	return &collectionImpl{httpClient}
}

func (c *collectionImpl) CreateOrUpdate(
	collName string,
	apiVersion string,
	data *models.Entity,
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateData(data)
	if err != nil {
		return nil, err
	}

	path := paths.CreateOrUpdateCollectionUrl(collName, apiVersion)

	res, err := c.Execute(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityMutationResponse(res)
}

func (c *collectionImpl) CreateOrUpdateBulk(
	collName string,
	apiVersion string,
	data *models.EntityCreateOrUpdateBulkRequest,
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateData(data)
	if err != nil {
		return nil, err
	}

	path := paths.CreateOrUpdateBulkCollectionUrl(collName, apiVersion)

	res, err := c.Execute(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityMutationResponse(res)
}

func (c *collectionImpl) MoveEntitiesToCollection(
	collName string,
	apiVersion string,
	guids []string,
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateGuids(guids)
	if err != nil {
		return nil, err
	}

	path := paths.MoveEntitiesHereUrl(collName, apiVersion)

	data := map[string][]string{"entityGuids": guids}

	res, err := c.Execute(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityMutationResponse(res)
}
