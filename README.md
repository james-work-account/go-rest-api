# REST API

## Running

```
go run main.go
```

If that fails, you may not have all dependencies installed. Run `go install` to install these.


## Schema

| Route          | Method |  Purpose                    |
|----------------|--------|-----------------------------|
| /api/cars      | GET    | Get all cars                |
| /api/cars      | POST   | Create a car                |
| /api/cars/{id} | GET    | Get a single car            |
| /api/cars/{id} | PUT    | Update a single car         |
| /api/cars/{id} | DELETE | Delete a single car         |