package main

import (
"fmt"
"time"
)

func Benchmark() {
start := time.Unix(0, 1585930075320274199)
if start.IsZero() {
fmt.Println("Couldn't benchmark start time!")
}
 	fmt.Println("rectangler service restarted in:", time.Since(start).Round(time.Millisecond))
}