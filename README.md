# GoBases
Teoría y práctica de GoBases ML

Requiere Go 1.27.3 o posterior. Con `GOTOOLCHAIN=auto`, Go descarga la
versión mínima indicada en `go.mod` cuando es necesario.

Para verificar el proyecto:

```sh
go test ./...
go run golang.org/x/vuln/cmd/govulncheck@v1.9.0 ./...
```


# Notaciones

Para poder utilizar paquetes tendremos que inicializar el proyecto con 
- go mod
Y luego traernos los paquetes con go get "nombre del repositorio" (por lo general github/...)

Para poder quitar paquetes que ya no querramos utilizar
go get "nombre del repositorio"@none
