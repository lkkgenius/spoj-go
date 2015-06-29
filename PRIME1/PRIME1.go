package main
import (
    "fmt"
    "math"
)

func is_prime(x int) bool {
    if x <= 1 {
        return false
    }
    n := int(math.Sqrt(float64(x)));
    for i := 2; i <= n; i++ {
        if 0 == x%i {
            return false
        }
    }
    return true
}

func main(){
    var r, m, n int;
    fmt.Scan(&r)
    for i := 1; i <= r; i++ {
        fmt.Scanf("%d %d", &m, &n);
        for k := m; k <= n; k++ {
            if is_prime(k) {
                fmt.Println(k)
            }
        }

        if i != r {
            fmt.Println()
        }
    }
}
