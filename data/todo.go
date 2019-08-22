package data

// Todo stores a task
type Todo struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func (t *Todo) checkTags(tags ...string) bool {
	seenTags := map[string]int{}

	for _, tag := range t.Tags {
		seenTags[tag]++
	}

	for _, tag := range tags {
		if tagCount, isTagSet := seenTags[tag]; !isTagSet || tagCount < 1 {
			return false
		}
	}

	return true
}
