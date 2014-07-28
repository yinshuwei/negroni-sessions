# negroni-sessions [![wercker status](https://app.wercker.com/status/988ab53fd546cb198ee5c4c530e0126b/s "wercker status")](https://app.wercker.com/project/bykey/988ab53fd546cb198ee5c4c530e0126b)
Negroni middleware/handler for easy session management.

[API Reference](http://godoc.org/github.com/goincremental/negroni-sessions)

## Usage

~~~ go
package main

import (
  "github.com/codegangsta/negroni"
  "github.com/goincremental/negroni-sessions"
  "github.com/gorilla/context"
  "net/http"
)

func main() {
  n := negroni.Classic()

  store := sessions.NewCookieStore([]byte("secret123"))  
  n.Use(sessions.Sessions("my_session", store))

  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    session := sessions.GetSession(req)
    session.Set("hello", "world")
  })

  n.UseHandler(mux)
  n.Run(":3000")
}

~~~

## Authors
* [David Bochenski](http://github.com/goincremental)
* [Jeremy Saenz](http://github.com/codegangsta)
