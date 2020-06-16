package rest

import (
	"context"
	"encoding/json"
	"github.com/movie-search/Gateway/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type IRoutes interface {
	InitRoutes(router *mux.Router)
	SearchMovie(w http.ResponseWriter, r *http.Request)
}

type Routes struct{}

func InitPackage() *Routes {
	return &Routes{}
}

func (r *Routes) InitRoutes(router *mux.Router) {
	router.HandleFunc("/search", r.SearchMovie).Methods("GET")
}

func (r *Routes) SearchMovie(w http.ResponseWriter, request *http.Request) {
	movieName := request.URL.Query().Get("title")
	if movieName == "" {
		log.Println("title not found")
		w.Write([]byte("title must be provided via QueryParams"))
		return
	}
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		//assume page will be 1 if it's not a valid integer
		page = 1
	}

	grpcConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		log.Println("error when dial grpc. err : ", err.Error())
	}

	defer grpcConn.Close()

	c := pb.NewMovieSearchClient(grpcConn)

	req := &pb.SearchRequest{
		MovieName: movieName,
		Page:      int32(page),
	}

	resp, err := c.SearchMovie(context.Background(), req)
	if err != nil {
		log.Println("error when call grpc. err : ", err.Error())
		w.Write([]byte("error when call grpc. Is GRPC up? See logs for more information or contact admin"))
	} else{
		json.NewEncoder(w).Encode(resp)
	}

}