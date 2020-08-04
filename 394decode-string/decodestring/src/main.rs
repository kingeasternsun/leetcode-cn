fn main() {
    // println!("Hello, world!");
    let b = Solution::decode_string(String::from("3[a]2[bc]"));
    println!("{}",b);
}

pub struct Solution;
// use std::collections::VecDeque;
impl Solution {
    pub fn decode_string(s: String) -> String {
        let bytes = s.as_bytes();
        let mut res: Vec<Vec<u8>> = Vec::new();

        // let mut beg = Some(0);
        // let mut end = 0;
        for i in 0..bytes.len() {
            if bytes[i] == '[' as u8 {
                continue;
            }

            if bytes[i] == ']' as u8 {
                let tmp_str = res.pop().unwrap();
                let num = String::from_utf8(res.pop().unwrap())
                    .unwrap()
                    .parse::<i32>()
                    .unwrap();
                let tmp_str = tmp_str.repeat(num as usize);

                //判断前一个还是是字符串 再次合并到前一个字符串中, 这里只需要向前合并一次
                if res.last().is_some() && res.last().unwrap()[0].is_ascii_alphabetic(){
                    res.last_mut().unwrap().extend(tmp_str.iter().copied());
                }else{
                    res.push(tmp_str);
                }
            }

            if bytes[i].is_ascii_alphabetic() {
                //如果当前是字母，前一个也是字母，追加到上一个vec
                if res.last().is_some() && res.last().unwrap()[0].is_ascii_alphabetic(){
                    res.last_mut().unwrap().push(bytes[i]);
                }else{
                    res.push(vec![bytes[i]]);
                }
            }

            if bytes[i].is_ascii_digit() {
                //如果当前是数字，前一个也是数字，追加到上一个vec
                if i >= 1 && bytes[i - 1].is_ascii_digit() {
                    res.last_mut().unwrap().push(bytes[i]);
                } else {
                    res.push(vec![bytes[i]]);
                }
            }
        }

        if let Some(res1)=res.pop() {
           return  String::from_utf8(res1).unwrap()
        }

        return s
        
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_104() {
        assert_eq!(Solution::decode_string(String::from("3[a]2[bc]")), String::from("aaabcbc"));
        assert_eq!(Solution::decode_string(String::from("3[a2[c]]")), String::from("accaccacc"));
        assert_eq!(Solution::decode_string(String::from("2[abc]3[cd]ef")), String::from("abcabccdcdcdef"));
        assert_eq!(Solution::decode_string(String::from("abc3[cd]xyz")), String::from("abccdcdcdxyz"));
        assert_eq!(Solution::decode_string(String::from("")), String::from(""));
    }
}