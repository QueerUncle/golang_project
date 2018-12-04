/**
 *  2018/11/1  lize
 */
package mysql

import (
  "database/sql"
  "encoding/json"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "app"
)

type sqlDate struct{

  userName string `json:"user_name"`

  password string `json:"password"`

  ip string `json:"ip"`

  port string `json:"port"`

  dbName string `json:"db_name"`

}


var DB *sql.DB

func init(){

  var SqlObj sqlDate

  SqlObj.userName = app.WebSiteConfig.Mysql.Default.Username

  SqlObj.password = app.WebSiteConfig.Mysql.Default.Password

  SqlObj.ip = app.WebSiteConfig.Mysql.Default.Host

  SqlObj.port = app.WebSiteConfig.Mysql.Default.Port

  SqlObj.dbName = app.WebSiteConfig.Mysql.Default.Dbname

  fmt.Println("open the database, mytest")

  //path := strings.Join([]string{userName,":",password,"@tcp(",ip,":",port,")/",dbName,"?charset=utf8"},"")

  path := SqlObj.userName+":"+SqlObj.password+"@tcp("+SqlObj.ip+":"+SqlObj.port+")/"+SqlObj.dbName+"?charset=utf8"

  fmt.Println(path)

  var err error

  DB,err = sql.Open("mysql",path)

  if err!=nil{

    fmt.Println("数据库启动失败",err)

    return

  }

  //设置数据库最大连接
  DB.SetConnMaxLifetime(100)
  //设置数据库最大闲置连接数
  DB.SetMaxIdleConns(10)

  //验证连接
  if err = DB.Ping(); err != nil {

    fmt.Println("数据库连接失败：",err)

    return

  }

  fmt.Println("数据库初始化成功.......")

  fmt.Println(DB,"我是DB")

}

//查
func Select(sql string) error {

  rows,err := DB.Query(sql)

  fmt.Println(rows)

  resData := make(map[string]interface{})

  var dataArr []interface{}

  for rows.Next() {

    single := make(map[string]interface{})

    var title string

    var id string

    var color string

    err = rows.Scan(&title, &id, &color)

    if err !=nil{

      fmt.Println(err)

      //return

    }

    single["title"] = title

    single["id"] = id

    single["color"] = color

    dataArr = append(dataArr,single)

  }

  resData["code"] = 200

  resData["Msg"] = "success"

  resData["data"] = dataArr

  b,_ := json.Marshal(resData)

  //fmt.Fprint(writer,string(b)

  fmt.Print(string(b))

  //c := string(b)

  return nil

}
//增
func Insert(db *sql.DB, name string, sex string, age int) error {

  sen := "INSERT INTO people(name, sex, age) VALUES (?, ?, ?)"

  stmt, err := db.Prepare(sen)

  if err != nil {

    return err

  }

  defer stmt.Close() //预处理记得关闭连接，使用defer在函数return前关闭

  rst, err := stmt.Exec(name, sex, age)

  if err != nil {

    return err

  }

  lastId, err := rst.LastInsertId()

  if err != nil {

    return err

  }

  fmt.Printf("last insert id %d.\n", lastId)

  return nil

}
//改
func Update(db *sql.DB, id int, age int) error {
  sen := "UPDATE people SET age = ? WHERE id = ?"
  stmt, err := db.Prepare(sen)
  if err != nil {
    return err
  }
  defer stmt.Close()
  rst, err := stmt.Exec(age, id)
  if err != nil {
    return err
  }
  rowsAffected, err := rst.RowsAffected()
  if err != nil {
    return err
  }
  fmt.Printf("%d records has been updated.\n", rowsAffected)
  return nil
}

