FROM golang:alpine

COPY . ./blog
WORKDIR /go/blog/
RUN apk update && apk upgrade && apk add build-base
# install the dependencies
RUN apk add sqlite  && apk add npm 
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
RUN  apk del npm 

CMD ["./main"]
