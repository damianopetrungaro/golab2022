# Steps

1. Add benchmarks for the logger
1. Run and analyze CPU profile
   1. go test --bench=. -cpuprofile=cpu.out 
1. Run and analyze memory profile
   1. go test --bench=. -memprofile=mem.out 
