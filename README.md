# Gif Engine
A gif search engine web service that allows people to build up their own set of Term to Gif responses

## Key Features
- Get a gif relative to your terms
- Save a set of terms and a gif
- Import gifs
- Fallback to another common Gif service

## Database
The database needs to contain the following information
- File path to the gif on the FS
- Terms linked to the gif

The terms attached to the gif need to be indexed so they can be searched. Due to this requirement
we will begin with MongoDB, potentially adding ElasticSearch later

## Requests

### Create a gif
```bash
curl --location 'http://localhost:5000/write' \
--header 'Content-Type: text/plain' \
--data '{
    "url": "https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExemNjajVvcGhqd2RnODZnYzBkamR3c3Q2bm00dmR3OGwybW1kaHE0MiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/CuMiNoTRz2bYc/giphy.gif",
    "tags": ["anakin", "its working"]
}'
```

