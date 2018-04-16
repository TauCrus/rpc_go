package user

import (
	"fmt"
	"net/http"
)

type User struct{}

type UserArgs struct {
	UserName string `json:"user_name"`
}

type UserReply struct {
	Data string `json:"data"`
}

func (u *User) SayHello(r *http.Request,params *UserArgs,reply *UserReply)error{
	userName := params.UserName

	data := fmt.Sprint("Hello ",userName)
	fmt.Println("data",data)
	reply.Data = data
	
	return nil
}

func (u *User) PrintHello(r *http.Request,params *UserArgs,reply *UserReply)error{
	fmt.Println("Hello")

	return nil
}