# TaskFlow API Documentation

Base URL: `http://localhost:8080`

All request and response bodies use JSON unless otherwise specified.

## 1. Health Check

**GET** `/health`

Checks whether the server is running.

Successful response: `200 OK`

```text
TaskFlow API is running
```

## 2. Create a Task

**POST** `/api/v1/tasks`

Request body:

```json
{
  "title": "Learn Go",
  "completed": false
}
```

Successful response: `201 Created`

```json
{
  "id": 1,
  "title": "Learn Go",
  "completed": false
}
```

Possible errors:
- `400 Bad Request` — Invalid JSON or empty title.

## 3. Get All Tasks

**GET** `/api/v1/tasks`

Successful response: `200 OK`

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "completed": false
  }
]
```

If no tasks exist, the response is an empty JSON array.

## 4. Get a Task by ID

**GET** `/api/v1/tasks/{id}`

Example: `/api/v1/tasks/1`

Successful response: `200 OK`

Possible errors:
- `400 Bad Request` — Invalid task ID.
- `404 Not Found` — Task does not exist.

## 5. Update a Task

**PUT** `/api/v1/tasks/{id}`

Request body:

```json
{
  "title": "Learn Go REST API",
  "completed": true
}
```

Successful response: `200 OK`, containing the updated task.

Possible errors:
- `400 Bad Request` — Invalid JSON, invalid ID, or empty title.
- `404 Not Found` — Task does not exist.

## 6. Delete a Task

**DELETE** `/api/v1/tasks/{id}`

Successful response: `204 No Content`

Possible errors:
- `400 Bad Request` — Invalid task ID.
- `404 Not Found` — Task does not exist.

## Error Response Format

Example:

```json
{
  "error": "Task not found"
}
```

## Notes

- Task data is stored in memory.
- Data is lost when the server stops.
- Task IDs are assigned by the server.
- The API currently has no authentication or database integration.
