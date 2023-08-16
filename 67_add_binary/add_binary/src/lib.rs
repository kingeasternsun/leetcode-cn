struct Solution;
impl Solution {
    // 两次旋转
    pub fn add_binary(a: String, b: String) -> String {
        let mut ret = Vec::new();
        let mut a_bytes = a.into_bytes();
        let mut b_bytes = b.into_bytes();
        a_bytes.reverse();
        b_bytes.reverse();
        let max_len = a_bytes.len().max(b_bytes.len());

        let mut jin = 0;
        for i in (0..max_len) {
            let mut sum = jin + a_bytes.get(i).unwrap_or(&b'0').clone() - b'0'
                + b_bytes.get(i).unwrap_or(&b'0').clone()
                - b'0';
            jin = sum >> 1;
            sum = sum & 1;
            ret.push(sum + b'0');
        }
        if jin > 0 {
            ret.push(jin + b'0');
        }
        ret.reverse();

        String::from_utf8(ret).unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::add_binary("11".to_string(), "1".to_string()),
            "100".to_string()
        );
        assert_eq!(
            Solution::add_binary("1010".to_string(), "1011".to_string()),
            "10101".to_string()
        );
    }
}
