package main
import "fmt"

func main() {
    var a, b, c int
    fmt.Print("Enter three numbers: ")
    fmt.Scan(&a, &b, &c)

    if a >= b && a >= c {
        fmt.Println("Largest is:", a)
    } else if b >= a && b >= c {
        fmt.Println("Largest is:", b)
    } else {
        fmt.Println("Largest is:", c)
    }
}
