package function

import "fmt"

func PrimeSearch() {
	kelipatan9 := 0
	for i := 2; i <= 100; i++ {
		same := 0
		for j := i; j > 1; j-- {
			if i%j == 0 {
				same += 1
			}
		}
		if i%9 == 0 {
			kelipatan9++
			fmt.Printf("Kelipatan ke 9 ke-%d\n", kelipatan9)
		} else if same > 1 {
			fmt.Println(i)
		} else {
			fmt.Printf("Bilangan Prima: %d\n", i)
		}
		same = 0
	}
}
