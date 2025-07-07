- Max is 359999 then just return 99:59:59 like the problem
- Calculate second is always mod 60, since we need to know what left behind after divided by 60
- If its below 1 hour or 3600 sec, don't need to guess the hour since its below 1 hour, minutes calculate by second divided by 60, since its int, there is no decimal in it
- if its after 1 hour, calculate hours with again second divided by 3600, but for minutes, they need to calculate total - hours before divided by 60
- `%d` means insert int, and `02` is to make sure, the number will be 2 digit
- Time Complexity O(1)

best solve
- assign s to mod 60, m to divide 60, it makes sense
- assign h with m / 60, means, the prior m that we calculate / 60 since 1 hour is 60 minutes
- assign new m with m % 60 since we need to know what left behind after divided for hour
- Time Complexity O(1)