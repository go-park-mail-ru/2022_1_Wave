FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN bash cmd/album/prod.sh

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/album /var/www/service/
WORKDIR /var/www/service/
EXPOSE 8081
EXPOSE 9081
RUN chmod +x album
ENTRYPOINT ["sh", "-c", "./album"]
