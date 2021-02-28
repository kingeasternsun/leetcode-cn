/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-28 13:59:47
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-28 14:17:45
 * @FilePath: /is_monotonic/src/lib.rs
 */

pub struct Solution;
impl Solution {
    pub fn is_monotonic(a: Vec<i32>) -> bool {
        let mut incre = true;
        let mut decre = true;
        a.windows(2).all(|w| {
            if w[0] > w[1] {
                incre = false;
            }
            if w[0] < w[1] {
                decre = false;
            }
            incre || decre
        })
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::is_monotonic(vec![1, 2]), true);
        assert_eq!(Solution::is_monotonic(vec![1, 2,3]), true);
        assert_eq!(Solution::is_monotonic(vec![1, 2, 1]), false);
        assert_eq!(Solution::is_monotonic(vec![2, 1, 2]), false);
        assert_eq!(Solution::is_monotonic(vec![ 2]), true);
    }
}
