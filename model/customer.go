package model

type Customer struct {
}

type CustomerList struct {
	list []*Customer // imutable
}

// REF: https://levelup.gitconnected.com/building-immutable-data-structures-in-go-56a1068c76b2
func (cl *CustomerList) ListAt(index uint) *Customer {
	return cl.list[index]
}
