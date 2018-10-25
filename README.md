# demoapp-pcf-healthcheck

App to Demo PCF Healthcheck

Simple Golang app that exposes `/` and `/healthz` endpoints on port 8080.

```sh
make build
``` 

to build locally. 

`curl localhost:8080/healthz` to test