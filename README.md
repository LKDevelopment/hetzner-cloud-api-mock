# [Unofficial] [Beta] Hetzner Cloud API Mock Server

I want to make the development of tools or integrations for the Hetzner Cloud as easy as possible.
Sometimes people needs to test there software against the API, but do not want to have any costs. 
So i created this little mock server.

It is proposed as a read only replacement for the Hetzner Cloud API in your tests.

It is quite easy! Just run:

```bash
docker run lkdevelopment/hetzner-cloud-api-mock:latest
```

Then, it runs a basic mock server that is out of the box available on Port 8080. So just browse to http://localhost:8080/servers and you a testable result. All methods that are available in the official documentation are useable on this mock.

### Notice
This is an unofficial tool, that downloads the information about the lastest stable api from the hetzner cloud and mock the results shown in the documentation. It is just really "stupid" mock server, that does not contain any business logic from the Hetzner Cloud.