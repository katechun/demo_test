package database


type Store interface{
	Get(ID int)(int,error)
}
