struct Solution;
impl Solution {
    // XOR:  0ms 2MB
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        let mut ret = 0;
        for n in 0..=nums.len() as i32 {
            ret ^= n;
        }

        for n in nums {
            ret ^= n;
        }

        ret
    }

    // the beauty of iterator
    // 4ms 2.17mb
    pub fn missing_number1(nums: Vec<i32>) -> i32 {
        let mut ret = (0..=nums.len() as i32)
            .into_iter()
            .fold(0, |acc, n| acc ^ n);
        nums.into_iter().fold(ret, |acc, n| acc ^ n)
    }

    // 4ms 2.13mb
    pub fn missing_number2(nums: Vec<i32>) -> i32 {
        (0..=nums.len() as i32).chain(nums.into_iter()).fold(0, |acc, n| acc ^ n)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::missing_number(vec![3, 0, 1]), 2);
        assert_eq!(Solution::missing_number(vec![0, 1]), 2);
        assert_eq!(Solution::missing_number(vec![0]), 1);
        assert_eq!(Solution::missing_number(vec![9, 6, 4, 2, 3, 5, 7, 0, 1]), 8);
    }
    #[test]
    fn it_works1() {
        assert_eq!(Solution::missing_number1(vec![3, 0, 1]), 2);
        assert_eq!(Solution::missing_number1(vec![0, 1]), 2);
        assert_eq!(Solution::missing_number1(vec![0]), 1);
        assert_eq!(
            Solution::missing_number1(vec![9, 6, 4, 2, 3, 5, 7, 0, 1]),
            8
        );
    }
    #[test]
    fn it_works2() {
        assert_eq!(Solution::missing_number2(vec![3, 0, 1]), 2);
        assert_eq!(Solution::missing_number2(vec![0, 1]), 2);
        assert_eq!(Solution::missing_number2(vec![0]), 1);
        assert_eq!(
            Solution::missing_number2(vec![9, 6, 4, 2, 3, 5, 7, 0, 1]),
            8
        );
    }
}
