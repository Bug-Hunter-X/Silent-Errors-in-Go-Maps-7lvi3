```go
package main

import (

        "fmt"

)

func main() {
        m := make(map[string]int)
        m["a"] = 1
        val, ok := m["b"]
        if ok {
                fmt.Println("Value found:", val)
        } else {
                fmt.Println("Key not found") //Handles the case where the key is not present
        }
}
```