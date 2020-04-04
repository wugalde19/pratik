# pratik
Go Microservices

## How to run services
In order to run all the services using docker you will need to
create a network.
This will allow containers to talk between them.
```
// docker network create network_name
docker network create myNetwork
```

Once you create your network, you need to build the DB image and run
the container.
```
// From the project's root
cd postgres

docker build . --build-arg db_pass_env=value_set_in_env_file -t postgres

docker run --rm --name pg-docker -p 5432:5432 --network myNetwork postgres:latest
```

Notice that when we run the container we specify what network we want this
container to use, using the `network` flag.

Then you'll need to build the MC1 image and run the container.
```
// From the project's root
cd mc1

docker build . -t goapp

docker run --rm --name goapp-docker -p 8000:8000 --network myNetwork goapp:latest
```

Finally, you'll need to build the MC2 image and run the container.
```
// From the project's root
cd mc2

docker build . -t goapp2

docker run --rm --name goapp2-docker -p 8001:8001 --network myNetwork goapp2:latest
```

Now with the containers up and running you are ready to start sending
requests.

## ENDPOINTS

### MC1

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

### MC2


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
