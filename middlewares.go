package etcdmiddle

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
)

func EtcdMiddleware(servers []string, ctxKey interface{}) negroni.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		cli, _ := Connect(servers)
		context.Set(req, ctxKey, cli)
		next(rw, req)
	}
}
