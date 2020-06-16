package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Service interface {
	Search(ctx context.Context, url string, req SearchRequest) ([]Movie, error)
}

type MovieSearchService struct{}

func NewService() Service {
	return MovieSearchService{}
}

func (MovieSearchService) Search(ctx context.Context, url string, req SearchRequest) ([]Movie, error) {

	var httpResp HTTPResponse

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("http request failed. error : %s\n", err)
		return httpResp.Result, err
	}

	if response.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&httpResp)
	} else {
		return httpResp.Result, errors.New("http response is not 200")
	}

	return httpResp.Result, nil
}
