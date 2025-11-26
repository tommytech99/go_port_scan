package main

import ( "fmt"
        "net" // The package is built-in to the golang I installed
        )

func main() {
    _, err := net.Dial("tcp", "localhost:8000") 
    if err == nil {
        fmt.Println("Connection Successful") // can I get net data like headers from this?
        }
    if err != nil {
        fmt.Println("Error on connection...") // This shows the error if it doesn't connect
        fmt.Println(err)
        }
}
