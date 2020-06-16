package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/movie-search/Movie/handler"
)

type Endpoints struct {
	SearchEndpoint endpoint.Endpoint
}

func MakeSearchEndpoint(srv handler.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(handler.SearchRequest)
		url := fmt.Sprintf(URL_FORMAT, OMDBAPIKEY, req.MovieName, req.Page)
		res, err := srv.Search(ctx, url, req)
		if err != nil {
			return handler.SearchResponse{nil, err.Error()}, errors.New(err.Error())
		}
		return handler.SearchResponse{res, ""}, nil
	}
}

func (e Endpoints) Search(ctx context.Context) ([]handler.Movie, error) {
	req := handler.SearchRequest{}
	resp, err := e.SearchEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	getResp := resp.(handler.SearchResponse)
	if getResp.ErrorMessage != "" {
		return nil, errors.New(getResp.ErrorMessage)
	}
	return getResp.Response, nil
}
