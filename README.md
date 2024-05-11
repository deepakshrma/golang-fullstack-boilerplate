# golang-fullstack-boilerplate

This is going to be awesome!!

## Project Structure

```
|   cmd         :=  Common commands, CLI Executables
|   config      :=  App configurations
|   controller  :=
|   env         :=  System environment configurations
|   route       :=
|   server      :=
|   service     :=
|   template    :=  App templates
|   ui          :=
|   util        :=  App utilities

```

## How to create multiple ENV files

Navigate to [env](pkg/env) folder and create files pattern application.{APP_MODE}.env
Here {APP_MODE} is different profile.

## How to run

1. Local

    ```bash
    go run server/main.go
    ```

2. Prod other env

    ```bash
    APP_MODE={env} go run server/main.go
    
    ## Example prod
    APP_MODE=prod go run server/main.go
    ```

3. Using `make` command

    ```bash
    make run
    
    ## Dev mode
    make dev
    
    ```

## How to set logger

You can set diff log level using env variable `LOG_LEVEL={debug,info,warn,error}`  in environment
files[application.env](application.env). The default is `LOG_LEVEL=info`.

```properties
APP_VERSION=1.0.0
LOG_LEVEL=error
```


### useful docker commands

```
docker compose stop
docker compose rm -f
docker-compose up --build -d 
docker exec -it 4c0f5f4acb0c /bin/sh 

```