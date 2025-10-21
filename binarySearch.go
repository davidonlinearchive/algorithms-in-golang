package main

func binSearch(nums []int, target int) int {
  low := 0
  high := len(nums) - 1

  for low <= high {
    mid := (low + high) / 2
    if target > nums[mid] {
      low = mid + 1 // Search right
    } else if target < nums[mid] {
       high = mid - 1 // Seach left 
    } else {
      return mid // target is mid
    }

  }
      return -1 // target not found
}
