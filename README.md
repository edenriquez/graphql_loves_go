# Graph loves Go

Repo with example shown for backs for the furut meetup talk.
This repo aim is to provide to the attendes the code base explained in these **[slides](https://docs.google.com/presentation/d/1ssDKywo4ReXP2GP59gc8l88PBFdB51K2jnnkhnyTe_4/edit?usp=sharing)**.

## Installation

download repository 

```bash
git clone git@github.com:edenriquez/graphql_loves_go.git
```

install backend dependencies 

```
cd backend
# enable go module if is not 
export  GO111MODULE="on"
# run gqlgen over the project
go run github.com/99designs/gqlgen
```

install frontend dependencies 

```
cd frontend
go get -u github.com/shurcooL/graphql
```

## Usage

run gqlgen server with provided code base


```
# under backend directory
go run server/server.go
```

open another terminal and run the following command 

```
go run main.go
```

The you will see client side consuming backend graphql server side.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
