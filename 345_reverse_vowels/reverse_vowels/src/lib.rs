struct Solution;

impl Solution {
    pub fn reverse_vowels(s: String) -> String {
        let vowels = String::from("aeiou").into_bytes();

        let mut s = s.into_bytes();
        let mut ids = vec![];
        s.iter().enumerate().for_each(|x| {
            if vowels.iter().any(|v| *v == x.1.to_ascii_lowercase()) {
                ids.push(x.0);
            }
        });

        println!("{:?}",ids);

        if ids.len() < 2 {
            return String::from_utf8(s).unwrap();
        }

        let mut left = 0;
        let mut right = ids.len() - 1;
        while left < right {
            s.swap(ids[left],ids[right]);
            left += 1;
            right -= 1;
        }

        return String::from_utf8(s).unwrap();
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::reverse_vowels(String::from("hello")), String::from("holle"));
        assert_eq!(Solution::reverse_vowels(String::from("leetcode")), String::from("leotcede"));
    }
}
