# minesweeper-api

Golang back end implementation of Minesweeper API. This implementation doesn't provide a front end. Game is finished when the state is `exploded` or `exploded`.

## Build

```
  $ go build
```

## Run

```
  $ go run
```

# REST API

The REST API is described below

## Create a New Game

### Request
`POST /minesweeper`

```
  $ curl http://localhost:8080/minesweeper \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"rows": 2, "cols": 2, "mines": 1}'
```

### Response
```
    HTTP/1.1 201 Created
    Content-Type: application/json; charset=utf-8
    Date: Thu, 30 Sep 2021 03:06:18 GMT
    Content-Length: 122

    {
        "id": "029e8973-e31a-4573-9246-ca073d36d748",
        "rows": 2,
        "cols": 2,
        "mines": 1,
        "state": "created"
    }
```

## Start the game

### Request
`POST /minesweeper/:id/start`

```
  curl http://localhost:8080/minesweeper/6d17c000-df15-4c7e-9215-7d0335043cb8/start \
    --include \
    --header "Content-Type: application/json" \
    --request "POST"
```

### Response
```
    HTTP/1.1 201 Created
    Content-Type: application/json; charset=utf-8
    Date: Thu, 30 Sep 2021 03:10:00 GMT
    Content-Length: 832

    {
        "id": "6d17c000-df15-4c7e-9215-7d0335043cb8",
        "rows": 2,
        "cols": 2,
        "board": [
            [
                {
                    "value": 0,
                    "mine": false,
                    "clicked": false,
                    "flagged": false
                },
                {
                    "value": 1,
                    "mine": false,
                    "clicked": false,
                    "flagged": false
                }
            ],
            [
                {
                    "value": 0,
                    "mine": true,
                    "clicked": false,
                    "flagged": false
                },
                {
                    "value": 0,
                    "mine": false,
                    "clicked": false,
                    "flagged": false
                }
            ]
        ],
        "mines": 1,
        "state": "started",
        "createTime": 1632971400
    }
```

## Get the game

### Request
`GET /minesweeper/:id`

```
  curl http://localhost:8080/minesweeper/6d17c000-df15-4c7e-9215-7d0335043cb8
```

### Response
```
    {
        "id": "6d17c000-df15-4c7e-9215-7d0335043cb8",
        "rows": 2,
        "cols": 2,
        "board": [
            [
                {
                    "value": 0,
                    "mine": false,
                    "clicked": false,
                    "flagged": false
                },
                {
                    "value": 1,
                    "mine": false,
                    "clicked": false,
                    "flagged": false
                }
            ],
            [
                {
                    "value": 0,
                    "mine": true,
                    "clicked": false,
                    "flagged": false
                },
                {
                    "value": 0,
                    "mine": false,
                    "clicked": false,
                    "flagged": false
                }
            ]
        ],
        "mines": 1,
        "state": "started",
        "createTime": 1632971400,
        "consumedTime": 147
    }

```

## Click a cell

### Request
`POST /minesweeper/:id/reveal`

```
  curl http://localhost:8080/minesweeper/6d17c000-df15-4c7e-9215-7d0335043cb8/reveal \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"row": 0, "col": 0}'
```

### Response
```
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 30 Sep 2021 03:14:38 GMT
    Content-Length: 875

    {
        "id": "6d17c000-df15-4c7e-9215-7d0335043cb8",
        "rows": 2,
        "cols": 2,
        "board": [
            [
                {
                    "value": 0,
                    "mine": false,
                    "clicked": true,
                    "flagged": false
                },
                {
                    "value": 1,
                    "mine": false,
                    "clicked": false,
                    "flagged": false
                }
            ],
            [
                {
                    "value": 0,
                    "mine": true,
                    "clicked": false,
                    "flagged": false
                },
                {
                    "value": 0,
                    "mine": false,
                    "clicked": true,
                    "flagged": false
                }
            ]
        ],
        "mines": 1,
        "numClicks": 2,
        "state": "started",
        "createTime": 1632971400,
        "consumedTime": 278
    }
```


## Flag a cell

### Request
`POST /minesweeper/:id/flag`

```
  curl http://localhost:8080/minesweeper/6d17c000-df15-4c7e-9215-7d0335043cb8/flag \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"row": 0, "col": 1}'
```

### Response
```
    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Thu, 30 Sep 2021 03:15:55 GMT
    Content-Length: 874

    {
        "id": "6d17c000-df15-4c7e-9215-7d0335043cb8",
        "rows": 2,
        "cols": 2,
        "board": [
            [
                {
                    "value": 0,
                    "mine": false,
                    "clicked": true,
                    "flagged": false
                },
                {
                    "value": 1,
                    "mine": false,
                    "clicked": false,
                    "flagged": true
                }
            ],
            [
                {
                    "value": 0,
                    "mine": true,
                    "clicked": false,
                    "flagged": false
                },
                {
                    "value": 0,
                    "mine": false,
                    "clicked": true,
                    "flagged": false
                }
            ]
        ],
        "mines": 1,
        "numClicks": 2,
        "state": "started",
        "createTime": 1632971400,
        "consumedTime": 355
    }

```

