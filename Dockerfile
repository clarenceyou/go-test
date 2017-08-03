FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/hashicorp/logutils \
    && go build -o main . 
CMD ["/app/main"]