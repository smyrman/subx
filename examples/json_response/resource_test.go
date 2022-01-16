package api_test

import (
	"strings"
	"testing"
	"bytes"

	"github.com/smyrman/subx"
	api "github.com/smyrman/subx/examples/json_response"
)

type AlbumView struct {
	ID          string `json:"id,omitempty"` // Deliberate bug: wrong type.
	Title       string `json:"title"`
	Description string `json:"description"`
}

func TestInsert(t *testing.T) {
	const data = `{
		"title": "Foo",
		"description": "bar"
	}`

	rsc := api.NewResource[api.Album]()

	var rec bytes.Buffer
	req := strings.NewReader(data)
	rsc.Insert(&rec, req)

	t.Logf("got JSON response: %s",rec.String())
	t.Run("Expect correct response", subx.Test(
		rec.Bytes,
		subx.DecodeJSON(func(vf func() AlbumView) error{
			var acc subx.Accumulator
			acc.SetPrefix("unexpected response fields:")
			a := vf()
			acc.Registerf("id: %w", subx.CompareNotEqual("")(subx.Value(a.ID)))
			acc.Registerf("title: %w", subx.CompareEqual("foo")(subx.Value(a.Title)))  // Deliberate bug: wrong case.
			return acc.Result()
		}),

	))
}
