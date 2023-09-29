package entity

import (
	"github.com/Alfagov/purview-go-sdk/src/models"
	"github.com/Alfagov/purview-go-sdk/src/paths"
	"github.com/Alfagov/purview-go-sdk/src/response"
	"github.com/Alfagov/purview-go-sdk/src/utils"
	"net/http"
)

func (e *entityImpl) GetByGuid(
	guid string,
	options *models.QueryParams,
) (*models.AtlasEntityWithExtInfo, error) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}

	path := paths.EntityByGuidPath(guid, options)

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityWithExtInfoResponse(res)
}

func (e *entityImpl) GetByUniqueAttributes(
	uniqueAttributes *models.QueryAttributes,
	params *models.QueryParams,
) (*models.AtlasEntityWithExtInfo, error) {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return nil, err
	}
	err = utils.ValidateData(uniqueAttributes)
	if err != nil {
		return nil, err
	}

	path := paths.EntityByUniqueAttributePath(uniqueAttributes, params)

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityWithExtInfoResponse(res)
}

func (e *entityImpl) CreateOrUpdateEntities(
	data *models.EntityCreateOrUpdateBulkRequest,
	options *models.QueryParams,
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateData(data)
	if err != nil {
		return nil, err
	}

	path := paths.CreateOrUpdateEntitiesPath(options)

	res, err := e.Execute(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityMutationResponse(res)
}

func (e *entityImpl) CreateOrUpdate(
	data *models.EntityCreateOrUpdateRequest,
	options *models.QueryParams,
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateData(data)
	if err != nil {
		return nil, err
	}

	path := paths.CreateOrUpdateEntityPath(options)

	res, err := e.Execute(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityMutationResponse(res)
}

func (e *entityImpl) DeleteByGuid(guid string) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}

	path := paths.EntityByGuidPath(guid, nil)

	resp, err := e.Execute(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) DeleteByGuids(guids []string) error {
	err := utils.ValidateGuids(guids)
	if err != nil {
		return err
	}

	path := paths.DeleteEntitiesByGuids(guids)

	_, err = e.Execute(http.MethodDelete, path, nil)
	return err
}

func (e *entityImpl) DeleteByUniqueAttribute(
	uniqueAttributes *models.
		QueryAttributes,
) error {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	path := paths.EntityByUniqueAttributePath(uniqueAttributes, nil)

	resp, err := e.Execute(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) GetEntitiesByUniqueAttributes(
	uniqueAttributes *models.
		QueryAttributes,
	queryOptions *models.QueryParams,
) (
	*models.AtlasEntityWithExtInfo, error,
) {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return nil, err
	}

	path := paths.GetEntitiesByUniqueAttributesPath(
		uniqueAttributes,
		queryOptions,
	)

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityWithExtInfoResponse(res)
}

func (e *entityImpl) GetHeader(guid string) (*models.AtlasEntityHeader, error) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}

	path := paths.HeaderPath(guid)

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessGetEntityHeaderResponse(res)
}

func (e *entityImpl) ListByGuids(guids []string, options *models.QueryParams) (
	*models.AtlasEntityWithExtInfo, error,
) {
	err := utils.ValidateGuids(guids)
	if err != nil {
		return nil, err
	}

	path := paths.ListEntitiesByGuidsPath(guids, options)

	resp, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityWithExtInfoResponse(resp)
}

func (e *entityImpl) PartialUpdateEntityAttributeByGuid(
	guid string, name string, value interface{},
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}
	err = utils.ValidateData(value)
	if err != nil {
		return nil, err
	}

	path := paths.PartialUpdateEntityAttributeByGuidPath(guid, name)

	resp, err := e.Execute(http.MethodPatch, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityMutationResponse(resp)
}

func (e *entityImpl) PartialUpdateEntityByUniqueAttributes(
	uniqueAttributes *models.QueryAttributes,
	entityData *models.AtlasEntityWithExtInfo,
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateData(entityData)
	if err != nil {
		return nil, err
	}

	path := paths.EntityByUniqueAttributePath(uniqueAttributes, nil)

	resp, err := e.Execute(http.MethodPut, path, entityData)
	if err != nil {
		return nil, err
	}

	return response.ProcessEntityMutationResponse(resp)
}
