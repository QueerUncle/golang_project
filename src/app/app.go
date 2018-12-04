/**
 *  2018/10/31  lize
 */
package app

import (

  "config"

  "fmt"

  "strconv"

  "net/http"
)

// 网站基本配置
var WebSiteConfig config.ConfigData

func init(){

  WebSiteConfig = WebSiteConfig.Init()

  fmt.Println(WebSiteConfig.WebPort)

}

func Server(){

  //该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连接请求。
  //该方法有两个参数：第一个参数 addr 即监听地址；第二个参数表示服务端处理程序，通常为空
  //第二个参数为空意味着服务端调用 http.DefaultServeMux 进行处理

  s := strconv.FormatInt(WebSiteConfig.WebPort,10)

  Addr := ":"+s

  http.ListenAndServe(Addr, nil)

}
