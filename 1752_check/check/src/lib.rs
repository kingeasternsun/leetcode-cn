struct Solution;

impl Solution {
    // 找到拐点，也就是第一个递减的元素，从这个位置再判断是否单调非递减
    pub fn check(nums: Vec<i32>) -> bool {
        // 找到拐点
        let mut min = 0;
        for i in 1..nums.len() {
            if nums[i] < nums[i - 1] {
                min = i;
                break;
            }
        }

        if min == 0 {
            return true;
        }

        for i in min..nums.len() {
            if nums[i] > nums[(i + 1) % nums.len()] {
                return false;
            }
        }

        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::check(vec![3, 4, 5, 1, 2]), true);
        assert_eq!(Solution::check(vec![2, 1, 3, 4]), false);
        assert_eq!(Solution::check(vec![1, 2, 3]), true);
        assert_eq!(Solution::check(vec![1, 2]), true);
        assert_eq!(Solution::check(vec![2, 1]), true);
        assert_eq!(Solution::check(vec![6, 10, 6]), true);
        assert_eq!(Solution::check(vec![6, 6, 10]), true);
        assert_eq!(Solution::check(vec![10, 6, 6]), true);
        assert_eq!(Solution::check(vec![6, 6, 6]), true);
    }
}
