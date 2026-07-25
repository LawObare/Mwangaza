module mwangaza

go 1.25.0

// Dependencies needed:
//   github.com/gin-gonic/gin      — HTTP router and middleware
//   github.com/mattn/go-sqlite3   — SQLite driver (CGO-enabled)
//
// Install with:
//   go get github.com/gin-gonic/gin
//   go get github.com/mattn/go-sqlite3
//
// Alternative (pure Go, no CGO):
//   modernc.org/sqlite

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	golang.org/x/crypto v0.54.0 // indirect
)
