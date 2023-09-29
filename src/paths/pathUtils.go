package paths

import (
	"net/url"
)

func createGuidsParamValues(guids []string) string {
	values := url.Values{}
	for _, guid := range guids {
		values.Add("guid", guid)
	}
	return values.Encode()
}
