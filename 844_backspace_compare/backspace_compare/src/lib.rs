struct Solution;
impl Solution {
    pub fn backspace_compare(s: String, t: String) -> bool {
        let mut s_ret = vec![];
        for &b in s.as_bytes() {
            if b == b'#' {
                s_ret.pop();
            } else {
                s_ret.push(b);
            }
        }

        let mut t_ret = vec![];
        for &b in t.as_bytes() {
            if b == b'#' {
                t_ret.pop();
            } else {
                t_ret.push(b);
            }
        }

        s_ret == t_ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::backspace_compare("ab#c".to_string(), "ad#c".to_string()),
            true
        );
        assert_eq!(
            Solution::backspace_compare("ab##".to_string(), "c#d#".to_string()),
            true
        );
        assert_eq!(
            Solution::backspace_compare("ab##".to_string(), "c#".to_string()),
            true
        );
        assert_eq!(
            Solution::backspace_compare("a#c".to_string(), "b".to_string()),
            false
        );
    }
}
