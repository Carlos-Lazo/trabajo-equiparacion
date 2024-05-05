# Trabajo de equiparación - Testing: Test Doubles

```html
Este trabajo consiste en el uso del lenguaje Go para explicar de manera sencilla y breve los tipos de tests doubles. Aquí se implementaran **mock**, **fake**, **dummy** y **stub**.
```
## Logica de negocio
```html
El proyecto contiene un dominio el cual es llamado 'motor'. El cual intenta replicar de manera parcial el funcionamiento de una interfaz de db. El dominio cuenta con metodos como 'BuscarPorNombre()', 'AgregarEntrada()', los cuales serán objetos de test.
Se intenta, a su manera, debido al ejemplo, replicar al patron Repository
```
## Sobre los test
```html
    ª En principio, la carpeta test contiene una carpeta homonima
    al dominio testeado. Dentro de cada carpeta de dominio
    abra una seccion para cada tipo de double.
    ª La capa sobre la cual se desarrollaran los test es service.
    º Los test tiene un desarrollo en comentarios.
    º Cada uno de los test esta definido en dos archivos;
        * _test.go # Donde se encuentran las funciones de test.
        * _<tipoTestDouble>.go # Donde se encuentra la definicion del tipo de mock.
```
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
. # Raiz
├── go.mod
├── go.sum
├── internal 
│   ├── domain
│   └── motor 
│       ├── repository.go # Definicion de la capa repository
│       ├── service.go # Definicion de la capa service
│       └── storage.go # Archivo que usaer para emular un almacen de datos.
├── main.go
├── mocks
│   └── motor # Aqui dejare un test por cada seccion
│       ├── dummy
│       │   └── service_dummy_test.go
│       ├── fake
│       │   ├── motor_fake.go # Definicion del tipo de Double
│       │   └── service_fake_test.go # Ejemplos de tests.
│       ├── mock
│       │   ├── motor_mock.go
│       │   └── service_mock_test.go
│       └── stub
│           ├── motor_stub.go
│           └── service_stub_test.go
└── README.md

10 directories, 14 files
```
## Arquitectura
#### Solo con fines prácticos, omitiremos la capa del controlador y, por supuesto, el almacen de datos.
![arquitectura del patron repository](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse3.mm.bing.net%2Fth%3Fid%3DOIP.mssrvZrbz8nFEPralCTO6AHaBW%26pid%3DApi&f=1&ipt=acddd3c3720b72912d5897d4f09682a7bb1c06a856bfe2be6a7456dc2ce87d0f&ipo=images)
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
