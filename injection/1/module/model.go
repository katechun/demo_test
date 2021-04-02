package module

import (
	"fmt"
	"injection/1/database"
)
type Module struct {
	Store database.Store
}



func (m *Module)GetNumber(ID int)error{
	result,err := m.Store.Get(ID)
	if err!= nil {
		return err
	}
	if result >10{
		return fmt.Errorf("result toot high: %d",result)
	}

	return nil
}