# serverless-url-shortener [![Build Status](https://travis-ci.org/michaelmcallister/serverless-url-shortener.svg?branch=master)](https://travis-ci.org/michaelmcallister/serverless-url-shortener) [![Coverage Status](https://coveralls.io/repos/github/michaelmcallister/serverless-url-shortener/badge.svg?branch=master)](https://coveralls.io/github/michaelmcallister/serverless-url-shortener?branch=master) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

a URL shortener API written in Golang built for AWS Lambda and AWS S3

![](https://sknk.ws/assets/serverless-url-shortener.gif)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

```
TODO
```

### Installing
```
TODO
```

## Running the tests

Tests are provided by Golangs testing framework and can be run with `go test`

```
☁  serverless-url-shortener [master] ⚡
go test
PASS
ok      _~/git/serverless-url-shortener    0.004s
☁  serverless-url-shortener [master] ⚡
```
Further information can be found on [Golang.org](https://golang.org/pkg/testing/)

## Built With

* [AWS S3](https://aws.amazon.com/s3/) - _The_ simple storage service
* [AWS Lambda](https://aws.amazon.com/lambda/) - Run code without thinking about the servers.
* [AWS API Gateway](https://aws.amazon.com/api-gateway/) - Proxy RESTful requests directly to the Lambda function.
* [eawsy/aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim) - Enables me to write my code in Go!


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
