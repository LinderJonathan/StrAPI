# StrAPI

## Example commands
```bash
curl http://localhost:5000/activities
```

```bash
curl http://localhost:5000/activities/0
```

```bash
curl http://localhost:5000/activities     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"id": 1,"title": "title1","description": "description1","durationHours":1, "durationMinutes": 1, "durationSeconds": 1, "activity": 1}'
```
## Links
[GO REST api](https://go.dev/doc/tutorial/web-service-gin)

## Questions
1. How does POST requests "magically" work with the handler function?