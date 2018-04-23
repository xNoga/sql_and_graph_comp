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

**SQL times are in seconds and graph times are en milliseconds**

|         	| SQL                	| GRAPH             	|
|---------	|--------------------	|-------------------	|
| Depth 1 	| Average: 1.34s <br> Median: 0.98s 	| Average: 8.268ms <br> Median: 8.123ms 	|
| Depth 2 	| Average: 1.26s<br> Median: 1.09s  	| Average: 40.166ms <br> Median: 39.998ms	|
| Depth 3 	| Average: 2.32s <br> Median: 2.17s  	| Average: 310.474ms <br> Median: 305.352ms	|
| Depth 4 	| Average: 5.58s <br> Median: 4.97s  	| Average: 4314.517ms <br> Median: 4254.345ms	|
| Depth 5 	| Average: 12.06s <br> Median: 11.97s  	| Average: 31685.433ms <br> Median: 29452.233ms	|

## Differences
As you can see the two databases are very equal in the first 4 queries. This could be due to whatever tiny technical difference at run time or whatever. In terms of performance they look very equal here. If you were to do every query 20 times and then calculate average and median this might look a little different, I don't know. It could yield a different result between the two.

In the fith query there are quite a big difference between the two. As you can see the sql query is much faster than the graph query.
This could be due to the fact that relational databases operate much faster on large datasets. If you would like to understand it better I highly suggest you read [this](https://stackoverflow.com/questions/13046442/comparison-of-relational-databases-and-graph-databases) StackOverflow post. It gives some great explainations and insights as to why this could be.

## Conclusion
When operating with such a large dataset the relational sql database might be the better choice here if you are looking for pure performance. 
Graph databases does not suffer from a large database per se but does suffer when the query that explores relationships gets really large. 
That is also the reason that data with a lot of specific relations is better suited in a graph database. But query number 5 retrieves almost 500.000 entries which makes the graph query really slow. PostgreSQL is really well optimized and that might be the reason as to why it outperforms the graph database in this specific case. 

If I needed a database that had a lot of very specific relationships I would still choose a graph database. 
The Cypher language used for the queries are also way easier to read in my opinion. Especially when it comes to finding relations between nodes. The queries for the sql looks horrible in this assignment. 
