# Go Name Parser

Nameparser API services, [python-nameparser](https://github.com/derek73/python-nameparser) implemented by Golang.

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

You can use the following commands to test the GRPC API.

```bash
cd example/grpc-client
go run main.go
```

## Tutorial

- [Using Golang and Docker to implement Python computing services](https://soulteary.com/2023/05/22/using-golang-and-docker-to-implement-python-computing-services.html)

## Dev

```bash
docker build -t soulteary/go-nameparser . -f docker/Dockerfile
```

## Credits

- [python-nameparser](https://github.com/derek73/python-nameparser), the main calculation function implementation.
- [Docker Python in Go](https://github.com/soulteary/docker-python-in-go), the principle.

## License

Keep the same with the [python-nameparser](https://github.com/derek73/python-nameparser), use [LGPL-2.1 license](https://github.com/soulteary/go-nameparser/blob/main/LICENSE).
