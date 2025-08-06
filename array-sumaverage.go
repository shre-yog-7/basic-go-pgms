package main
import "fmt"

func main() {
    var nums = [5]int{10, 20, 30, 40, 50}
    var sum int

    for _, value := range nums {
        sum += value
    }

    avg := float64(sum) / float64(len(nums))
    fmt.Println("Sum:", sum)
    fmt.Println("Average:", avg)
}
