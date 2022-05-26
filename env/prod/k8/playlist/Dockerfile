FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN go get github.com/mailru/easyjson
RUN go install github.com/mailru/easyjson/...@latest
RUN bash gRPC.sh
#RUN bash easyjsonGenerate.sh
RUN go build -o ./tmp/playlist ./cmd/playlist/playlist.go

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/playlist /var/www/service/
WORKDIR /var/www/service/
EXPOSE 8084
EXPOSE 9084
RUN chmod +x playlist
ENTRYPOINT ["sh", "-c", "./playlist"]