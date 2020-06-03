# RemLog - remote logging service

This project provides Go package for logs sending (client side) and standalone application for logs collection (server side).

## Introduction

The purpose of this project is to provide simple way of remote logs collection. The idea is to have multiple applications, potentialy launched on different hosts, and one single point of logs saving.

Sample usage scenario can be group of sensor nodes (ex. set of RPi Zero W boards) and a laptop as a monitoring device shown on picture below.

![Diagram of RPi nodes, wireles router and laptop](./data/RemLog-network.png)

Provided Go sources contains client side package, server side package (for custom server implementation) and sample server implementation for out-of-a-box log collection. All of that allow to set up environment as on picture below.

![Diagram of package usage](./data/RemLog-usage.png)

## Installation and usage

There's no sophisticated installation procedure needed to use RemLog packages. If provided standalone applications will not be used there's no need to implicitly download repository and build sources.

When one wants to use provided server or tester application (simple dummy client) then it enough just to download repository:

```bash
git clone github.com/markamdev/remlog
```

and build applications:

```bash
cd remlog
make
```

Compiled binaries will be placed in *./build* directory.

Since v0.2.0 release there are additional Makefile targets allowing simple cross compilation for selected platforms:

* *amd64* for x64 architecture (no OS forced so host used)
* *intel32* for x86 architecture (no OS forced so host used)
* *pizero* for ARMv6 architecture (Linux set for OS), compatible with Raspberri Pi Zero board
* *pi3* for ARMv7 architecture (Linux set for OS), compatible with Raspberry Pi 3 board (and probably other boards)

These targets are prepared for simpler testing on multiple devices/boards.

### Client side and server side packages usage

This project is splitted into two parts that can be used in third party applications: client and server. With use of these parts one can easily send logs from application or implement it's own log collecting server.

#### Client side package

To be able to send log messages to server application has to import client package:

```go
import (
    "github.com/markamdev/remlog/client"
)
```

Before any message sending is possible *remlog* client has to be initialized and registered on server:

```go
cnf := client.RLCconfig{Server: "10.0.0.1:9999", Name: "clientName"}
if err := client.Init(&cnf); err != nil {
    fmt.Println("Failed to initialize client: ", err.Error())
    return
}

if err := client.Register(); err != nil {
    fmt.Println("Failed to register client: ", err.Error())
    return
}
```

In code above *client.RLConfig* is a structure that contains client's configuration: address (with port) of log collection server and name (some string identifier) of this client instance.

It is highly recommended to not use application binary name as a client's name. It can make difficult to analyze logs if more than one node will be running same application sending log to same server. Any string containing hostname and/or ip address would be more useful.

#### Server side package

### Sample log collection server (rlserver) usage

**rlserver** application provided by this project is a ready-to-use simple log collecting server compatible with [rltester](#sample-log-sending-client-(rltester)-usage) and client side package. Building instructions are availble at the beginning of [Install and usage section](#installation-and-usage).

**rlserver** is listening on default (*9999*) or user defined UDP port and saves all received log messages in specified output file (*default.log* by default). Default values can be modified using following command line params:

* *-d port* to set UDP listening port
* *-o path* to set output file for received logs

As usual *-h* option will print help message.

### Sample log sending client (rltester) usage

If **rltester** is build (see beginning of [Install and usage section](#installation-and-usage)) it can be used to check working server instance.

When launched without any params **rltester** assumes server is running at *localhost:9999* and client's name is set to *TesterApp*. By default one set of log messages is sent to server (with 200ms interval)

To change default values command line params can be used:

* *-s servername:port* changes default RemLog server location
* *-n somename* changes default client's name sent in registration request
* *-l N* sets number of log packets sent to server (each packet contains one log message of each severity level)

As usual *-h* option will print help message.

## License

Code is published under [MIT License](https://opensource.org/licenses/MIT) as it seems to be the most permissive license. If for some reason you need to have this code published with other license (ex. to reuse the code in your project) please contact [author](#author-/-contact) directly.

## Author / Contact

If you need to contact me feel free to write me an email:  
[markamdev.84#dontwantSPAM#gmail.com](maitlo:)
