package main


type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}


func main(){
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Val
	i = s2Ptr
	i=i
}
