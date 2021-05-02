/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-01 20:36:59
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-01 21:43:10
 * @FilePath: /fib/src/lib.rs
 */

pub struct Solution;

impl Solution {
    pub fn multiply(a: [[i64; 2]; 2], b: [[i64; 2]; 2]) -> [[i64; 2]; 2] {
        let mut c:[[i64; 2]; 2] = [[0,0],[0,0]];
        for i in 0..2 {
            for j in 0..2 {
                c[i][j] = (a[i][0] * b[0][j] + a[i][1] * b[1][j]) % (1000000007)
            }
        }

        c
    }

    pub fn pow(mut a :[[i64;2];2], n:i32)-> [[i64; 2]; 2] {

        let mut c:[[i64; 2]; 2]= [[1,0],[0,1]];
        let mut n = n;
        while n>0{
            if n&1 ==1{
                c = Self::multiply(c, a);
            }

            a = Self::multiply(a, a);
            // println!("{:?}",a);
            n = n>>1;
        }

        c
    }

    pub fn fib(n: i32) -> i32 {
        if n <2{
            return n;
        }

        let r:[[i64; 2]; 2] = Self::pow([[1,1],[1,0]], n-1);
        return r[0][0] as i32

    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(Solution::fib(5),5);
        assert_eq!(Solution::fib(48),807526948);
        assert_eq!(Solution::fib(6),8);
    }
}

// 0 1 1 2 3 5