# Elements of Programming Interviews solutions in Go

This repo contains my solutions in Go for the problems from ["Elements of Programming Interview"](https://elementsofprogramminginterviews.com/). The testing code is provided in https://github.com/stefantds/go-epi-judge, check the Readme there for help to get it running.

For anyone who wants to solve the EPI problems in Go, consider forking https://github.com/stefantds/go-epi-judge as a clean project.

Update: finally completed all the solutions.

```
+--------------------------------+--------------------------------+----------------+
|            CHAPTER             |            PROBLEM             |     TESTS      |
+--------------------------------+--------------------------------+----------------+
| Chapter 04: Primitive Types    | 4.00 Bootcamp: Primitive Types | Passed         |
+                                +--------------------------------+----------------+
|                                | 4.01 Computing the parity of a | Passed         |
|                                | word                           |                |
+                                +--------------------------------+----------------+
|                                | 4.02 Swap bits                 | Passed         |
+                                +--------------------------------+----------------+
|                                | 4.03 Reverse bits              | Passed         |
+                                +--------------------------------+----------------+
|                                | 4.04 Find a closest integer    | Passed         |
|                                | with the same weight           |                |
+                                +--------------------------------+----------------+
|                                | 4.05 Compute product without   | Passed         |
|                                | arithmetical operators         |                |
+                                +--------------------------------+----------------+
|                                | 4.06 Compute quotient without  | Passed         |
|                                | arithmetical operators         |                |
+                                +--------------------------------+----------------+
|                                | 4.07 Compute pow(x,y)          | Passed         |
+                                +--------------------------------+----------------+
|                                | 4.08 Reverse digits            | Passed         |
+                                +--------------------------------+----------------+
|                                | 4.09 Check if a decimal        | Passed         |
|                                | integer is a palindrome        |                |
+                                +--------------------------------+----------------+
|                                | 4.10 Generate uniform random   | Passed         |
|                                | numbers                        |                |
+                                +--------------------------------+----------------+
|                                | 4.11 Rectangle intersection    | Passed         |
+--------------------------------+--------------------------------+----------------+
| Chapter 05: Arrays             | 5.00 Bootcamp: Arrays          | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.01 The Dutch national flag   | Passed         |
|                                | problem                        |                |
+                                +--------------------------------+----------------+
|                                | 5.02 Increment an              | Passed         |
|                                | arbitrary-precision integer    |                |
+                                +--------------------------------+----------------+
|                                | 5.03 Multiply two              | Passed         |
|                                | arbitrary-precision integers   |                |
+                                +--------------------------------+----------------+
|                                | 5.04 Advancing through an      | Passed         |
|                                | array                          |                |
+                                +--------------------------------+----------------+
|                                | 5.05 Delete duplicates from a  | Passed         |
|                                | sorted array                   |                |
+                                +--------------------------------+----------------+
|                                | 5.06 Buy and sell a stock once | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.07 Buy and sell a stock      | Passed         |
|                                | twice                          |                |
+                                +--------------------------------+----------------+
|                                | 5.08 Computing an alternation  | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.09 Enumerate all primes to n | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.10 Permute the elements of   | Passed         |
|                                | an array                       |                |
+                                +--------------------------------+----------------+
|                                | 5.11 Compute the next          | Passed         |
|                                | permutation                    |                |
+                                +--------------------------------+----------------+
|                                | 5.12 Sample offline data       | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.13 Sample online data        | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.14 Compute a random          | Passed         |
|                                | permutation                    |                |
+                                +--------------------------------+----------------+
|                                | 5.15 Compute a random subset   | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.16 Generate nonuniform       | Passed         |
|                                | random numbers                 |                |
+                                +--------------------------------+----------------+
|                                | 5.17 The Sudoku checker        | Passed         |
|                                | problem                        |                |
+                                +--------------------------------+----------------+
|                                | 5.18 Compute the spiral        | Passed         |
|                                | ordering of a 2D array         |                |
+                                +--------------------------------+----------------+
|                                | 5.19 Rotate a 2D array         | Passed         |
+                                +--------------------------------+----------------+
|                                | 5.20 Compute rows in Pascal's  | Passed         |
|                                | Triangle                       |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 06: Strings            | 6.00 Bootcamp: Strings         | Passed         |
+                                +--------------------------------+----------------+
|                                | 6.01 Interconvert strings and  | Passed         |
|                                | integers                       |                |
+                                +--------------------------------+----------------+
|                                | 6.02 Base conversion           | Passed         |
+                                +--------------------------------+----------------+
|                                | 6.03 Compute the spreadsheet   | Passed         |
|                                | column encoding                |                |
+                                +--------------------------------+----------------+
|                                | 6.04 Replace and remove        | Passed         |
+                                +--------------------------------+----------------+
|                                | 6.05 Test palindromicity       | Passed         |
+                                +--------------------------------+----------------+
|                                | 6.06 Reverse all the words in  | Passed         |
|                                | a sentence                     |                |
+                                +--------------------------------+----------------+
|                                | 6.07 The look-and-say problem  | Passed         |
+                                +--------------------------------+----------------+
|                                | 6.08 Convert from Roman to     | Passed         |
|                                | decimal                        |                |
+                                +--------------------------------+----------------+
|                                | 6.09 Compute all valid IP      | Passed         |
|                                | addresses                      |                |
+                                +--------------------------------+----------------+
|                                | 6.10 Write a string            | Passed         |
|                                | sinusoidally                   |                |
+                                +--------------------------------+----------------+
|                                | 6.11 Implement run-length      | Passed         |
|                                | encoding                       |                |
+                                +--------------------------------+----------------+
|                                | 6.12 Find the first occurrence | Passed         |
|                                | of a substring                 |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 07: Linked Lists       | 7.00 Bootcamp: Delete From     | Passed         |
|                                | List                           |                |
+                                +--------------------------------+----------------+
|                                | 7.00 Bootcamp: Insert In List  | Passed         |
+                                +--------------------------------+----------------+
|                                | 7.00 Bootcamp: Search In List  | Passed         |
+                                +--------------------------------+----------------+
|                                | 7.01 Merge two sorted lists    | Passed         |
+                                +--------------------------------+----------------+
|                                | 7.02 Reverse a single sublist  | Passed         |
+                                +--------------------------------+----------------+
|                                | 7.03 Test for cyclicity        | Passed         |
+                                +--------------------------------+----------------+
|                                | 7.04 Test for overlapping      | Passed         |
|                                | lists---lists are cycle-free   |                |
+                                +--------------------------------+----------------+
|                                | 7.05 Test for overlapping      | Passed         |
|                                | lists---lists may have cycles  |                |
+                                +--------------------------------+----------------+
|                                | 7.06 Delete a node from a      | Passed         |
|                                | singly linked list             |                |
+                                +--------------------------------+----------------+
|                                | 7.07 Remove the kth last       | Passed         |
|                                | element from a list            |                |
+                                +--------------------------------+----------------+
|                                | 7.08 Remove duplicates from a  | Passed         |
|                                | sorted list                    |                |
+                                +--------------------------------+----------------+
|                                | 7.09 Implement cyclic right    | Passed         |
|                                | shift for singly linked lists  |                |
+                                +--------------------------------+----------------+
|                                | 7.10 Implement even-odd merge  | Passed         |
+                                +--------------------------------+----------------+
|                                | 7.11 Test whether a singly     | Passed         |
|                                | linked list is palindromic     |                |
+                                +--------------------------------+----------------+
|                                | 7.12 Implement list pivoting   | Passed         |
+                                +--------------------------------+----------------+
|                                | 7.13 Add list-based integers   | Passed         |
+--------------------------------+--------------------------------+----------------+
| Chapter 08: Stacks and Queues  | 8.01 Implement a stack with    | Passed         |
|                                | max API                        |                |
+                                +--------------------------------+----------------+
|                                | 8.02 Evaluate RPN expressions  | Passed         |
+                                +--------------------------------+----------------+
|                                | 8.03 Is a string well-formed?  | Passed         |
+                                +--------------------------------+----------------+
|                                | 8.04 Normalize pathnames       | Passed         |
+                                +--------------------------------+----------------+
|                                | 8.05 Compute buildings with a  | Passed         |
|                                | sunset view                    |                |
+                                +--------------------------------+----------------+
|                                | 8.06 Compute binary tree nodes | Passed         |
|                                | in order of increasing depth   |                |
+                                +--------------------------------+----------------+
|                                | 8.07 Implement a circular      | Passed         |
|                                | queue                          |                |
+                                +--------------------------------+----------------+
|                                | 8.08 Implement a queue using   | Passed         |
|                                | stacks                         |                |
+                                +--------------------------------+----------------+
|                                | 8.09 Implement a queue with    | Passed         |
|                                | max API                        |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 09: Binary Trees       | 9.01 Test if a binary tree is  | Passed         |
|                                | height-balanced                |                |
+                                +--------------------------------+----------------+
|                                | 9.02 Test if a binary tree is  | Passed         |
|                                | symmetric                      |                |
+                                +--------------------------------+----------------+
|                                | 9.03 Compute the lowest common | Passed         |
|                                | ancestor in a binary tree      |                |
+                                +--------------------------------+----------------+
|                                | 9.04 Compute the LCA when      | Passed         |
|                                | nodes have parent pointers     |                |
+                                +--------------------------------+----------------+
|                                | 9.05 Sum the root-to-leaf      | Passed         |
|                                | paths in a binary tree         |                |
+                                +--------------------------------+----------------+
|                                | 9.06 Find a root to leaf path  | Passed         |
|                                | with specified sum             |                |
+                                +--------------------------------+----------------+
|                                | 9.07 Implement an inorder      | Passed         |
|                                | traversal without recursion    |                |
+                                +--------------------------------+----------------+
|                                | 9.08 Compute the kth node in   | Passed         |
|                                | an inorder traversal           |                |
+                                +--------------------------------+----------------+
|                                | 9.09 Compute the successor     | Passed         |
+                                +--------------------------------+----------------+
|                                | 9.10 Implement an inorder      | Passed         |
|                                | traversal with constant space  |                |
+                                +--------------------------------+----------------+
|                                | 9.11 Reconstruct a binary tree | Passed         |
|                                | from traversal data            |                |
+                                +--------------------------------+----------------+
|                                | 9.12 Reconstruct a binary tree | Passed         |
|                                | from a preorder traversal with |                |
|                                | markers                        |                |
+                                +--------------------------------+----------------+
|                                | 9.13 Compute the leaves of a   | Passed         |
|                                | binary tree                    |                |
+                                +--------------------------------+----------------+
|                                | 9.14 Compute the exterior of a | Passed         |
|                                | binary tree                    |                |
+                                +--------------------------------+----------------+
|                                | 9.15 Compute the right sibling | Passed         |
|                                | tree                           |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 10: Heaps              | 10.01 Merge sorted files       | Passed         |
+                                +--------------------------------+----------------+
|                                | 10.02 Sort an                  | Passed         |
|                                | increasing-decreasing array    |                |
+                                +--------------------------------+----------------+
|                                | 10.03 Sort an almost-sorted    | Passed         |
|                                | array                          |                |
+                                +--------------------------------+----------------+
|                                | 10.04 Compute the k closest    | Passed         |
|                                | stars                          |                |
+                                +--------------------------------+----------------+
|                                | 10.05 Compute the median of    | Passed         |
|                                | online data                    |                |
+                                +--------------------------------+----------------+
|                                | 10.06 Compute the k largest    | Passed         |
|                                | elements in a max-heap         |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 11: Searching          | 11.01 Search a sorted array    | Passed         |
|                                | for first occurrence of k      |                |
+                                +--------------------------------+----------------+
|                                | 11.02 Search a sorted array    | Passed         |
|                                | for entry equal to its index   |                |
+                                +--------------------------------+----------------+
|                                | 11.03 Search a cyclically      | Passed         |
|                                | sorted array                   |                |
+                                +--------------------------------+----------------+
|                                | 11.04 Compute the integer      | Passed         |
|                                | square root                    |                |
+                                +--------------------------------+----------------+
|                                | 11.05 Compute the real square  | Passed         |
|                                | root                           |                |
+                                +--------------------------------+----------------+
|                                | 11.06 Search in a 2D sorted    | Passed         |
|                                | array                          |                |
+                                +--------------------------------+----------------+
|                                | 11.07 Find the min and max     | Passed         |
|                                | simultaneously                 |                |
+                                +--------------------------------+----------------+
|                                | 11.08 Find the kth largest     | Passed         |
|                                | element                        |                |
+                                +--------------------------------+----------------+
|                                | 11.09 Find the missing IP      | Passed         |
|                                | address                        |                |
+                                +--------------------------------+----------------+
|                                | 11.10 Find the duplicate and   | Passed         |
|                                | missing elements               |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 12: Hash Tables        | 12.00 Bootcamp: Hash Tables    | Passed         |
+                                +--------------------------------+----------------+
|                                | 12.01 Test for palindromic     | Passed         |
|                                | permutations                   |                |
+                                +--------------------------------+----------------+
|                                | 12.02 Is an anonymous letter   | Passed         |
|                                | constructible?                 |                |
+                                +--------------------------------+----------------+
|                                | 12.03 Implement an ISBN cache  | Passed         |
+                                +--------------------------------+----------------+
|                                | 12.04 Compute the LCA,         | Passed         |
|                                | optimizing for close ancestors |                |
+                                +--------------------------------+----------------+
|                                | 12.05 Find the nearest         | Passed         |
|                                | repeated entries in an array   |                |
+                                +--------------------------------+----------------+
|                                | 12.06 Find the smallest        | Passed         |
|                                | subarray covering all values   |                |
+                                +--------------------------------+----------------+
|                                | 12.07 Find smallest subarray   | Passed         |
|                                | sequentially covering all      |                |
|                                | values                         |                |
+                                +--------------------------------+----------------+
|                                | 12.08 Find the longest         | Passed         |
|                                | subarray with distinct entries |                |
+                                +--------------------------------+----------------+
|                                | 12.09 Find the length of a     | Passed         |
|                                | longest contained interval     |                |
+                                +--------------------------------+----------------+
|                                | 12.10 Compute all string       | Passed         |
|                                | decompositions                 |                |
+                                +--------------------------------+----------------+
|                                | 12.11 Test the Collatz         | Passed         |
|                                | conjecture                     |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 13: Sorting            | 13.01 Compute the intersection | Passed         |
|                                | of two sorted arrays           |                |
+                                +--------------------------------+----------------+
|                                | 13.02 Merge two sorted arrays  | Passed         |
+                                +--------------------------------+----------------+
|                                | 13.03 Computing the h-index    | Passed         |
+                                +--------------------------------+----------------+
|                                | 13.04 Remove first-name        | Passed         |
|                                | duplicates                     |                |
+                                +--------------------------------+----------------+
|                                | 13.05 Smallest                 | Passed         |
|                                | nonconstructible value         |                |
+                                +--------------------------------+----------------+
|                                | 13.06 Render a calendar        | Passed         |
+                                +--------------------------------+----------------+
|                                | 13.07 Merging intervals        | Passed         |
+                                +--------------------------------+----------------+
|                                | 13.08 Compute the union of     | Passed         |
|                                | intervals                      |                |
+                                +--------------------------------+----------------+
|                                | 13.09 Partitioning and sorting | Passed         |
|                                | an array with many repeated    |                |
|                                | entries                        |                |
+                                +--------------------------------+----------------+
|                                | 13.10 Team photo day---1       | Passed         |
+                                +--------------------------------+----------------+
|                                | 13.11 Implement a fast sorting | Passed         |
|                                | algorithm for lists            |                |
+                                +--------------------------------+----------------+
|                                | 13.12 Compute a salary         | Passed         |
|                                | threshold                      |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 14: Binary Search      | 14.00 Bootcamp: Binary Search  | Passed         |
| Trees                          | Trees                          |                |
+                                +--------------------------------+----------------+
|                                | 14.01 Test if a binary tree    | Passed         |
|                                | satisfies the BST property     |                |
+                                +--------------------------------+----------------+
|                                | 14.02 Find the first key       | Passed         |
|                                | greater than a given value in  |                |
|                                | a BST                          |                |
+                                +--------------------------------+----------------+
|                                | 14.03 Find the k largest       | Passed         |
|                                | elements in a BST              |                |
+                                +--------------------------------+----------------+
|                                | 14.04 Compute the LCA in a BST | Passed         |
|                                |                                |                |
+                                +--------------------------------+----------------+
|                                | 14.05 Reconstruct a BST from   | Passed         |
|                                | traversal data                 |                |
+                                +--------------------------------+----------------+
|                                | 14.06 Find the closest entries | Passed         |
|                                | in three sorted arrays         |                |
+                                +--------------------------------+----------------+
|                                | 14.07 Enumerate extended       | Passed         |
|                                | integers                       |                |
+                                +--------------------------------+----------------+
|                                | 14.08 Build a minimum height   | Passed         |
|                                | BST from a sorted array        |                |
+                                +--------------------------------+----------------+
|                                | 14.09 Test if three BST nodes  | Passed         |
|                                | are totally ordered            |                |
+                                +--------------------------------+----------------+
|                                | 14.10 The range lookup problem | Passed         |
|                                |                                |                |
+                                +--------------------------------+----------------+
|                                | 14.11 Add credits              | Passed         |
|                                |                                |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 15: Recursion          | 15.00 Bootcamp: Recursion      | Passed         |
+                                +--------------------------------+----------------+
|                                | 15.01 The Towers of Hanoi      | Passed         |
|                                | problem                        |                |
+                                +--------------------------------+----------------+
|                                | 15.02 Compute all mnemonics    | Passed         |
|                                | for a phone number             |                |
+                                +--------------------------------+----------------+
|                                | 15.03 Generate all             | Passed         |
|                                | nonattacking placements of     |                |
|                                | n-Queens                       |                |
+                                +--------------------------------+----------------+
|                                | 15.04 Generate permutations    | Passed         |
+                                +--------------------------------+----------------+
|                                | 15.05 Generate the power set   | Passed         |
+                                +--------------------------------+----------------+
|                                | 15.06 Generate all subsets of  | Passed         |
|                                | size k                         |                |
+                                +--------------------------------+----------------+
|                                | 15.07 Generate strings of      | Passed         |
|                                | matched parens                 |                |
+                                +--------------------------------+----------------+
|                                | 15.08 Generate palindromic     | Passed         |
|                                | decompositions                 |                |
+                                +--------------------------------+----------------+
|                                | 15.09 Generate binary trees    | Passed         |
+                                +--------------------------------+----------------+
|                                | 15.10 Implement a Sudoku       | Passed         |
|                                | solver                         |                |
+                                +--------------------------------+----------------+
|                                | 15.11 Compute a Gray code      | Passed         |
+--------------------------------+--------------------------------+----------------+
| Chapter 16: Dynamic            | 16.00 Bootcamp: Max Sum        | Passed         |
| Programming                    | Subarray                       |                |
+                                +--------------------------------+----------------+
|                                | 16.00 Bootcamp: Fibonacci      | Passed         |
|                                |                                |                |
+                                +--------------------------------+----------------+
|                                | 16.01 Count the number of      | Passed         |
|                                | score combinations             |                |
+                                +--------------------------------+----------------+
|                                | 16.02 Compute the Levenshtein  | Passed         |
|                                | distance                       |                |
+                                +--------------------------------+----------------+
|                                | 16.03 Count the number of ways | Passed         |
|                                | to traverse a 2D array         |                |
+                                +--------------------------------+----------------+
|                                | 16.04 Compute the binomial     | Passed         |
|                                | coefficients                   |                |
+                                +--------------------------------+----------------+
|                                | 16.05 Search for a sequence in | Passed         |
|                                | a 2D array                     |                |
+                                +--------------------------------+----------------+
|                                | 16.06 The knapsack problem     | Passed         |
|                                |                                |                |
+                                +--------------------------------+----------------+
|                                | 16.07 Building a search index  | Passed         |
|                                | for domains                    |                |
+                                +--------------------------------+----------------+
|                                | 16.08 Find the minimum weight  | Passed         |
|                                | path in a triangle             |                |
+                                +--------------------------------+----------------+
|                                | 16.09 Pick up coins for        | Passed         |
|                                | maximum gain                   |                |
+                                +--------------------------------+----------------+
|                                | 16.10 Count the number of      | Passed         |
|                                | moves to climb stairs          |                |
+                                +--------------------------------+----------------+
|                                | 16.11 The pretty printing      | Passed         |
|                                | problem                        |                |
+                                +--------------------------------+----------------+
|                                | 16.12 Find the longest         | Passed         |
|                                | nondecreasing subsequence      |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 17: Greedy Algorithms  | 17.00 Bootcamp: Greedy         | Passed         |
| and Invariants                 | Algorithms and Invariants      |                |
+                                +--------------------------------+----------------+
|                                | 17.01 Compute an optimum       | Passed         |
|                                | assignment of tasks            |                |
+                                +--------------------------------+----------------+
|                                | 17.02 Schedule to minimize     | Passed         |
|                                | waiting time                   |                |
+                                +--------------------------------+----------------+
|                                | 17.03 The interval covering    | Passed         |
|                                | problem                        |                |
+                                +--------------------------------+----------------+
|                                | 17.03 Invariant Bootcamp: Two  | Passed         |
|                                | Sum                            |                |
+                                +--------------------------------+----------------+
|                                | 17.04 The 3-sum problem        | Passed         |
|                                |                                |                |
+                                +--------------------------------+----------------+
|                                | 17.05 Find the majority        | Passed         |
|                                | element                        |                |
+                                +--------------------------------+----------------+
|                                | 17.06 The gasup problem        | Passed         |
|                                |                                |                |
+                                +--------------------------------+----------------+
|                                | 17.07 Compute the maximum      | Passed         |
|                                | water trapped by a pair of     |                |
|                                | vertical lines                 |                |
+                                +--------------------------------+----------------+
|                                | 17.08 Compute the largest      | Passed         |
|                                | rectangle under the skyline    |                |
+--------------------------------+--------------------------------+----------------+
| Chapter 18: Graphs             | 18.01 Search a maze            | Passed         |
+                                +--------------------------------+----------------+
|                                | 18.02 Paint a Boolean matrix   | Passed         |
+                                +--------------------------------+----------------+
|                                | 18.03 Compute enclosed regions | Passed         |
+                                +--------------------------------+----------------+
|                                | 18.04 Deadlock detection       | Passed         |
+                                +--------------------------------+----------------+
|                                | 18.05 Clone a graph            | Passed         |
+                                +--------------------------------+----------------+
|                                | 18.06 Making wired connections | Passed         |
+                                +--------------------------------+----------------+
|                                | 18.07 Transform one string to  | Passed         |
|                                | another                        |                |
+                                +--------------------------------+----------------+
|                                | 18.08 Team photo day---2       | Passed         |
+--------------------------------+--------------------------------+----------------+
| Chapter 24: Honors Class       | 24.01 Compute the greatest     | Passed         |
|                                | common divisor                 |                |
+                                +--------------------------------+----------------+
|                                | 24.02 Find the first missing   | Passed         |
|                                | positive entry                 |                |
+                                +--------------------------------+----------------+
|                                | 24.03 Buy and sell a stock at  | Passed         |
|                                | most k times                   |                |
+                                +--------------------------------+----------------+
|                                | 24.04 Compute the maximum      | Passed         |
|                                | product of all entries but one |                |
+                                +--------------------------------+----------------+
|                                | 24.05 Compute the longest      | Passed         |
|                                | contiguous increasing subarray |                |
+                                +--------------------------------+----------------+
|                                | 24.06 Rotate an array          | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.07 Identify positions       | Passed         |
|                                | attacked by rooks              |                |
+                                +--------------------------------+----------------+
|                                | 24.08 Justify text             | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.09 Implement list zipping   | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.10 Copy a postings list     | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.11 Compute the longest      | Passed         |
|                                | substring with matching parens |                |
+                                +--------------------------------+----------------+
|                                | 24.12 Compute the maximum of a | Passed         |
|                                | sliding window                 |                |
+                                +--------------------------------+----------------+
|                                | 24.13 Compute fair bonuses     | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.14 Search a sorted array of | Passed         |
|                                | unknown length                 |                |
+                                +--------------------------------+----------------+
|                                | 24.15 Search in two sorted     | Passed         |
|                                | arrays                         |                |
+                                +--------------------------------+----------------+
|                                | 24.16 Find the kth largest     | Passed         |
|                                | element---large n, small k     |                |
+                                +--------------------------------+----------------+
|                                | 24.17 Find an element that     | Passed         |
|                                | appears only once              |                |
+                                +--------------------------------+----------------+
|                                | 24.18 Find the line through    | Passed         |
|                                | the most points                |                |
+                                +--------------------------------+----------------+
|                                | 24.19 Convert a sorted doubly  | Passed         |
|                                | linked list into a BST         |                |
+                                +--------------------------------+----------------+
|                                | 24.20 Convert a BST to a       | Passed         |
|                                | sorted doubly linked list      |                |
+                                +--------------------------------+----------------+
|                                | 24.21 Merge two BSTs           | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.22 Implement regular        | Passed         |
|                                | expression matching            |                |
+                                +--------------------------------+----------------+
|                                | 24.23 Synthesize an expression | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.24 Count inversions         | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.25 Draw the skyline         | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.26 Measure with defective   | Passed         |
|                                | jugs                           |                |
+                                +--------------------------------+----------------+
|                                | 24.27 Compute the maximum      | Passed         |
|                                | subarray sum in a circular     |                |
|                                | array                          |                |
+                                +--------------------------------+----------------+
|                                | 24.28 Determine the critical   | Passed         |
|                                | height                         |                |
+                                +--------------------------------+----------------+
|                                | 24.29 Max Square Submatrix     | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.29 Max Submatrix            | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.30 Implement Huffman coding | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.31 Trapping water           | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.32 The heavy hitter problem | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.33 Find the longest         | Passed         |
|                                | subarray with sum constraint   |                |
+                                +--------------------------------+----------------+
|                                | 24.34 Road network             | Passed         |
+                                +--------------------------------+----------------+
|                                | 24.35 Test if arbitrage is     | Passed         |
|                                | possible                       |                |
+--------------------------------+--------------------------------+----------------+
|                                              TOTAL              | 218/218 (100%) |
+--------------------------------+--------------------------------+----------------+
```
