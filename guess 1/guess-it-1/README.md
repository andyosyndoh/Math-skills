# Guess It Right

## Overview

This Go program reads a series of numbers from standard input and, for each number, prints out a range in which the next number should fall. The range is calculated based on a specific algorithm that uses previous inputs to determine an appropriate range for the next input.

## Objectives

The main objective of this program is to demonstrate the use of mathematical skills to predict and provide a range for subsequent numbers based on a sequence of inputs. 

## Instructions

### Input

The program expects a sequence of numbers as input, provided one per line. The input represents a graph where the x-axis is the line number (0, 1, 2, 3, ...) and the y-axis is the actual number.

Example input:

```
189
113
121
114
145
110
```


### Output

For each input number, the program outputs the number followed by the calculated range for the next number. The range is displayed as two space-separated values representing the lower and upper bounds.

Example output:
```
189 --> 
113 --> 37 265
121 --> 39 243
114 --> 39 230
145 --> 50 223
110 --> 48 216
```
### Running The Program

Ensure you have Go installed on your system.

```
cd student
go run .
```
After the program launches succesfully, you can now type in your digits one by one on the console.
