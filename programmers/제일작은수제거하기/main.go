func solution(arr []int) []int {
    if len(arr) == 1 {
        return []int{-1}
    }
    var (
        removeIndex int = 0
        minValue int = arr[0]
    )
    for i := range arr {
        if arr[i] < minValue {
            minValue = arr[i]
            removeIndex = i
        }
    }
    
    return append(arr[:removeIndex], arr[removeIndex+1:]...)
}
