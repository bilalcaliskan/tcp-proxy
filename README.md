## TCP Proxy
[![CI](https://github.com/bilalcaliskan/tcp-proxy/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/tcp-proxy/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/tcp-proxy)](https://hub.docker.com/r/bilalcaliskan/tcp-proxy/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/tcp-proxy)](https://goreportcard.com/report/github.com/bilalcaliskan/tcp-proxy)

Simple TCP proxy written Golang. 

### Installation
#### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/tcp-proxy/releases) page.

#### Source
Source installation method requires [Golang 1.16](https://golang.org/doc/go1.16)
```shell
$ go get github.com/bilalcaliskan/tcp-proxy
```
Then with the help of the [Makefile](Makefile), can be run easily with below command:
```shell
$ make run
```

#### Docker
Docker image can be downloaded with below command:
```shell
$ docker run bilalcaliskan/tcp-proxy:latest
```

### Usage
tcp-proxy can be customized with several command line arguments:
```shell
$ tcp-proxy --help
Usage of ./bin/main:
      --proxyPort int        Provide a port to run proxy server on (default 3000)
      --proxyProto string    Provide a proxy server protocol (default "tcp")
      --targetDns string     Provide a target DNS to proxy (default "en.wikipedia.org")
      --targetPort int       Provide a target port to proxy (default 443)
      --targetProto string   Provide a target protocol to proxy (default "tcp")
```

### Example
To test tcp-proxy, all you need is curl command:
```shell
$ tcp-proxy --targetDns en.wikipedia.org --targetPort=443 --proxyPort=3000 &
$ curl -v -k https://localhost:3000/wiki/OSI_model -H "Host: en.wikipedia.org"
```
