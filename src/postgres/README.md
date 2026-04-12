# notes

```sh
$ docker run --name playground-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -d postgres:latest
$ docker ps
$ docker exec -it <container_name_or_id> /bin/bash
> createdb playground
> which psql
> psql -U root
>> \l
```
