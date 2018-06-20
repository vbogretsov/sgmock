FROM alpine:3.7
LABEL author="bogrecov@gmail.com"

ENV SGMOCK_KEY=only-for-tests

COPY ./sgmock /bin/sgmock
COPY ./docker-entrypoint.sh /bin/docker-entrypoint.sh

RUN adduser -D sgmock

USER sgmock

EXPOSE 9001

ENTRYPOINT ["docker-entrypoint.sh"]