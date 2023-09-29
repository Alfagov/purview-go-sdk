package discovery

import (
	"github.com/Alfagov/purview-go-sdk/src/client"
	"github.com/Alfagov/purview-go-sdk/src/models"
	"github.com/Alfagov/purview-go-sdk/src/paths"
	"github.com/Alfagov/purview-go-sdk/src/response"
	"github.com/Alfagov/purview-go-sdk/src/utils"
	"net/http"
)

type Discovery interface {
	AutoComplete(
		apiVersion string,
		request *models.AutoCompleteRequest,
	) (*models.AutoCompleteResult, error)
	Browse(
		apiVersion string,
		request *models.BrowseRequest,
	) (*models.BrowseResult, error)
	Query(
		apiVersion string,
		request *models.QueryRequest,
	) (*models.SearchResult, error)
	Suggest(
		apiVersion string,
		request *models.SuggestRequest,
	) (*models.SuggestResult, error)
}

type discoveryImpl struct {
	client.HttpClient
}

func NewDiscovery(httpClient client.HttpClient) Discovery {
	return &discoveryImpl{httpClient}
}

func (d *discoveryImpl) AutoComplete(
	apiVersion string,
	request *models.AutoCompleteRequest,
) (*models.AutoCompleteResult, error) {
	err := utils.ValidateData(request)
	if err != nil {
		return nil, err
	}

	path := paths.AutoCompleteUrl(apiVersion)

	res, err := d.Execute(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	return response.ProcessAutoCompleteResult(res)
}

func (d *discoveryImpl) Browse(
	apiVersion string,
	request *models.BrowseRequest,
) (*models.BrowseResult, error) {
	err := utils.ValidateData(request)
	if err != nil {
		return nil, err
	}

	path := paths.BrowseUrl(apiVersion)

	res, err := d.Execute(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	return response.ProcessBrowseResult(res)
}

func (d *discoveryImpl) Query(
	apiVersion string,
	request *models.QueryRequest,
) (*models.SearchResult, error) {
	err := utils.ValidateData(request)
	if err != nil {
		return nil, err
	}

	path := paths.QueryUrl(apiVersion)

	res, err := d.Execute(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	return response.ProcessSearchResult(res)
}

func (d *discoveryImpl) Suggest(
	apiVersion string,
	request *models.SuggestRequest,
) (*models.SuggestResult, error) {
	err := utils.ValidateData(request)
	if err != nil {
		return nil, err
	}

	path := paths.SuggestUrl(apiVersion)

	res, err := d.Execute(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	return response.ProcessSuggestResult(res)
}
