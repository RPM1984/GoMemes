FROM golang:latest 
RUN mkdir /app 
ADD ./src/ /app/ 
WORKDIR /app
RUN go get -d
RUN go build -o main . 
CMD ["/app/main"]
EXPOSE 80