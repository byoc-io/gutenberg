FROM debian:jessie
MAINTAINER Daniel Negri <danielnegri@byoc.io>

RUN apt-get update && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /usr/local/share/gutenberg
WORKDIR /usr/local/share/gutenberg

COPY release/bin /usr/local/share/gutenberg/bin
COPY config /etc/gutenberg

RUN ln -s /usr/local/share/gutenberg/bin/gutenberg /usr/local/bin/gutenberg

EXPOSE 8080

ENTRYPOINT ["gutenberg"]
CMD ["serve"]