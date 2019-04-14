package slice_factor


func sliceFactor(input []int) []int {
    retSlice := make([]int, len(input))

    for i, iVal := range input {
        for r, rVal := range retSlice {
            if i == r {
                continue
            }
            if rVal == 0 {
                retSlice[r] = iVal
                continue
            }
            retSlice[r] = iVal * rVal
        }
    }
    return retSlice
}
