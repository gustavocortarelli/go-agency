# go-agency

The meaning of this project is learn GO in phases. On each phase, the complexity and libraries that are being used on 
this project are changes. I started with native queries and `go-chi` library to create my APIs (until release 0.1.1).

On version 2.0, `go-chi` was replaced by fiber, and native queries are also replaced by `GORM`. Moreover, the database
complexity was increased, in order to bring a better comprehension about how GORM manage and query data.

## GORM

Currently, this project is using GORM, and the logger is configured to write all queries for the purpose of understand
how queries are built.