# StrAPI

## Example commands

### GET request
```bash
curl http://localhost:5000/activities
```

```bash
curl http://localhost:5000/activities/0
```

### POST request
```bash
curl -X POST http://localhost:5000/activities \
-H "Content-Type: application/json" \
-d '{
  "title": "title1",
  "description": "description1",
  "durationHours": 1,
  "durationMinutes": 1,
  "durationSeconds": 1,
  "activity": 1
}'
```

### PUT request
```bash
curl -X PUT http://localhost:5000/activities/1 \
-H "Content-Type: application/json" \
-d '{
  "id": 1,
  "title": "Morning Run",
  "description": "5km jog around the park",
  "durationHours": 0,
  "durationMinutes": 30,
  "durationSeconds": 0,
  "activity": 2
}'
```
### DELETE request
```bash
curl -X DELETE http://localhost:5000/activities/2
```

## Database

Write about SQLite here

## Links
[GO REST api](https://go.dev/doc/tutorial/web-service-gin)

## Questions
1. How does POST requests "magically" work with the handler function?

2. How are "race-conditions" handled, i.e. possible PUT requests towards the same endpoint?