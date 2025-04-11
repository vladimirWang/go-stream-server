package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m

}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// router.POST("/user", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 	w.Write([]byte("post user success"))
	// })
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("login success"))
	})
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)

	http.ListenAndServe(":8000", mh)
}
