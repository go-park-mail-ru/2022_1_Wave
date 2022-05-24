FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
#RUN go get github.com/mailru/easyjson
#RUN go install github.com/mailru/easyjson/...@latest
#RUN bash easyjsonGenerate.sh
RUN bash cmd/api/prod.sh


FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/api /var/www/service/
COPY config/config.toml /var/www/service/
WORKDIR /var/www/service/
RUN ls .
RUN chmod +x api
ENTRYPOINT ["sh", "-c", "./api"]
