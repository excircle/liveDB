FROM mysql:5.7
RUN apt update -y
COPY ./static/mysql/my.cnf /etc/
COPY ./static/mysql/init.sql /etc/
RUN ln -s /var/lib/mysql/mysqld.sock /var/run/mysqld/mysqld.sock
