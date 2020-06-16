# OMDB Movie Searcher

Welcome to my simple application to search movie title from Open Movie Database (OMDB).
This application entirely written in Golang and built upon 2 main services as explained below.

### Api Gateway
First service is the main API Gateway. This API will directly accessed by the client when they want 
to search any movie on any page from OMDB

Currently, we only support 1 API :

Method : Get<br/>
URL : `/search`<br/>
Supported URL Query Parameters : `title` and `page`<br/>
Example : `localhost:8080/search?title=superman&page=1`<br/>
The API then call our next service via gRPC.

You can access the sourcecode from this service in the folder `Gateway`

### GRPC Service
This is our second service to demonstrate how a GRPC Server can be created
and used by our first service. Basically, this service is responsible to make
the API call to OMDB API and return the result to the first service and later
served to the user.

The service is created based on [go-kit](https://github.com/go-kit/kit), which make it
easier since this service is intended to be used as microservice. You can access the sourcecode of this service in the folder `Movie`

### How to run
To run both the service, you can simply open 2 terminals and do `go run cmd/main.go` on both folder `Gateway` and `Movie`. 
By default, the GRPC service will run via port `8081` and API Gateway service will run via port `8080`