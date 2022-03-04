**Question :**

Introduction
The first century spans from the year 1 up to and including the year 100, the second century - from the year 101 up to and including the year 200, etc.

Task
Given a year, return the century it is in.

Examples
1705 --> 18
1900 --> 19
1601 --> 17
2000 --> 20


**Solution Explanation :**

Calculation of century.

If the remainder from the division by 10 and 100 of the given year is not 0, that year falls within the next century.


**Time Complexity : O(n)**