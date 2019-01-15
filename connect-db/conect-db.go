package main

import (
  "database/sql"
  "log"
  "fmt"
  _ "github.com/mattn/go-sqlite3"

)

type Person struct {
  Name string
  Age  int
}

var DbConection *sql.DB

func main()  {
  DbConection, _ := sql.Open("sqlite3", "./example.sql")
  defer DbConection.Close()
  // db create
  cmd := `CREATE TABLE IF NOT EXISTS person(
            name STRING,
            age  INT)`
  // ↓コマンドうつやつ
  _, err := DbConection.Exec(cmd)
  if err != nil {
    log.Fatalln(err)
  }

  // // insert文
  // cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
  // _, err = DbConection.Exec(cmd, "Kosuke", 21)
  // if err != nil {
  //   log.Fatalln(err)
  // }

  // //update文
  // cmd = "UPDATE person SET age = ? WHERE name = ?"
  // _, err = DbConection.Exec(cmd, 29, "Kosuke")
  // if err != nil {
  //   log.Fatalln(err)
  // }


  // DBのデータを表示させる
  cmd = "SELECT * FROM person"
  rows, _ := DbConection.Query(cmd)
  defer rows.Close()
  var pp []Person
  // ここのfor文でDBから引っ張ってきたデータをPersonのstructに入れている
  for rows.Next() {
    var p Person
    err := rows.Scan(&p.Name, &p.Age)
    if err != nil {
      fmt.Println(err)
    }
    pp = append(pp, p)
  }
  for _, p := range pp {
    fmt.Println(p.Name, p.Age)
  }

  // //特定のレコードだけ持ってくる
  // cmd = "SELECT * FROM person where age = ?"
  // row := DbConection.QueryRow(cmd, 290)
  // var p Person
  // err = row.Scan(&p.Name, &p.Age)
  // if err != nil {
  //   if err == sql.ErrNoRows {
  //     log.Println("No row")
  //   } else {
  //     fmt.Println(err)
  //   }
  // }
  // fmt.Println(p.Name, p.Age)

  // // delete
  // cmd = "DELETE FROM person WHERE age = ?"
  // _, err = DbConection.Exec(cmd, 29)
  // if err != nil {
  //   fmt.Println(err)
  // }
}
