/**
 *  2018/11/5  lize
 */
package pulic2

import (
  "encoding/json"
  "errors"
  "fmt"
  "mysql"
  "net/http"
  "os"
  "os/exec"
  "path/filepath"
  "strconv"
  "strings"
  "time"

  "html/template"
)

type Message struct {

  Code int `json:"code"`

  Msg  string `json:"msg"`

}

type PagingType struct {

  pageNo int  `json:"page_no"`

  pageSize int `json:"page_size"`

  name string `json:"name"`

  title string `json:"title"`

  libid string `json:"libid"`

  sqlparagraph string `json:"sqlparagraph"`

}

type Pspn struct {
  
  start int `json:"pn"`
  
  end int `json:"ps"`
  
  count int `json:"count"`
  
}

//获取路径
func getCurrentPath() (string, error) {

  file, err := exec.LookPath(os.Args[0])

  if err != nil {

    return "", err

  }

  path, err := filepath.Abs(file)

  if err != nil {

    return "", err

  }

  i := strings.LastIndex(path, "/")

  if i < 0 {

    i = strings.LastIndex(path, "\\")

  }

  if i < 0 {

    return "", errors.New(`error: Can't find "/" or "\".`)

  }

  return string(path[0 : i+1]), nil

}

//分页分隔器

func PagingProcessor (p *PagingType) ( psn Pspn){

  fmt.Println(p.sqlparagraph,"我是拼接sql")

  var sql string
  
  //if p.libid == ""{
  //
  //  sql = "SELECT COUNT("+p.title+") FROM "+p.name+""
  //
  //}else{
  //
  //  sql = "SELECT COUNT("+p.title+") FROM "+p.name+" where id = '"+p.libid+"'"
  //
  //}

  if p.sqlparagraph == ""{

    sql = "SELECT COUNT("+p.title+") FROM "+p.name+""

  }else{

    sql = "SELECT COUNT("+p.title+") FROM "+p.name+" where "+p.sqlparagraph+""

  }

  fmt.Println(sql,"我分页sql")

  var qqq int

  count,err := mysql.DB.Query(sql)

  for count.Next() {

    var title string

    err = count.Scan(&title)

    if err !=nil{

      fmt.Println(err)

      return

    }

    qqq,_ = strconv.Atoi(title)

    psn.count = qqq
    
  }

  if p.pageNo-1 == 0 {

    psn.start = 0

    if p.pageSize < qqq {

      psn.end = p.pageSize

    }else{

      //psn.end = qqq+1
      psn.end = p.pageSize

    }

  }else if p.pageNo-1 > 0 {

    psn.start = p.pageSize*(p.pageNo-1)

    if p.pageSize*p.pageNo-1 < qqq{

      psn.end =p.pageSize*p.pageNo-1

    }else{

      //psn.end = qqq+1
      psn.end = p.pageSize

    }

  }else if p.pageNo-1 < 0{

    psn.start = 0

    psn.end = 0

  }

  return psn

}

//解决跨域
func ResponseWithOrigin(w http.ResponseWriter, r *http.Request, code int, json []byte) {

  w.Header().Set("Content-Type", "application/json; charset=utf-8")

  w.Header().Set("Access-Control-Allow-Origin", "*")
  //"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名

  w.WriteHeader(code)

  w.Write(json)

}

//服务模板 index.html
func TemplateHandler(w http.ResponseWriter, r *http.Request) {

  path,err := getCurrentPath()

  if err !=nil{

    fmt.Println("路径获取失败！")

    return

  }

  fmt.Println(path,"我是path")

  t, _ := template.ParseFiles(path+"build/index.html")

  fmt.Println(t.Name())

  t.Execute(w, nil)

}
//服务模板 c
func TemplateHandler_c(w http.ResponseWriter, r *http.Request) {

  path,err := getCurrentPath()

  if err !=nil{

    fmt.Println("路径获取失败！")

    return

  }

  fmt.Println(path,"我是path")

  t, _ := template.ParseFiles(path+"dist/c.html")

  fmt.Println(t.Name())

  t.Execute(w, nil)

}

//服务模板_d.html
func TemplateHandler_d(w http.ResponseWriter, r *http.Request) {

  path,err := getCurrentPath()

  if err !=nil{

    fmt.Println("路径获取失败！")

    return

  }

  fmt.Println(path,"我是path")

  t, _ := template.ParseFiles(path+"dist/d.html")

  fmt.Println(t.Name())

  t.Execute(w, nil)

}

