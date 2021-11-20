pub struct Solution;
impl Solution {
    // learn from https://leetcode-cn.com/problems/integer-replacement/solution/gong-shui-san-xie-yi-ti-san-jie-dfsbfs-t-373h/
    pub fn integer_replacement(n: i32) -> i32 {
        let mut n = n as u32;
        let mut cnt = 0;
        while n > 1 {
            if n == 1 {
                return cnt;
            }

            if n & 1 == 0 {
                n = n >> 1;
            } else if (n & 0b11 == 0b01) || n == 3 {
                n = n - 1;
            } else {
                n = n + 1;
            }
            cnt += 1;
        }

        cnt
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::integer_replacement(8), 3);
        assert_eq!(Solution::integer_replacement(7), 4);
        assert_eq!(Solution::integer_replacement(3), 2);
        assert_eq!(Solution::integer_replacement(2), 1);
        assert_eq!(Solution::integer_replacement(4), 2);
        assert_eq!(Solution::integer_replacement(2147483647), 32);
    }
}
