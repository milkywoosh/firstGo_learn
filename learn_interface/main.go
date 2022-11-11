package main

import (
    "fmt"
    _"reflect"
    _"strconv"
)
// try to implement interface

type Looper interface {
    FruitPrefix(prefix string)
}

type Fruit struct {
  names []string
}

func (f Fruit) FruitPrefix(prefix string) {
    var tmp string
    for i := 0; i < len(f.names); i++ {
        tmp = string(f.names[i][0])

         if tmp == prefix {
            fmt.Printf("%s. %s\n", "ok", tmp)
         } else {
            fmt.Printf("%s. %s\n", "false", tmp)
         }
    }
         

        
    
}

func ExecInterface(l Looper) {
    fmt.Println(l)
    l.FruitPrefix("m")
}

func main() {

    burjois := Fruit{
        names : []string{"mangga", "markisa", "melon", "manggis", "jeruk", "apel"},
    }
    ExecInterface(burjois)


}