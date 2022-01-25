# Gopher Translator Service

An interview assignment project.
To see the full assignment click [here](ASSIGNMENT.md).

## Overview

An HTTP API that accepts english word or sentences and translates them to Gopher language.

### Endpoints

* **POST /word**

    Example request:
    ```JSON
        {
            "english_word": "apple"
        }
    ```

    Example response:
    ```JSON
        {
            "gopher_word": "gapple"
        }
    ```

* **POST /sentence**

    Example request:
    ```JSON
        {
            "english_sentence": "Coding in Go is awesome!"
        }
    ```

    Example response:
    ```JSON
        {
            "gopher_sentence": "odingCogo gin oGogo gis gawesome!"
        }
    ```

* **GET /history**

    Example response:
    ```JSON
    [
        {
            "apple": "gapple"
        },
        {
            "Coding in Go is awesome!": "odingCogo gin oGogo gis gawesome!"
        }
    ]
    ```

## Running the server

### Help
```sh
Usage of gopher-translate:
  -port int
        port number the server should listen on (default 8080)
  -specfile string
        json file with translation instructions
```
For example on the specfile check [this](configs/gopher.json) file.

### Run via Go:
```sh
$ go run cmd/gopher-translate/main.go --port 8080 --specfile configs/gopher.json
```

### Run in Docker container:

```sh
$ docker build -t gopher .   
$ docker run --rm -p 8080:8080 gopher 
```

### Run via docker-compose:
```sh
$ docker-compose up
```