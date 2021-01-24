fn main() {
    println!("Hello, world!");
}

pub struct Solution;
impl Solution {
    pub fn reverse_left_words(s: String, n: i32) -> String {
        let mut s = s.into_bytes();
        // s.rotate_left(n);

        // return    String::from(&s[n as usize..])+&s[..n as usize];
         
    }

    fn reverse()
}



#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn find() {
        assert_eq!(Solution::reverse_left_words("abcdefg".to_string(), 2),"cdefgab".to_string());
    }
}