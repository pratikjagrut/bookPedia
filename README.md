# bookPedia
This is a sample go application demonstrating postgres database and rest api services.

## Quick start

### Prerequisites
- [go](https://golang.org/dl/) version v1.10+
- go get gopkg.in/src-d/go-git.v4

### Steps to run

#### Setup postgres DB
`Using docker for easy installation and standalone env`

```
$ docker pull postgres
$ docker run --rm  --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres
```
Get the app and run
```
$ mkdir $GOPATH/src/github.com/pratikjagrut
$ cd $GOPATH/src/github.com/pratikjagrut
$ git clone https://github.com/pratikjagrut/bookPedia.git
$ cd bookPedia
$ dep ensure -v
$ go run main.go
```
Note: I'm using `httpie` but you can use `curl` or any other clinet you wish.

Output:
Insert a book:
```
$ http POST localhost:8080/book/add < books.json
HTTP/1.1 200 OK
Content-Length: 34
Content-Type: text/plain; charset=utf-8
Date: Mon, 10 Jun 2019 06:25:24 GMT

{
    "ID": 1234,
    "Name": "Go in Action"
}
```
Get all books in DB: 
```
$ http GET localhost:8080/books
HTTP/1.1 200 OK
Content-Length: 36
Content-Type: text/plain; charset=utf-8
Date: Mon, 10 Jun 2019 06:33:45 GMT

[
    {
        "ID": 1234,
        "Name": "Go in Action"
    }
]
```
Get perticular book by id: 
```
$ http GET localhost:8080/book/1234
HTTP/1.1 200 OK
Content-Length: 34
Content-Type: text/plain; charset=utf-8
Date: Mon, 10 Jun 2019 06:35:35 GMT

{
    "ID": 1234,
    "Name": "Go in Action"
}
```
Delete perticular book by id: 
```
$ http DELETE localhost:8080/book/1234
HTTP/1.1 200 OK
Content-Length: 22
Content-Type: text/plain; charset=utf-8
Date: Mon, 10 Jun 2019 06:44:20 GMT

{
    "ID": 1234,
    "Name": ""
}
```
