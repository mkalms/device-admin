FROM alpine:latest

USER root
RUN mkdir -p /mnt/shared_data && chown -R 1000 /mnt/shared_data

USER 1000

COPY /build/app /

EXPOSE 8080

CMD ["/app"]
