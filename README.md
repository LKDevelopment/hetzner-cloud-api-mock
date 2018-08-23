# [Unofficial] [Beta] Hetzner Cloud API Mock Server

I want to make the development of tools or integrations for the Hetzner Cloud as easy as possible.
Sometimes people need to test their software against the API, but do not want to have any costs. 
So i created this little mock server.

It is proposed as a read only replacement for the Hetzner Cloud API in your tests.

It is quite easy! Just run:

```bash
docker run -p 8080:8080 -it lkdevelopment/hetzner-cloud-api-mock:latest
```

`-p 8080:8080` maps the internal port of the docker image (8080) to your machine's port (8080). `-it` makes the session interactive so that you can stop it with `CTRL` + `C`.

Then, it runs a basic mock server that is available on Port 8080 by default. So just browse to http://localhost:8080/servers and you get a testable result. All methods that are available in the official documentation are useable on this mock.

### Notice
This is an unofficial tool, that downloads the information about the lastest stable api from the hetzner cloud and mocks the results shown in the documentation. It is just a really "stupid" mock server, that does not contain any business logic from the Hetzner Cloud.
