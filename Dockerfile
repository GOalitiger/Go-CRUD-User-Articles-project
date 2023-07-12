FROM golang:1.19  AS builder

ENV GO111MODULE=on 
ENV APP_HOME /go/src/my-app
ENV APP_MAIN /go/main


RUN mkdir -p "$APP_HOME"
RUN mkdir -p "$APP_MAIN"

WORKDIR "$APP_HOME"

# COPY go.mod "$APP_HOME"
# COPY go.sum "$APP_HOME"

COPY . "$APP_HOME"
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# RUN cp main "$APP_MAIN"
# WORKDIR "$APP_MAIN"
# RUN rm -r "$APP_HOME"

# COPY "D:/Ali-Demo-Project/echo-User-Articles-project/main.go" "$APP_HOME"
# EXPOSE 8080



# CMD ["go","run","."]

FROM alpine:latest  

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/my-app/main ./
COPY --from=builder /go/src/my-app/.env ./

# CMD ["./app"]  
CMD ["./main"]
