# Sso Service

This is the Sso service

Generated with

```
micro new sso --namespace=com.hxlb --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.hxlb.srv.sso
- Type: srv
- Alias: sso

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./sso-srv
```

Build a docker image
```
make docker
```