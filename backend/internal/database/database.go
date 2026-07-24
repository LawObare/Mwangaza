// Package database manages the SQLite connection and schema lifecycle.
//
// Connect() opens a SQLite database at the given path using
// github.com/mattn/go-sqlite3 (CGO-enabled driver).
//
// Call pattern:
//   db, err := database.Connect("./data/lakenet.db")
//   defer db.Close()
//
// The *sql.DB handle is passed to handlers via routes.Setup(db)
// and stored in handler/service structs for queries.
package database