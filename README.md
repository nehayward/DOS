# DOS API

## Overview
This is a simple DOS attack api, built to help test server load.

## Getting Started
1.


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

## Credits
- [Nick Hayward](https://github.com/nehayward)
