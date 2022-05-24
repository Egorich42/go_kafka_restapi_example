# Go-Kafka simple example

Example of basic connection Golang-app REST-API to Kafka.
Analog of API for get coordinates from run tracker - longitude and lattitude

#### requirements
  - Go 1.18
  - Docker
  - docker-compose


# How to run

1. Run Kafka/Zookeeper
```sh
$ sudo docker-compose up --build
```
2. Add necessary env-variables

```sh 
export AppPort=10000 #port for application
export kafkaPort=9092
```
3. Run app

```sh
$ go run main.go
```  

# Request example
```
curl -X POST http://localhost:8089/v1/coordinates -H 'Content-Type: application/json' -d '{"user_id":"d23432","lat":56.86,"lon":67.97}'
```

