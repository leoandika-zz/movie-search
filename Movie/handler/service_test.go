package handler

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUp() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}

func TestSearch(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			   "Search":[
				  {
					 "Title":"test title 1",
					 "Year":"2016",
					 "imdbID":"tt4853102",
					 "Type":"movie",
					 "Poster":"test1.jpg"
				  },
				  {
					 "Title":"test title 2",
					 "Year":"2013",
					 "imdbID":"tt2166834",
					 "Type":"movie",
					 "Poster":"test2.jpg"
				  }],
			   "totalResults":"100",
			   "Response":"True"
		}`)
	}))
	defer mockServer.Close()

	expected := []Movie{
		{
			Title:  "test title 1",
			Year:   "2016",
			IMDBID: "tt4853102",
			Type:   "movie",
			Poster: "test1.jpg",
		},
		{
			Title:  "test title 2",
			Year:   "2013",
			IMDBID: "tt2166834",
			Type:   "movie",
			Poster: "test2.jpg",
		},
	}

	service, ctx := setUp()

	actual, _ := service.Search(ctx, mockServer.URL, SearchRequest{
		MovieName: "batman",
		Page:      1,
	})

	assert.Equal(t, expected, actual)
}
