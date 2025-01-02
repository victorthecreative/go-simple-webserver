# Simple web server in go

This project is an introducing to the net/http package, is simple little file serve with html static website.

## Routes

```mermaid
graph TD
    Server --> / --> index.html
    Server --> /hello --> hello func
    Server --> /form --> form func --> form.html
```

to execute it you must have [**golang**](https://golang.org/)

run

```
go build
./go-simple-webserver
```
