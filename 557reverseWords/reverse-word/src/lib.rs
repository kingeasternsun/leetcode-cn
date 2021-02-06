/*
 * @Description: 557
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-06 09:10:55
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 09:22:28
 * @FilePath: \557reverseWords\reverse-word\src\lib.rs
 */
pub struct Solution;
impl Solution {
    pub fn reverse_words(s: String) -> String {
        let mut s = s.into_bytes();
        s.split_mut(|c| c == &(' ' as u8)).for_each(|w| w.reverse());
        unsafe { String::from_utf8_unchecked(s) }
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::reverse_words(String::from("Let's take LeetCode contest")),
            String::from("s'teL ekat edoCteeL tsetnoc")
        );
    }
}
