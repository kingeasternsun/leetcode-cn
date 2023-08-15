struct Solution;

enum Direction {
    Up,
    Down,
    Equal,
}

impl Solution {
    // 0ms 1.98mb
    pub fn longest_mountain(arr: Vec<i32>) -> i32 {
        if arr.len() < 3 {
            return 0;
        }
        let mut ret = 0;
        let mut pre_dir = Direction::Equal;
        let mut pre_len = 0;
        arr.windows(2).for_each(|x| {
            match pre_dir {
                Direction::Equal => {
                    if x[0] < x[1] {
                        // 上山开始
                        pre_len = 2;
                        pre_dir = Direction::Up;
                    }
                }
                Direction::Up => {
                    if x[0] < x[1] {
                        // 继续上山
                        pre_len += 1;
                    } else if x[0] == x[1] {
                        pre_len = 0;
                        pre_dir = Direction::Equal;
                    } else {
                        // 到达山顶，开始下山
                        pre_len += 1;
                        pre_dir = Direction::Down;
                    }
                }
                Direction::Down => {
                    if x[0] < x[1] {
                        // 到达山谷，开始爬新的山
                        ret = ret.max(pre_len);
                        pre_len = 2;
                        pre_dir = Direction::Up;
                    } else if x[0] == x[1] {
                        ret = ret.max(pre_len);
                        pre_len = 0;
                        pre_dir = Direction::Equal;
                    } else {
                        // 继续下山
                        pre_len += 1;
                    }
                }
            }
        });

        // 特殊case的处理
        match pre_dir {
            Direction::Down => ret = ret.max(pre_len),
            _ => {}
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::longest_mountain(vec![2, 1, 4, 7, 3, 2, 5]), 5);

        assert_eq!(Solution::longest_mountain(vec![2, 2, 1, 2]), 0);
        assert_eq!(Solution::longest_mountain(vec![2, 2, 4, 2]), 3);
        assert_eq!(Solution::longest_mountain(vec![2, 2, 4, 4, 2]), 0);
        assert_eq!(Solution::longest_mountain(vec![2, 2, 2]), 0);
        assert_eq!(Solution::longest_mountain(vec![2, 2, 4]), 0);
        assert_eq!(Solution::longest_mountain(vec![2, 2, 1]), 0);
        assert_eq!(Solution::longest_mountain(vec![2, 4]), 0);
        assert_eq!(Solution::longest_mountain(vec![2, 1]), 0);
        assert_eq!(Solution::longest_mountain(vec![1, 2]), 0);
        assert_eq!(
            Solution::longest_mountain(vec![875, 884, 239, 731, 723, 685]),
            4
        );
    }
}
