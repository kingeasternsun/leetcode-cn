struct Solution;
impl Solution {
    /*
    这个题目，可以等价于 从nums中选取两个不同元素，这两个元素的差值绝对值为K 的组合有几种
    1. 由于 1<= nums[i] <= 100, 可以采用 数组map 加 K长度的滑动窗口
     */
    pub fn count_k_difference(nums: Vec<i32>, k: i32) -> i32 {
        if k > 100 {
            return 0;
        }
        let mut vec_map = vec![0; 101];
        // beauty of iterator
        nums.into_iter().for_each(|n| vec_map[n as usize] += 1);
        vec_map
            .windows(k as usize + 1)
            .map(|m| m[0] * m[k as usize])
            .sum()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::count_k_difference(vec![1, 2, 2, 1], 1), 4);
        assert_eq!(Solution::count_k_difference(vec![1,3], 3), 0);
        assert_eq!(Solution::count_k_difference(vec![3,2,1,5,4], 2), 3);
        assert_eq!(Solution::count_k_difference(vec![3,3,6,9,6], 3), 6);
    }
}
