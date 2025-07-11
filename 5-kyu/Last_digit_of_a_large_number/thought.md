- Use this [url](https://brilliant.org/wiki/finding-the-last-digit-of-a-power/) or [video](https://www.youtube.com/watch?v=bBvVeZZ3xkI)
- Based on that i made my solve

best solve
- Using math/big
- big.NewInt(0) is initially set big int with 0
- a.SetString(n1, 10) means, a will parse string n1 into base-10 big integers, This allows you to work with very large numbers that don't fit into int64
- a.Exp(a, b, big.NewInt(10)) basically means a^b mod 10
- return it as int