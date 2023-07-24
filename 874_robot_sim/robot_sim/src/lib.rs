struct Solution;
impl Solution {
    pub fn robot_sim(commands: Vec<i32>, obstacles: Vec<Vec<i32>>) -> i32 {
        let mut obs_set = std::collections::HashSet::new();
        for obs in obstacles {
            obs_set.insert((obs[0], obs[1]));
        }
        let mut ret = 0;
        let dirs = vec![(0, 1), (1, 0), (0, -1), (-1, 0)];
        let mut cur_dir_id = 0;
        let mut cur_pos = (0, 0);
        for cmd in commands {
            if cmd == -1 {
                // turn left
                cur_dir_id = (cur_dir_id + 1) % 4;
            } else if cmd == -2 {
                cur_dir_id = (cur_dir_id + 3) % 4;
            } else {
                for _ in 0..cmd {
                    let next_pos = (
                        cur_pos.0 + dirs[cur_dir_id].0,
                        cur_pos.1 + dirs[cur_dir_id].1,
                    );
                    if next_pos.0 < 0 || next_pos.1 < 0 || obs_set.get(&next_pos).is_some() {
                        break;
                    }

                    cur_pos = next_pos;
                    let new_dis = cur_pos.0 * cur_pos.0 + cur_pos.1 * cur_pos.1;
                    if new_dis > ret {
                        ret = new_dis;
                    }
                }
            }
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::robot_sim(vec![4, -1, 3], vec![]), 25);
        assert_eq!(
            Solution::robot_sim(vec![4, -1, 4, -2, 4], vec![vec![2, 4]]),
            65
        );
    }
}
