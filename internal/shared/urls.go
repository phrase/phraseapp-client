package shared

import "fmt"

const (
	DocsBaseUrl   = "https://phraseapp.com/docs"
	DocsConfigUrl = DocsBaseUrl + "/developers/cli/configuration"
)

func DocsFormatsUrl(formatName string) string {
	return fmt.Sprintf("%s/guides/formats/%s", DocsBaseUrl, formatName)
}
