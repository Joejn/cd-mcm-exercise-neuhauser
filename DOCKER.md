# Aufgabe 3

Das Dockerfile besteht aus zwei Stages: Builder und Runtime. Der Vorteil eines Multistage-Dockerfiles besteht darin, dass nur die letzte Stage in das Docker-Image übernommen wird, wodurch die Größe des Images reduziert werden kann.

## Builder Stage

Die Build-Stage ist dafür verantwortlich, dass die Anwendung gebaut wird.

```Dockerfile
# Verwendet ein Golang-Image
FROM golang:1.26-alpine AS builder

# Setzt das Verzeichnis, in dem RUN-, CMD- etc. Befehle ausgeführt werden
WORKDIR /app

# Kopiert Dateien in das Image
COPY go.mod go.sum ./

# Lädt Abhängigkeiten herunter
RUN go mod download

# Kopiert den gesamten Ordner im Kontext in das Image
COPY . .

# Baut die Anwendung
# CGO_ENABLED=0 -> Verwendet rein statische Binärdateien, es werden also keine Bibliotheken zur Laufzeit geladen
#                   Wird benötigt, damit die Anwendung portabel ist
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-server ./cmd/api
```

## Runtime Stage

Die Runtime-Stage kopiert die gebaute Version aus der Build-Stage und startet diese.

```Dockerfile
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /api-server .

EXPOSE 8080

ENTRYPOINT ["./api-server"]
```

## Dateigrößen

| Image       | Größe   |
| ----------- | ------- |
| build-stage | 509 MB  |
| full        | 29,3 MB |

# Aufgabe 4

## CRUD Tests

### Produkte erstellen

```bash
root@docker-desktop:/# curl -X POST http://localhost:8080/products -d '{"name":"Saft","price":1.10}'
{"id":1,"name":"Saft","price":1.1}
```

```bash
root@docker-desktop:/# curl -X POST http://localhost:8080/products -d '{"name":"Reis","price":0.99}'
{"id":2,"name":"Reis","price":0.99}
```

```bash
root@docker-desktop:/# curl -X POST http://localhost:8080/products -d '{"name":"Butter","price":2.40}'
{"id":3,"name":"Butter","price":2.4}
```

### Produkte abfragen

```bash
root@docker-desktop:/# curl http://localhost:8080/products
[{"id":1,"name":"Saft","price":1.1},{"id":2,"name":"Reis","price":0.99},{"id":3,"name":"Butter","price":2.4}]
```

```bash
root@docker-desktop:/# curl http://localhost:8080/products/1
{"id":1,"name":"Saft","price":1.1}
```

### Produkte updaten

```bash
root@docker-desktop:/# curl -X PUT http://localhost:8080/products/1 -d '{"name":"Bier","price":1.40}'
{"id":1,"name":"Bier","price":1.4}
root@docker-desktop:/# curl http://localhost:8080/products/1
{"id":1,"name":"Bier","price":1.4}

```

### Produkte lösche

```bash
root@docker-desktop:/# curl -X DELETE http://localhost:8080/products/1
{"result":"success"}
root@docker-desktop:/# curl http://localhost:8080/products/1
{"error":"Product not found"}
```

## Docker Container neu starten

Produkte vor dem Neustart

```bash
root@docker-desktop:/# curl http://localhost:8080/products
[{"id":2,"name":"Reis","price":0.99},{"id":3,"name":"Butter","price":2.4}]
```

Produkte nach dem Neustart

```bash
root@docker-desktop:/# curl http://localhost:8080/products
[{"id":2,"name":"Reis","price":0.99},{"id":3,"name":"Butter","price":2.4}]
```

## Tests

```
PS D:\github\cd-mcm-exercise-neuhauser> go test -v ./internal/handler/
=== RUN   TestHealthEndpoint
--- PASS: TestHealthEndpoint (0.00s)
=== RUN   TestGetProductsEmpty
--- PASS: TestGetProductsEmpty (0.00s)
=== RUN   TestCreateAndGetProduct
--- PASS: TestCreateAndGetProduct (0.00s)
=== RUN   TestGetProductNotFound
--- PASS: TestGetProductNotFound (0.00s)
=== RUN   TestUpdateProduct
--- PASS: TestUpdateProduct (0.00s)
=== RUN   TestDeleteProduct
--- PASS: TestDeleteProduct (0.00s)
=== RUN   TestCreateInvalidProduct
--- PASS: TestCreateInvalidProduct (0.00s)
PASS
ok      github.com/mrckurz/CI-CD-MCM/internal/handler   0.587s
```
