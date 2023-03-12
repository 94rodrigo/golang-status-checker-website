FROM golang
WORKDIR /app
EXPOSE 8080
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY . ./
COPY /middleware ./middleware
RUN go build -o /main
CMD [ "/main" ]