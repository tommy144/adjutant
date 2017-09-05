FROM ubuntu:latest
ADD adjutant_linux_amd64 .
ADD conf.yml .
ENTRYPOINT ["./adjutant_linux_amd64"]
