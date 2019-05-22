### Base de datos

```sql
 - CREATE DATABASE ptruoradb;
 - use ptruoradb;
 - CREATE TABLE ptruoradb.tbl_domain (name STRING,server_change BOOL,ssl_grade  STRING(200),previus_ssl_grade STRING(200),logo STRING(200),title STRING(200),is_down BOOL);
 - CREATE TABLE ptruoradb.tbl_server (id_domain STRING, address STRING, ssl_grade STRING, country STRING, owner STRING); 
- CREATE TABLE ptruoradb.tbl_traces (id_domain STRING, datetime TIMESTAMP);
```