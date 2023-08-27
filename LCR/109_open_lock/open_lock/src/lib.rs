struct Solution;
use std::collections::{HashMap, HashSet};
impl Solution {
    pub fn open_lock(deadends: Vec<String>, target: String) -> i32 {
        let mut dead_map = std::collections::HashSet::new();
        for dead in deadends {
            dead_map.insert(dead.into_bytes());
        }

        let mut visited = HashMap::new();
        let cnt = 0;
        let mut min = i32::MAX;
        let mut cur = vec![b'0', b'0', b'0', b'0'];

        Self::dfs(
            &mut cur,
            &target.into_bytes(),
            &dead_map,
            &mut visited,
            cnt,
            &mut min,
        );
        min - 1
    }

    // 即将访问 cur
    fn dfs(
        cur: &mut Vec<u8>,
        target: &Vec<u8>,
        dead_map: &HashSet<Vec<u8>>,
        visited: &mut HashMap<Vec<u8>, i32>,
        cnt: i32,
        min: &mut i32,
    ) {
        let mut cnt = cnt;
        cnt += 1;

        if dead_map.contains(cur) {
            return;
        }

        if cur == target {
            *min = cnt.min(*min);
            return;
        }

        // 如果当前次数已经超过了已知的最小次数，就可以退出了
        if cnt >= *min {
            return;
        }

        // 如果这个状态之前访问过
        if let Some(&pre_cnt) = visited.get(cur) {
            // 次数比之前还多 或 一样，就不需要再尝试了
            if cnt >= pre_cnt {
                return;
            }
        }

        visited.insert(cur.clone(), cnt);

        for i in 0..4 {
            let tmp = cur[i];

            cur[i] = Self::pre(tmp);
            Self::dfs(cur, target, dead_map, visited, cnt, min);
            cur[i] = tmp;

            cur[i] = Self::next(tmp);
            Self::dfs(cur, target, dead_map, visited, cnt, min);
            cur[i] = tmp
        }
    }

    fn next(b: u8) -> u8 {
        if b < b'9' {
            return b + 1;
        }
        return b'0';
    }

    fn pre(b: u8) -> u8 {
        if b > b'0' {
            return b - 1;
        }
        return b'9';
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::open_lock(vec!["8888".to_string()], "0009".to_string()),
            1
        );
        assert_eq!(
            Solution::open_lock(
                vec![
                    "0201".to_string(),
                    "0101".to_string(),
                    "0102".to_string(),
                    "1212".to_string(),
                    "2002".to_string()
                ],
                "0202".to_string()
            ),
            1
        );
    }
}
