# sql-service

## This service features
- Real time query to database using Rest API.
- This can use as backend for BI tools for external data.
- Native Application use this as database.

## Supported databases.
- MySQL
- SQL Lite
- Postgre SQL

## Resquest Body

```
type Request struct {
	SQL  string `json:"sql"` // Database Query Like DML, DDL, DCL
	Type string `json:"type"`
}
``` 
## Response Body
```
type Response struct {
	Data string `json:"data"` // Data in stringify json format
	Msg  string `json:"msg"`  // Query Error or Success messages
}
```

