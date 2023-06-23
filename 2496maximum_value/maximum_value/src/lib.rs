struct Solution;
impl Solution {
    pub fn maximum_value1(strs: Vec<String>) -> i32 {
        let mut ret = 0;
        for str in &strs {
            let n = match str.parse::<i32>() {
                Ok(n) => n,
                Err(_) => str.len() as i32,
            };
            if n > ret {
                ret = n;
            }
        }

        ret
    }
    pub fn maximum_value(strs: Vec<String>) -> i32 {
        strs.iter().fold(0, |ans, item| {
            ans.max(item.parse::<i32>().unwrap_or(item.len() as i32))
        })
    }
}

#[cfg(test)]
mod tests {
    use std::{assert_eq, vec};

    use crate::Solution;
    #[test]
    fn it_works1() {
        assert_eq!(
            Solution::maximum_value1(vec![
                "alic3".to_string(),
                "bob".to_string(),
                "3".to_string(),
                "4".to_string(),
                "00000".to_string()
            ]),
            5
        );
        assert_eq!(
            Solution::maximum_value1(vec![
                "1".to_string(),
                "01".to_string(),
                "001".to_string(),
                "0001".to_string()
            ]),
            1
        );
    }
    fn it_works() {
        assert_eq!(
            Solution::maximum_value(vec![
                "alic3".to_string(),
                "bob".to_string(),
                "3".to_string(),
                "4".to_string(),
                "00000".to_string()
            ]),
            5
        );
        assert_eq!(
            Solution::maximum_value(vec![
                "1".to_string(),
                "01".to_string(),
                "001".to_string(),
                "0001".to_string()
            ]),
            1
        );
    }
}
