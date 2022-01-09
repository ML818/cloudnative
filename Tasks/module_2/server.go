package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {

	testHandler := wrapHandlerWithLogging(http.HandlerFunc(testDemo))
	http.Handle("/test", testHandler)

	healthzHandler := wrapHandlerWithLogging(http.HandlerFunc(healthz))
	http.Handle("/healthz", healthzHandler)

	indexHandler := wrapHandlerWithLogging(http.HandlerFunc(index))
	http.Handle("/", indexHandler)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}

// HANDLERS
// index
func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello, %v", get_client_ip(r))
}

// test handler
func testDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing\n")
}

// healthz
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to healthz, %v\n", get_client_ip(r))
}

// OTHER FUNCTIONS
//get specific environment
func get_specific_env(env string) string {
	e := os.Getenv(env)
	if e == "" {
		return "nil"
	}

	return e
}

// get IP
func get_client_ip(r *http.Request) net.IP {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}

	userIP := net.ParseIP(ip)

	// fmt.Printf("IP: %s, PORT: %s\n", userIP, port)

	return userIP
}

// modifyheader
func modifyHeader(w http.ResponseWriter, r *http.Request, env string) {

	for name, headers := range r.Header {
		for _, h := range headers {
			// fmt.Fprintf(w, "%v: %v\n", name, h)
			w.Header().Add(name, h)
		}
	}

	w.Header().Add(env, get_specific_env(env))

	// fmt.Println(w.Header())
}

// Logging
// Handler mid
func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("--> %s from %s%s", r.Method, get_client_ip(r), r.URL.Path)

		modifyHeader(rw, r, "VERSION")

		lrw := NewLoggingResponseWriter(rw)
		wrappedHandler.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode
		log.Printf("<-- %d : %s", statusCode, http.StatusText(statusCode))

	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
