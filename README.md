## Layer 4 Proxy
[![CI](https://github.com/bilalcaliskan/layer4-proxy/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/layer4-proxy/actions?query=workflow%3ACI)

Simple Layer 4 proxy written Golang. 

### Installation
#### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/layer4-proxy/releases) page.

#### Source
Source installation method requires [Golang 1.16](https://golang.org/doc/go1.16)
```shell
$ go get github.com/bilalcaliskan/layer4-proxy
```
Then with the help of the [Makefile](Makefile), can be run easily with below command:
```shell
$ make run
```

#### Docker
Docker image can be downloaded with below command:
```shell
$ docker run bilalcaliskan/layer4-proxy:latest
```

### Usage
Layer4-proxy can be customized with several command line arguments:
```shell
$ layer4-proxy --help
Usage of ./bin/main:
      --proxyPort int        Provide a port to run proxy server on (default 3000)
      --proxyProto string    Provide a proxy server protocol (default "tcp")
      --targetDns string     Provide a target DNS to proxy (default "en.wikipedia.org")
      --targetPort int       Provide a target port to proxy (default 443)
      --targetProto string   Provide a target protocol to proxy (default "tcp")
```

### Example
To test layer4-proxy, all you need is curl command:
```shell
$ layer4-proxy --targetDns en.wikipedia.org --targetPort=443 --proxyPort=3000 &
$ curl -v -k https://localhost:3000/wiki/OSI_model -H "Host: en.wikipedia.org"
```