
struct Solution;
impl Solution {
    // beautf of iterator
    // 0ms 1.96mb
    pub fn can_split_array(nums: Vec<i32>, m: i32) -> bool {
        if nums.len()<=2{
            return true
        }
        nums.windows(2).any(|win|win[0]+win[1]>=m)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::can_split_array(vec![2, 2, 1], 4), true);
        assert_eq!(Solution::can_split_array(vec![2, 1, 3], 5), false);
        assert_eq!(Solution::can_split_array(vec![2, 3, 3, 2, 3], 6), true);
        assert_eq!(Solution::can_split_array(vec![1], 6), true);
        assert_eq!(Solution::can_split_array(vec![6], 6), true);
        assert_eq!(Solution::can_split_array(vec![8], 6), true);

        assert_eq!(Solution::can_split_array(vec![1, 2], 6), true);
        assert_eq!(Solution::can_split_array(vec![2, 4], 6), true);
        assert_eq!(Solution::can_split_array(vec![2, 1,1,3], 4), true);
    }
}
