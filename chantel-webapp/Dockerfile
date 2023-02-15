FROM node:lts-slim AS builder
WORKDIR /app
COPY ./package.json ./package.json
RUN yarn
COPY . .
RUN yarn build

FROM nginx:stable
WORKDIR /code
RUN mkdir /logs
COPY --from=builder /app/build /code/dist
COPY ./nginx/nginx.conf /etc/nginx
COPY ./nginx/app.conf /etc/nginx/conf.d/default.conf 