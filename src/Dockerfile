FROM golang:latest

COPY . /silent-songs/
WORKDIR /go/silent-songs/
RUN apt update -y && apt upgrade -y
# install the dependencies
RUN apt-get install sqlite3 -y
RUN apt-get install nodejs -y && apt-get install npm -y
RUN npm install -g typescript
# init the database
RUN cat database/inicialize.sql | sqlite3 database/publications.db
#install the dependencies
RUN go mod download
# compile some stuff
RUN tsc public/script/*.ts
RUN go build main.go
# delete some things
RUN rm -rf public/script/*.ts
RUN  rm -rf database/inicialize.sql
RUN apt-get remove nodejs -y && apt-get remove npm -y
CMD ["./main"]
