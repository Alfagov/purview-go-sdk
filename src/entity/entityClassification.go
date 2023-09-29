package entity

import (
	"github.com/Alfagov/purview-go-sdk/src/models"
	"github.com/Alfagov/purview-go-sdk/src/paths"
	"github.com/Alfagov/purview-go-sdk/src/response"
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

	path := paths.AddClassificationPath

	data := map[string]interface{}{
		"classification": classification,
		"entityGuids":    entityGuids,
	}

	resp, err := e.Execute(http.MethodPost, path, data)
	if err != nil {
		return err
	}

	return response.ProcessEmptyResponse(resp)
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

	path := paths.ClassificationsPath(guid)

	resp, err := e.Execute(http.MethodPost, path, classifications)
	if err != nil {
		return err
	}

	return response.ProcessEmptyResponse(resp)
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

	path := paths.ClassesByUniqueAttributePath(attributes)

	resp, err := e.Execute(http.MethodPost, path, classifications)
	if err != nil {
		return err
	}

	return response.ProcessEmptyResponse(resp)
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

	path := paths.ClassificationPath(guid, classificationName)

	resp, err := e.Execute(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) DeleteClassificationByUniqueAttribute(
	name string, uniqueAttributes *models.QueryAttributes,
) error {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	path := paths.DeleteClassificationByUniqueAttributePath(
		uniqueAttributes,
		name,
	)

	resp, err := e.Execute(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
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

	path := paths.ClassificationPath(guid, name)

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessClassificationResponse(res)
}

func (e *entityImpl) GetClassifications(guid string) (
	*models.AtlasClassifications, error,
) {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return nil, err
	}

	path := paths.ClassificationsPath(guid)

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessClassificationsResponse(res)
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

	path := paths.SetClassificationsPath

	resp, err := e.Execute(http.MethodPost, path, classifications)
	return response.ProcessStringListResponse(resp)
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

	path := paths.ClassificationsPath(guid)

	resp, err := e.Execute(http.MethodPost, path, classifications)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) UpdateClassificationsByUniqueAttribute(
	classifications []*models.AtlasClassification,
	uniqueAttributes *models.QueryAttributes,
) error {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	err = utils.ValidateData(classifications)
	if err != nil {
		return err
	}

	path := paths.ClassesByUniqueAttributePath(uniqueAttributes)

	resp, err := e.Execute(http.MethodPost, path, classifications)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}
