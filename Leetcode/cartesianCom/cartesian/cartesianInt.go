// https://github.com/schwarmco/go-cartesian-product
// 笛卡尔积

package cartesian

import "sync"

// IterInt takes interface-slices and returns a channel, receiving cartesian products
func IterInt(params ...[]int) chan []int {
	// create channel
	c := make(chan []int)
	// create waitgroup
	var wg sync.WaitGroup
	// call iterator
	wg.Add(1)
	iterateInt(&wg, c, []int{}, params...)
	// call channel-closing go-func
	go func() { wg.Wait(); close(c) }()
	// return channel
	return c
}

// private, recursive Iteration-Function
func iterateInt(wg *sync.WaitGroup, channel chan []int, result []int, params ...[]int) {
	// dec WaitGroup when finished
	defer wg.Done()
	// no more params left?
	if len(params) == 0 {
		// send result to channel
		channel <- result
		return
	}
	// shift first param
	p, params := params[0], params[1:]
	// iterate over it
	for i := 0; i < len(p); i++ {
		// inc WaitGroup
		wg.Add(1)
		// create copy of result
		resultCopy := append([]int{}, result...)
		// call self with remaining params
		go iterateInt(wg, channel, append(resultCopy, p[i]), params...)
	}
}
