package url

type Url struct {
	Short string
	Full  string
	Hide  bool
}

func RemoveDuplicates(in []Url) []Url {
	seen := make(map[string]bool)
	filtered := []Url{}

	for _, x := range in {
		if seen[x.Full] {
			continue
		}
		seen[x.Full] = true

		filtered = append(filtered, x)
	}

	return filtered
}
