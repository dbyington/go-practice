package two_num_sum

func twoNumSum(nums []int, i int) bool {
    set := make(map[int]bool, len(nums))
    for _, num := range nums {
        if _, ok := set[num]; ok {
            return true
        }
        set[i-num] = true
    }
    return false
}
