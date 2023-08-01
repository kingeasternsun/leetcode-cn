struct Solution;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut set = std::collections::HashSet::new();
        let mut ret = 0;
        s.as_bytes().into_iter().for_each(|b| {
            if set.contains(b) {
                ret =ret.max(set.len());
                set.clear();
                set.insert(b);
            } else {
                set.insert(b);
            }
        });
        ret.max(set.len()) as i32
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::length_of_longest_substring("abcabcbb".to_string()), 3);
        assert_eq!(Solution::length_of_longest_substring("bbbb".to_string()), 1);
        assert_eq!(Solution::length_of_longest_substring("pwwkew".to_string()), 3);
        assert_eq!(Solution::length_of_longest_substring("".to_string()), 0);
        assert_eq!(Solution::length_of_longest_substring("a".to_string()), 1);
        assert_eq!(Solution::length_of_longest_substring("aa".to_string()), 1);
        assert_eq!(Solution::length_of_longest_substring("ab".to_string()), 2);
        assert_eq!(Solution::length_of_longest_substring("abc".to_string()), 3);
    }
}
