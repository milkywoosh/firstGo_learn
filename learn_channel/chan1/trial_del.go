package main

import (
    "fmt"
    "strings"
)

func main() {

 done := make(chan struct{})
 histogram := make(map[string]int)
 source := []string{
    "one two three four",
    "two four three",
    "three two one",
}

go func() {
 
    for _, val := range source {
        each_slice := strings.Split(val, " ")
        for _, num := range each_slice {
            value, ok := histogram[num]
            if !ok && value == 0 {
                histogram[num]+=1
            } else {
                histogram[num]+=1
            }
        }
    }
    close(done)
}()

<- done
fmt.Println(histogram)





}