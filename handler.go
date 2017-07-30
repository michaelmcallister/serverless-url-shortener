package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
)

const invalidBucketError = "S3_BUCKET environment variable is not set"
const invalidURLError = "'%s' is not a valid URL."
const defaultProtocol = "http://"

var Handle apigatewayproxy.Handler
var S3Handle *s3.S3

func init() {
	ln := net.Listen()
	Handle = apigatewayproxy.New(ln, nil).Handle
	go http.Serve(ln, http.HandlerFunc(handle))
	S3Handle = initS3()
}

func handle(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	hash := getHash(longURL)
	resp := Payload{LongURL: longURL}
	bucketS3, ok := os.LookupEnv("S3_BUCKET")

	if !ok {
		resp.Header = http.StatusBadRequest
		resp.Message = invalidBucketError
		returnResponse(&resp, w, r)
		return
	}

	if !isValidProtocol(longURL) {
		//default to HTTP protocol.
		longURL = defaultProtocol + longURL
	}

	if _, err := url.ParseRequestURI(longURL); err != nil {
		resp.Header = http.StatusBadRequest
		resp.Message = fmt.Sprintf(invalidURLError, longURL)
		returnResponse(&resp, w, r)
		return
	}

	redirResponse, err := createRedirObject(S3Handle, longURL, bucketS3, hash)

	if err != nil {
		resp.Header = http.StatusInternalServerError
		resp.Message = err.Error()
		returnResponse(&resp, w, r)
		return
	}
	resp.Header = http.StatusCreated
	resp.ShortURL = redirResponse.ShortURL
	resp.Message = fmt.Sprintf("URL '%s' created for '%s'", resp.ShortURL, resp.LongURL)
	returnResponse(&resp, w, r)
	return
}
