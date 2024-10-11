# go-live-reloading

Golang live-reloading using Docker and Air

## Prerequisite

With local go and air next files has to be created

- go.mod
- go.sum
- .air.toml

## Using

- To run (and dev with live reload) golang server use **docker compose up**
- To implement migration use **make db/migration/up**

## Service architecture

![Service architecture](/docs/pics/Service.png 'Service')
