- get maximal number with calculate all the number
- Using for loop to get all data
- When in index 1 or 2 or later, just decrease the max from [i-1] or prior number

best solve
- Starts from the end of ls and moves backward.
- At each step, v[i] is the sum of ls[i] and v[i+1].
- This avoids re-summing the whole slice every time — it reuses the previous sum.