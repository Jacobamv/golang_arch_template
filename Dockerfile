

#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.22.4-bullseye AS build_stage

# Create app folder
RUN mkdir /app 

# Copy the application files from the previous stage
COPY . /app/

# Set the working directory
WORKDIR /app/cmd

ENV GOCACHE=/app/.cache/go-build
ENV GOPROXY=http://nexus-test.humo.tj/repository/go-proxy/

# Кэшируем зависимости и выполняем go mod tidy
RUN --mount=type=cache,target=/go/pkg/mod,id=test go mod download -x

ENV GOCACHE=/app/.cache/go-build

# Build the Go application
RUN --mount=type=cache,target=/go/pkg/mod,id=test --mount=type=cache,target="/app/.cache/go-build",id=test2 GOOS=linux GOARCH=amd64 go build -o app .

#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM oracle/oraclelinux7-instantclient:21 AS final_stage

COPY --from=build_stage /app/ /srv/

WORKDIR /srv/cmd/

RUN chmod +x app

ENTRYPOINT ["/srv/cmd/app"]
