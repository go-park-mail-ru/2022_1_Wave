FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN bash cmd/api/prod.sh

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/api /var/www/service/
COPY --from=build_step /var/www/service/db /var/www/service/db
COPY config/config.toml /var/www/service/
WORKDIR /var/www/service/
EXPOSE 5000
RUN chmod +x api
ENTRYPOINT ["sh", "-c", "./api"]
