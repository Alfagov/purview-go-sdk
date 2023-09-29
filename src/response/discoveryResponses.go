package response

import (
	"github.com/Alfagov/purview-go-sdk/src/models"
	"io"
	"net/http"
)

func ProcessAutoCompleteResult(
	response *http.Response,
) (*models.AutoCompleteResult, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.AutoCompleteResult{}.UnmarshalJson(body)
}

func ProcessBrowseResult(
	response *http.Response,
) (*models.BrowseResult, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.BrowseResult{}.UnmarshalJson(body)
}

func ProcessSearchResult(
	response *http.Response,
) (*models.SearchResult, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.SearchResult{}.UnmarshalJson(body)
}

func ProcessSuggestResult(
	response *http.Response,
) (*models.SuggestResult, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.SuggestResult{}.UnmarshalJson(body)
}
