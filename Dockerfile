FROM docker.io/cupcakearmy/autorestic
ENTRYPOINT []
CMD [ "autorestic" ]
COPY autorestic-datadog-statsd /usr/bin/autorestic-datadog-statsd
