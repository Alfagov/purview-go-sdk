package response

import (
	"github.com/Alfagov/pwCatalog/src/models"
	"io"
	"net/http"
)

func ProcessEntityWithExtInfoResponse(
	response *http.Response,
) (*models.AtlasEntityWithExtInfo, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.AtlasEntityWithExtInfo{}.UnmarshalJson(body)
}

func ProcessGetEntityHeaderResponse(
	response *http.Response,
) (*models.AtlasEntityHeader, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.AtlasEntityHeader{}.UnmarshalJson(body)
}

func ProcessBulkImportResponse(
	response *http.Response,
) (*models.BulkImportResponse, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.BulkImportResponse{}.UnmarshalJson(body)
}

func ProcessEntityMutationResponse(
	response *http.Response,
) (*models.EntityMutationResponse, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	return models.EntityMutationResponse{}.UnmarshalJson(body)
}
