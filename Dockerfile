FROM debian:stretch-slim

EXPOSE 80

WORKDIR /app

COPY main /app

CMD ["./main"]