# food-delivery
Basic low-level Design for food-delivery  
Reference: [link](https://medium.com/@mayankbansal933/food-delivery-app-lld-c1409ef49266)

Docker images commands
docker network create food-delivery-networks  
docker-compose -f .\deployment\docker-compose.yaml -p food-delivery up --force-recreate  

### Run Migration
go run .\cmd\migration\main.go up
### Run Web Server
go run .\cmd\web\main.go
