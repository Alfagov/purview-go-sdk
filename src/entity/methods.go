package entity

import (
	"github.com/Alfagov/purview-go-sdk/src/models"
	"github.com/Alfagov/purview-go-sdk/src/utils"
	"net/http"
)

func (e *entityImpl) AddClassification(
	classification *models.AtlasClassification, entityGuids []string,
) error {
	err := utils.ValidateGuids(entityGuids)
	if err != nil {
		return err
	}
	err = utils.ValidateData(classification)
	if err != nil {
		return err
	}

	path := ""

	data := map[string]interface{}{
		"classification": classification,
		"entityGuids":    entityGuids,
	}

	resp, err := e.Execute(http.MethodPost, path, data)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) AddClassifications(
	guid string, classifications []*models.AtlasClassification,
) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateData(classifications)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, classifications)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) AddClassificationsByUniqueAttribute(
	classifications []*models.AtlasClassification,
	attributes *models.QueryAttributes,
) error {
	err := utils.ValidateQueryAttributes(attributes)
	if err != nil {
		return err
	}
	err = utils.ValidateData(classifications)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, classifications)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) AddLabel(guid string, labels []string) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, labels)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) AddLabelsByUniqueAttribute(
	labels []string,
	attributes *models.QueryAttributes,
) error {
	err := utils.ValidateQueryAttributes(attributes)
	if err != nil {
		return err
	}
	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, labels)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) AddOrUpdateBusinessMetadata(
	guid string,
	businessMetadata map[string]interface{},
) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateData(businessMetadata)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, businessMetadata)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) AddOrUpdateBusinessMetadataAttributes(
	businessMetadataAttributes map[string]interface{},
	attributes *models.QueryAttributes,
) error {
	err := utils.ValidateQueryAttributes(attributes)
	if err != nil {
		return err
	}
	err = utils.ValidateData(businessMetadataAttributes)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, businessMetadataAttributes)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) GetByGuid(
	guid string,
	customUnmarshaler func(data []byte) (*models.AtlasEntityWithExtInfo, error),
) (*models.AtlasEntityWithExtInfo, error) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageGetEntityResponse(res, customUnmarshaler)
}

func (e *entityImpl) GetByUniqueAttributes(
	uniqueAttributes *models.QueryAttributes,
	params *models.QueryParams,
	customUnmarshaler func(data []byte) (*models.AtlasEntityWithExtInfo, error),
) (*models.AtlasEntityWithExtInfo, error) {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return nil, err
	}
	err = utils.ValidateData(uniqueAttributes)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageGetEntityResponse(res, customUnmarshaler)
}

func (e *entityImpl) CreateOrUpdateEntities(
	data *models.EntityCreateOrUpdateBulkRequest,
	customUnmarshaler func(data []byte) (*models.EntityMutationResponse, error),
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateData(data)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return ManageENtityMutationResponse(res, customUnmarshaler)
}

func (e *entityImpl) CreateOrUpdate(
	data *models.EntityCreateOrUpdateRequest,
	customUnmarshaler func(data []byte) (*models.EntityMutationResponse, error),
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateData(data)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return ManageENtityMutationResponse(res, customUnmarshaler)
}

func (e *entityImpl) DeleteBusinessMetadata(
	guid string, businessMetadata map[string]interface{},
) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateData(businessMetadata)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, businessMetadata)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) DeleteBusinessMetadataAttributes(
	attributes map[string]string, businessMetadata map[string]interface{},
) error {
	err := utils.ValidateAttributes(attributes)
	if err != nil {
		return err
	}

	err = utils.ValidateData(businessMetadata)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, businessMetadata)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) DeleteByGuid(guid string) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, nil)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) DeleteByGuids(guids []string) error {
	err := utils.ValidateGuids(guids)
	if err != nil {
		return err
	}

	path := ""

	_, err = e.Execute(http.MethodDelete, path, nil)
	return err
}

