package paths

import "net/url"

var (
	collectionsBaseUrl = &url.URL{Path: "/catalog/api/collections"}
)

func CreateOrUpdateCollectionUrl(collName string, apiVersion string) *url.URL {
	path := collectionsBaseUrl.JoinPath(collName, "entity")
	path.RawQuery = "api-version=" + apiVersion

	return path
}

func CreateOrUpdateBulkCollectionUrl(
	collName string,
	apiVersion string,
) *url.URL {
	path := collectionsBaseUrl.JoinPath(collName, "entity", "bulk")
	path.RawQuery = "api-version=" + apiVersion

	return path
}

func MoveEntitiesHereUrl(
	collName string,
	apiVersion string,
) *url.URL {
	path := collectionsBaseUrl.JoinPath(collName, "entity", "moveHere")
	path.RawQuery = "api-version=" + apiVersion

	return path
}
