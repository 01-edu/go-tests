FROM docker.01-edu.org/golang:1.18.0-alpine3.15

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git

RUN go install mvdan.cc/gofumpt@v0.1.1
RUN go install github.com/01-edu/rc@v0.1.0

WORKDIR /piscine-go
RUN go mod init piscine-go
RUN go get github.com/01-edu/z01@v0.1.0

WORKDIR /go-tests
COPY go.* ./
RUN go mod download
COPY lib lib
COPY solutions solutions
COPY tests tests
RUN go install $(grep -rl 'challenge.Program' ./tests | rev | cut -d/ -f2- | rev)

RUN rm -rf /piscine-go

COPY entrypoint.sh /usr/local/bin
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
