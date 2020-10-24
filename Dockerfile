FROM golang:1.7
 
RUN mkdir /app
 
WORKDIR /app
 
COPY . /app
 
RUN go build ./main.go
 
CMD ["./main"]
