[![License: AGPL](https://img.shields.io/badge/license-AGPL-blue.svg)][license]
[![Documentation](https://godoc.org/github.com/jmbarzee/dominion?status.svg)][docs]
[![Code Quality](https://goreportcard.com/badge/github.com/jmbarzee/dominion)][goreport]
[![Build](https://github.com/jmbarzee/dominion/workflows/build/badge.svg)][status]
[![Coverage](https://codecov.io/gh/jmbarzee/dominion/branch/master/graph/badge.svg)][coverage]



# Dominion
Golang Distributed System and Home Automation

This library serves on an IoT network were services (lights, speakers, thermostat, cameras, processing...) will be auto-started, auto-distributed, and (maybe) auto-scaled. The Dominion & Domains are sorts of managers of all devices which are participating. They manage service start, maintance, and discovery.


## Abstractions
### Dominion (leader)
Run the Domain with `go run cmd/dominion/main.go`

Don't forget to set `DOMINION_CONFIG_FILE` [example](../main/cmd/dominion/ex.config.toml)

Listen for new Domains by:
1. Wait for ZeroConf Broadcasts advertizing a Domain
2. Dial Domain & establish lasting connection

Review Domains by:
1. Repeadetly send heartbeat to Domains to keep connection alive
2. Update Domains service list from heartbeat reply

Review Domain's Services by:
1. Routinely reviewing Domains and checking that required services are available/started 
2. Routinely reviewing Services and checking that dependencies are available/started



### Domain (follower)
Run the Domain with `go run cmd/domain/main.go`

Don't forget to set `DOMAIN_CONFIG_FILE` [example](../main/cmd/domain/ex.config.toml)

Domains find a Dominion by:
1. Identifying that they are lonely (no history of a dominion or heartbeats stopped)
2. Broadcasting to the network using ZeroConf

Domains remain in a Dominion by:
1. Listening for incomming Heartbeat RPCs
2. Update stored Dominion's identity 

Domains start Services by:
1. Listening for incomming StartService RPCs
2. Calling make in the specified Service Directory



### Service (Ecosystem) 
Services do whatever you want them too. Services are language agnostic. They can locate other services through the Dominion's GRPC server. Service dependencies are defined in the `DOMINION_CONFIG_FILE`

Services I use -> [ExMachina](github.com/jmbarzee/exmachina)


## Utilized Libraries

`github.com/blang/semver`

`google.golang.org/grpc`

`github.com/grandcat/zeroconf`

`github.com/BurntSushi/toml`



## Planned Development

1. Connection encryption - encrypt RPCs
2. Identity verification - sign communication with preestablished keypairs



[license]: https://opensource.org/licenses/GPL-3.0/
[docs]: (https://godoc.org/github.com/jmbarzee/dominion)

[status]: (https://github.com/jmbarzee/dominion/actions)
[coverage]: (https://codecov.io/gh/jmbarzee/dominion)
[goreport]: https://goreportcard.com/report/github.com/jmbarzee/dominion