package two_sum

/*
N^2 방법
*/
func twoSum1(nums []int, target int) []int {
	for i, vi := range nums {
		for j, vj := range nums {
			if j <= i {
				continue
			}
			if vi+vj == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*
시간복잡도 NlogN 방법
map으로 만든 다음, (n)
for문을 돌면서 합이 target이 되는 값을 찾아서 존재하면 반환한다. (n)

공간복잡도
*/
func twoSum2(nums []int, target int) []int {
	numsMap := make(map[int]int, 0)
	for i, v := range nums {
		numsMap[v] = i
	}
	for j, v := range nums {
		if i, ok := numsMap[target-v]; ok {
			if j == i {
				continue
			}
			return []int{i, j}
		}
	}
	return []int{}
}

/*
leetcode의 어느 사람의 코드
for문을 한번만 돌면서 map[subtract] = index 를 만들어준다. 두번의 for 문을 돌지 않는다.
*/
func twoSum3(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, vi := range nums {
		subtract := target - vi
		if j, ok := numsMap[subtract]; ok {
			return []int{i, j}
		}
		numsMap[vi] = i
	}
	return []int{}
}

/*
내 코드
*/
func twoSum4(nums []int, target int) []int {
	checkList := make(map[int]int)

	for i, v := range nums {
		diff := target - v
		if diffIndex, ok := checkList[diff]; ok {
			return []int{diffIndex, i}
		}
		checkList[v] = i
	}
	return []int{}
}
