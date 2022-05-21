FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN go get github.com/mailru/easyjson
RUN go install github.com/mailru/easyjson/...@latest
RUN bash easyjsonGenerate.sh
RUN bash cmd/playlist/prod.sh

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/playlist /var/www/service/
WORKDIR /var/www/service/
EXPOSE 8084
EXPOSE 9084
RUN chmod +x playlist
ENTRYPOINT ["sh", "-c", "./playlist"]