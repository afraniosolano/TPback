# Prueba Tru

Api rest desarrollada en go version go1.12.4 windows/amd64 y con base de datos cockroach-v19.1.0

# Antes de ejecutar se deben instalar las siguientes librerias:

  - go get -u github.com/gorilla/mux
  - go get gopkg.in/mgo.v2
  - go get gopkg.in/mgo.v2/bson
  - go get github.com/heppu/simple-cors
  - go get -u github.com/lib/pq


# Ejecutar la aplicación:
  - go run .\main.go

# Codigo base generado desde:
  - https://editor.swagger.io/
La definicion se encontrara en la carpeta api (swagger.yaml)


### Base de datos

```sql
 - CREATE DATABASE ptruoradb;
 - use ptruoradb;
 - CREATE TABLE ptruoradb.tbl_domain (name STRING,server_change BOOL,ssl_grade  STRING(200),previus_ssl_grade STRING(200),logo STRING(200),title STRING(200),is_down BOOL);
 - CREATE TABLE ptruoradb.tbl_server (id_domain STRING, address STRING, ssl_grade STRING, country STRING, owner STRING); 
- CREATE TABLE ptruoradb.tbl_traces (id_domain STRING, datetime TIMESTAMP);
```