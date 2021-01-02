# Backend API

## Start backend in DEV mode
Here for we will start our MongoDB, Express and frontend in Docker.
Next we have to set the necessary environment variables which are used by the backend.
At last we can start our backend service locally.

```
$ docker-compose up -d
$ . ./set-local-vars.sh
$ go run main.go
```

## Access UI
Visit http://localhost

## API
The API supports the following cases:
- add a movie (POST)
- remove a movie (DELETE)
- list all movies (GET)
- mark a movie as 'watched' (PUT)
- get a specific movie (optional) (GET)

Add a movie
```
$ curl localhost:8080/movie -d '{"title": "The Dark Knight", "year": 2008, "Watched":false }'
{"id":"5fe8d41419678cfd2ab2c6e1"}
```

List all movies
```
$ curl localhost:8080/movie
{"movies":[{"_id":"5fe8d41419678cfd2ab2c6e1","title":"The Dark Knight","year":2008,"watched":false},{"_id":"5fe8d4a319678cfd2ab2c6e2","title":"Zodiac","year":2007,"watched":false}]}
```

Put a movie as watched
```
$ curl -X PUT -H "Content-Type: application/json" -d '{"_id":"5fe8d4a319678cfd2ab2c6e2"}' localhost:8080/movies
""
```

Delete a movie
```
$ curl -XDELETE localhost:8080/movie/5fe8d41419678cfd2ab2c6e1
""
```

Optional: Get a specific movie (in comment because it's currently not used by the frontend)
```
$ curl localhost:8080/movie/5fe8d4a319678cfd2ab2c6e2
{"id":"5fe8d4a319678cfd2ab2c6e2","title":"Zodiac","watched":true,"year":2007}
```