FROM golang:1.20

# membuat direktori folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable
RUN go build -o mytaskapi

# RUN go build -o main .
CMD ["./mytaskapi"]