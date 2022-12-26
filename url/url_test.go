package url

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		data []Url
		want []Url
	}{
		{
			name: "no dupes",
			data: []Url{
				{"a", "a"},
				{"b", "b"},
				{"c", "c"},
			},
			want: []Url{
				{"a", "a"},
				{"b", "b"},
				{"c", "c"},
			},
		},
		{
			name: "have dupes",
			data: []Url{
				{"a", "a"},
				{"b", "a"},
				{"c", "c"},
			},
			want: []Url{
				{"a", "a"},
				{"c", "c"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RemoveDuplicates(test.data)
			assertEquals(t, got, test.want)
		})
	}
}

func assertEquals(t testing.TB, got, want []Url) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}
