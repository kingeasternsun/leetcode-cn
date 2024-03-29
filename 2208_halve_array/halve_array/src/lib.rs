struct Solution;

use std::cmp::Ordering;
#[derive(PartialEq, PartialOrd)]
struct Float64(f64);

impl Eq for Float64 {}
impl Ord for Float64 {
    fn cmp(&self, other: &Self) -> Ordering {
        self.partial_cmp(&other).unwrap_or(Ordering::Less)
    }
}

impl Solution {
    pub fn halve_array(nums: Vec<i32>) -> i32 {
        let mut ret = 0;
        // NOTE: the sum of the nums maybe overflow, so we need to convert it to f64 when sum
        let sum: f64 = nums.iter().map(|&x| x as f64).sum();
        let half_sum = sum / 2.0;
        let mut remove_sum = 0_f64;

        let mut queue = nums
            .iter()
            .map(|&x| Float64(x as f64))
            .collect::<std::collections::BinaryHeap<_>>();
        while let Some(Float64(m)) = queue.pop() {
            ret += 1;
            remove_sum = remove_sum + m / 2.0;
            if remove_sum < half_sum {
                queue.push(Float64(m / 2.0));
            } else {
                break;
            }
        }

        ret
    }
    pub fn halve_array2(nums: Vec<i32>) -> i32 {
        let mut ret = 0;
        let sum: f64 = nums.iter().map(|&x| x as f64).sum();
        let half_sum = sum / 2.0;
        let mut remove_sum = 0_f64;

        let mut queue = nums
            .iter()
            .map(|&x| Float64(x as f64))
            .collect::<std::collections::BinaryHeap<_>>();

        while remove_sum < half_sum {
            let x = queue.pop().unwrap().0 / 2.0;
            remove_sum += x;
            ret += 1;
            queue.push(Float64(x));
        }

        ret
    }

    // Optimization: modified the peek in place instead of pop, modify and push
    pub fn halve_array3(nums: Vec<i32>) -> i32 {
        let mut ret = 0;
        // NOTE: the sum of the nums maybe overflow, so we need to convert it to f64 when sum
        let sum: f64 = nums.iter().map(|&x| x as f64).sum();
        let half_sum = sum / 2.0;
        let mut remove_sum = 0_f64;

        let mut queue = nums
            .iter()
            .map(|&x| Float64(x as f64))
            .collect::<std::collections::BinaryHeap<_>>();
        while let Some(mut m) = queue.peek_mut() {
            ret += 1;
            remove_sum = remove_sum + m.0 / 2.0;
            if remove_sum < half_sum {
                *m = Float64(m.0 / 2.0);
            } else {
                break;
            }
        }

        ret
    }

    pub fn halve_array4(nums: Vec<i32>) -> i32 {
        let mut ret = 0;
        let sum: f64 = nums.iter().map(|&x| x as f64).sum();
        let half_sum = sum / 2.0;
        let mut left_sum = sum;

        let mut queue = nums
            .iter()
            .map(|&x| Float64(x as f64))
            .collect::<std::collections::BinaryHeap<_>>();

        while left_sum > half_sum {
            let mut m = queue.peek_mut().unwrap();
            left_sum -= m.0 / 2.0;
            ret += 1;
            *m = Float64(m.0 / 2.0);
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::halve_array(vec![5, 19, 8, 1]), 3);
        assert_eq!(Solution::halve_array(vec![3, 8, 20]), 3);
        assert_eq!(Solution::halve_array(vec![2, 16, 20]), 3);
        assert_eq!(Solution::halve_array(vec![2, 20]), 2);
    }

    #[test]
    fn it_works2() {
        assert_eq!(Solution::halve_array2(vec![5, 19, 8, 1]), 3);
        assert_eq!(Solution::halve_array2(vec![3, 8, 20]), 3);
        assert_eq!(Solution::halve_array2(vec![2, 16, 20]), 3);
        assert_eq!(Solution::halve_array2(vec![2, 20]), 2);
    }

    #[test]
    fn it_works3() {
        assert_eq!(Solution::halve_array3(vec![5, 19, 8, 1]), 3);
        assert_eq!(Solution::halve_array3(vec![3, 8, 20]), 3);
        assert_eq!(Solution::halve_array3(vec![2, 16, 20]), 3);
        assert_eq!(Solution::halve_array3(vec![2, 20]), 2);
    }

    #[test]
    fn it_works4() {
        assert_eq!(Solution::halve_array4(vec![5, 19, 8, 1]), 3);
        assert_eq!(Solution::halve_array4(vec![3, 8, 20]), 3);
        assert_eq!(Solution::halve_array4(vec![2, 16, 20]), 3);
        assert_eq!(Solution::halve_array4(vec![2, 20]), 2);
    }
}
