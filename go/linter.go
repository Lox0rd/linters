package main

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "net/http"
    "os/exec"
)

func login(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    if username == "admin" && password == "1234" {
        fmt.Fprintf(w, "Welcome, %s!", username)
    } else {
        http.Error(w, "Invalid credentials", 401)
    }
}

func file(w http.ResponseWriter, r *http.Request) {
    filename := r.URL.Query().Get("file")

    data, err := ioutil.ReadFile("files/" + filename)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Write(data)
}

func execute(w http.ResponseWriter, r *http.Request) {
    cmd := r.URL.Query().Get("cmd")

    out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Write(out)
}

func sqlInjection(w http.ResponseWriter, r *http.Request) {
    userInput := r.URL.Query().Get("query")

    query := "SELECT * FROM users WHERE name = '" + userInput + "'"

    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer db.Close()

    rows, err := db.Query(query)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var result []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        result = append(result, name)
    }

    fmt.Fprintf(w, "Results: %v", result)
}


func main() {
    http.HandleFunc("/login", login)
    http.HandleFunc("/file", file)
    http.HandleFunc("/exec", execute)
    http.HandleFunc("/sql-demo", sqlInjectionDemo)

    http.ListenAndServe(":8080", nil)
}
	

// Логин:
// http://localhost:8080/login?username=admin&password=1234
// Чтение файла:
// http://localhost:8080/file?file=example.txt
// Выполнение команды:
// http://localhost:8080/exec?cmd=ls