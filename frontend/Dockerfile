# syntax=docker/dockerfile:1

FROM node:14.17.5-alpine3.14 AS builder

WORKDIR /frontend

COPY package*.json ./
RUN npm install

COPY frontend/. .

RUN npm run build

ENV NODE_ENV production

EXPOSE 3000

CMD [ "npx", "serve", "-s", "build" ]

