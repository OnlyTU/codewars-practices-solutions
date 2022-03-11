### Question :

Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.


### Solution Explanation :

First I returned its value based on the condition of being 0.
I made a string definition called "temp" to keep the repeat count.
I found how many times it was written with a for loop.
I converted every value I found to integer to string, then returned the result.

**Time Complexity : O(n)**