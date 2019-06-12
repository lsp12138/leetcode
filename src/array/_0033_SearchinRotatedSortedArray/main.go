package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {

}

// class Solution {
//     public int search(int[] nums, int target) {
//         if (nums == null || nums.length == 0) return -1;
//         int left = 0;
//         int right = nums.length - 1;
//         while (left < right) {
//             int mid = left + (right - left) / 2;
//             if (nums[mid] > nums[right]) left = mid + 1;
//             else right = mid;
//         }
//         //System.out.println(left);
//         int split_t = left;
//         left = 0;
//         right = nums.length - 1;
//         if (nums[split_t] <= target && target <= nums[right]) left = split_t;
//         else right = split_t;
//         while (left <= right) {
//             int mid = left + (right - left) / 2;
//             if (nums[mid] == target) return mid;
//             else if (nums[mid] > target) right = mid - 1;
//             else left = mid + 1;
//         }
//         return -1;

//     }
// }
