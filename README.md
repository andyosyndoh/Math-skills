# Math-Skills
The program is able to read from a file and print the result of Average, Median, Variance and Standard deviation. In other words, the program is be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:

```console
189
113
121
114
145
110
...

```

This data represents a [statistical population](https://en.wikipedia.org/wiki/Statistical_population): each line contains one value.

To run the program :

```sh
 go run main.go data.txt

```

After reading the file, the program must executes each of the calculations asked above and print the results in the following manner (the following numbers are only examples):

```console
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65

```