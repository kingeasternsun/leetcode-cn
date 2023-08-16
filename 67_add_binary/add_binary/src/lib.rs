struct Solution;
impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        

    }

}



#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::add_binary("11".to_string(),"1".to_string()), "100".to_string());
        assert_eq!(Solution::add_binary("1010".to_string(),"1011".to_string()), "10101".to_string());
    }
}
