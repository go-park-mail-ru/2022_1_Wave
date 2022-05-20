FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN bash cmd/artist/prod.sh

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/artist /var/www/service/
WORKDIR /var/www/service/
EXPOSE 8082
EXPOSE 9082
RUN chmod +x artist
ENTRYPOINT ["sh", "-c", "./artist"]
