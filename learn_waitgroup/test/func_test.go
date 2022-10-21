// using sleep is not IDEAL for running goroutine bcs it doesnt represent 'real condition'
/*
	running golang test
	--> go test
	--> go test -v
	--> go test -v -run TestNamaFunction
	
	nice reference: 
		- https://stackoverflow.com/questions/36056615/what-is-the-advantage-of-sync-waitgroup-over-channels

*/


func RunAsync(group *sync.WaitGroup, i int, flag string) {
	defer group.Done()

		

	group.Add(1)

	if flag == "go" {
		fmt.Printf("%v\n", i)
	} else {
		fmt.Printf("%v\n", "no")
	}	
	time.Sleep(100 * time.Millisecond)	


}


func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}
	// var wg sync.WaitGroup
	// teswg := *wg{}
	for i:=0; i<20; i++ {
		go RunAsync(group, i, "no")
	}
	group.Wait()
	fmt.Println("selesai")
}

// ************************************************************************

func GoodExampleWG() {
        var wg sync.WaitGroup
      	var total = 0
        for i:=0; i<100; i++ {
                // Increment the WaitGroup counter.
                wg.Add(1)
                // Launch a goroutine to fetch the URL.
                go func(n int) {
                        // Decrement the counter when the goroutine completes.
                        defer wg.Done()
                        // Fetch the URL.
                        // fmt.Printf("%v\n", n)
                        total+=1

                }(i)
        }
        // Wait for all HTTP fetches to complete.
        wg.Wait()
        fmt.Println(total)
}

// ************************************************************************

func GoodExampleWG_1() {
        var wg sync.WaitGroup
        var total = 0
		var cek = 0



        for i:=0; i<100; i++ {
                // Increment the WaitGroup counter.
                wg.Add(2)
                // Launch a goroutine to fetch the URL.
                go func(n int) {
                        // Decrement the counter when the goroutine completes.
                        defer wg.Done()
                        // Fetch the URL.
                        // fmt.Printf("%v\n", n)
                        total+=1


                }(i)
                 go func(n int) {
                        // Decrement the counter when the goroutine completes.
                        defer wg.Done()
                        // Fetch the URL.
                        // fmt.Printf("%v\n", n)
                        cek+=1


                }(i)
        }
        // Wait for all HTTP fetches to complete.
        wg.Wait()
        fmt.Println("total: ", total)
        fmt.Println("cek: ", cek)
}



// ************************************************************************
func GoodExampleWG_2() {
        var wg sync.WaitGroup
        var total = 0
		var cek = 0

		caller_total := func (n int, val_ref *int, wg *sync.WaitGroup) {
			   defer wg.Done()
                        *val_ref+=1
		}
		caller_cek := func (n int, val_ref *int, wg *sync.WaitGroup) {
			   defer wg.Done()
                        *val_ref+=1
		}

        for i:=0; i<100; i++ {
                // Increment the WaitGroup counter.
                wg.Add(2)
                // Launch a goroutine to fetch the URL.
              	go caller_total(i, &total, &wg)
              	go caller_cek(i, &cek, &wg)
        }
        // Wait for all HTTP fetches to complete.
        wg.Wait()
        fmt.Println("total: ", total)
        fmt.Println("cek: ", cek)
}

func main() {
	for i:=0; i<10; i++ {
		GoodExampleWG()
	}
}

// ************************************************************************

