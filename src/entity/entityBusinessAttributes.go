package entity

import (
	"github.com/Alfagov/purview-go-sdk/src/models"
	"github.com/Alfagov/purview-go-sdk/src/paths"
	"github.com/Alfagov/purview-go-sdk/src/response"
	"github.com/Alfagov/purview-go-sdk/src/utils"
	"net/http"
)

func (e *entityImpl) AddOrUpdateBusinessMetadata(
	guid string,
	businessMetadata map[string]interface{},
	options *models.QueryParams,
) error {
	err := utils.ValidateGuid(guid)
	if err != nil {
		return err
	}
	err = utils.ValidateData(businessMetadata)
	if err != nil {
		return err
	}

	path := paths.BusinessMetadataPath(guid, options)

	resp, err := e.Execute(http.MethodPost, path, businessMetadata)
	if err != nil {
		return err
	}

	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) AddOrUpdateBusinessMetadataAttributes(
	businessMetadataAttributes map[string]interface{},
	guid string, bmName string,
) error {
	err := utils.ValidateData(businessMetadataAttributes)
	if err != nil {
		return err
	}

	path := paths.BusinessMetadataAttributePath(guid, bmName)

	resp, err := e.Execute(http.MethodPost, path, businessMetadataAttributes)
	if err != nil {
		return err
	}

	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) DeleteBusinessMetadata(
	guid string, options *models.QueryParams,
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

	path := paths.BusinessMetadataPath(guid, options)

	resp, err := e.Execute(http.MethodDelete, path, businessMetadata)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}

func (e *entityImpl) DeleteBusinessMetadataAttributes(
	guid string, bmName string, businessMetadata map[string]interface{},
) error {
	err := utils.ValidateData(businessMetadata)
	if err != nil {
		return err
	}

	path := paths.BusinessMetadataAttributePath(guid, bmName)

	resp, err := e.Execute(http.MethodDelete, path, businessMetadata)
	if err != nil {
		return err
	}
	return response.ProcessEmptyResponse(resp)
}
func (e *entityImpl) GetSampleBusinessMetadataTemplate() (
	[]byte, error,
) {

	path := paths.SampleBusinessMetadataTemplatePath

	res, err := e.Execute(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessFile(res)
}

func (e *entityImpl) ImportBusinessMetadata() (
	*models.BulkImportResponse,
	error,
) {

	path := paths.ImportBusinessMetadataPath

	resp, err := e.Execute(http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}

	return response.ProcessBulkImportResponse(resp)
}
