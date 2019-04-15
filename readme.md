Hello server is a sample http server for dev and test.

## Usage

This app is usually running as docker container. Docker repository is below.
https://cloud.docker.com/repository/docker/shin1x1/hello

```
$ docker run --rm -p 8000:8000 shin1x1/hello 
```

```
$ curl -v http://localhost:8000/
> GET / HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Mon, 15 Apr 2019 06:58:28 GMT
< Content-Length: 8
<
"Hello"

$ curl -v http://localhost:8000/health
> GET /health HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Mon, 15 Apr 2019 06:59:01 GMT
< Content-Length: 2
<
ok

$ curl -v http://localhost:8000/notfound
> GET /notfound HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 404 Not Found
< Content-Type: application/json; charset=UTF-8
< Date: Mon, 15 Apr 2019 06:59:27 GMT
< Content-Length: 24
<
{"message":"Not Found"}
```
