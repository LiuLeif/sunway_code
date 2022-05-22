# README #

## deploy from scratch ##

install go 1.18.2, then

```
$> cd <code path>
$> export GO111MODULE=off
$> export GOPATH=$PWD
$> go install hexintek.com/license_server
$> ./bin/license_server

```

## deploy with docker ##

```
$> cd <code path>
$> docker build -t license_server .
$> docker run -it --net=host --rm license_server
```

## demo

visit http://localhost:8080

user: v1@v1.com / v2@v2.com
pass: admin

