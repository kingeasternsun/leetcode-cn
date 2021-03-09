/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-09 14:08:22
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-09 14:10:18
 * @FilePath: /191hammingWeight/hammingWeight/src/lib.rs
 */
pub struct Solution;
impl Solution {
    pub fn hammingWeight1 (n: u32) -> i32 {


        return n.count_ones() as i32
        
    }
    pub fn hammingWeight (n: u32) -> i32 {

        let mut res = 0;
        let mut n = n;
        while n !=0{

            n = n &(n -1);
            res+=1;

        }

        res
        
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(Solution::hammingWeight1(3),2);
        assert_eq!(Solution::hammingWeight(3),2);
    }
}
