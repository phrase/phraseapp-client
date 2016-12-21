package paclient

import (
	"os"
	"regexp"
)

var (
	separator         = string(os.PathSeparator)
	PlaceholderRegexp = regexp.MustCompile("<(locale_name|tag|locale_code)>")
)

const (
	DocsBaseUrl   = "https://phraseapp.com/docs"
	DocsConfigUrl = DocsBaseUrl + "/developers/cli/configuration"
)