func (e *entityImpl) DeleteByUniqueAttribute(uniqueAttributes map[string]string) error {
	err := utils.ValidateAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, nil)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) DeleteClassification(
	guid string, classificationName string,
) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateClassificationName(classificationName)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, nil)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) DeleteClassificationByUniqueAttribute(uniqueAttributes map[string]string) error {
	err := utils.ValidateAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, nil)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) DeleteLabels(guid string, labels []string) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, nil)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) DeleteLabelsByUniqueAttribute(
	uniqueAttributes map[string]string, labels []string,
) error {
	err := utils.ValidateAttributes(uniqueAttributes)
	if err != nil {
		return err
	}
	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodDelete, path, nil)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) GetClassification(guid string, name string) (
	*models.AtlasClassification, error,
) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}
	err = utils.ValidateClassificationName(name)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageGetClassificationResponse(res, nil)
}

func (e *entityImpl) GetClassifications(guid string) (
	*models.AtlasClassifications, error,
) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageGetClassificationsResponse(res, nil)
}

func (e *entityImpl) GetEntitiesByUniqueAttributes(uniqueAttributes map[string]string) (
	*models.AtlasEntityWithExtInfo, error,
) {
	err := utils.ValidateAttributes(uniqueAttributes)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageGetEntityResponse(res, nil)
}

func (e *entityImpl) GetHeader(guid string) (*models.AtlasEntityHeader, error) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}

	path := ""

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageGetEntityHeaderResponse(res, nil)
}

func (e *entityImpl) GetSampleBusinessMetadataTemplate(templateName string) (
	interface{}, error,
) {
	//TODO implement me
	panic("implement me")
}

func (e *entityImpl) ImportBusinessMetadata() (
	*models.BulkImportResponse,
	error,
) {

	path := ""

	resp, err := e.Execute(http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageBulkImportResponse(resp, nil)
}

func (e *entityImpl) ListByGuids(guids []string) (
	*models.AtlasEntityWithExtInfo, error,
) {
	err := utils.ValidateGuids(guids)
	if err != nil {
		return nil, err
	}

	path := ""

	resp, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageGetEntityResponse(resp, nil)
}

func (e *entityImpl) PartialUpdateEntityAttributeByGuid(
	guid string, name string, value interface{},
	customUnmarshaler func(data []byte) (*models.EntityMutationResponse, error),
) (*models.EntityMutationResponse, error) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}
	err = utils.ValidateData(value)
	if err != nil {
		return nil, err
	}

	path := ""

	resp, err := e.Execute(http.MethodPatch, path, nil)
	if err != nil {
		return nil, err
	}

	return ManageENtityMutationResponse(resp, customUnmarshaler)
}

func (e *entityImpl) PartialUpdateEntityByUniqueAttributes(
	uniqueAttributes map[string]string, entityData map[string]interface{},
	customUnmarshaler func(data []byte) (*models.AtlasEntityWithExtInfo, error),
) (*models.AtlasEntityWithExtInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (e *entityImpl) SetClassifications(
	guid string, classifications map[string]models.AtlasEntityHeader,
) ([]string, error) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateData(classifications)
	if err != nil {
		return nil, err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, classifications)
	return ManageStringListResponse(resp, nil)
}

func (e *entityImpl) SetLabels(guid string, labels []string) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}

	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, labels)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) SetLabelsByUniqueAttribute(
	uniqueAttributes map[string]string, labels []string,
) error {
	err := utils.ValidateAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, labels)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) UpdateClassifications(
	guid string, classifications []*models.AtlasClassification,
) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}

	err = utils.ValidateData(classifications)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, classifications)
	return ManageEmptyResponse(resp)
}

func (e *entityImpl) UpdateClassificationsByUniqueAttribute(
	classifications []*models.AtlasClassification,
	uniqueAttributes map[string]string,
) error {
	err := utils.ValidateAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	err = utils.ValidateData(classifications)
	if err != nil {
		return err
	}

	path := ""

	resp, err := e.Execute(http.MethodPost, path, classifications)
	return ManageEmptyResponse(resp)
}
