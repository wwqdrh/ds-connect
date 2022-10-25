FROM ubuntu:22.10

COPY artifacts/shadow-controller/shadow-controller-linux-amd64 /usr/sbin/shadow-controller

EXPOSE 18080

ENTRYPOINT ["/usr/sbin/shadow-controller"]
