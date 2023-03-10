package yaml

import (
	"log"

	"github.com/seetohjinwei/links/url"
	"gopkg.in/yaml.v3"
)

type T struct {
	Shorts []string `yaml:"shorts"`
	Full   string   `yaml:"full"`
	Hide   bool     `yaml:"hide"`
}

func Parse(data []byte) []url.Url {
	ts := []T{}
	err := yaml.Unmarshal(data, &ts)
	if err != nil {
		log.Fatalf("error while un-marshalling yaml file: %v", err)
	}

	links := []url.Url{}

	for _, t := range ts {
		for _, short := range t.Shorts {
			link := url.Url{Short: short, Full: t.Full, Hide: t.Hide}
			links = append(links, link)
		}
	}

	return links
}
