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

    DATABASE_DIALECT= [dialect]
    DATABASE_CHARSET= [charset used: utf8]
    DATABASE_NAME=  [database name]

    DATABASE_HOST= [database endpoint]
    DATABASE_PORT= [database exposed port]

    DATABASE_USERNAME= [database user]
    DATABASE_PASSWORD= [database password]
    DATABASE_PASSWORD_ROOT= [database root password]

To run locally:

    go run main.go

To run on Docker (WIP) you have to modify the Dockerfile $database and the $port with the proper information and:

    docker build -t openvino-api:latest .
    docker run -ti openvino-api:latest -e ".env" 

To run using Docker-compose:

    docker-compose up
