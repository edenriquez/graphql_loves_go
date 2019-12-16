# Graph loves Go

Repo with example shown for backs for the furut meetup talk.
This repo aim is to provide to the attendes the code base explained in these **[slides](https://docs.google.com/presentation/d/1ssDKywo4ReXP2GP59gc8l88PBFdB51K2jnnkhnyTe_4/edit?usp=sharing)**.

## Installation

download repository 

```bash
git clone git@github.com:edenriquez/graphql_loves_go.git
```


Install dependencies 

```
cd server
export  GO111MODULE="on"
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen
```

install frontend dependencies 

```
cd frontend
go get -u github.com/shurcooL/graphql
```


## Usage

To run server in your local:


```
go run server/server.go
```


To run server in Docker:

```
cd server

docker-compose up -d
```


visit graphql playground http://localhost:8080

visit netdata monitor   http://localhost:19999


## Run stress benchmark

```
./server_stress.sh
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
