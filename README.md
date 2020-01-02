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

TBD

### Sample log collection server (rlserver) usage

**rlserver** application provided by this project is a ready-to-use simple log collecting server compatible with [rltester](#sample-log-sending-client-(rltester)-usage) and client side package. Building instructions are availble at the beginning of [Install and usage section](#installation-and-usage).

**rlserver** is listening on default (*9999*) or user defined UDP port and saves all received log messages in specified output file (*default.log* by default). Default values can be modified using following command line params:

* *-d port* to set UDP listening port
* *-o path* to set output file for received logs

As usual *-h* option will print help message.

### Sample log sending client (rltester) usage

If **rltester** is build (see beginning of [Install and usage section](#installation-and-usage)) it can be used to check working server instance.

When launched without any params **rltester** assumes server is running at *localhost:9999* and client's name is set to *TesterApp*.

To change default values command line params can be used:

* *-s servername:port* changes default RemLog server location
* *-n somename* changes default client's name sent in registration request

## License

TBD (MIT or Apache 2-clause)

## Contact

If you need to contact me feel free to write me an email:  
[markamdev.84#dontwantSPAM#gmail.com](maitlo:)
