struct Solution;
impl Solution {
    pub fn max_power(s: String) -> i32 {
        let mut ret = 1;
        let mut pre = None;
        s.into_bytes().windows(2).for_each(|win| {
            if win[0] == win[1] {
                match pre {
                    None => pre = Some(2),
                    Some(pre_len) => {
                        pre.replace(pre_len + 1);
                    }
                }
            } else if let Some(pre_len) = pre {
                ret = ret.max(pre_len);
                pre = None;
            }
        });

        if let Some(pre_len) = pre {
            ret = ret.max(pre_len);
            pre = None;
        }

        ret
    }

    pub fn max_power1(s: String) -> i32 {
        let mut ret = 1;
        let mut pre = None;
        s.into_bytes().windows(2).for_each(|win| {
            if win[0] == win[1] {
                match pre {
                    None => {
                        pre = Some(2);
                        ret = ret.max(2);
                    }
                    Some(pre_len) => {
                        pre.replace(pre_len + 1);
                        ret = ret.max(pre_len + 1);
                    }
                }
            } else if let Some(pre_len) = pre {
                pre = None;
            }
        });

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::max_power("leetcode".to_string()), 2);
        assert_eq!(Solution::max_power("abbcccdddeeeeedcba".to_string()), 5);
        assert_eq!(Solution::max_power("abaaa".to_string()), 3);
        assert_eq!(Solution::max_power("bbaaa".to_string()), 3);
        assert_eq!(Solution::max_power("aa".to_string()), 2);
        assert_eq!(Solution::max_power("a".to_string()), 1);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::max_power1("leetcode".to_string()), 2);
        assert_eq!(Solution::max_power1("abbcccdddeeeeedcba".to_string()), 5);
        assert_eq!(Solution::max_power1("abaaa".to_string()), 3);
        assert_eq!(Solution::max_power1("bbaaa".to_string()), 3);
        assert_eq!(Solution::max_power1("aa".to_string()), 2);
        assert_eq!(Solution::max_power1("a".to_string()), 1);
    }
}
