FROM golang:1.20
WORKDIR /app
COPY ./backend ./
RUN go mod download
RUN #go get github.com/githubnemo/CompileDaemon
#ENTRYPOINT CompileDaemon --build="go build main.go" --command=main
#ENTRYPOINT go run /home/alexey/go/src/git.iu7.bmstu.ru/kaa20u554/main.go
#EXPOSE 5259
CMD ["go", "run", "main.go"]
#RUN CGO_ENABLED=0 GOOS=linux go build -o /
#

