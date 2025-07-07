Overview of Functions
You have two main functions:

NextBigger(n int) int

PivotSorter(s string) (string, string)

1. NextBigger(n int) int
go
Copy
Edit
d := strconv.Itoa(n)
Converts an integer to a string: O(log n) (number of digits is proportional to log n).

go
Copy
Edit
for i := len(d) - 1; i >= 1; i-- {
    if int(d[i-1]) < int(d[i]) {
        pivot = i - 1
        break
    }
}
Worst case: loop runs over all digits once from the right. So O(log n) time.

go
Copy
Edit
left := d[:pivot]
right := d[pivot:]
String slicing: O(log n) in size, but constant time for slicing.

go
Copy
Edit
p, r := PivotSorter(right)
This is the main call we need to analyze next.

go
Copy
Edit
res := left + p + r
resInt, _ := strconv.Atoi(res)
String concatenation and conversion back to int: again, O(log n) time and space.

2. PivotSorter(s string)
This takes the right substring, which can be up to log n in length.

go
Copy
Edit
pivot := s[0]
for i := 1; i <= len(s)-1; i++ {
    ...
}
Loop over all digits except the first: O(log n) in the worst case.

go
Copy
Edit
digit[0], digit[index] = digit[index], digit[0]
Constant time.

go
Copy
Edit
sl := digit[1:]
sort.Slice(sl, ...)
Sorting a slice of at most log n characters: O(log n * log log n) using Go’s quicksort (which is O(n log n) for n elements).

### Final Time Complexity
NextBigger is O(log n) for most of its steps.

The heaviest part is the sort in PivotSorter, which is O(log n * log log n).

So, the overall time complexity of the code is:

O(log n * log log n)

Where n is the input number, and log n is the number of digits in n.