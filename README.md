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
