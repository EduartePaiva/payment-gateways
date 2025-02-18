FROM nginx:1.27-alpine

WORKDIR /etc/nginx/conf.d

COPY ./nginx.conf ./default.conf