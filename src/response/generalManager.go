package response

import (
	"encoding/json"
	"errors"
	"github.com/Alfagov/pwCatalog/src/models"
	"io"
	"net/http"
)

func ProcessStringListResponse(
	response *http.Response,
) ([]string, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, processErrorResponse(body)
	}

	var out []string
	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func ProcessEmptyResponse(response *http.Response) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return processErrorResponse(body)
	}

	return nil
}

func ProcessFile(response *http.Response) ([]byte, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return nil, processErrorResponse(body)
	}

	return body, nil
}

func processErrorResponse(errMsg []byte) error {
	responseError, err := models.ErrorResponse{}.UnmarshalJson(errMsg)
	if err != nil {
		return err
	}

	return errors.New(responseError.ErrorMessage)
}
