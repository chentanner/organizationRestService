# Golang organization rest server

## Overview

This package contains everything needed to run the rest server which supports connecting to a PostgreSQL database.
It is composed of:
- a REST api service documented by OpenAPI
- a PostgreSQL database for storage

## Installation
#### Download and install
1) Clone the repo 
2) ```go get -v```
    - If it complains about go modules set the terminal environment variable GO111MODULE='on'
3) Update swagger generated interfaces by running this command in the project directory
    - ``` oapi-codegen.exe --generate server,spec,types api/api.yml  > api/gen/organizations.gen.go ```

#### Configuration
Configuration will require a .env file placed in the root of the cloned repository. This will define environment variables for running the server. 
- db_name - Postgres database name
- db_pass = Database password
- db_user = Database username
- db_host = Database host name
- db_port = Database port
- CORS = Comma seperated string listing cors enabled domains

#### Run
1) ```go run main.go```
2) Connect to services, the default url is http://localhost:8000



## Rest API
The rest API is documented by swagger
#### Organizations
- /organizations
  - Get - Fetches a list of all organizations
  - Post - Creates a new organization
