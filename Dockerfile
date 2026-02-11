FROM alpine:latest

RUN apk --no-cache add tzdata
RUN cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
RUN echo "Asia/Ho_Chi_Minh" > /etc/timezone
WORKDIR /app
COPY application/migration ./application/migration
COPY application/etc/casx ./application/etc/casx
COPY SSO ./SSO

EXPOSE 9091
CMD ["sh", "-c", "/app/SSO"]
