//Using a Map
//Time: O(N); Space: O(N)
func twoSum(nums []int, target int) []int {
    indexMap := make(map[int]int)
    result := make([]int, 2)
    for i, num := range nums{
        if val, isInMap := indexMap[target - num]; isInMap {
            result[0] = val 
            result[1] = i
            break
        }else {
            indexMap[num] = i
        }
    } 
    return result
}
