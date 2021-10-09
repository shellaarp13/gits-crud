# GITS-CRUD



## Contributors

1. Pramaishella Ardiani Regita Putri - Universitas Telkom

2. Risky Kurniawan - Universitas Adhirajasa Reswara Sanjaya

## About Application
- This application can develop in golang version 1.7

## Requirement Application
- Golang version 1.7
- Postgresql

## Directory Application

- [DBSchema](https://gitlab.com/riskykurniawan15/gits-crud/-/tree/main/dbschema)  directory about prototype database (ERD)
- [MYapp](https://gitlab.com/riskykurniawan15/gits-crud/-/tree/main/myapp)  directory about appliaction
- [Postman-Collection](https://gitlab.com/riskykurniawan15/gits-crud/-/tree/main/postman-collection)  directory about collection in postman for testing this appliaction

## Configure Application

```
before configure this application you can create database in postgresql
open directory myapp
duplicate file ".env.example" and rename to ".env"
open ".env" and configuring environtment application
if first time open application you can open terminal and execute command "go get ./..." for install all go modules
```

## Migration Database Application

```
open directory myapp
open terminal and please running command "go run script/migration/create_migration_script.go"
please wait the process application
```

## Run Application

```
open directory myapp
open terminal and please running command "go run app/main.go"
dont close terminal if you can running application, for exit you can closing terminal or press ctrl + c in terminal
```

## Build Application

```
open directory myapp
open terminal and please running command "go mod vendor"
wait process and if success to create vendor this application can copy all importing file in folder vendor
open terminal and please running command "go build app/main.go"
wait process and if success application was created file in myapp directiori, this file is "main"
running application from your terminal and open file "main"
dont close terminal if you can running application, for exit you can closing terminal or press ctrl + c in terminal
```