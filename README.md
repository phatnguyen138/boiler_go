# Project bolier_go

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## First after get
## Update the module name to suit your project
```bash
./scripts/update_modules.sh
```
Then type in your desired name

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```
``

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
create database for using
```bash
make createdb
```
drop database
```bash
make createdb
```
create new migration
```bash
make migratecreate name=name
```
migrate up
```bash
make migrateup
```
migrate down
```bash
make migratedown
```
