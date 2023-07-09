struct Solution;
impl Solution {
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut i: usize = 0;
        let mut j = numbers.len() - 1;

        while i < j {
            let sum = numbers[i] + numbers[j];
            if sum == target {
                return vec![i as i32 + 1, j as i32 + 1];
            }
            if sum > target {
                j -= 1;
            } else {
                i += 1;
            }
        }
        vec![0, 0]
    }

    // 加个快速判断，如果首两个之和大于target，就肯定不存在，或如果尾两个之和小于target，就肯定不存在
    pub fn two_sum1(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        if numbers[0] + numbers[1] > target
            || numbers[numbers.len() - 2] + numbers[numbers.len() - 1] < target
        {
            return vec![0, 0];
        }

        let mut i: usize = 0;
        let mut j = numbers.len() - 1;

        while i < j {
            let sum = numbers[i] + numbers[j];
            if sum == target {
                return vec![i as i32 + 1, j as i32 + 1];
            }
            if sum > target {
                j -= 1;
            } else {
                i += 1;
            }
        }
        vec![0, 0]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![1, 2]);
        assert_eq!(Solution::two_sum(vec![2, 3, 4], 6), vec![1, 3]);
        assert_eq!(Solution::two_sum(vec![-1, 0], -1), vec![1, 2]);
        assert_eq!(Solution::two_sum(vec![-1, 2], -1), vec![0, 0]);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::two_sum1(vec![2, 7, 11, 15], 9), vec![1, 2]);
        assert_eq!(Solution::two_sum1(vec![2, 3, 4], 6), vec![1, 3]);
        assert_eq!(Solution::two_sum1(vec![-1, 0], -1), vec![1, 2]);
        assert_eq!(Solution::two_sum1(vec![-1, 2], -1), vec![0, 0]);
    }
}
