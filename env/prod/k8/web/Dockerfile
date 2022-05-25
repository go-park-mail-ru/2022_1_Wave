FROM nginx

COPY config/nginx/nginx.prod.conf /etc/nginx/nginx.conf
COPY env/prod/certs /var/www/certs
RUN mkdir -p /var/lib/nginx/proxy_cache
WORKDIR /var/www/wave/
