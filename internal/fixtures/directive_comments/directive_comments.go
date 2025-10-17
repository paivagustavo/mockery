package directive_comments

import "net/http"

// Requester is an interface that defines a method for making HTTP requests.
//
//mockery:generate: true
type Requester interface {
	Get(path string) (string, error)
}

type RequesterWithoutAnnotation interface {
	Get(path string) (string, error)
}

// MatryerRequester is an interface that should be mocked with matryer's template
//
//mockery:generate: true
//mockery:template: matryer
type MatryerRequester interface {
	Get(path string) (string, error)
}

// Server is an interface that defines a method for handling HTTP requests.
//
//mockery:generate: true
//mockery:structname: FunServer
type Server interface {
	HandleRequest(path string, handler http.Handler)
}

// ServerWithDifferentFile is an interface that defines a method for handling HTTP requests.
//
//mockery:generate: true
//mockery:filename: server_with_different_file.go
type ServerWithDifferentFile interface {
	HandleRequest(path string, handler http.Handler)
}
