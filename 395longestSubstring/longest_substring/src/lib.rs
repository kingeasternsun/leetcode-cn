/*
 * @Description: 395
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-28 14:22:15
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-28 15:14:46
 * @FilePath: /longest_substring/src/lib.rs
 */

pub struct Solution;

impl Solution {
    pub fn longest_substring(s: String, k: i32) -> i32 {
        let a = 'a' as u8;
        let bs = s.into_bytes();
        let mut cnt = [0; 26];
        bs.iter().for_each(|x| cnt[(x - a) as usize] += 1);

        // println!("{:?}",cnt);
        for win_len in (1..bs.len()+1).rev() {
            if Self::winlen_match(&bs[..], cnt, k, win_len) {
                return win_len as i32;
            }

            cnt[(bs[win_len-1] - a) as usize]-=1;
        }

        0
    }

    pub fn winlen_match(bs: &[u8], mut cnt: [i32; 26], k: i32, win_len: usize) -> bool {
        // println!("{:?}",cnt);
        if Self::window_match(&cnt[..], k) {
            return true;
        }

        //开始滑动窗口　
        for end in win_len..bs.len() {
            cnt[(bs[end] - 'a' as u8) as usize] += 1;
            cnt[(bs[end - win_len] - 'a' as u8) as usize] -= 1;
            if Self::window_match(&cnt[..], k) {
                return true;
            }
        }

        false
    }

    //判断是否满足
    pub fn window_match(cnt: &[i32], k: i32) -> bool {
        cnt.into_iter().all(|&x| (x == 0) || (x >= k))
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::longest_substring(String::from("abc"), 2),0);
        assert_eq!(Solution::longest_substring(String::from("aaabb"), 3), 3);
        assert_eq!(Solution::longest_substring(String::from("ababbc"), 2),5);
        assert_eq!(Solution::longest_substring(String::from("aaabb"), 4),0);
    }
}
