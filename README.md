# pratik
Go Microservices


#ENDPOINTS

#MC1

-------- Register User
```
PATH: localhost:8000/v1/registration/
METHOD: POST
HEADER: No needed
BODY:
{
  "name": "Test",
  "mobile_number": "1234567890",
  "email": "test@test.com",
  "password": "test123"
}
```

-------- Login
```
PATH: localhost:8000/v1/login/
METHOD: POST
HEADER: No needed
BODY:
{
  "mobile_number": "1234567890",
  "password": "test123"
}
```

-------- Get User Details
```
PATH: localhost:8000/v1/user/
METHOD: GET
HEADER: Authorization
BODY: No needed
```

#MC2


-------- Registered Users Count
```
PATH: localhost:8001/v1/user/count
METHOD: GET
HEADER: Authorization
BODY: No needed
```

-------- Update Password
```
PATH: localhost:8001/v1/user/update-password
METHOD: POST
HEADER: Authorization
BODY:
{
  "user_id": 1,
  "old_password": "test123",
  "new_password": "test456"
}
```
