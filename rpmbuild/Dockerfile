FROM rpmbuild/centos7
USER root

WORKDIR /usr/local

RUN yum install -y wget git mercurial
RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
RUN tar xvzf go1.9.2.linux-amd64.tar.gz
RUN rm go1.9.2.linux-amd64.tar.gz

RUN yum install -y rpm-build redhat-rpm-config rpmdevtools glibc glibc-common glibc-headers glibc-dev

ENV PATH /root/go/bin:/usr/local/go/bin:$PATH
RUN go get github.com/golang/dep/cmd/dep
RUN go install github.com/golang/dep/cmd/dep

WORKDIR /usr/local/go/src
COPY . .
