#!/bin/sh
mongoimport -d imdb-api -c movies --file imdb.json --jsonArray --uri "mongodb://mongo:27017"
