struct Solution;
impl Solution {
    // 4ms 2.8mb  排序 加 滑动窗口
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        if nums.len() < 2 {
            return nums.len() as i32;
        }
        let mut nums = nums;
        let mut pre = None;
        let mut ret = 1;
        nums.sort_unstable();
        // 滑动窗口
        nums.windows(2).for_each(|win| {
            if win[0] == win[1] {
                // just do nothing
            } else if win[1] == win[0] + 1 {
                match pre {
                    None => pre = Some(2),
                    Some(pre_len) => pre = Some(pre_len + 1),
                }
            } else {
                match pre {
                    None => pre = None,
                    Some(pre_len) => {
                        ret = ret.max(pre_len);
                        pre = None;
                    }
                }
            }
        });

        if let Some(pre_len) = pre {
            ret = ret.max(pre_len);
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::longest_consecutive(vec![0, 3, 7, 2, 5, 8, 4, 6, 0, 1]),
            9
        );
        assert_eq!(Solution::longest_consecutive(vec![100, 4, 200, 1, 3, 2]), 4);

        assert_eq!(Solution::longest_consecutive(vec![0, 1]), 2);
        assert_eq!(Solution::longest_consecutive(vec![0, 0]), 1);
        assert_eq!(Solution::longest_consecutive(vec![1, 1]), 1);
        assert_eq!(Solution::longest_consecutive(vec![1, 1, 2]), 2);
        assert_eq!(Solution::longest_consecutive(vec![0]), 1);
        assert_eq!(Solution::longest_consecutive(vec![1, 2, 0, 1]), 3);
    }
}
