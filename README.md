# Go Practice & Learning Repository
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)

This repository is a personal collection of solutions to various coding challenges, implemented in Go. It serves as a log of my practice and learning journey with the Go language, featuring problems from platforms like LeetCode, CodeWars, and others.

## Repository Structure

The solutions are organized into directories based on their source platform or category:

*   **`CodeWars/`**: Solutions for problems from CodeWars.
*   **`LeetCode/`**: Solutions for problems from LeetCode.
*   **`NeetCode/`**: Solutions for problems from the NeetCode roadmap.
*   **`Hexlet/`**: Solutions for exercises from Hexlet.
*   **`T-Bank/`**: Implementations for programming challenges from T-Bank.
*   **`Other/`**: Miscellaneous Go exercises, including algorithm implementations and language feature explorations.
*   **`usecases/`**: Helper structs for test cases used across different solutions.

## Solved Problems

### CodeWars

*   **`digitalRoot.go`**: Calculates the recursive sum of digits of a number until a single-digit number is achieved.
*   **`findOutlier.go`**: Finds the single integer in an array that has a different parity (odd or even) from the rest.
*   **`reverseWords.go`**: Reverses each word in a given string while preserving whitespace.
*   **`splitString.go`**: Splits a string into pairs of two characters. If the string has an odd length, the final pair is completed with an underscore.

### LeetCode

*   **`removeElement.go`**: Removes all instances of a given value in-place from an array and returns the new length of the array.

### NeetCode

*   **`groupAnagrams.go`**: Groups a slice of strings by their anagrams.

### Hexlet

*   **`palyndrome.go`**: Checks if a given string can be rearranged to form a palindrome.
*   **`titleString.go`**: Validates if a string adheres to title case capitalization rules.

### T-Bank Challenges

*   **`kateAndDocs.go`**: Calculates the minimum number of stair flights Kate must climb to deliver documents, given time constraints and floor locations.
*   **`kostyaAndPaper.go`**: Determines the maximum possible increase in the sum of numbers by changing up to `k` digits to '9'.
*   **`sashaAndTests.go`**: Counts the maximum number of distinct test numbers (composed of identical digits) that can be created within a given range `[l, r]`.
*   **`yashaAndSport.go`**: Finds if a single swap of two people can arrange a line such that people with even height are at even positions, and odd height at odd positions.

### Other Exercises

*   **`arraysMatrix.go`**: Defines a `Matrix` type and methods for basic matrix arithmetic (Addition, Subtraction).
*   **`iotaUsing.go`**: Demonstrates the use of `iota` to create a `Direction` enum (North, East, South, West).
*   **`squareRootMIT.go`**: Implements the Newton-Raphson method to find the square root of a number.

## Usage

Many of the solution files contain a `main` function or dedicated test functions with sample cases. You can execute these files directly to see the code in action.

```bash
# Navigate to the file's directory and run it
go run <file_name>.go
```

For example, to run the anagram grouping problem:
```bash
go run NeetCode/groupAnagrams.go
```
Note that some files depend on the `usecases` package within this repository.