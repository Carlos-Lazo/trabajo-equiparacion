# Trabajo de equiparación - Testing: Test Doubles

Este trabajo consiste en el uso del lenguaje Go para explicar de manera sencilla y breve los tests doubles. Aquí se implementan **mock**, **fake**, **dummy** y **stub**.

## Requisitos

+ #### [Go](https://go.dev/doc/install)
+ #### Paquete testify
```bash
      go get github.com/stretchr/testify 
```
### Utils
- #### Ejecutar un proyecto de go desde cero - Lo cual no es netamente necesario, es solo a título informativo.
```bash
    go mod init <nombre/del/modulo>
    go mod tidy
    go run main.go
```
1. Genera los archivos *go.mod* y *go.sum*.
2. Actualiza, verifica y borra las dependencias según corresponda.
3. Ejecuta el proyecto sin generar un compilado.
## Estructura
```bash
├── directorio
│   ├── db.go
│   ├── motor_dummy_test.go
│   ├── motor_fake_test.go
│   ├── motor.go
│   ├── motor_mock_test.go
│   └── motor_stub_test.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```
## Uso
- #### Una vez que tenga los requisitos previos instalados, puede ejecutar las pruebas unitarias ejecutando el siguiente comando en la carpeta del proyecto:
```Go
# Desde la raiz
go test ./... -v
```
+ Para más información sobre los test

```Go
go help test 
```
# trabajo-equiparacion
