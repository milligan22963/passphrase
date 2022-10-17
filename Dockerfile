FROM scratch
COPY passphrase /
ENTRYPOINT ["/passphrase"]