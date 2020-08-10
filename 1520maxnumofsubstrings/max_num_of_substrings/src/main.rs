fn main() {
    println!("Hello, world!");
    println!(
        "{:?}",
        Solution::max_num_of_substrings(String::from("abab"))
    );
}

// https://www.bilibili.com/video/BV1yz4y1D7p3
pub struct Solution {}
use std::cmp;
impl Solution {
    pub fn max_num_of_substrings(s: String) -> Vec<String> {
        let mut res: Vec<String> = Vec::new();
        let mut left: Vec<i32> = vec![2147483647; 26];
        let mut right: Vec<i32> = vec![-2147483648; 26];
        let base = 'a' as u8;

        let s_b = s.as_bytes();

        for i in 0..s_b.len() {
            left[(s_b[i] - base) as usize] =
                cmp::min(i as i32, left[(s_b[i] as u8 - base) as usize]);
            right[(s_b[i] - base) as usize] =
                cmp::max(i as i32, right[(s_b[i] as u8 - base) as usize]);
        }

        let extention = |i: i32| -> i32 {
            let mut p = right[(s_b[i as usize] - base) as usize];
            let mut cur = i as usize;
            while cur < p as usize {
                if left[(s_b[cur] - base) as usize] < i {
                    return -1;
                }

                if right[(s_b[cur] - base) as usize] > p {
                    p = right[(s_b[cur] - base) as usize]
                }

                cur += 1;
            }

            return p;
        };

        let mut last: i32 = -1;
        for i in 0..s_b.len() {
            if i != left[(s_b[i] as u8 - base) as usize] as usize {
                continue;
            }

            let p = extention(i as i32);
            if p == -1 {
                continue;
            }

            if i as i32 > last {
                res.push(String::from(&s[i..(p as usize + 1)]));
            } else {
                res.pop();
                res.push(String::from(&s[i..(p as usize + 1)]));
            }

            last = p
        }

        res
    }
}
