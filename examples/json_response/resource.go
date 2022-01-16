package api

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type Album struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (a Album) WithID(id int) Album {
	a.ID = id
	return a
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Model[T any] interface {
	WithID(i int) T
}

type Resource[M Model[M]] struct {
	lock sync.RWMutex
	db   map[int]M
	incr int
}

func NewResource[M Model[M]]() Resource[M] {
	return Resource[M]{
		db: make(map[int]M),
	}
}

func (rsc *Resource[M]) Insert(w io.Writer, r io.Reader) {
	var target M

	enc := json.NewEncoder(w)
	dec := json.NewDecoder(r)

	dec.DisallowUnknownFields()
	if err := dec.Decode(&target); err != nil {
		must(enc.Encode(Error{
			Code:    http.StatusBadRequest,
			Message: "could not decode request body: " + err.Error(),
		}))
		return
	}

	rsc.lock.Lock()
	defer rsc.lock.Unlock()

	rsc.incr++
	rsc.db[rsc.incr] = target.WithID(rsc.incr)

	must(enc.Encode(target))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
