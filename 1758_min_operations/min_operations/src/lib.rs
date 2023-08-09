struct Solution;
impl Solution {
    // 生成的字符串无非 0101 或 1010， 只需要和期望的进行对比即可
    pub fn min_operations(s: String) -> i32 {
        let exp = "01".as_bytes();
        let bytes = s.as_bytes();
        let mut ret0 = 0;
        for (id, &b) in bytes.iter().enumerate() {
            if b != exp[id & 1] {
                ret0 += 1;
            }
        }

        let mut ret1 = 0;
        for (id, &b) in bytes.iter().enumerate() {
            if b != exp[(id + 1) & 1] {
                ret1 += 1;
            }
        }

        ret1.min(ret0)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::min_operations(String::from("0100")), 1);
        assert_eq!(Solution::min_operations(String::from("1000")), 1);
        assert_eq!(Solution::min_operations(String::from("10")), 0);
        assert_eq!(Solution::min_operations(String::from("01")), 0);
        assert_eq!(Solution::min_operations(String::from("101")), 0);
        assert_eq!(Solution::min_operations(String::from("010")), 0);
        assert_eq!(Solution::min_operations(String::from("1111")), 2);
        assert_eq!(Solution::min_operations(String::from("0")), 0);
        assert_eq!(Solution::min_operations(String::from("1")), 0);
    }
}
