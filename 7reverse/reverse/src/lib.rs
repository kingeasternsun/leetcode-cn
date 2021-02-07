/*
 * @Description:  参考 <浅谈rust在算法题和竞赛中的应用>https://www.bilibili.com/video/BV1Yy4y1e7zR?p=27
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-04 11:42:10
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-04 13:59:08
 * @FilePath: \7reverse\reverse\src\lib.rs
 */
pub struct Solution;
impl Solution {
    pub fn reverse(x: i32) -> i32 {

        fn helper(mut n:i32)->Option<i32>{
            let mut res = 0i32;
            while n.abs()!=0{
                res = res.checked_mul(10)?.checked_add(n%10)?;
                n = n/10;
            }

            Some(res)
        }

        helper(x).unwrap_or_default()

    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        // assert_eq!(Solution::reverse(123),321);
        assert_eq!(Solution::reverse(-123),-321);
        // assert_eq!(Solution::reverse(120),21);
        // assert_eq!(Solution::reverse(0),0);
    }
}
