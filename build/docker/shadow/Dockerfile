FROM ubuntu:22.10

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y openssh-server dnsutils iputils-ping net-tools iproute2 curl lsof && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /var/run/sshd && \
    # SSH login fix. Otherwise user is kicked off after login
    sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

COPY build/docker/shadow/sshd_config /etc/ssh/sshd_config
RUN chmod +rw /etc/ssh/sshd_config

RUN rm -rf /etc/ssh/ssh_host_rsa_key && ssh-keygen -N '' -t rsa -f /etc/ssh/ssh_host_rsa_key
	
# RUN echo '123456' | passwd --stdin root 
RUN echo 'root:123456' | chpasswd

EXPOSE 22

COPY artifacts/shadow/shadow-linux-amd64 /usr/sbin/shadow
COPY build/docker/shadow/run.sh /run.sh
COPY build/docker/shadow/disconnect.sh /disconnect.sh

RUN chmod 755 /disconnect.sh && \
    chmod 755 /run.sh

ENTRYPOINT ["/run.sh"]