//服务模板_e.html
func TemplateHandler_e(w http.ResponseWriter, r *http.Request) {

  path,err := getCurrentPath();

  if err !=nil{

    fmt.Println("路径获取失败！")

    return

  }

  fmt.Println(path,"我是path")

  t, _ := template.ParseFiles(path+"dist/e.html")

  fmt.Println(t.Name())

  t.Execute(w, nil)

}

//查询标签
func QueryLabel(w http.ResponseWriter, r *http.Request){

  pNo,_ := strconv.Atoi(r.FormValue("pageNo"))

  pSize,_ := strconv.Atoi(r.FormValue("pageSize"))

  var a PagingType

  a.pageNo = pNo

  a.pageSize = pSize

  a.name = "r_column"

  a.title = "title"

  psn := PagingProcessor(&a)

  start := strconv.Itoa(psn.start)

  end := strconv.Itoa(psn.end)

  count := strconv.Itoa(psn.count)

  sql := "SELECT * FROM r_column limit "+start+","+end+""

  fmt.Println(sql)

  rows,err := mysql.DB.Query(sql)

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

      return

    }

    single["title"] = title

    single["id"] = id

    single["color"] = color

    dataArr = append(dataArr,single)

  }

  fmt.Println(dataArr)

  resData["code"] = 200

  resData["Msg"] = "success"

  resData["data"] = dataArr

  resData["pageNo"] = a.pageNo

  resData["pageSize"] = a.pageSize

  resData["count"] = count

  result, err := json.Marshal(resData)

  if err != nil {

    fmt.Println(err)

    return

  }

  ResponseWithOrigin(w, r, http.StatusOK, result)

  return

}

//添加、或修改标签标签
func AddLabel(w http.ResponseWriter, r *http.Request){

  id := r.PostFormValue("id")

  color :=r.PostFormValue("color")

  title := r.PostFormValue("title")

  var sql string

  if id == "" {

    fmt.Println("我没有id")

    sqlid := title+color

    sql = "INSERT INTO r_column (title,id,color) VALUES ('"+title+"','"+sqlid+"','"+color+"')"

  }else{

    fmt.Println("我有id")

    sql = "UPDATE r_column SET title = '"+title+"', color = '"+color+"'WHERE id = '"+id+"'"

  }

  _,err := mysql.DB.Exec(sql)

  resData := make(map[string]interface{})

  if err !=nil{

    resData["code"] = 500

    resData["Msg"] = "error"

  }else{

    resData["code"] = 200

    resData["Msg"] = "success"

    resData["data"] = true

  }

  result, err := json.Marshal(resData)

  if err != nil {

    fmt.Println(err)

    return

  }

  ResponseWithOrigin(w, r, http.StatusOK, result)

  return

}

//删除标签
func DelLabel(w http.ResponseWriter, r *http.Request){

  id:= r.FormValue("id")

  fmt.Println(id)

  if id == ""{

    panic("id为空")

    return

  }

  var sqls string

  var sql string

  resData := make(map[string]interface{})

  sqls = "SELECT name FROM articleInfo WHERE id = '"+id+"'"

  list,err := mysql.DB.Query(sqls)

  var dataArr []interface{}

  for list.Next() {

    single := make(map[string]interface{})

    var name string

    err = list.Scan(&name)

    if err !=nil{

      fmt.Println(err)

      return

    }

    single["name"] = name

    dataArr = append(dataArr,single)

  }

  if len(dataArr) !=0{

    resData["code"] = 200

    resData["Msg"] = "success"

    resData["data"] = "删除失败"

  }else{

    sql = "DELETE FROM r_column WHERE id = '"+id+"'"

    _,err := mysql.DB.Exec(sql)

    if err !=nil{

      resData["code"] = 400

      resData["Msg"] = "error"

    }else{

      resData["code"] = 200

      resData["Msg"] = "success"

      resData["data"] = true

    }

  }

  result, err := json.Marshal(resData)

  if err != nil {

    fmt.Println(err)

    return

  }

  ResponseWithOrigin(w, r, http.StatusOK, result)

  return

}

