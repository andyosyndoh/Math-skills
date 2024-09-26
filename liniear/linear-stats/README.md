# Linear-Stats
The program is able to read from a file and print the Linier Regression Linr formula and the Pearson Correlation Coefficient of the data provided. In other words, the program is be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:

```console
189
113
121
114
145
110
...

```
This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

### *IMPORTANT*

*Make sure the file with the data is in the root directory i.e Linear Stats*

To run the program :

`cd cmd`

` go run main.go data.txt`

After reading the file, the program executes each of the calculations asked above and print the results in the following manner (the following expressions are only examples):

```console
Linear Regression Line: y = 0.001174x + -2252.335949
Pearson Correlation Coefficient: 0.0026653118

```

## Notion

[Linear Regression](https://en.wikipedia.org/wiki/Linear_regression)

[Pearson Correlation Coefficient](https://en.wikipedia.org/wiki/Pearson_correlation_coefficient)