## Request example

```sh
go run main.go
```

```sh
curl --location --request POST 'localhost:8080/reservation' \
--header 'Content-Type: application/json' \
--data-raw '{
    "guestId":1,
    "startDate":"2024-07-04T01:00:00Z",
    "endDate":"2024-07-05T00:00:00Z",
    "roomId":1
}'
```
