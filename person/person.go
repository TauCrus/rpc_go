package person

import (
	"net/http"
	"rpc_go/sql"
	"log"
)

type Person struct{}

type PersonArgs struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type PersonData struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type PersonReply struct {
	Data []PersonData
}
type PersonReplyInsert struct {
	Id int `json:"id"`
}
func (p *Person) QueryAll(r *http.Request,params *PersonArgs,reply *PersonReply) error{

	rows, err := sql.DB.Query(`SELECT first_name,last_name FROM t_person`)
	if nil != err{
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		var tmp PersonData
		rows.Scan(&tmp.FirstName,&tmp.LastName)

		reply.Data = append(reply.Data,tmp)
	}

	return nil
}
func (p *Person) Insert(r *http.Request,params *PersonArgs,reply *PersonReplyInsert) error{

	res, err := sql.DB.Exec(`INSERT INTO test.t_person(first_name,last_name) VALUES (?,?)`,params.FirstName,params.LastName)
	if nil != err{
		log.Fatal(err.Error())
	}

	id, err := res.LastInsertId()
	if nil != err{
		log.Fatal(err.Error())
	}

	reply.Id = int(id)

	return nil
}
