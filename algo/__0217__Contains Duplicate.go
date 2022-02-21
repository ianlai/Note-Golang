func containsDuplicate(nums []int) bool {
    if len(nums) == 0 {
        return false
    }
    numsMap := make(map[int]bool)
    for _, num := range(nums){
        
        if _, ok := numsMap[num]; ok {
            return true
        } else {
            numsMap[num] = true
        }
        // fmt.Printf("%v\n", numsMap)
    }
    return false
}