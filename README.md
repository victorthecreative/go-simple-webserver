# Simple web server in go

This project is an introducing to the net/http package, is simple little file serve with html static website.

## Routes

```mermaid
graph TD
    Server --> hello["/hello"] --> helloFunction["hello function"]
    Server --> form["/form"] --> formFunction["form function"] --> formHtml["form.html"]
    Server --> root["/"] --> indexHtml["index.html"]
```

to execute it you must have [**golang**](https://golang.org/)

run

```
go build
./go-simple-webserver
```
