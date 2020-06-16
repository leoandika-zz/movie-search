package rest

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/movie-search/Gateway/pb"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type pkgMock struct {
	MovieSearch       *pb.MockMovieSearchClient
}

func mockPackage(t *testing.T) (*Routes, pkgMock){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMovieSearchClient := pb.NewMockMovieSearchClient(mockCtrl)

	p := &Routes{}

	return p, pkgMock{
		MovieSearch: mockMovieSearchClient,
	}
}

func TestInitPackage(t *testing.T) {
	rest := InitPackage()

	assert.NotNil(t, rest)
}

func TestInitRoutes(t *testing.T) {
	rest := InitPackage()

	router := mux.NewRouter()
	rest.InitRoutes(router)
}

func TestSearchMovie_AbleToGetMovieList(t *testing.T) {

	p, m := mockPackage(t)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://www.testing.com?title=anything&page=1", nil)



	req := &pb.SearchRequest{
		MovieName: "anything",
		Page:      int32(1),
	}

	m.MovieSearch.EXPECT().SearchMovie(gomock.Any(), req).Return(&pb.SearchResponse{
		MovieList: []*pb.Movie{
			{
				Title:  "test title",
				Year:   "test year",
				Type:   "test type",
				ImdbID: "test imdbid",
				Poster: "test poster",
			},
		},
		Err: "",
	}, nil)

	p.SearchMovie(w, r)
}
