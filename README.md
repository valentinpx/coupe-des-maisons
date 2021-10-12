# CDM-API
API de la coupe des maisons Harry Potter de la piscine 2021 d'Epitech Lille.

## Installation and build
### Prerequisite
In order to use this project, you first must have installed go. The steps are described [here](https://golang.org/doc/install).

### Step 1
Install the required dependencies via go command line and clone the repo:
```sh
$ go get github.com/gin-gonic/gin
$ go get github.com/mattn/go-sqlite3
```
```sh
$ git clone https://github.com/valentinpx/cdm-api.git
$ cd cdm-api/src
```

### Step 2
Build the project and run it.
```sh
$ go build .
$ mv src ../a.out
$ cd ..
$ ./a.out "<db_path>" "<host>:<port>"
```
***Arguments***
- `"db_path":string` Path to the database
- `"host":string` Host of the API (ex: localhost)
- `"port":string` Port of the host (ex: 4242)
You can retrieve the post key after executing the program, Have fun!

## API
###  Get a house total
***Definition***
- `GET /api/houses/<house_name>/total`

***Arguments***
- `"house_name":string` Name of the house

***Response***
- `200 OK` on success
```json
{
    "name" : "Serpentard",
    "total" : 42
}
```

### Get the transations list
***Definition***
- `GET /api/transactions`

***Response***
- `200 OK` on success
```json
[
    {
        "house": "Serpentard",
        "description": "Illusiooon",
        "amount": 42,
        "author": "Connor",
        "date": "12/10/2021 01:41:48"
    },
    {
        "house": "Serdaigle",
        "description": "Dumb",
        "amount": 64,
        "author": "Ledore",
        "date": "12/10/2021 01:42:25"
    },
]
```

### Add a transation
***Definition***
- `POST /api/transactions`

***Body Arguments***
- `"house_name":string` Name of the house
- `"description":string` Description of the action that was worth these points
- `"amount":string` Amount of points
- `"author":string` Giver
- `"key":string` Key generated when the program is started

***Body***
```json
{
    "transaction": {
        "house": "Serpentard",
        "description": "Illusiooon",
        "amount": 42,
        "author": "Connor",
    },
    "key": "yakak"
}
```

***Response***
- `200 OK` on success
- `403 Forbidden` on wrong key
