# notes

```sh
$ docker run --name playground-mysql -e MYSQL_ROOT_PASSWORD=mypassword -p 3306:3306 -d mysql:latest  # pin version for production
$ docker ps
$ docker exec -it <container_name_or_id> /bin/bash
$ mysql --user=root --password playground
> CREATE DATABASE IF NOT EXISTS playground;
> SHOW DATABASES;
> SHOW TABLES;
```
