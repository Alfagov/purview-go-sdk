package response

import (
	"github.com/Alfagov/purview-go-sdk/src/models"
	"io"
	"net/http"
)

func ProcessClassificationResponse(
	response *http.Response,
) (*models.AtlasClassification, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.AtlasClassification{}.UnmarshalJson(body)
}

func ProcessClassificationsResponse(
	response *http.Response,
) (*models.AtlasClassifications, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.AtlasClassifications{}.UnmarshalJson(body)
}
