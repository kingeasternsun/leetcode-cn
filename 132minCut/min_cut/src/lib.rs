/*
 * @Description: 132 min_cut
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-08 11:18:50
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-08 11:38:36
 * @FilePath: \132minCut\min_cut\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn min_cut(s: String) -> i32 {
        if s.len() <= 1 {
            return 0;
        }
        let mut isPali = vec![vec![false; s.len()]; s.len()];

        let bs = s.into_bytes();
        //预先计算任意区间是否是回文字符串
        let mut judge = |mut x: usize, mut y: usize| {
            while y < bs.len() && bs[x] == bs[y] {
                isPali[x][y] = true;
                if x == 0 {
                    break;
                };
                x -= 1;
                y += 1;
            }
        };
        (0..bs.len()).for_each(|x| {
            judge(x, x);
            judge(x, x + 1);
        });

        //dp计算
        let mut dp = vec![0 as i32; bs.len()];
        for end in 1..bs.len() {

            dp[end] = bs.len() as i32 - 1;

            for beg in 0..end + 1 {
                if isPali[beg][end] {
                    if beg == 0 {
                        dp[end] = 0
                    } else if dp[beg - 1] + 1 < dp[end] {
                        dp[end] = dp[beg - 1] + 1
                    }
                }
            }
        }

        dp[dp.len() - 1]
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::min_cut(String::from("aab")),1);
        assert_eq!(Solution::min_cut(String::from("aba")),0);
    }
}
