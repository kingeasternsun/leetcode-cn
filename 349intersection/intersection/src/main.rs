fn main() {
    println!("Hello, world!");
    assert_eq!(Solution::intersection(vec![1,2,2,1], vec![2,3]),vec![2]);
    assert_eq!(Solution::intersection(vec![4,9,5], vec![9,4,9,8,4]),vec![4,9]);
}
pub struct Solution;
use std::collections::HashSet;
impl Solution {
    pub fn intersection(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let m1: HashSet<_> = nums1.iter().cloned().collect();
        let m2: HashSet<_> = nums2.iter().cloned().collect();

        let res : Vec<i32> = m1.intersection(&m2).cloned().collect();

        res
    }
}
