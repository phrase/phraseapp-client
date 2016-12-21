package main

import "github.com/phrase/phraseapp-go/phraseapp"

var Debug bool

type ProjectLocales interface {
	ProjectIds() []string
}

func LocalesForProjects(client *phraseapp.Client, projectLocales ProjectLocales) (map[string][]*phraseapp.Locale, error) {
	projectIdToLocales := map[string][]*phraseapp.Locale{}
	for _, pid := range projectLocales.ProjectIds() {
		if _, ok := projectIdToLocales[pid]; !ok {
			remoteLocales, err := RemoteLocales(client, pid)
			if err != nil {
				return nil, err
			}
			projectIdToLocales[pid] = remoteLocales
		}
	}
	return projectIdToLocales, nil
}

func RemoteLocales(client *phraseapp.Client, projectId string) ([]*phraseapp.Locale, error) {
	page := 1
	locales, err := client.LocalesList(projectId, page, 25)
	if err != nil {
		return nil, err
	}
	result := locales
	for len(locales) == 25 {
		page = page + 1
		locales, err = client.LocalesList(projectId, page, 25)
		if err != nil {
			return nil, err
		}
		result = append(result, locales...)
	}
	return result, nil
}

func GetFormats(client *phraseapp.Client) (map[string]*phraseapp.Format, error) {
	formats, err := client.FormatsList(1, 25)
	if err != nil {
		return nil, err
	}
	formatMap := map[string]*phraseapp.Format{}
	for _, format := range formats {
		formatMap[format.ApiName] = format
	}
	return formatMap, nil
}
