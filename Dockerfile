FROM alpine:latest

WORKDIR /app
COPY application/migration ./application/migration
COPY application/etc/casx ./application/etc/casx
COPY SSO ./SSO

EXPOSE 9091
CMD ["sh", "-c", "/app/SSO"]
