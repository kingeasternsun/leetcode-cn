fn main() {
    println!("Hello, world!");
}

pub struct Solution;

impl Solution {

    fn reverse(s: &mut [u8]) {
        if s.len() == 0 {
            return;
        }
        for i in 0..s.len() / 2 {
            s.swap(i, s.len() - 1 - i);
        }
        return;
    }
    
    pub fn reverse_str1(s: String, k: i32) -> String {
        let mut s = s.into_bytes();
        let k = k as usize;

        let len = s.len();
        for i in 0..s.len() / (2 * k) {
            Self::reverse_range(&mut s[i * 2 * k..i * 2 * k + k], k);
        }
        let left = s.len() % (2 * k);
        if left > 0 {
            Self::reverse_range(&mut s[len - left..], k);
        }

        unsafe { String::from_utf8_unchecked(s) }
    }
    pub fn reverse_str(s: String, k: i32) -> String {
        let mut s = s.into_bytes();
        let k = k as usize;
        //上面的代码可以简写为
        for ch in s.chunks_mut(2 * k) {
            Self::reverse_range(ch, k)
        }
        unsafe { String::from_utf8_unchecked(s) }
    }

    fn reverse_range(s: &mut [u8], k: usize) {
        if s.len() < k {
            // Self::reverse(s)
            s.reverse();
        } else {
            // Self::reverse(&mut s[..k])
            (&mut s[..k]).reverse();
        }
    }


}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn find() {
        assert_eq!(
            Solution::reverse_str("abcdefg".to_string(), 2),
            "bacdfeg".to_string()
        );
        assert_eq!(
            Solution::reverse_str("abcdefgh".to_string(), 2),
            "bacdfegh".to_string()
        );
        assert_eq!(
            Solution::reverse_str("abcde".to_string(), 2),
            "bacde".to_string()
        );
    }
}
