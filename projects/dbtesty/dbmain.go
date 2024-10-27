package main

import (
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
)
 
func createTable(db *sql.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER NOT NULL
    );`
    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Table created successfully!")
}

func insertData(db *sql.DB, user string, age int) {
    query := "INSERT INTO users(name, age) VALUES(?, ?)"
    _, err := db.Exec(query, user,  age)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Data inserted successfully!")
}

func queryData(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        var age int
        if err := rows.Scan(&id, &name, &age); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }
}

func updateData(db *sql.DB, user string) {
    query := "UPDATE users SET age = ? WHERE name = ?"
    _, err := db.Exec(query, 35, user)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Data updated successfully!")
}

func deleteData(db *sql.DB, user string) {
    query := "DELETE FROM users WHERE name = ?"
    _, err := db.Exec(query, user)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Data deleted successfully!")
}


 func main() {
   fmt.Println("dbtesty")
    // Connect to SQLite database (or create it if it doesn't exist)
    db, err := sql.Open("sqlite3", "./example.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Verify the connection
    if err := db.Ping(); err != nil {
        log.Fatal("Failed to connect:", err)
    }
    fmt.Println("Connected to SQLite database!")

    createTable(db)

    insertData(db, "Alice", 27)

    queryData(db)

    updateData(db, "Alice")

    //deleteData(db)

}
