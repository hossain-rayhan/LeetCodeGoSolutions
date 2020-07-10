func smallerNumbersThanCurrent(nums []int) []int {
    var count [101]int
    
    for i := 0; i < len(nums); i++ {
        count[nums[i]]++
    }
    
    for i := 1; i < len(count); i++ {
        count[i] = count[i] + count[i-1]
    }
    
    for i := 0; i < len(nums); i++ {
        if(nums[i] != 0){
            nums[i] = count[nums[i] - 1]
        }
    }
    return nums
}
