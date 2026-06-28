FROM debian:bookworm-slim

WORKDIR /app

COPY bin/api .

EXPOSE 3001

CMD ["/app/api"]