package story

import "encoding/json"

type option struct {
	arc  string
	text string
}

type arc struct {
	title   string
	story   []string
	options []option
}

func parseStory(story []byte) (map[string]arc, error) {
	var storyMap map[string]arc
	err := json.Unmarshal(story, &storyMap)
	if err != nil {
		return nil, err
	}
	return storyMap, nil
}
