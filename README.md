# FlexQL
FlexQL is an open-source database lib for flexible data inquiries in relational database. 

Inspired by [GraphQL](https://graphql.org), a query language that allows you to get data for your exact needs. 

The idea is to implement the same logic in relational databases, which can save the cost and increase your db performance. 

## Install 

Import go modules
```go
import "github.com/cymon1997/flexql"
``` 

## Getting Started

Sample: 
```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/cymon1997/flexql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var db *flexql.DB

func init() {
	pgDB, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			"db_user", "db_pass", "127.0.0.1", "5432", "db_table", "disable"))
	if err != nil {
		log.Fatalf("error open db: %v", err)
	}
	db = flexql.NewDB(pgDB)
}

func main() {
	// Table definition
	person := flexql.NewTable("person")
	person.AddField(flexql.Field{Name: "id", Type: flexql.TypeString})
	person.AddField(flexql.Field{Name: "name"}) //default is string

	address := flexql.NewTable("address")
	address.AddField(flexql.Field{Name: "id"})
	address.AddField(flexql.Field{Name: "person_id"})
	address.AddField(flexql.Field{Name: "value"})

	// Retrieve specific data
	fields := []*flexql.Field{
		person.Get("id"),
		person.Get("name"),
		address.Get("value"),
	}

	// Define source table
	tables := []*flexql.Table{
		person, address,
	}

	// Define condition if needed
	cds := []*flexql.Condition{
		{
			A:        person.Get("id"),
			B:        address.Get("person_id"),
			Operator: "=",
		},
	}

	// Execute operation
	result, err := db.SelectContext(context.Background(),
		fields, tables, cds)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse the result
	for _, row := range result {
		fmt.Printf("result: name: %v address: %v",
			row.Get(person.Get("name")),
			row.Get(address.Get("value")))
	}
}
```

## Background

Even if you use GraphQL, the actual query to your relational database is still using full (not compacted) query. 
Also, you need to re-define new query in order to have specific needs for the same table/logic. 

Using this library, solve that problem and provides you flexibility for your queries. 

## How to Contribute 

If you have suggestions, you can create new issues and describe your concerns. 
If you want to help improve this library, you can create an issue (if not exists) to describe your concern and create PR and tag your issues. 

## Creator Note 

I hope this repository can initiate a new perspective on using relational database, I do really appreciate any feedback or suggestion though!

If you have anything, you can contact me at [cymon1997@gmail.com](cymon1997@gmail.com)

Hope this library can help everyone!