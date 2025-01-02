# Simple web server in go

This project is an introducing to the net/http package, is simple little file serve with html static website.

## Routes

```mermaid
A["Server"] --> B["/hello"] --> C["hello function"]
A --> D["/form"] --> E["form function"] --> F["form.html"]
A --> G["/"] --> H["index.html"]
```

to execute it you must have [**golang**](https://golang.org/)

run

```
go build
./go-simple-webserver
```
