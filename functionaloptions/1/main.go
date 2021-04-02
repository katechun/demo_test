package main

import (
	"crypto/tls"
	"time"
)

func main(){
	s1,_:=NewServer("localhost",1024)
	s2,_:=NewServer("localhost",1024,Protocol("udp"))
	s3,_:=NewServer("localhost",1024,Protocol("udp"),MaxConns(1000),Timeout(300*time.Second))
	s1=s1
	s2=s2
	s3=s3


}

type Option func(*Server)

func Protocol(p string)Option{
	return func(s *Server){
		s.protocol=p
	}
}

func Timeout(t time.Duration)Option{
	return func(s *Server){
		s.timeout=t
	}
}

func MaxConns(m int)Option{
	return func(s *Server){
		s.maxconns=m
	}
}

func TLS(t *tls.Config)Option{
	return func(s *Server){
		s.tls=t
	}
}

type Server struct {
	addr string
	port int
	protocol string
	timeout time.Duration
	maxconns int
	tls *tls.Config
}


func NewServer(addr string,port int,options ...func(server *Server))(*Server,error){
	srv := Server{
		addr: addr,
		port: port,
		protocol: "tcp",
		timeout: 30*time.Second,
		maxconns: 1000,
		tls: nil,
	}

	for _,option := range options{
		option(&srv)
	}
	return &srv,nil

}