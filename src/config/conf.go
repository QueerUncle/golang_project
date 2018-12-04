/**
 *  2018/10/31  lize
 */

package config

import (

  "encoding/json"
  "fmt"

  "io/ioutil"

  "log"

)

type tlsData struct {
  Open int64 `json:"open"`
  Domain string `json:"domain"`
  Email string `json:"email"`
}
type mysqlSetOpt struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Host string `json:"host"`
  Port string `json:"port"`
  Dbname string `json:"dbname"`
  Socket string `json:"socket"`
}
type mysqlData struct {
  RW int64 `json:"RW"`
  Default mysqlSetOpt `json:"default"`
  Write mysqlSetOpt `json:"write"`
  Read mysqlSetOpt `json:"read"`
}
type RedisData struct {
  Host string `json:"host"`
  Port string `json:"port"`
  Dbname string `json:"dbname"`
  Password string `json:"password"`
}
type CasData struct {
  Server string `json:"server"`
  WhiteList []string `json:"whiteList"`
}


type ConfigData struct {
  WebTitle string `json:"webTitle"`
  TemplateUrl string `json:"templateUrl"`
  StaticFilePath string `json:"staticFilePath"`
  IsAllStatic int64 `json:"isAllStatic"`
  WebPort int64 `json:"webPort"`
  TLS tlsData `json:"tls"`
  Mysql mysqlData `json:"mysql"`
  Redis RedisData `json:"redis"`
  SessionType string `json:"sessionType"`
  SessionName string `json:"sessionName"`
  SessionLifeTime int64 `json:"sessionLifeTime"`
  SessionRedis RedisData `json:"sessionRedis"`
  Cas CasData `json:"cas"`
  Custom map[string]interface{} `json:"custom"`
}

func (_self ConfigData) Init() ConfigData{

  log.Println("读取配置文件 ...\n")

  var ret ConfigData

  filePath := ""

  cont, err := ioutil.ReadFile(filePath + "config.json")

  if err!=nil {

    log.Fatal(err)

  }

  if err := json.Unmarshal(cont, &ret); err != nil {

    log.Fatal(err)

  }

  fmt.Println("文件读取成功",ret)

  return ret

}
