# Architecture

- Client: ESP32(ESP-IDF)
- DB: Sqlite3
- Message API: Go(Gin)

::: mermaid
sequenceDiagram
	participant client as ESP32
    participant api as Message API
    participant db as Sqlite3
    
	client->>api: POST
	Note left of api: {id:123,from:mu}
