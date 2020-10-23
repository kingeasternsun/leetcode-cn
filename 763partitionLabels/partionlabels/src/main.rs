fn main() {
    println!("Hello, world!");
    Solution::partition_labels(String::from("ababcbacadefegdehijhklij"));

    let hello = "Здравствуйте";

    let s = &hello[0..2];
    println!("{:?}", s);
    //
}

pub struct Solution;

impl Solution {
    pub fn partition_labels(s: String) -> Vec<i32> {
        let mut part = vec![(2147483647, -1); 26];
        let mut i = 0;
        for b in s.bytes() {
            let part_id = (b - 'a' as u8) as usize;
            let (left, right) = part[part_id];
            if i < left {
                part[part_id].0 = i;
            }

            if i > right {
                part[part_id].1 = i;
            }
            i = i + 1;
        }

        part.sort_by(|a, b| a.0.cmp(&b.0));
        // println!("{:?}",part);

        for i in 1..26 {
            if part[i].1 == -1 {
                continue;
            }

            if part[i].0 > part[i - 1].1 {
                continue;
            }

            if part[i].1 < part[i - 1].1 {
                part[i].1 = part[i - 1].1;
                // part[i].0 = part[i-1].0;
                // part[i-1].1 = -1;
                // continue
            }

            part[i].0 = part[i - 1].0;
            part[i - 1].1 = -1;
        }

        let mut res = Vec::new();
        for i in 0..26 {
            if part[i].1 == -1 {
                continue;
            }
            res.push(part[i].1 - part[i].0 as i32 + 1);
        }
        return res;
    }
}
