package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Alfagov/purview-go-sdk/src/models"
	"io"
	"net/http"
)

func ManageEntityWithExtInfoResponse(
	response *http.Response,
	customUnmarshaler func(data []byte) (*models.AtlasEntityWithExtInfo, error),
) (*models.AtlasEntityWithExtInfo, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseError.ErrorMessage)
	}

	if customUnmarshaler != nil {
		return customUnmarshaler(body)
	}

	return models.AtlasEntityWithExtInfo{}.UnmarshalJson(body)
}

func ManageGetEntityHeaderResponse(
	response *http.Response,
	customUnmarshaler func(data []byte) (*models.AtlasEntityHeader, error),
) (*models.AtlasEntityHeader, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseError.ErrorMessage)
	}

	if customUnmarshaler != nil {
		return customUnmarshaler(body)
	}

	return models.AtlasEntityHeader{}.UnmarshalJson(body)
}

func ManageBulkImportResponse(
	response *http.Response,
	customUnmarshaler func(data []byte) (*models.BulkImportResponse, error),
) (*models.BulkImportResponse, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseError.ErrorMessage)
	}

	if customUnmarshaler != nil {
		return customUnmarshaler(body)
	}

	return models.BulkImportResponse{}.UnmarshalJson(body)
}

func ManageEntityMutationResponse(
	response *http.Response,
	customUnmarshaler func(data []byte) (*models.EntityMutationResponse, error),
) (*models.EntityMutationResponse, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseError.ErrorMessage)
	}

	if customUnmarshaler != nil {
		return customUnmarshaler(body)
	}

	return models.EntityMutationResponse{}.UnmarshalJson(body)
}

func ManageClassificationResponse(
	response *http.Response,
	customUnmarshaler func(data []byte) (*models.AtlasClassification, error),
) (*models.AtlasClassification, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseError.ErrorMessage)
	}

	if customUnmarshaler != nil {
		return customUnmarshaler(body)
	}

	return models.AtlasClassification{}.UnmarshalJson(body)
}

func ManageClassificationsResponse(
	response *http.Response,
	customUnmarshaler func(data []byte) (*models.AtlasClassifications, error),
) (*models.AtlasClassifications, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseError.ErrorMessage)
	}

	if customUnmarshaler != nil {
		return customUnmarshaler(body)
	}

	return models.AtlasClassifications{}.UnmarshalJson(body)
}

func ManageStringListResponse(
	response *http.Response,
	customUnmarshaler func(data []byte) ([]string, error),
) ([]string, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseError.ErrorMessage)
	}

	if customUnmarshaler != nil {
		return customUnmarshaler(body)
	}

	var out []string
	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func ManageEmptyResponse(response *http.Response) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		responseError, err := models.ErrorResponse{}.UnmarshalJson(body)
		if err != nil {
			return err
		}

		message := fmt.Sprintf(
			"Error code: %s, message: %s", responseError.ErrorCode,
			responseError.ErrorMessage,
		)

		return errors.New(message)
	}

	return nil
}
