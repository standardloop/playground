# notes


```sh
$ docker run --name playground-mongo -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=mypassword -e MONGO_INITDB_DATABASE=playground -v ${PWD}/mongo-init.js:/docker-entrypoint-initdb.d/mongo-init.js:ro mongo:6.0
```

```sh
$ docker-compose up --build -d mongodb
```

```sh
$ mongosh
$ use playground
$ db.auth("root", "mypassword")
$ show collections randNum
$ db.randNum.find()
```
