package interface_annotations

import "net/http"

// Requester is an interface that defines a method for making HTTP requests.
//
// mockery_generate: true
type Requester interface {
	Get(path string) (string, error)
}

type RequesterWithoutAnnotation interface {
	Get(path string) (string, error)
}

// MatryerRequester is an interface that should be mocked with matryer's template
//
// mockery_generate: true
// mockery_template: matryer
type MatryerRequester interface {
	Get(path string) (string, error)
}

// Server is an interface that defines a method for handling HTTP requests.
//
// mockery_generate: true
// mockery_struct_name: FunServer
type Server interface {
	HandleRequest(path string, handler http.Handler)
}

// ServerWithDifferentFile is an interface that defines a method for handling HTTP requests.
//
// mockery_generate: true
// mockery_file_name: server_with_different_file.go
type ServerWithDifferentFile interface {
	HandleRequest(path string, handler http.Handler)
}
