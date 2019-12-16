FROM golang:1.12

# Set the WORKDIR to the project path in your GOPATH, e.g. /go/src/github.com/go-martini/martini/
WORKDIR $GOPATH/src/github.com/edenriquez/graphql_loves_go/server/

# Copy the content of your repository into the image


COPY . ./
# RUN rm go.mod go.sum
RUN export GO111MODULE="on"
RUN go get github.com/ziutek/mymysql 
RUN go get github.com/ziutek/mymysql/native
RUN go get github.com/joho/godotenv
ENV GO111MODULE="on"
RUN go get github.com/99designs/gqlgen
RUN ls
