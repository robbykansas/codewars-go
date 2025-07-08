- have func for check prime
- loop in isPrime works like this, n=11, loop `i=2` to `i*i <= 11` up to `i=3` since `4*4 = 16 > 11`, so 11 is prime, since its not divisible by 2 and 3
- sum the total of prime number
- time complexity is O(n√n)

| n      | √n   | log₂(n) | n√n       | n log n |
| ------ | ---- | ------- | --------- | ------- |
| 10     | 3.16 | \~3.32  | 31.6      | 33.2    |
| 100    | 10   | \~6.64  | 1000      | 664     |
| 1000   | 31.6 | \~9.97  | 31623     | 9966    |
| 10,000 | 100  | \~13.3  | 1,000,000 | 133,000 |


best solve
- Use ProbablyPrime from [math/big](https://pkg.go.dev/math/big)
- time complexity is O(n log³ n) which is better than O(n√n)