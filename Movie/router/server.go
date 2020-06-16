package router

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/movie-search/Movie/handler"
	"github.com/movie-search/Movie/pb"
)

type grpcServer struct {
	movie grpctransport.Handler
}

func (s *grpcServer) SearchMovie(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	_, resp, err := s.movie.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SearchResponse), nil
}

func NewGRPCServer(_ context.Context, endpoint Endpoints) pb.MovieSearchServer {
	return &grpcServer{
		movie: grpctransport.NewServer(
			endpoint.SearchEndpoint,
			handler.DecodeGRPCSearchRequest,
			handler.EncodeGRPCSearchResponse,
		),
	}
}
