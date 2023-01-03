# notes

```sh
$ docker run --name playground-mysql -e MYSQL_ROOT_PASSWORD=mypassword -p 3306:3306 -d mysql:8.0
```


After connecting to docker
```sh
$ mysql --user=root --password playground
$ SHOW TABLES;
$ SELECT * FROM rand_nums;
```

Create DB
```SQL
CREATE DATABASE IF NOT EXISTS playground;
SHOW DATABASES;
```