//给标签下增加文章或者修改文章
func AddAarticleOfLabel(w http.ResponseWriter,r *http.Request){

  labelId := r.PostFormValue("labelId")

  content := r.PostFormValue("content")

  articleid := r.PostFormValue("articleid")

  name := r.PostFormValue("name")

  var sql string

  var time = time.Now().Format("2006-01-02 15:04:05")

  fmt.Println(labelId,content,name,time)

  if articleid == ""{

    sql = "INSERT INTO articleInfo (name,content,id,time) VALUES ('"+name+"','"+content+"','"+labelId+"','"+time+"')"

  }else{

    sql = "UPDATE articleInfo SET name = '"+name+"', content = '"+content+"' , id='"+labelId+"',time = '"+time+"' WHERE articleid = '"+articleid+"'"

  }

  _,err := mysql.DB.Exec(sql)

  resData := make(map[string]interface{})

  if err !=nil{

    fmt.Println(err)

    resData["code"] = 500

    resData["Msg"] = "error"



  }else{

    resData["code"] = 200

    resData["Msg"] = "success"

    resData["data"] = true

  }

  result, err := json.Marshal(resData)

  if err != nil {

    fmt.Println(err)

    return

  }

  ResponseWithOrigin(w, r, http.StatusOK, result)

  return


}

//查询某个标签下或者全部的文章
func QueryAarticle(w http.ResponseWriter,r *http.Request){

 labelId := r.FormValue("id")

 fmt.Println(labelId,"labelId")

 pNo,_ := strconv.Atoi(r.FormValue("pageNo"))

 pSize,_ := strconv.Atoi(r.FormValue("pageSize"))

  var a PagingType

  a.pageNo = pNo

  a.pageSize = pSize

  a.name = "articleInfo"

  a.title = "name"

  a.libid = labelId

  if labelId !=""{

    a.sqlparagraph = "id = '"+labelId+"'"

  }

  psn := PagingProcessor(&a)

  start := strconv.Itoa(psn.start)

  end := strconv.Itoa(psn.end)

  count := strconv.Itoa(psn.count)

 var sql string

 if labelId == ""{

   sql = "SELECT a.*, b.title , b.color FROM articleInfo as a  join r_column as b on a.id=b.id limit "+start+","+end+""
   //sql = "SELECT a.*, b.title , b.color FROM articleInfo as a  join r_column as b on a.id=b.id "

 }else{

   sql = "SELECT a.*, b.title , b.color FROM articleInfo as a  join r_column as b on a.id=b.id where a.id = '"+labelId+"' limit "+start+","+end+""

   //sql = "SELECT a.name, b.title FROM articleInfo ,r_column as a,b  where articleInfo.id = '"+labelId+"' and r_column.id = articleInfo.id"

   //stmr,err := db.DB.Query("SELECT tagId , id , tagName , title , createTime , content FROM article , taglist where article.tagId = taglist.id;")

 }
 //sql = "SELECT a.*, b.* FROM articleInfo as a ,r_column as b  where a.id = '460e5f50d8d911e892b4ddd9b9497828' and a.id = b.id"

 fmt.Println(sql,"我是查询数据sql")

 rows,err := mysql.DB.Query(sql)

 resData := make(map[string]interface{})

 var dataArr []interface{}

 for rows.Next() {

   single := make(map[string]interface{})

   var articleid int

   var name string

   var content string

   var labelid string

   var time string

   var title string

   var color string

   err = rows.Scan(&name, &content, &labelid,&articleid,&time,&title,&color)

   if err !=nil{

     fmt.Println(err)

     return

   }

   single["name"] = name

   single["title"] = title

   single["content"] = content

   single["labelid"] = labelid

   single["articleid"] = articleid

   single["time"] = time

   single["color"] = color

   dataArr = append(dataArr,single)

 }

 if len(dataArr) == 0 {

   var nullAry [0]int

   resData["data"] = nullAry

 }else{

   resData["data"] = dataArr

   resData["pageNo"] = a.pageNo

   resData["pageSize"] = a.pageSize

   resData["count"] = count

 }

 resData["code"] = 200

 resData["Msg"] = "success"

 result, err := json.Marshal(resData)

 if err != nil {

   fmt.Println(err)

   return

 }

 ResponseWithOrigin(w, r, http.StatusOK, result)

 return

}

//删除文章
func DelAarticle(w http.ResponseWriter, r *http.Request){

 articleid := r.FormValue("articleid")

 fmt.Println(articleid)

 if articleid == ""{

   panic("articleid为空")

   return

 }

 var sql string

 sql = "DELETE FROM articleInfo WHERE articleid = '"+articleid+"'"

 resData := make(map[string]interface{})

 _,err := mysql.DB.Exec(sql)

 if err !=nil{

   fmt.Println(err)

   resData["code"] = 400

   resData["Msg"] = "error"

 }else{

   resData["code"] = 200

   resData["Msg"] = "success"

   resData["data"] = true

 }

 result, err := json.Marshal(resData)

 if err != nil {

   fmt.Println(err)

   return

 }

 ResponseWithOrigin(w, r, http.StatusOK, result)

 return

}
