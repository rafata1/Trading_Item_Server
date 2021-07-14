package core

import (
	"context"
	"github.com/jmoiron/sqlx"
	protopack "tradingitem.server/generated_proto"
)

type Server struct {
	protopack.UnimplementedAccountServicesServer
	myDB * sqlx.DB
}

func NewServer() *Server {
	return &Server{
		myDB: connectToDB(),
	}
}

func (s *Server) SignupService(ctx context.Context, req *protopack.SignupReq) ( res *protopack.SignupRes,e error)  {

	statusCode, detail := signUp(req.GetEmail(), req.GetUsername(), req.GetPassword(), req.GetPhoneNumber(),
		req.GetGender(), req.GetDob(), s.myDB)
	return &protopack.SignupRes{
		StatusCode: statusCode,
		Detail: detail,
	}, nil

}

func (s *Server) LoginService(ctx context.Context, req *protopack.LoginReq) (res *protopack.LoginRes, e error) {
	statusCode, detail := login(req.GetEmail(), req.GetPassword(), s.myDB)
	return &protopack.LoginRes{
		StatusCode: statusCode,
		Detail: detail,
	}, nil
}
