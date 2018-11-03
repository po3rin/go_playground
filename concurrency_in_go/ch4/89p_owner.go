package ch4

import "fmt"

func P89() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 4; i++ {
				fmt.Printf("========%+v\n", i)
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Println(result)
		}
		fmt.Println("Done")
	}

	results := chanOwner()
	fmt.Println("yes")
	consumer(results)
}
