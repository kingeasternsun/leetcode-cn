fn main() {
    println!("Hello, world!");
}

pub struct Solution;
use std::cmp;
impl Solution {
    pub fn longest_awesome(s: String) -> i32 {
        let mut pos = [2_147_483_647i32; 1024];
        pos[0] = -1;

        let mut status = 0;
        let base = '0' as u8;
        let mut ans = 0;

        for i in 0..s.len() {
            let mask = 1 << (s.as_bytes()[i] - base);
            status ^= mask;
            ans = cmp::max(ans, i as i32 - pos[status]);

            pos[status] = cmp::min(pos[status], i as i32);

            //如果是0，表示当前s[0:i]就是一个回文串，就可以提前返回
            if status == 0 {
                continue;
            }

            for j in 0..10 {
                ans = cmp::max(ans, i as i32 - pos[status ^ (1 << j)]);

                //如果是0，表示当前s[0:i]就是一个回文串，就可以提前返回
                if status ^ (1 << j) == 0 {
                    break;
                }
                // mask = 1<<j;
            }
        }

        ans
    }
}
