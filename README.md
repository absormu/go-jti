# go-jti

IMPLEMENTASI
------------

- [x] Gmail (Sign In With Google)
- [x] Metode ENCRYPT/DECRYPT
- [ ] WEB SOCKET
- [ ] POSTGRE SQL 

INTRODUCTION
------------
* [LINK GITHUB](https://www.github.com/absormu/go-jti)
* DOCUMENTATION API postman collection (folder /docs/go-jti.postman_collection.json)
* REQUEST & RESPONSE API (folder /docs/Input.md & Output.md)
* DATABASE : create database db_gojti, (folder /scripts/db_gojti.sql)

INSTALLATION, CONFIGURATION AND RUNING
------------

 * Import database from folder /scripts/db_gojti.sql
 * Config from /pkg/configuration/config.go
 * go build
 * ./go-jti or ./go-jti.exe
 * Import postman collection FROM /docs/go-jti.postman_collection
  
 TESTING
------------
* http://localhost:9670/ & sign google account


## POSTGRES
```
    https://api.elephantsql.com/console/3ff65ee8-6ede-4d41-a458-aedcd5f5c1e4/details
    postgres://atdlgvno:iOI0Wu5FczrsFfOEu-xaWTZ7lXLP4l1D@rain.db.elephantsql.com/atdlgvno
```


