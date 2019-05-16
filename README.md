## [Unofficial] Hetzner Cloud API Mock

This mock allows developers to test against the Hetzner Cloud API without creating real resources or having a real account on the Hetzner Cloud. 
It uses the API Blueprint from the documentation.

Differences to the first version:
- Rewrite in go (instead of js)
- HTTP header validation (Authorization and Content-Type)
- Body parameter validation (When the docs require a number we validate if the request contains a number)
- Using [drafter (c++ api blueprint parser)](https://github.com/apiaryio/drafter) directly instead of using a wrapper

The mock is a complete rewrite of https://github.com/LKDevelopment/hetzner-cloud-api-mock. 

## Specs:
- Drafter v3.x 
- Go 1.12

## How to use?
To start the mock as a container listing to port 8080 run:
```bash
docker run -d -p 8080:8080 lkdevelopment/hetzner-cloud-api-mock:latest
```

### Notice
This is an unofficial tool, that downloads the information about the lastest stable api from the hetzner cloud and mocks the results shown in the documentation. It is just a really "stupid" mock server, that does not contain any business logic from the Hetzner Cloud.

