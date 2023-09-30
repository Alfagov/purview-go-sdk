package entity

import (
	"github.com/Alfagov/pwCatalog/src/models"
	"github.com/Alfagov/pwCatalog/src/paths"
	"github.com/Alfagov/pwCatalog/src/response"
	"github.com/Alfagov/pwCatalog/src/utils"
	"net/http"
)

func (e *entityImpl) AddLabel(guid string, labels []string) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := paths.LabelPath(guid)

	resp, err := e.Execute(http.MethodPost, path, labels)
	if err != nil {
		return err
	}

	return response.ProcessEmptyResponse(resp)
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

	path := paths.LabelsByUniqueAttributePath(attributes)

	resp, err := e.Execute(http.MethodPost, path, labels)
	if err != nil {
		return err
	}

	return response.ProcessEmptyResponse(resp)
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

	path := paths.LabelPath(guid)

	resp, err := e.Execute(http.MethodDelete, path, labels)
	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) DeleteLabelsByUniqueAttribute(
	uniqueAttributes *models.QueryAttributes, labels []string,
) error {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return err
	}
	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := paths.LabelsByUniqueAttributePath(uniqueAttributes)

	resp, err := e.Execute(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
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

	path := paths.LabelPath(guid)

	resp, err := e.Execute(http.MethodPost, path, labels)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) SetLabelsByUniqueAttribute(
	uniqueAttributes *models.QueryAttributes, labels []string,
) error {
	err := utils.ValidateQueryAttributes(uniqueAttributes)
	if err != nil {
		return err
	}

	err = utils.ValidateData(labels)
	if err != nil {
		return err
	}

	path := paths.LabelsByUniqueAttributePath(uniqueAttributes)

	resp, err := e.Execute(http.MethodPost, path, labels)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}
