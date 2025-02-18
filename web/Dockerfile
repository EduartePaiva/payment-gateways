FROM node:22-alpine

WORKDIR /app

COPY package.json .
COPY pnpm-lock.yaml .

RUN npm install -g pnpm@latest-10
RUN pnpm install

EXPOSE 5173

CMD [ "pnpm", "dev", "--host" ]