Requirements:
Go : https://golang.org/dl/
Go Swagger : go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
dep : https://github.com/golang/dep/blob/master/README.md

### generate swagger files
swagger generate server -f swagger.yml

### download dependences
dep init
dep ensure

### build the server
go build ./cmd/planet-auth-server/

### start the server
./planet-auth-server.exe --port=8082

### sample usage
curl -X GET http://localhost:8082/users/jsmith
curl -X POST http://localhost:8082/users -H "accept: application/json" -H "Content-Type: application/json" -d @input/CreateUser1.txt
curl -X DELETE http://localhost:8082/users/jsmith
curl -X PUT http://localhost:8082/users/jsmith -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateUser1.txt
curl -X PUT http://localhost:8082/users/jdoe -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateUser2.txt
curl -X PUT http://localhost:8082/users/spervez -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateUser3.txt
curl -X PUT http://localhost:8082/users/jsmith2 -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateUser4.txt

curl -X GET http://localhost:8082/groups/frisbeeplayers
curl -X POST http://localhost:8082/groups -H "accept: application/json" -H "Content-Type: application/json" -d @input/CreateGroup1.txt
curl -X DELETE http://localhost:8082/groups/frisbeeplayers
curl -X PUT http://localhost:8082/groups/admins -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateGroup1.txt
curl -X PUT http://localhost:8082/groups/admins -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateGroup2.txt
curl -X PUT http://localhost:8082/groups/users -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateGroup3.txt
curl -X PUT http://localhost:8082/groups/frisbeeplayers -H "accept: application/json" -H "Content-Type: application/json" -d @input/UpdateGroup4.txt

### Persistence
This is a persistent server, which means that a subsequent execution will maintain state from the current execution.
Data is stored on disk in flat files. The two files are "groupData" and "userData". The location of these files is determined by the environment varriable DATA_DIR.

e.g. on windows - SET DATA_DIR="C:\\Users\\jdoe\\go\\src\\PlanetLabs\\AuthService\\data"

This will create the two files
C:\Users\jdoe\go\src\PlanetLabs\AuthService\data\groupData
C:\Users\jdoe\go\src\PlanetLabs\AuthService\data\userData

When the server starts up it reads state from these two files. If the files are not found or are empty, the server initializes in the empty state.
