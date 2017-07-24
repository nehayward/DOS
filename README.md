# DOS API
This is a simple DOS attack api, built to help test server load.

[![Build Status](https://travis-ci.com/nehayward/DOS.svg?token=VaqoffAmWMpp9yR53aWy&branch=master)](https://travis-ci.com/nehayward/DOS)
![Version](https://img.shields.io/badge/version-0.1-brightgreen.svg)



## Installation
**Minimum Go version:** Go 1.7

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get github.com/gorilla/mux
$ go get -u github.com/nehayward/dos/
```
#### Alternative
```sh
./start.sh
```

## Usage
```sh
$ go install github.com/nehayward/dos/
$ dos
```

## Docker Installation
1. Install [Docker](https://www.docker.com/).
2. Build Container and Run

```sh
docker build -t dos .
docker run -d -p 8080:8080 dos:0.1
```

## Minikube Getting Started
1. Install [minikube](https://github.com/kubernetes/minikube) on Mac on Linux
### macOS
```shell
brew cask install minikube
```
### Linux
```shell
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
```
2. Ensure minikube is running with `minikube status`

3. `kubectl create -f pod.yaml`



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

## TODO
- [ ] Add Authentication

## Testing
`go test github.com/nehayward/dos/core -v`

## Requirements
* [Go](https://github.com/golang/example)

### License
This work is published under the MIT license.

Please see the [LICENSE](https://github.com/nehayward/DOS/blob/master/LICENSE) file for details.
