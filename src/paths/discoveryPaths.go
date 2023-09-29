package paths

import "net/url"

var (
	apiUrl          = &url.URL{Path: "/catalog/api/"}
	autoCompleteUrl = apiUrl.JoinPath("search", "autocomplete")
	browseUrl       = apiUrl.JoinPath("browse")
	queryUrl        = apiUrl.JoinPath("search", "query")
	suggestUrl      = apiUrl.JoinPath("search", "suggest")
)

func AutoCompleteUrl(version string) *url.URL {
	path := &url.URL{Path: autoCompleteUrl.Path}
	path.RawQuery = "api-version=" + version

	return path
}

func BrowseUrl(version string) *url.URL {
	path := &url.URL{Path: browseUrl.Path}
	path.RawQuery = "api-version=" + version

	return path
}

func QueryUrl(version string) *url.URL {
	path := &url.URL{Path: queryUrl.Path}
	path.RawQuery = "api-version=" + version

	return path
}

func SuggestUrl(version string) *url.URL {
	path := &url.URL{Path: suggestUrl.Path}
	path.RawQuery = "api-version=" + version

	return path
}
