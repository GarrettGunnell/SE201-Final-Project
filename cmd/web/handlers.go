package main

import (
    "html/template"
    "log"
    "net/http"
    "time"

	"github.com/gorilla/websocket"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    http.Redirect(w, r, "/players", 301)
}

func players(w http.ResponseWriter, r *http.Request) {
    cookie, cookieerr := r.Cookie("Callsign")
    if cookieerr != nil {
        log.Println(cookieerr.Error())
    }
    if cookie != nil {
        http.Redirect(w, r, "/game", http.StatusSeeOther)
    }
    if r.Method == http.MethodGet {
        ts, err := template.ParseFiles("ui/html/index.html")
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

        cookiecoordX := http.Cookie{
            Name: "XCoordinate",
            Value: "5",
            Expires: time.Now().AddDate(0, 0, 1),
            Path: "/",
        }

        http.SetCookie(w, &cookiecoordX)
        cookiecoordY := http.Cookie{
            Name: "YCoordinate",
            Value: "5",
            Expires: time.Now().AddDate(0, 0, 1),
            Path: "/",
        }

        http.SetCookie(w, &cookiecoordY)

        http.Redirect(w, r, "/game", 301)
    }

}

func game(w http.ResponseWriter, r *http.Request) {
  ts, err := template.ParseFiles("ui/html/game.html")
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
  w.Write([]byte(callsign))
}

// Websocket stuff below
var upgrader= websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn){
	for{
		messageType, p, err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p ); err != nil{
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}
	
	log.Println("Client Successfully Connected...")

	reader(ws)
}
