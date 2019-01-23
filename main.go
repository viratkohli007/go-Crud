package main
import (
         "net/http"
         "html/template"
         "fmt"
         "database/sql"
       _ "github.com/lib/pq"
       )

type homest struct{}

type formst struct{
    Age string
    FirstName string
    LastName  string
    Email string
}

type formst2 struct{
    Id string
    Age string
    FirstName string
    LastName  string
    Email string
}

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "test123"
  dbname   = "postgres"
         )

func main() {

	const Port = os.Getenv("PORT")
    http.HandleFunc("/", Home)
    http.HandleFunc("/form", Form)
    http.HandleFunc("/display", Display)
    http.HandleFunc("/list", List)
    err := http.ListenAndServe(Port, nil)
    if err != nil{
    	fmt.Println(err)
    }
}

func Home(w http.ResponseWriter, r *http.Request) {

   t, _ := template.ParseFiles("home.html")
   t.Execute(w, "")
}

func Form(w http.ResponseWriter, r *http.Request) {
   t, _ := template.ParseFiles("form.html")
   t.Execute(w, "")
}

func Display(w http.ResponseWriter, r *http.Request) {
    db2 := dbconn()
     formobj := new(formst)
     formobj.FirstName = r.FormValue("first_name")
     formobj.LastName = r.FormValue("last_name")
     formobj.Age = r.FormValue("age")
     formobj.Email = r.FormValue("email")

 insert:= `insert into users(age, first_name, last_name, email) values($1, $2, $3, $4)`
 _, err := db2.Exec(insert,formobj.Age, formobj.FirstName, formobj.LastName, formobj.Email)
 if err != nil{
  panic (err)
 }

  t, _ := template.ParseFiles("display.html")
  t.Execute(w,formobj)
}

func dbconn() *sql.DB{

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    // connstr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil{
      fmt.Println(err)
    }
    //defer db.Close()
    //fmt.Println(db)

     fmt.Println("successfully connected")
     return db
}

func List(w http.ResponseWriter, r *http.Request) {

     db2 := dbconn()
     read := `select * from users`
     rows, err := db2.Query(read)
     if err != nil{
      panic(err)
     }

    var Table formst2
    var Table2 []formst2

    for rows.Next() {

      err = rows.Scan(&Table.Id, &Table.Age, &Table.FirstName, &Table.LastName, &Table.Email)
      if err != nil{
        panic(err)
      }
       // fmt.Println("ID | AGE | FirstName | LastName | Email")
       // fmt.Printf("%v | %v | %v | %v | %v\n", Table.Id, Table.Age, Table.FirstName, Table.LastName, Table.Email)
       Table2 = append(Table2, Table)
    }

     t, _ := template.ParseFiles("list.html")
     t.Execute(w, Table2)
}
