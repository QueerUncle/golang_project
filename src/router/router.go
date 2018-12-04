/**
 *  2018/10/31  lize
 */
package router

import (
  "net/http"
  "pulic2"
)

func init(){

  fsh := http.FileServer(http.Dir("./dist/"))

  http.Handle("/dist/", http.StripPrefix("/dist/", fsh))

  sfsh := http.FileServer(http.Dir("./build/static/"))

  http.Handle("/static/", http.StripPrefix("/static/", sfsh))

  ssh := http.FileServer(http.Dir("./"))

  http.Handle("/statics/", http.StripPrefix("/statics/", ssh))

  http.HandleFunc("/", pulic2.TemplateHandler)

  http.HandleFunc("/c", pulic2.TemplateHandler_c)

  http.HandleFunc("/d", pulic2.TemplateHandler_d)

  http.HandleFunc("/e", pulic2.TemplateHandler_e)

  http.HandleFunc("/queryLabel", pulic2.QueryLabel)

  http.HandleFunc("/addLabel", pulic2.AddLabel)

  http.HandleFunc("/delLabel", pulic2.DelLabel)

  http.HandleFunc("/addAarticleOfLabel", pulic2.AddAarticleOfLabel)

  http.HandleFunc("/delAarticle", pulic2.DelAarticle)

  http.HandleFunc("/queryAarticle", pulic2.QueryAarticle)

}
