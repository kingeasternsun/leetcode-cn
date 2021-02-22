use std::collections;

use collections::HashMap;

/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-20 11:55:00
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-20 16:12:45
 * @FilePath: \697findShortestSubArray\find_shortest_sub_array\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn find_shortest_sub_array(nums: Vec<i32>) -> i32 {
        let mut num_map = std::collections::HashMap::new();
        for i in 0..nums.len() {
            let e = num_map.entry(&nums[i]).or_insert((i, i, 0));
            e.1 = i; //记录最后一次出现的位置
            e.2 += 1; //记录出现的个数
        }

        let mut res = (0, 0, 0);
        for (_, v) in &num_map {
            if v.2 > res.2 {
                res = (v.0, v.1, v.2);
                continue;
            }

            if (v.2 == res.2) && (v.1 - v.0) < (res.1 - res.0) {
                res = (v.0, v.1, v.2);
            }
        }

        (res.1 - res.0 + 1) as i32
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_shortest_sub_array(vec![1, 2, 2, 3, 1]),2);
        assert_eq!(Solution::find_shortest_sub_array(vec![1,2,2,3,1,4,2]),6);
        assert_eq!(Solution::find_shortest_sub_array(vec![1,1]),2);
        assert_eq!(Solution::find_shortest_sub_array(vec![1]),1);
    }
}
