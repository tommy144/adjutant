FROM ubuntu:latest
RUN apt-get update && apt-get install ca-certificates -y
ADD adjutant_linux_amd64 .
ADD conf.yml .
ENTRYPOINT ["./adjutant_linux_amd64"]
