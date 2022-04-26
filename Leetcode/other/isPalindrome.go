package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var i, y int
	var reverseNum int = 0
	input:=x
	for input > 0 {
		y = input % 10
		input = input / 10
		reverseNum = reverseNum*10 + y
		i++
		println(reverseNum)
	}
	return reverseNum==x
}

func main() {
    var input int =456789
    ret := isPalindrome(input)
    println(ret)
}
