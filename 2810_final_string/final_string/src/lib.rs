struct Solution;
impl Solution {
    // 0ms 1.96mb
    pub fn final_string(s: String) -> String {
        let mut ret = Vec::with_capacity(s.as_bytes().len());
        for b in s.as_bytes(){
            if *b == b'i'{
                ret.reverse();
            }else{
                ret.push(b.clone());
            }
        }
        String::from_utf8(ret).unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::final_string("string".to_string()), "rtsng".to_string());
        assert_eq!(Solution::final_string("poiinter".to_string()), "ponter".to_string());
    }
}
