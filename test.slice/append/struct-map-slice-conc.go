package main

import (
	"sync"
)

type merger struct {
	mu                    sync.Mutex
	fieldsRBOfOneParentRB map[string]*FieldsRBs // map[descUID-parentRBName]*FieldsRBs
}

// FieldsRBs contains all RB from mCls in a parentRB
//
//	type FieldsRBs struct {
//		countCls        int
//		NamesOfFieldRBs []string
//		rbsOfFields     []*ClustersRBs
//	}
type FieldsRBs struct {
	sync.Mutex
	countCls int
	nameSli  []string
	idSli    []int
}

func main() {

}
