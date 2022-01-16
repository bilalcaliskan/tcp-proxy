# TCP Proxy
[![CI](https://github.com/bilalcaliskan/tcp-proxy/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/tcp-proxy/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/tcp-proxy)](https://hub.docker.com/r/bilalcaliskan/tcp-proxy/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/tcp-proxy)](https://goreportcard.com/report/github.com/bilalcaliskan/tcp-proxy)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_tcp-proxy&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_tcp-proxy)
[![codecov](https://codecov.io/gh/bilalcaliskan/tcp-proxy/branch/master/graph/badge.svg)](https://codecov.io/gh/bilalcaliskan/tcp-proxy)
[![Release](https://img.shields.io/github/release/bilalcaliskan/tcp-proxy.svg)](https://github.com/bilalcaliskan/tcp-proxy/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/tcp-proxy)](https://github.com/bilalcaliskan/tcp-proxy)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Simple TCP proxy written with Golang using built-in [net package](https://pkg.go.dev/net).

## Configuration
tcp-proxy can be customized with several command line arguments:
```
--proxyPort         int         Provide a port to run proxy server on, defaults to 3000
--proxyProto        string      Provide a proxy server protocol, defaults to tcp
--targetDns         string      Provide a target DNS to proxy, defaults to en.wikipedia.org
--targetPort        int         Provide a target port to proxy, defaults to 443
--targetProto       string      Provide a target protocol to proxy, defaults to tcp
```

## Installation
### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/tcp-proxy/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ ./tcp-proxy --targetDns en.wikipedia.org --targetPort=443 --proxyPort=3000
```

### Source
Currently, source installation method requires [Golang 1.16](https://golang.org/doc/go1.16)
```shell
$ go get github.com/bilalcaliskan/tcp-proxy
```
Then with the help of the [Makefile](Makefile), can be run easily with below command:
```shell
$ make run
```

### Docker
You can simply run docker image with default configuration:
```shell
$ docker run bilalcaliskan/tcp-proxy:latest
```

## Example Usage
To test tcp-proxy, you can run the binary with provided arguments:
```shell
$ tcp-proxy --targetDns en.wikipedia.org --targetPort=443 --proxyPort=3000 &
```

And then test it with curl command:
```shell
$ curl -v -k https://localhost:3000/wiki/OSI_model -H "Host: en.wikipedia.org"
```

## Development
This project requires below tools while developing:
- [Golang 1.17](https://golang.org/doc/go1.17)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)

## License
Apache License 2.0
