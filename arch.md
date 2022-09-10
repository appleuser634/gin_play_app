# Architecture

- Client: ESP32(ESP-IDF)
- DB: Sqlite3
- Message API: Go(Gin)


### メッセージの送受信 

::: mermaid
sequenceDiagram
	title Messaging Flow
	participant client as ESP32
    participant api as Message API
    participant db as Sqlite3
    
	client->>api: POST
	Note left of api: {message:hello,from:mu,to:ha,pass_token:sha256}	
	api->>db: SELECT * FROM USER where USER="from"
	db->>api: 
	Note over api: Check Credential
	api->>db: INSERT into message
	
