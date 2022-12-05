# notes

```sh
$ docker run --name playground -e POSTGRES_USER=myusername -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -d postgres
```

```sh
$ go get github.com/jackc/pgservicefile                                    
go: downloading github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b
go: github.com/jackc/pgservicefile@v0.0.0-20200714003250-2b9c44734f2b: verifying module: github.com/jackc/pgservicefile@v0.0.0-20200714003250-2b9c44734f2b: cannot authenticate record data in server response
```

Guess I'll use MySQL for now, have this one later
