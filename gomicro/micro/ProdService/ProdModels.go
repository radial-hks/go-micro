package ProdService

import "strconv"

type ProdMOdel struct {
	ProdID   int
	ProdName string
}

func NewProd(id int, pname string) *ProdMOdel {
	return &ProdMOdel{
		ProdID:   id,
		ProdName: pname,
	}
}

func NewProdList(n int) []*ProdMOdel {
	ProdList := make([]*ProdMOdel, 0)
	for i := 0; i < n; i++ {
		name := "service" + strconv.Itoa(i)
		ProdList = append(ProdList, NewProd(100+i, name))
	}
	return ProdList

}
