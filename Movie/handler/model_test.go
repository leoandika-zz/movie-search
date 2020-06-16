package handler

import (
	"context"
	"github.com/movie-search/Movie/pb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeGRPCSearchRequest(t *testing.T) {

	param := &pb.SearchRequest{
		MovieName: "test title",
		Page:      1,
	}

	expected := SearchRequest{
		MovieName: "test title",
		Page:      1,
	}

	actual, err := DecodeGRPCSearchRequest(context.TODO(), param)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected, actual.(SearchRequest))
}

func TestEncodeGRPCSearchResponse(t *testing.T) {
	param := SearchResponse{
		Response: []Movie{
			{
				Title:  "test1",
				Year:   "2000",
				Type:   "movie",
				IMDBID: "12345",
				Poster: "test1.jpg",
			},
			{
				Title:  "test2",
				Year:   "3000",
				Type:   "movie",
				IMDBID: "67891",
				Poster: "test2.jpg",
			},
		},
		ErrorMessage: "",
	}

	expected := &pb.SearchResponse{
		MovieList: []*pb.Movie{
			{
				Title:  "test1",
				Year:   "2000",
				Type:   "movie",
				ImdbID: "12345",
				Poster: "test1.jpg",
			},
			{
				Title:  "test2",
				Year:   "3000",
				Type:   "movie",
				ImdbID: "67891",
				Poster: "test2.jpg",
			},
		},
		Err: "",
	}

	actual, err := EncodeGRPCSearchResponse(context.TODO(), param)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected, actual)
}
