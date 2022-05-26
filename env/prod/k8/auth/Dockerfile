FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN go get github.com/mailru/easyjson
RUN go install github.com/mailru/easyjson/...@latest
RUN bash gRPC.sh
RUN bash easyjsonGenerate.sh
RUN go build -o ./tmp/auth ./cmd/auth/auth.go

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/auth /var/www/service/
WORKDIR /var/www/service/
EXPOSE 8085
EXPOSE 9085
RUN chmod +x auth
ENTRYPOINT ["sh", "-c", "./auth"]