package main

import (
    "html/template"
    "log"
    "net/http"
    "time"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    http.Redirect(w, r, "/players", 301)
}

func players(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        ts, err := template.ParseFiles("ui/html/index.tmpl")
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
            return
        }
        err = ts.Execute(w, nil)
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }
    } else if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Failed to parse form", 500)
            return
        }
        callsign := r.Form.Get("callsign")

        cookie := http.Cookie{
            Name: "Callsign",
            Value: callsign,
            Expires: time.Now().AddDate(0, 0, 1),
            Path: "/",
        }
        http.SetCookie(w, &cookie)

        log.Println("Callsign: " + callsign)
        http.Redirect(w, r, "/game", 301)
    }

}

func game(w http.ResponseWriter, r *http.Request) {
  ts, err := template.ParseFiles("ui/html/game.tmpl")
  if err != nil {
      log.Println(err.Error())
      http.Error(w, "Internal Server Error", 500)
      return
  }

  err = ts.Execute(w, nil)
  if err != nil {
      log.Println(err.Error())
      http.Error(w, "Internal Server Error", 500)
  }

  cookie, cookieerr := r.Cookie("Callsign")
  if cookieerr != nil {
      log.Println(cookieerr.Error())
      http.Error(w, "Internal Server Error: Could not obtain callsign from cookie", 500)
      return
  }

  callsign := cookie.Value
  log.Println(callsign)
  //w.Write([]byte(callsign))
}
