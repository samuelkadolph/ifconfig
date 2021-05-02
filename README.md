## samuelkadolph/ifconfig

[![Build Status](https://img.shields.io/travis/com/samuelkadolph/ifconfig.svg?style=for-the-badge)](https://travis-ci.com/samuelkadolph/ifconfig/ "Build Status")
[![Docker Pulls](https://img.shields.io/docker/pulls/samuelkadolph/ifconfig.svg?style=for-the-badge)](https://hub.docker.com/r/samuelkadolph/ifconfig/ "Docker Pulls")
[![Docker Stars](https://img.shields.io/docker/stars/samuelkadolph/ifconfig.svg?style=for-the-badge)](https://hub.docker.com/r/samuelkadolph/ifconfig/ "Docker Stars")
[![MIT License](https://img.shields.io/github/license/samuelkadolph/ifconfig.svg?style=for-the-badge)](https://github.com/samuelkadolph/ifconfig/blob/master/LICENSE "MIT License")

ifconfig is a simple Go HTTP app that returns the IP of the client. Used to run [ifconfig.ca](https://ifconfig.ca) and
[ifconfig.so](https://ifconfig.so) for a no-nonsense simple use in scripting.

### Usage

#### Go

```
go get github.com/samuelkadolph/ifconfig
go build github.com/samuelkadolph/ifconfig
$GOPATH/bin/ifconfig
```

#### Docker

```
docker run --detach --name=ifconfig --restart always --publish 7570:7570 samuelkadolph/ifconfig
```

### Parameters

These parameters are for running the app directly. There are no parameters for the docker image.

| Parameter | Function |
| :----: | --- |
| `-listen ip:port` | Listens on the specified ip and port. Can be called multiple times |
