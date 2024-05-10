# Gif Engine
A gif search engine web service that allows people to build up their own set of Term to Gif responses

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

### Request a gif
```bash
curl --location --globoff 'http://localhost:5000/gif?tags[]=something'
```

### Concatenate two gifs
```bash
curl --location 'http://localhost:5000/join' \
--header 'Content-Type: text/plain' \
--data '{
    "urls": ["https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExemNjajVvcGhqd2RnODZnYzBkamR3c3Q2bm00dmR3OGwybW1kaHE0MiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/CuMiNoTRz2bYc/giphy.gif", "https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExemMxM3lhbHRrdDdvMGtpYncwZHA4cDVudHJqMTJ2eWd0bDNlYXl2dCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/RfEbMBTPQ7MOY/giphy.gif"]
}'
```