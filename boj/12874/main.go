package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	   ()()() : 100개
	   ()
	   ()()
	   (())
	   ()()()

	   ()()()(): (((( 4개, )))) 4개
	   ()
	   ()()
	   ()()()
	   ()()()()

	   (())
	   (()())

	   ##
	   )))) (((() : 유효한 괄호 ( 1개, )1개

	   ##
		유효한 골호수를 센다.
		유효한 괄호수로 나올 수 있는 경우의 조합을 구한다.
		문자열 비교를 해서 찾는다.

	   (())(()) : 유효한 괄호 수 ( 4개, )4개
		() : 유효한 괄호 수를 센다.
		(()) :
		((())) (()()) ()()()

		()()()
		(())

	   ()
	   (())
	   (())(())
	   ()()
	   (()())
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string
	fmt.Fscanln(reader, &s)
	writer.WriteString(s)
	return
}

func ValidElem(s string) int {
	maxCount := 0
	validCount := 0
	for i := range s {
		if maxCount > 0 && s[i] == ')' {
			maxCount -= 1
			validCount++
		} else if s[i] == '(' {
			maxCount += 1
		}
	}
	return validCount
}

/*
n1 = () => 1
n2 = ()() / (()) => 2*n1
n3 = ()()() / (())() / ()(()) / ((())) => 2*n2
*/

/*
- 닫는 괄호가 먼저 나오면 제외
(())()))) => (())()
()()))() => ()()()
)()(())) => ()(())
*/

//func GetCase(n int) map[int]int {
//	dp := make(map[int]int)
//
//	return dp
//}

/*
1-100까지 가능한 것들을 counting한다.
*/

func solution(s string) {

}

func solution(s string) int {

	return 0
}
