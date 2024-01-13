# GoLang RESTful API with MongoDB 🚀

This project is a simple RESTful API built using GoLang with MongoDB.

## Getting Started🛠️

Follow these steps to run the project locally:

1. **Start the Server:**

```bash
go run main.go
```

The server will start at http://localhost:8000.

## API Endpoints🚪

### Get Users

- **Endpoint:** `GET /user`
- **Description:** Retrieve a list of all users.

### Get User by ID

- **Endpoint:** `GET /user/:id`
- **Description:** Retrieve a specific user by ID.

### Create User

- **Endpoint:** `POST /user`
- **Description:** Create a new user.

### Update User by ID

- **Endpoint:** `PUT /user/:id`
- **Description:** Update a specific user by ID.

### Delete User by ID

- **Endpoint:** `DELETE /user/:id`
- **Description:** Delete a specific user by ID.

The API Postman Collection is available in the /postman-collection/ directory.

## Dependencies📦

- [httprouter](https://github.com/julienschmidt/httprouter): HTTP request router
- [mgo](https://gopkg.in/mgo.v2): MongoDB driver for Go

## Closing Notes📝

If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

Happy coding!🚀👨‍💻
