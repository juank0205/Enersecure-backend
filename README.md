# EnerSecure-backend
Ultra seguro

## API Documentation

### Base URL

```
http://localhost:8080/api/v1
```
### PUBLIC endpoints

use the base url `http://localhost:8080/api/v1/public` and expose the following endpoints.

#### POST /login

Authenticates a client and returns their profile information (excluding password).

##### Request Body

```json
{
  "email": "juan.perez@example.com",
  "password": "secure123"
}
```

##### Response (200 OK)

```json
{
  "token": "..."
}
```

##### Response (401 Unauthorized)

```json
{
  "success": "false"
}
```

#### POST /register

Registers a new client.

##### Request Body

```json
{
  "first_name": "Juan",
  "last_name": "Pérez",
  "government_id": 123456789,
  "phone_number": "1234567890",
  "password": "secure123",
  "email": "juan.perez@example.com"
}
```

##### Response (200 OK)

```json
{
  "success": true
}
```
