package engine

import (
	"io"
	"mime/multipart"
	"time"

	"github.com/labstack/gommon/log"
	"net")

type (
	// Engine defines the interface for HTTP server.
	Engine interface {
		// SetHandler sets the handler for the HTTP server.
		SetHandler(Handler)

		// SetLogger sets the logger for the HTTP server.
		SetLogger(*log.Logger)

		// SetListener sets custom listener for the HTTP server.
		SetListener(net.Listener)

		// Start starts the HTTP server.
		Start()
	}

	// Request defines the interface for HTTP request.
	Request interface {
		// TLS returns true if HTTP connection is TLS otherwise false.
		TLS() bool

		// Scheme returns the HTTP protocol scheme, `http` or `https`.
		Scheme() string

		// Host returns HTTP request host. Per RFC 2616, this is either the value of
		// the `Host` header or the host name given in the URL itself.
		Host() string

		// URI returns the unmodified `Request-URI` sent by the client.
		URI() string

		// URL returns `engine.URL`.
		URL() URL

		// Header returns `engine.Header`.
		Header() Header

		// Proto() string
		// ProtoMajor() int
		// ProtoMinor() int

		// UserAgent returns the client's `User-Agent`.
		UserAgent() string

		// RemoteAddress returns the client's network address.
		RemoteAddress() string

		// Method returns the request's HTTP method.
		Method() string

		// SetMethod sets the HTTP method of the request.
		SetMethod(string)

		// Body returns request's body.
		Body() io.Reader

		// FormValue returns form field value for the provided name.
		FormValue(string) string

		// FormFile returns form file for the provided name.
		FormFile(string) (*multipart.FileHeader, error)

		// MultipartForm returns multipart form.
		MultipartForm() (*multipart.Form, error)
	}

	// Response defines the interface for HTTP response.
	Response interface {
		// Header returns `engine.Header`
		Header() Header

		// WriteHeader sends an HTTP response header with status code.
		WriteHeader(int)

		// Write writes the data to the connection as part of an HTTP reply.
		Write(b []byte) (int, error)

		// Status returns the HTTP response status.
		Status() int

		// Size returns the number of bytes written to HTTP response.
		Size() int64

		// Committed returns true if HTTP response header is written, otherwise false.
		Committed() bool

		// Write returns the HTTP response writer.
		Writer() io.Writer

		// SetWriter sets the HTTP response writer.
		SetWriter(io.Writer)
	}

	// Header defines the interface for HTTP header.
	Header interface {
		// Add adds the key, value pair to the header. It appends to any existing values
		// associated with key.
		Add(string, string)

		// Del deletes the values associated with key.
		Del(string)

		// Set sets the header entries associated with key to the single element value.
		// It replaces any existing values associated with key.
		Set(string, string)

		// Get gets the first value associated with the given key. If there are
		// no values associated with the key, Get returns "".
		Get(string) string

		// Keys returns header keys.
		Keys() []string
	}

	// URL defines the interface for HTTP request url.
	URL interface {
		// Path returns the request URL path.
		Path() string

		// SetPath sets the request URL path.
		SetPath(string)

		// QueryValue returns query parameter value for the provided name.
		QueryValue(string) string

		// QueryString returns the URL query string.
		QueryString() string
	}

	// Config defines engine configuration.
	Config struct {
		Address      string        // TCP address to listen on.
		TLSCertfile  string        // TLS certificate file path.
		TLSKeyfile   string        // TLS key file path.
		ReadTimeout  time.Duration // Maximum duration before timing out read of the request.
		WriteTimeout time.Duration // Maximum duration before timing out write of the response.
	}

	// Handler defines an interface to server HTTP requests via `ServeHTTP(Request, Response)`
	// function.
	Handler interface {
		ServeHTTP(Request, Response)
	}

	// HandlerFunc is an adapter to allow the use of `func(Request, Response)` as HTTP handlers.
	HandlerFunc func(Request, Response)
)

// ServeHTTP serves HTTP request.
func (h HandlerFunc) ServeHTTP(req Request, res Response) {
	h(req, res)
}
