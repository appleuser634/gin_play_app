module local.package/controller

go 1.19

replace local.package/model => ../model

replace local.package/requests => ../requests

require (
	local.package/model v0.0.0-00010101000000-000000000000
	local.package/requests v0.0.0-00010101000000-000000000000
)

require (
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
)
