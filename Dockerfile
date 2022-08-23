FROM node:alpine
RUN apk add --no-cache go
COPY frontend/dispnote /app/
COPY server/dispnote.go /app/
WORKDIR /app/
RUN go build dispnote.go
RUN rm /app/dispnote.go
WORKDIR /app/frontend/dispnote
RUN npm install crypto-js
RUN npm run build
WORKDIR /app/
ENV DISPNOTE_BIND_ADDR="0.0.0.0:8085"