package main

import (
    "fmt"
    "math/big"
)

func factorial(n uint64) (r *big.Int) {
    one, bn := big.NewInt(1), new(big.Int).SetUint64(n)
    r = big.NewInt(1)
    if bn.Cmp(one) <= 0 {
        return
    }

    for i := big.NewInt(2); i.Cmp(bn) <= 0; i.Add(i,one) {
        r.Mul(r,i)
    }
    return
}

func main() {
    fmt.Println(factorial(100))
}

