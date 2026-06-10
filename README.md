# OpenVino - Golang API

Minimal GoLang API Project Structure with Docker used for Bloock APIs

## Requirements

 - Docker & Docker-compose installed
 - GoLang CLI

## Setup
Clone this repository

    git clone https://github.com/openvino/openvino-api
    cd openvino-api

Install Go dependencies

    go mod download

Setup environment variables

    cp .env.example .env

Fill the required parameters in the .env file. Some of them can be left empty depending on the case.

    ENVIRONMENT= [dev/prod environment]
    API_PORT= [exposed api port]

    DB_NAME=  [database name]
    DB_HOST= [database endpoint]
    DB_PORT= [database exposed port, 5433 by default for local development]
    DB_USERNAME= [database user]
    DB_PASSWORD= [database password]
    SKIP_AUTO_MIGRATE= [true to skip automatic schema migration on startup]

To run locally:

    docker compose --profile local-db up -d openvino-db
    set -a; source .env; set +a
    go run main.go

To run on Docker (WIP) you have to modify the Dockerfile $database and the $port with the proper information and:

    docker build -t openvino-api:latest .
    docker run -ti openvino-api:latest -e ".env" 

To run using Docker-compose:

    docker compose up --build openvino-api

To run using Docker-compose with the local database:

    docker compose --profile local-db up --build
