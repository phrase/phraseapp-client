package main

import "github.com/phrase/phraseapp-go/phraseapp"

var Debug bool

type ProjectLocales interface {
	ProjectIds() []string
}

type LocaleCacheKey struct {
	ProjectID string
	Branch    string
}

type LocaleCache map[LocaleCacheKey][]*phraseapp.Locale

func LocalesForProjects(client *phraseapp.Client, projectLocales ProjectLocales, branch string) (LocaleCache, error) {

	projectIdToLocales := LocaleCache{}
	for _, pid := range projectLocales.ProjectIds() {
		key := LocaleCacheKey{
			ProjectID: pid,
			Branch:    branch,
		}

		if _, ok := projectIdToLocales[key]; !ok {
			remoteLocales, err := RemoteLocales(client, key)
			if err != nil {
				return nil, err
			}

			projectIdToLocales[key] = remoteLocales
		}
	}
	return projectIdToLocales, nil

}

func RemoteLocales(client *phraseapp.Client, key LocaleCacheKey) ([]*phraseapp.Locale, error) {
	page := 1
	locales, err := client.LocalesList(key.ProjectID, page, 25, &phraseapp.LocalesListParams{Branch: &key.Branch})
	if err != nil {
		return nil, err
	}
	result := locales
	for len(locales) == 25 {
		page = page + 1
		locales, err = client.LocalesList(key.ProjectID, page, 25, &phraseapp.LocalesListParams{Branch: &key.Branch})
		if err != nil {
			return nil, err
		}
		result = append(result, locales...)
	}
	return result, nil
}
