FROM nginx

COPY config/nginx/nginx.prod.conf /etc/nginx/nginx.conf
COPY env/prod/certs /var/www/certs

WORKDIR /var/www/wave/

EXPOSE 80
EXPOSE 443
