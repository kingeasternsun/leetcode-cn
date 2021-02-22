
/*
 * @Description:1438
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-22 09:16:09
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-22 09:50:29
 * @FilePath: \1438longestSubarray\longest_subarray\src\lib.rs
 */
pub struct Solution;

impl Solution {
    //376 ms  3.7MB
    pub fn longest_subarray_slow(nums: Vec<i32>, limit: i32) -> i32 {
        if nums.len() <= 1 {
            return nums.len() as i32;
        }

        let mut res = 1;
        //双指针滑窗
        let mut beg = 0;
        let mut end = 0;
        let mut monoMax = vec![nums[0]]; //维护窗口中的最大值，因为从左边移除元素，所以最大值在最左边
        let mut monoMin = vec![nums[0]]; //维护窗口中的最小值，因为从左边移除元素，所以最小值在最左边
        while end < nums.len() && beg <= end {
            //窗口中最大值 最小值之差满足条件
            if (monoMax[0] - monoMin[0]) <= limit {
                if end - beg + 1 > res {
                    res = end - beg + 1
                }

                end += 1;

                if end == nums.len() {
                    return res as i32;
                }

                //更新 monoMax，最大值要在最左边
                while let Some(e) = monoMax.last() {
                    if *e >= nums[end] {
                        break;
                    }
                    monoMax.pop();
                }
                monoMax.push(nums[end]);

                //更新 monoMin，最小值要在最左边
                while let Some(e) = monoMin.last() {
                    if *e <= nums[end] {
                        break;
                    }
                    monoMin.pop();
                }
                monoMin.push(nums[end]);
            } else {
                //移除 beg
                if nums[beg] >= monoMax[0] {
                    monoMax.remove(0);
                }

                if nums[beg] <= monoMin[0] {
                    monoMin.remove(0);
                }

                beg += 1;
            }
        }

        res as i32
    }

    pub fn longest_subarray(nums: Vec<i32>, limit: i32) -> i32 {
        if nums.len() <= 1 {
            return nums.len() as i32;
        }

        let mut res = 1;
        //双指针滑窗
        let mut beg = 0;
        let mut end = 0;
        let mut monoMax = std::collections::VecDeque::new(); //维护窗口中的最大值，因为从左边移除元素，所以最大值在最左边
        monoMax.push_back(nums[0]);
        let mut monoMin = std::collections::VecDeque::new(); //维护窗口中的最小值，因为从左边移除元素，所以最小值在最左边
        monoMin.push_back(nums[0]);

        while end < nums.len() && beg <= end {
            //窗口中最大值 最小值之差满足条件
            if (monoMax[0] - monoMin[0]) <= limit {
                if end - beg + 1 > res {
                    res = end - beg + 1
                }

                end += 1;

                if end == nums.len() {
                    return res as i32;
                }

                //更新 monoMax，最大值要在最左边
                while let Some(e) = monoMax.back(){
                    if *e >= nums[end] {
                        break;
                    }
                    monoMax.pop_back();
                }
                monoMax.push_back(nums[end]);

                //更新 monoMin，最小值要在最左边
                while let Some(e) = monoMin.back() {
                    if *e <= nums[end] {
                        break;
                    }
                    monoMin.pop_back();
                }
                monoMin.push_back(nums[end]);
            } else {
                //移除 beg
                if nums[beg] >= monoMax[0] {
                    monoMax.pop_front();
                }

                if nums[beg] <= monoMin[0] {
                    monoMin.pop_front();
                }

                beg += 1;
            }
        }

        res as i32
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]

    fn it_works_slow() {

        assert_eq!(Solution::longest_subarray_slow(vec![8, 2, 4, 7], 4), 2);
        assert_eq!(Solution::longest_subarray_slow(vec![10, 1, 2, 4, 7, 2], 5), 4);
        assert_eq!(
            Solution::longest_subarray_slow(vec![4, 2, 2, 2, 4, 4, 2, 2], 0),
            3
        );
    }
    fn it_works() {
        assert_eq!(Solution::longest_subarray(vec![8, 2, 4, 7], 4), 2);
        assert_eq!(Solution::longest_subarray(vec![10, 1, 2, 4, 7, 2], 5), 4);
        assert_eq!(
            Solution::longest_subarray(vec![4, 2, 2, 2, 4, 4, 2, 2], 0),
            3
        );
    }

}
