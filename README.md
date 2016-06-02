dumprequest
-----------

# .

dump HTTP/1.x wire representation

- see [https://golang.org/pkg/net/http/httputil/#DumpRequest]

# Usage

## get and run 

```
git clone https://github.com/heshed/dumprequest.git
cd dumprequest

go run main.go
http://127.0.0.1:47767
```


## send request and dump

### curl 

```
curl -X POST -H "Content-Type: application/json" -d '{"lable":"dumprequest"}' "http://127.0.0.1:47767/graphs/one/two/three" -v
* About to connect() to 127.0.0.1 port 47767 (#0)
*   Trying 127.0.0.1... connected
* Connected to 127.0.0.1 (127.0.0.1) port 47767 (#0)
> POST /graphs/one/two/three HTTP/1.1
> User-Agent: curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.15.3 zlib/1.2.3 libidn/1.18 libssh2/1.4.2
> Host: 127.0.0.1:47767
> Accept: */*
> Content-Type: application/json
> Content-Length: 23
>
< HTTP/1.1 200 OK
< Date: Wed, 01 Jun 2016 15:14:46 GMT
< Content-Length: 262
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 127.0.0.1 left intact
* Closing connection #0
"POST /graphs/one/two/three HTTP/1.1\r\nHost: 127.0.0.1:47767\r\nAccept: */*\r\nContent-Type: application/json\r\nUser-Agent: curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.15.3 zlib/1.2.3 libidn/1.18 libssh2/1.4.2\r\n\r\n{\"lable\":\"dumprequest\"}"
```

### dumprequest


```
POST /graphs/one/two/three HTTP/1.1
Host: 127.0.0.1:47767
Accept: */*
Content-Type: application/json
User-Agent: curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.15.3 zlib/1.2.3 libidn/1.18 libssh2/1.4.2

{"lable":"dumprequest"}
```

