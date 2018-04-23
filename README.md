# sql_and_graph_comp

## Setup

The code is split up into two different folders. The sql-folder contains the postgresql code in Go and the graph-folder contains the Neo4J code in Node.
That means you have to have Node and Go installed on your computer in order the run the code.

When both is setup on your machine you simply just navigate to either folder and run the following commands:

/graph
```
node neo4j.js
```
/sql
```
go test
```

## Execution time
I have chosen to not run every query 20 times as that would take forever. I guess the point of the assignment is to look at the time difference which we can easily do with only running every query 5 times or so.

|         	| SQL                	| GRAPH             	|
|---------	|--------------------	|-------------------	|
| Depth 1 	| Average: 1.34s <br> Median: 0.98s 	| Average: 8.268ms <br> Median: 8.123ms 	|
| Depth 2 	| Average: 1.26s<br> Median: 1.09s  	| Average: 40.166ms <br> Median: 39.998ms	|
| Depth 3 	| Average: 2.32s <br> Median: 2.17s  	| Average: 310.474ms <br> Median: 305.352ms	|
| Depth 4 	| Average: 5.58s <br> Median: 4.97s  	| Average: 4314.517ms <br> Median: 4254.345ms	|
| Depth 5 	| Average: 12.06s <br> Median: 11.97s  	| Average: 31685.433ms <br> Median: 29452.233ms	|
