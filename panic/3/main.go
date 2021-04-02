package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"context"
	"github.com/pkg/errors"
)


func main(){
	g,ctx := errgroup.WithContext(context.Background())
	for i:=0;i<4;i++{
		g.Go(func()error {
		if err := action(ctx);err != nil {
			return err
		}
		if err :=action(ctx);err != nil {
			return err
		}
		if err := action(ctx);err != nil {
			return err
		}
		return nil
	})
	if err := g.Wait();err != nil {
		log.Printf(err.Error())
	}
}
}

func  action(ctx context.Context)error{
	return  errors.New("xxxx")
}
