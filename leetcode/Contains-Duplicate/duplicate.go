package Contains_Duplicate

func containsDuplicate(nums []int) bool {
	check := make(map[int]struct{})
	/*
		bool을 사용하면 초기화하면서 메모리를 차지하지만,
		struct{}는 따로 메모리를 차지하지 않기 때문에 많이들 사용
	*/
	for _, v := range nums {
		if _, ok := check[v]; ok {
			return true
		}
		check[v] = struct{}{}
	}
	return false
}
