FROM golang:latest

COPY . ./blog
WORKDIR /go/blog/
RUN apt update -y && apt upgrade -y
# install the dependencies
RUN apt-get install sqlite3 -y  && apt-get install npm -y
RUN npm install -g typescript
# init the database
RUN cat database/inicialize.sql | sqlite3 database/publications.db

# compile some stuff
WORKDIR /go/blog/public/script

RUN tsc *.ts
# delete some things
RUN rm -rf *.ts
WORKDIR /go/blog/
#compile
RUN go build main.go
# delete
RUN  rm -rf database/inicialize.sql
RUN  apt-get remove npm -y

CMD ["./main"]
