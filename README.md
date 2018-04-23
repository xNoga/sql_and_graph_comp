# sql_and_graph_comp

## Setup

The code is split up into two different folders. The sql-folder contains the postgresql code in Go and the graph-folder contains the Neo4J code in Node.
That means you have to have Node and Go installed on your computer in order the run the code.

When both is setup on your machine you simply just nagigate to either folder and run the following commands:

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
