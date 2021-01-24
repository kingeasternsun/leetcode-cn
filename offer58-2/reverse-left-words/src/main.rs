fn main() {
    println!("Hello, world!");
}

pub struct Solution;

impl Solution {
    pub fn reverse_left_words(s: String, n: i32) -> String {
        let mut s = s.into_bytes();
        let n = n as usize;
        // return    String::from(&s[n as usize..])+&s[..n as usize];

       Self::reverse(&mut s[..n]);
       Self::reverse(&mut s[n..]);
       Self::reverse(&mut s[..]);


        unsafe{ String::from_utf8_unchecked(s)}
         
    }

    fn reverse(s: &mut [u8]){
        if s.len()==0{
            return 
        }

        for i in 0..s.len()/2{
            s.swap(i, s.len()-1-i);
        }

        return 
    }
}



#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn find() {
        assert_eq!(Solution::reverse_left_words("abcdefg".to_string(), 2),"cdefgab".to_string());
    }
}