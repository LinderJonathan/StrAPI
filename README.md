# StrAPI

## Backend
The backend server is written in Go, and uses the Gin HTTP Web framework for HTTP routing- and request handling.

For database interaction, the ```modernc.org/sqlite``` and ```database/sql``` packages were used to open and establish SQLite database connections to send queries.

### Example of supported HTTP methods

**GET request**
```bash
curl http://localhost:5000/activities
```
**GET request**

Activities are assigned a unique identifier through increments on POST requests. The current GET request on the frontend will display all activities, and identifiers are at the time not necessary
```bash
curl http://localhost:5000/activities/0
```

**POST request**
```bash
curl -X POST http://localhost:5000/activities \
-H "Content-Type: application/json" \
-d '{
  "title": "Morning Run",
  "description": "5km jog around the park",
  "durationHours": 0,
  "durationMinutes": 30,
  "durationSeconds": 15,
  "activity": 2
}'
```

**PUT request**
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
**DELETE request**
```bash
curl -X DELETE http://localhost:5000/activities/2
```

## React with TypeScript frontend
The Frontend is written in TypeScript with the React framework, together with node for development. The UI currently holds a home page and a page for the users activities, where all posted activities are shown.

## Database
StrAPI uses SQLite for managing data due to the smaller scale of the project.

## How to use
### Clone the repository
```https://github.com/LinderJonathan/StrAPI.git```

### Start the backend- and frontend server
Open two separate terminal windows. In the first terminal, navigate to the backend folder, build the program and run the server:
```bash
cd backend/

go build .

go run .
```

This will start the backend server on port ``5000``.

In the second terminal, navigate to the React frontend app and start the frontend server:
```bash
cd frontend/app/

npm run dev
```