# Angular Go MongoDB Movie application

This a very basic movie application where you can add movies you want to watch.
You can mark movies as "watched". You can also remove movies from the list.
The fronted is using the Angular framework. The backend is written in Go and uses the [gin-gonic framework](https://github.com/gin-gonic/gin). The movies are stored in a NoSQL MongoDB.

## Run the application locally
You can run the application locally using docker-compose.
```
$ docker-compose up -d --build
```

This will spin up 4 Docker containers:
  * MongoDB container (NoSQL DB)
  * Mongo Express container (UI in front of the MongoDB)
  * Go container (to host our backend API)
  * Nginx container (to host our frontend)

```
$ docker ps
CONTAINER ID   IMAGE                         COMMAND                  CREATED          STATUS          PORTS                    NAMES
96650b3b1c0f   angular-go-mongodb_frontend   "/docker-entrypoint.…"   20 seconds ago   Up 19 seconds   0.0.0.0:80->80/tcp       frontend
e7a290c24c08   angular-go-mongodb_backend    "./main"                 20 seconds ago   Up 19 seconds   0.0.0.0:8080->8080/tcp   backend
765230b549a2   mongo-express                 "tini -- /docker-ent…"   21 seconds ago   Up 20 seconds   0.0.0.0:8081->8081/tcp   mongo-express
ba93362b410e   mongo                         "docker-entrypoint.s…"   21 seconds ago   Up 20 seconds   27017/tcp                mongo
```

## Application
Visit http://localhost
![Screenshot 2021-01-02 at 13 47 30](https://user-images.githubusercontent.com/14105387/103457614-18b8c380-4d01-11eb-9408-aa5e34f769b0.png)

## Database
Visit http://localhost:8081
![Screenshot 2021-01-02 at 13 49 26](https://user-images.githubusercontent.com/14105387/103457658-69302100-4d01-11eb-8167-01ccd94510dc.png)

## Frontend
See [here](./ui)

## Backend
See [here](./server)
