struct Solution;
impl Solution {
    pub fn is_circular_sentence1(sentence: String) -> bool {
        let mut iter = sentence.split_whitespace();
        let mut first = iter.next().unwrap();
        for word in iter {
            if word.as_bytes().first() != first.as_bytes().last() {
                return false;
            }
            first = word;
        }

        if sentence
            .split_whitespace()
            .next()
            .unwrap()
            .as_bytes()
            .first()
            != first.as_bytes().last()
        {
            return false;
        }

        true
    }

    // 由于题目中提到两个单词中只有一个空格，所以只需要判断空格前后的两个字符是否相同即可，
    // 这里就可以利用Rust中的windows方法得到长度为3的滑动窗口直接进行判断
    pub fn is_circular_sentence(sentence: String) -> bool {
        if sentence.as_bytes().first() != sentence.as_bytes().last() {
            return false;
        }
        sentence.as_bytes().windows(3).all(|x| {
            if x.len() == 3 && x[1] == b' ' && x[0] != x[2] {
                return false;
            }
            return true;
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works1() {
        assert_eq!(
            Solution::is_circular_sentence1("leetcode exercises sound delightful".to_string()),
            true
        );
        assert_eq!(Solution::is_circular_sentence1("eetcode".to_string()), true);
        assert_eq!(Solution::is_circular_sentence1("dog".to_string()), false);
        assert_eq!(
            Solution::is_circular_sentence1("Leetcode is cool".to_string()),
            false
        );
    }

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::is_circular_sentence("leetcode exercises sound delightful".to_string()),
            true
        );
        assert_eq!(Solution::is_circular_sentence("eetcode".to_string()), true);
        assert_eq!(Solution::is_circular_sentence("dog".to_string()), false);
        assert_eq!(
            Solution::is_circular_sentence("Leetcode is cool".to_string()),
            false
        );
    }
}
