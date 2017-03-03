package etcdmiddle

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/coreos/etcd/client"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/require"
	goctx "golang.org/x/net/context"
)

func setupServeHTTP(t *testing.T) (negroni.ResponseWriter, *http.Request) {
	r := require.New(t)
	req, err := http.NewRequest(http.MethodGet, "http://test-etcd.com/test", nil)
	r.NoError(err)
	return negroni.NewResponseWriter(httptest.NewRecorder()), req
}

func TestMiddleware(t *testing.T) {
	r := require.New(t)
	mware := EtcdMiddleware([]string{"http://127.0.0.1:2379"}, "test_ETCD")
	resp, req := setupServeHTTP(t)

	mware.ServeHTTP(resp, req, func(wrIn http.ResponseWriter, reqIn *http.Request) {
		cIface := context.Get(req, "test_ETCD")
		r.NotNil(cIface)
		cli := cIface.(*client.Client)
		kapi := client.NewKeysAPI(*cli)
		// set "/foo" key with "bar" value
		log.Print("Setting '/foo' key with 'bar' value")
		resp, err := kapi.Set(goctx.Background(), "/foo", "bar", nil)
		if err != nil {
			log.Fatal(err)
		} else {
			// print common key info
			log.Printf("Set is done. Metadata is %q\n", resp)
		}
	})

}
