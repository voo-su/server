FROM alpine:3.19

ARG GOLANG_VERSION=1.22.2
ARG SSH_PASSWORD=root

# Установка зависимостей и очистка временных файлов
RUN apk update && \
    apk add --no-cache make gcc openssh bash musl-dev openssl-dev ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

# Установка Go
RUN wget https://dl.google.com/go/go$GOLANG_VERSION.src.tar.gz && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.src.tar.gz && \
    rm go$GOLANG_VERSION.src.tar.gz

# Настройка SSH
RUN sed -i -e 's/#PermitRootLogin.*/PermitRootLogin\ yes/' \
    -e 's/#Port 22/Port 22/' /etc/ssh/sshd_config && \
    echo "root:${SSH_PASSWORD}" | chpasswd && \
    /usr/bin/ssh-keygen -A && \
    ssh-keygen -t rsa -b 4096 -f  /etc/ssh/ssh_host_key && \
    echo "export VISIBLE=now" >> /etc/profile

# Добавление Go в PATH
ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir /opt/voo-su

WORKDIR /opt/voo-su

COPY . /opt/voo-su

RUN ln -sf /opt/voo-su/configs/voo-su.yaml /opt/voo-su/runtime/voo-su.yaml

EXPOSE 22
EXPOSE 8000
EXPOSE 8001
EXPOSE 8002

CMD ["/usr/sbin/sshd", "-D"]
