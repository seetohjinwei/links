package page

import (
	"reflect"
	"testing"

	"github.com/seetohjinwei/links/url"
)

func TestGenerate(t *testing.T) {
	names := []string{"test.tmpl"}
	links := []url.Url{
		{Short: "a1", Full: "full_a"},
		{Short: "a2", Full: "full_a"},
		{Short: "c1", Full: "full_c"},
	}

	got := Generate("../templates", "test", names, links)
	want := []byte(`<div>
    <ul>
        <li><a href=full_a>a1</a></li>
        <li><a href=full_c>c1</a></li>
    </ul>
</div>
`)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %q want %q", string(got), string(want))
	}
}
