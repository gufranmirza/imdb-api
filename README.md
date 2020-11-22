# imdb-api
Provides Search functionality for Popular Movies from IMDB 


## Build Project
```
make
```

## Apply Migrations
```
make migrate-up
```
or to create new migration
```
make migrate-new
```


## Load Data
Load the data from JSON file it will look for `./data/imdb.json`
```
make dataloader
```

## Running in Production 
```
make start
```

## Development
```
make run
```

## Search API
[http://localhost:8001/imdb-api/v1/movies/search?q=Lucas&limit=5](http://localhost:8001/imdb-api/v1/movies/search?q=Lucas&limit=5)

## API Documentation
[Swagger API documetation](http://localhost:8001/imdb-api/v1/swagger/index.html)

## Testing
Running test
```
make test
```

## Generate Mocks
Generate package mocks
```
make mock
```

## Running Docker image
Build and run docker image
```
docker-compose up -d --build
```
