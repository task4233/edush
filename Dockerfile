FROM centos:centos7

RUN yum -y install httpd
# CMD ["/usr/sbin/httpd", "-DFOREGROUND"]
COPY access_log.txt /tmp