module tui

go 1.19

replace local.package/model => ../model

require github.com/rivo/tview v0.0.0-20220916081518-2e69b7385a37

require local.package/model v0.0.0-00010101000000-000000000000

require (
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.4.1-0.20210905002822-f057f0a857a1 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2 // indirect
	golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d // indirect
	golang.org/x/text v0.3.7 // indirect
)
