FROM 61.145.163.124:8084/library/ubuntu:20.04
LABEL authors="ye"
WORKDIR /app
COPY ./metrics /app/metrics
RUN chmod +x /app/metrics
EXPOSE 10019
ENTRYPOINT ["./metrics"]