
# Enchainte - GoLang API Sample

Minimal GoLang API Project Structure with Docker

## Requirements

 - Docker and docker-compose installed
 - GoLang CLI

## Setup
Clone this repository

    git clone https://github.com/openvino/openvino-api.git
    cd openvino-api

Install Go dependencies

    go mod download

Setup environment variables
    
    cp .env.example .env
    cp .env.yml.example .env.yml

Fill the required parameters in the .env file

    MYSQL_ROOT_PASSWORD= [root_password]
    MYSQL_DATABASE= [database_name]
    MYSQL_USER= [user_name]
    MYSQL_PASSWORD= [user_password] 

Fill the required parameters in the .env.yml file

    port: [api port]
    database:
	    dialect: [db dialect]
	    host: [db host]
	    port: [db port]
	    username: [db username]
	    password: [db password]
	    name: [db name]
	    charset: utf8

To run locally:

    go run main.go

To run on Docker (WIP):

    docker-compose up