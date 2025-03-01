FROM node:22-alpine as build

WORKDIR /app

RUN npm install -g pnpm@latest-10

COPY ./package.json .
COPY ./pnpm-lock.yaml .

RUN pnpm install

COPY . .

RUN pnpm build

FROM nginx:1.27-alpine

# WORKDIR /etc/nginx/conf.d

COPY prod.nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /bin/www

EXPOSE 80