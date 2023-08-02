struct Solution;

impl Solution {
    /*
    这个实现方法中，利用HashSet来记录当前的最长子串的所有字符，然后用来判断字符是否在子串中已经存在
    但是发现相同字符后，需要把之前的相同字符b以及b之前的字符从set中移除
     */
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut set = std::collections::HashSet::new();
        let mut left = 0;
        let mut right = 0;
        let mut ret = 0;
        let bytes = s.as_bytes();

        while right < bytes.len() {
            let b = bytes[right];
            if set.contains(&b) {
                ret = ret.max(set.len());
                while left < right {
                    set.remove(&bytes[left]);
                    left += 1;

                    if bytes[left - 1] == b {
                        break;
                    }
                }
            }
            right += 1;
            set.insert(b);
        }
        ret.max(set.len()) as i32
    }

    /*
        在这个实现方法中，通过记录每个字符最近一次在s中的位置，从而快速判断字符是否在子串中存在
    ┌─┬─┬─┬─┬─┬─┐
    │ │i│ │x│ │j│
    └─┴─┴─┴─┴─┴─┘
    如图，假设双指标分别指向i和j，记录s从i到j-1的子串(不包含j)，s[j]最近一次在s中出现的位置为x，
    如果 x>=i, 那么说明 s[j] 在子串s[i:j)中存在
         */
    pub fn length_of_longest_substring2(s: String) -> i32 {
        if s.as_bytes().len() == 0 {
            return 0;
        }
        let mut pre_pos = vec![-1; 128];
        let mut left = -1;
        let mut right = 0;
        let mut ret = 0;
        let bytes = s.as_bytes();

        while right < bytes.len() {
            let b = bytes[right];
            let x = pre_pos[b as usize];
            if x >= left {
                // 注意 pre_pos,left初始化为-1， right 为0的时候会进入这里，保证正确性
                ret = ret.max(right as i32 - left);
                left = x + 1;
            }

            pre_pos[b as usize] = right as i32;
            right += 1;
        }
        ret.max(right as i32 - left)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::length_of_longest_substring("abcaefg".to_string()),
            6
        );
        assert_eq!(
            Solution::length_of_longest_substring("abcabcbb".to_string()),
            3
        );
        assert_eq!(Solution::length_of_longest_substring("bbbb".to_string()), 1);
        assert_eq!(
            Solution::length_of_longest_substring("pwwkew".to_string()),
            3
        );
        assert_eq!(Solution::length_of_longest_substring("".to_string()), 0);
        assert_eq!(Solution::length_of_longest_substring("a".to_string()), 1);
        assert_eq!(Solution::length_of_longest_substring("aa".to_string()), 1);
        assert_eq!(Solution::length_of_longest_substring("aab".to_string()), 2);
        assert_eq!(Solution::length_of_longest_substring("ab".to_string()), 2);
        assert_eq!(Solution::length_of_longest_substring("abc".to_string()), 3);
    }

    #[test]
    fn it_works2() {
        assert_eq!(
            Solution::length_of_longest_substring2("abcaefg".to_string()),
            6
        );
        assert_eq!(
            Solution::length_of_longest_substring2("abcabcbb".to_string()),
            3
        );
        assert_eq!(
            Solution::length_of_longest_substring2("bbbb".to_string()),
            1
        );
        assert_eq!(
            Solution::length_of_longest_substring2("pwwkew".to_string()),
            3
        );
        assert_eq!(Solution::length_of_longest_substring2("".to_string()), 0);
        assert_eq!(Solution::length_of_longest_substring2("a".to_string()), 1);
        assert_eq!(Solution::length_of_longest_substring2("aa".to_string()), 1);
        assert_eq!(Solution::length_of_longest_substring2("ab".to_string()), 2);
        assert_eq!(Solution::length_of_longest_substring2("aab".to_string()), 2);
        assert_eq!(Solution::length_of_longest_substring2("abc".to_string()), 3);
    }
}
