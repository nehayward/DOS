# DOS API
This is a simple DOS attack api, built to help test server load.

![GoDoc]()
![Version](https://img.shields.io/badge/version-0.1-brightgreen.svg)



## Installation
**Minimum Go version:** Go 1.6

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/nehayward/dos/
```

## Usage
```sh
$ go install
$ dos
```

## Endpoints
|URL | HTTP Method | Functionality |
|:---:|:---:|:---:|
|[/attack](#attack) | POST | Create a new attack, pass URL |
|[/attacks](#attacks)| GET | List current attacks |
|[/attacks/{attack-ID}](#attacksattackid)| GET | Look up attack by ID |
|[/stop](#stop)| DELETE | Stop an attack |

----------------------------

### /attack
```
curl -H "Content-Type: application/json" -d '{"url":"http://localhost:8888"}' http://localhost:8080/attack
```

### /attacks
```
curl http://localhost:8080/attacks
```

###  /attacks/{attackID}
```
curl http://localhost:8080/attacks/10
```

###  /stop
```
curl -X DELETE -d '{"ID": 10}' http://localhost:8080/stop
```

## Wishlist
- [ ]

## Requirements
* [Go](https://github.com/golang/example)


### License
This work is published under the MIT license.

Please see the [LICENSE](https://github.com/nehayward/DOS/blob/master/LICENSE) file for details.
