## Lottery-Counter
The following program renders the selection of millions of bets in a lottery.

winning bet: 5

| Numbers matching | Winners |
| ------------- |:-------------:
| 5             | 0             |
| 4             | 12            |
| 3             | 818           |
| 2             | 22613         |

A bet is made up of 5 distinct numbers from 1 to 90. 
The lottery reward then depends on how many of the picks match with the ones selected from a bet picks. 
A player wins if he/she has 2, 3, 4 or 5 matching numbers.


### Constraints:
1. A bet is a set of unique values
2. A bet pick is bound between 1 and 90
3. A bet is a set with a limit from 2 picks to 5 picks
4. The report shall be generated within a couple of seconds after picking the winner number.
5. The bets counts can be assumed between 5 million bets and 10 million.

### Libraries outside of Go stdlib used:
#### bitset: https://github.com/bits-and-blooms/bitset
   The type struct `Counter` in the `library` package makes use of the bitset from the bits-and-blooms package as a `s` buffer, to allocate all the bets records for each number pick bound between 1 and 90.

### Benchmarks

Time Complexity: O(X * N) where X is the length of the bets and N is the maximum number of bets, which can be assumed to be 10million maximum.

Space Complexity: O(M * N) where M is the length of the counter slice that holds the bitset elements. This can be assumed to be 90. 
N is the maximum number of bets, which can be assumed to be 10million.

```shell
PS C:\Users\gnusosa\lottery-counter> .\main.exe data.txt
2025/03/10 19:30:33 record on line 1240016: wrong number of fields
2025/03/10 19:30:33 record on line 1240028: wrong number of fields
Invalid bet found: [5 34 24 42]
Invalid bet found: [85 21 67 93 2]
Invalid bet found: [10 85 69 18446744073709551609 50]
READY
11 12 13 14 15
Winning bet:
[11 12 13 14 15]
2025/03/10 19:30:39 took: 85 ms
2025/03/10 19:30:39
5: 1
4: 99
3: 8174
2: 225397
```

```shell
PS C:\Users\gnusosa\lottery-counter> .\main.exe data.txt
2025/03/10 19:30:33 record on line 1240016: wrong number of fields
2025/03/10 19:30:33 record on line 1240028: wrong number of fields
Invalid bet found: [5 34 24 42]
Invalid bet found: [85 21 67 93 2]
Invalid bet found: [10 85 69 18446744073709551609 50]
READY
11 12 13 14 15
Winning bet:
[11 12 13 14 15]
2025/03/10 19:30:39 took: 75 ms
2025/03/10 19:30:39
5: 1
4: 99
3: 8174
2: 225397
```


On a Thinkpad X1, I saw rates from 75ms to 85ms.
This is a good result given the fact that there is no eccentric overcomplicated algorithm, but a classic divide-and-conquer (a lá MapReduce) concurrent funnel with go routines. This is a vainilla Go project using only the stdlib. This is a big win.

> Rule 4. Fancy algorithms are buggier than simple ones, and they're much harder to implement. Use simple algorithms as well as simple data structures.
> Rob Pike's 5 Rules of Programming

### How to run?
- `make run-with-file`
- `make file=<file_name_at_root> run` to use a custom file
