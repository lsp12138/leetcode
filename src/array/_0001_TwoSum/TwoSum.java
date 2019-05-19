package array._0001_TwoSum;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    public static void main(String[] args) {
        int[] nums = { 0, 4, 2, 0 };
        int target = 0;
        Solution s = new Solution();
        System.out.println(Arrays.toString(s.twoSum(nums, target)));
    }
}

// class Solution {
//     public int[] twoSum(int[] nums, int target) {
//         int[] result = { 0, 0 };
//         for (int i = 0; i < nums.length; i++) {
//             for (int j = i + 1; j < nums.length; j++) {
//                 if (nums[i] + nums[j] == target) {
//                     result[0] = i;
//                     result[1] = j;
//                 }
//             }
//         }
//         return result;
//     }
// }

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int tmp = target - nums[i];
            if (map.containsKey(tmp)) {
                return new int[] { map.get(tmp), i };
            }
            map.put(nums[i], i);
        }
        throw new IllegalArgumentException("no solution");
    }
}