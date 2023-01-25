package repository

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	// "github.com/Mikaelemmmm/sql2pb/core"
// 	"github.com/Mikaelemmmm/sql2pb/core"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {

// 	DBType := "mysql"
// 	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "127.0.0.1", 3306, "zero-demo")
// 	pkg := "my_package"
// 	goPkg := "./my_package"
// 	table := "*"
// 	serviceName := "usersrv"
// 	fieldStyle := "sqlPb"

// 	DB, err := sql.Open(DBType, connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer DB.Close()
// 	s, err := core.GenerateSchema(DB, table, nil, nil, serviceName, goPkg, pkg, fieldStyle)

// 	if nil != err {
// 		log.Fatal(err)
// 	}

// 	if nil != s {
// 		fmt.Println(s)
// 	}
// }
