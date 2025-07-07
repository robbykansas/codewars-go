How my thought behind my solve is

- get all string unique into w
- when check the unique string, if happens that there is already inside the w, insert to r, only once and count it
- the third or forth later that already in w, and r, won't count, hence the result
- Time Complexity O(n)

How best solve

- using hash table to get rune for each string
- h[r]++ means, the key is each rune string, the value is count, so they insert unique string into key
- h[r] == 2, if the value of that key is 2, they will increase the count
- Time Complexity O(n)