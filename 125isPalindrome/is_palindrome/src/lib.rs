/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-16 14:29:53
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-16 14:56:15
 * @FilePath: /125isPalindrome/is_palindrome/src/lib.rs
 */
pub struct  Solution;
impl Solution {
    pub fn is_palindrome1(s: String) -> bool {
        let b = s.into_bytes();
        let mut i = 0;
        let mut j = b.len()-1;

        while i<j {

            //跳过非字母或数字
            while i<j && (!b[i].is_ascii_alphanumeric()) {i+=1;};
            //跳过非字母或数字
            while i<j && (!b[j].is_ascii_alphanumeric()) {j-=1;};
            if i<j && b[i].to_ascii_uppercase()!=b[j].to_ascii_uppercase(){
                return false
            }
            //由于 j 是usize，所以j有可能变为0，然后 j-=1 后溢出
            if i ==j {
                return true
            }
            i+=1;
            j-=1;
            
        }

        true
    }

    //更巧妙的写法
    pub fn is_palindrome(s: String) -> bool {
        let b = s.into_bytes();
        let mut i = 0;
        let mut j = b.len()-1;

        while i<j {

            //跳过非字母或数字
            if !b[i].is_ascii_alphanumeric() {i+=1;continue;};
            //跳过非字母或数字
            if !b[j].is_ascii_alphanumeric() {j-=1;continue;};

            if b[i].to_ascii_uppercase()!=b[j].to_ascii_uppercase(){
                return false
            }
            i+=1;
            j-=1;
            
        }

        true
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(Solution::is_palindrome("A man, a plan, a canal: Panama".to_string()),true);
        assert_eq!(Solution::is_palindrome(String::from("race a car")),false);
        assert_eq!(Solution::is_palindrome(String::from("a.")),true);
        assert_eq!(Solution::is_palindrome(String::from(".")),true);
    }
}
