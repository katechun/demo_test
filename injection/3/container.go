package main

import "database/sql"

type myRepository struct {
	db *sql.DB
}

type container struct {
	db *sql.DB
	myRepository *myRepository
}


func (c *container)getDB()(*sql.DB,error){
	if c.db == nil {
		db,err := sql.Open("mysql","dsn")
		if err != nil {
			return nil, err
		}
		c.db = db
	}

	return c.db,nil
}


func (c *container)getMyRepository()(*myRepository,error){
	if c.myRepository == nil {
		db,err := c.getDB()
		if err != nil {
			return nil, err
		}
		c.myRepository = &myRepository{db}
	}
	return c.myRepository,nil

}