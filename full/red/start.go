package main

import (
"fmt"
"time"
)

func Benchmark() {
start := time.Unix(0, 1596482573769986000)
if start.IsZero() {
fmt.Println("Couldn't benchmark start time!")
}
 	fmt.Println("red service restarted in:", time.Since(start).Round(time.Millisecond))
}
