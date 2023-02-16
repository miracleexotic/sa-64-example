# syntax=docker/dockerfile:1

FROM node:14.17.5-alpine3.14

WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend/. .
EXPOSE 3000
CMD [ "npm", "start" ]
