package main

import (
	"fmt"
	"strconv"
	"sync"
)

type merger struct {
	mu                    sync.Mutex
	fieldsRBOfOneParentRB map[string]*FieldsRBs // map[descUID-parentRBName]*FieldsRBs
}

type FieldsRBs struct {
	sync.Mutex
	countCls int
	nameSli  []string // slice
	idSli    []int    // slice
}

// FieldsRBs contains all RB from mCls in a parentRB
//
//	type FieldsRBs struct {
//		countCls        int
//		NamesOfFieldRBs []string
//		rbsOfFields     []*ClustersRBs
//	}

func main() {
	var wg sync.WaitGroup
	me := &merger{
		fieldsRBOfOneParentRB: make(map[string]*FieldsRBs),
	}

	// fieldsRBOfOneParentRB["uid"]
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go me.handle(strconv.Itoa(i), i, &wg)
	}
	for j := 1; j <= 100000; j++ {
		wg.Add(1)
		go me.handle(strconv.Itoa(j), j, &wg)
	}
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go me.handle(strconv.Itoa(i), i, &wg)
	}
	for j := 1; j <= 10000; j++ {
		wg.Add(1)
		go me.handle(strconv.Itoa(j), j, &wg)
	}

	wg.Wait()
	fmt.Println("==============> concurrency append")
	fmt.Printf("merger.fieldsRBOfOneParentRB  map: %+v\n", me.fieldsRBOfOneParentRB)

	fmt.Printf("len(nameSli): %d \n", len(me.fieldsRBOfOneParentRB["uid"].nameSli))
	fmt.Printf("len(idSli): %d \n", len(me.fieldsRBOfOneParentRB["uid"].idSli))

	// fmt.Printf("fieldsRBOfOneParentRB[uid]:\n %+v\n", me.fieldsRBOfOneParentRB["uid"])

}

// 仿造handle
func (m *merger) handle(s string, i int, wg *sync.WaitGroup) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.fieldsRBOfOneParentRB["uid"] == nil {
		m.fieldsRBOfOneParentRB["uid"] = &FieldsRBs{}
		m.fieldsRBOfOneParentRB["uid"].countCls = 10000
	}
	// time.Sleep(1 * time.Millisecond)
	m.fieldsRBOfOneParentRB["uid"].Lock()
	m.fieldsRBOfOneParentRB["uid"].nameSli = append(m.fieldsRBOfOneParentRB["uid"].nameSli, s)
	m.fieldsRBOfOneParentRB["uid"].idSli = append(m.fieldsRBOfOneParentRB["uid"].idSli, i)
	m.fieldsRBOfOneParentRB["uid"].Unlock()

	wg.Done()
}
