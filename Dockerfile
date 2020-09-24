FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 80
RUN go build -o main .
CMD ["/bin/sh","./run.sh"]