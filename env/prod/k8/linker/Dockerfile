FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN go get github.com/mailru/easyjson
RUN go install github.com/mailru/easyjson/...@latest
RUN bash generateAll.sh
RUN go build -o ./tmp/linker ./cmd/linker/linker.go

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/linker /var/www/service/
WORKDIR /var/www/service/
RUN chmod +x linker
ENTRYPOINT ["sh", "-c", "./linker"]
