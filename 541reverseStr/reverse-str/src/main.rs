fn main() {
    println!("Hello, world!");
}

pub struct Solution;

impl Solution {
    pub fn reverse_str(s: String, k: i32) -> String {

        let mut s = s.into_bytes();
        let len = s.len();
        let k = k as usize;

        for i in 0..s.len()/(2*k){
            Self::reverse_range(&mut s[i*2*k..i*2*k+k],k);
        }

        let left = s.len()%(2*k);
        if left>0{
            Self::reverse_range(&mut s[len-left..],k);
        }

        


        unsafe{ String::from_utf8_unchecked(s)}
        
    }

    fn reverse_range(s:&mut[u8], k:usize){
        if s.len()<k{
            return Self::reverse(s)
        } 

        return Self::reverse(&mut s[..k])
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
        assert_eq!(Solution::reverse_str("abcdefg".to_string(), 2),"bacdfeg".to_string());
        assert_eq!(Solution::reverse_str("abcdefgh".to_string(), 2),"bacdfegh".to_string());
        assert_eq!(Solution::reverse_str("abcde".to_string(), 2),"bacde".to_string());
    }
}