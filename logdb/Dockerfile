FROM mariadb:10.5.9-focal

ADD logsdatabase.sql /docker-entrypoint-initdb.d

EXPOSE 3306

RUN chmod 755 /docker-entrypoint-initdb.d/logsdatabase.sql