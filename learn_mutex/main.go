    package main
    import (
        "fmt"
        _"math/rand"
        _"strings"
         "sync"
        _"time"
    )


    type Asset map[string]int

    type Money struct {
        mutx sync.Mutex
        v int
    }

    func (m *Money) AddSaving(val int) {

            m.mutx.Lock()

            m.v += val

            defer m.mutx.Unlock()


    }

    var John Money
    var Cash Asset
    var Store int
    var m sync.Mutex

    func main() {

        Cash = make(Asset)
        // Cash["Saving"]=1000
        // Cash["Spending"]=1999
        John = Money{
            v : 0,
        }
        for i := 0; i < 1000; i++ {
              go John.AddSaving(i)     
            
        }
        
        

        
        // increment John based on needed with function struct
        // time.Sleep(500 * time.Millisecond)
        fmt.Printf("%v\n", John.v)


    }