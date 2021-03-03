/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-03 11:10:53
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-03 11:22:34
 * @FilePath: \338countBits\count_bits\src\lib.rs
 */
pub struct Solution;
impl Solution {
    pub fn count_bits(num: i32) -> Vec<i32> {
        let num = num as usize;
        let mut res = vec![0i32; num+1];
        for i in 1..res.len() {
            res[i] = res[i >> 1] + (i  & 1) as i32;
        }
        res
    }
    pub fn count_bits1(num: i32) -> Vec<i32> {
        
        (0..=num).map(|x|{
            x.count_ones() as i32
        }).collect()
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::count_bits(5), vec![0, 1, 1, 2, 1, 2]);
        assert_eq!(Solution::count_bits1(5), vec![0, 1, 1, 2, 1, 2]);
    }
}
