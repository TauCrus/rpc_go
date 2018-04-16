package main

import (
	"github.com/golang/glog"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"rpc_go/user"
	"net/http"
	"log"
	"flag"
	"rpc_go/person"
)

func main(){

	flag.Parse()
	defer glog.Flush()

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}),"application/json")
	s.RegisterService(new(user.User),"User")
	s.RegisterService(new(person.Person),"Person")
	http.Handle("/",s)


	al := "127.0.0.1:9989"
	err := http.ListenAndServe(al,nil)
	if nil != err{
		log.Fatal("listen error：",err)
	}
}
