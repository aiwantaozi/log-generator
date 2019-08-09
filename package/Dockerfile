FROM ubuntu:16.04
COPY log-generator /usr/bin/
RUN mkdir -p /var/mylogs/json && mkdir -p /var/mylogs/text && chmod 777 /var/mylogs
CMD ["log-generator"]
