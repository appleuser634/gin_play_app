# Architecture

- Client: ESP32(ESP-IDF)
- Message API: Go(Gin)
- DB: Sqlite3

### メッセージの送信 

``` mermaid
sequenceDiagram

	title Send Message Flow
	participant client as ESP32
    participant api as Message API
    participant db as Sqlite3
    
	client->>api: POST
	Note left of api: {message:hello,from:mu,to:ha,pass_token:sha256}	
	api->>db: SELECT * FROM USER where USER="from"
	db->>api: 
	Note over api: Check Credential
	api->>db: INSERT into message
```

### メッセージの受信 

``` mermaid
sequenceDiagram

	title Receive Message Flow
	participant client as ESP32
    participant api as Message API
    participant db as Sqlite3
    
	client->>api: GET (Interval 10s-)
	Note left of api: {from:mu,pass_token:sha256,message_id:id}
	api->>db: SELECT * FROM USER where USER="from"
	db->>api: 
	Note over api: Check Credential
	api->>db: SELECT * FROM message where id > "message_id"
	db->>api: 
	Note over api: Get Unread Messages
	api->>client: Return Messages
	Note left of api: {messages[{from:ha,message1,from:mimoc:message2}]}
```
