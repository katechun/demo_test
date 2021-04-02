package main


type House struct {
	Material string
	HasFireplace bool
	Floors int
}

type HouseOption func(*House)

func WithConcrete()HouseOption{
	return func(h *House){
		h.Material = "concrete"
	}
}

func WithoutFireplace() HouseOption{
	return func(h *House){
		h.HasFireplace = false
	}
}


func Withfloors(floors int)HouseOption{
	return func(h *House){
		h.Floors = floors
	}
}


func NewHouse(opts ...HouseOption)*House{
	const (
		defaultFloors =2
		defaultHasFireplace = true
		defaultMaterial ="wood"
	)

	h:=&House{
		Material: defaultMaterial,
		HasFireplace: defaultHasFireplace,
		Floors: defaultFloors,
	}

	for _,opt := range opts{
		opt(h)
	}
	return h
}


func main(){
	h := NewHouse(
		WithConcrete(),
		WithoutFireplace(),
		Withfloors(3),
		)

	h=h
}


