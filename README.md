# Go Name Parser

Nameparser API services, python-nameparser implemented by golang.

## Usage

Start NameParser service, `8080` for HTTP client and `8081` for GRPC client.

```bash
docker run --rm -it -p 8080:8080 -p 8081:8081 soulteary/go-nameparser
```

You can use the following commands to test the HTTP API.

```bash
# curl --request POST 'http://127.0.0.1:8080/api/convert' --header 'Content-Type: application/json' --data-raw '{"name": "Dr. Juan Q. Xavier de la Vega III (Doc Vega)"}'

{"text":"Dr. Juan Q. Xavier de la Vega III (Doc Vega)","detail":{"title":"Dr.","first":"Juan","middle":"Q. Xavier","last":"de la Vega","suffix":"III","nickname":"Doc Vega"}}
```

## Build

```bash
docker build -t soulteary/go-nameparser . -f docker/Dockerfile
```
