FROM mausved/wave_music_api_v1:latest AS build_step
COPY ./ /var/www/service
WORKDIR /var/www/service/
RUN go get github.com/mailru/easyjson
RUN go install github.com/mailru/easyjson/...@latest
RUN cd bash/ && bash gRPC.sh
RUN bash easyjsonGenerate.sh && cd ..
RUN go build -o ./tmp/user ./cmd/user/user.go

FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=build_step /var/www/service/tmp/user /var/www/service/
WORKDIR /var/www/service/
RUN chmod +x user
ENTRYPOINT ["sh", "-c", "./user"]
