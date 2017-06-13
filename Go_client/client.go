package main

import (
  "fmt"
  "github.com/go-redis/redis"
)

func ConnectNewClient() {
  client := redis.NewClient(&redis.Options {
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })

  pubsub := client.Subscribe("test1") // El argumento es el canal al cual se subscribe
  defer pubsub.Close() // Cierra la conexión cuando no se requiera
  fmt.Println("=> Cliente subscrito")

  for { // Bucle para dejar a la escucha el servidor recibiendo mensajes
    msg, err := pubsub.ReceiveMessage()
    if err != nil {
      fmt.Println("=> No es posible recibir mensajes: ", err)
    }
    fmt.Println("=> Mensaje recibido")
    fmt.Println("=> Canal: ", msg.Channel, "| Mensaje: ", msg.Payload)
  }
}

func main() {
  ConnectNewClient()
}