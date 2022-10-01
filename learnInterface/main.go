package main

import (
    "fmt"
)
// try to implement interface

type Looper interface {
    FruitPrefix(prefix string)
}

type Fruit struct {
  names []string
}

func (f Fruit) FruitPrefix(prefix string) {

    for i := 0; i < len(f.names); i++ {
           if f.names[i][0] == prefix {
            fmt.Println(data)    
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