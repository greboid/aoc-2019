FROM golang:1.13

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

VOLUME /app
WORKDIR /app
CMD ["/entrypoint.sh"]