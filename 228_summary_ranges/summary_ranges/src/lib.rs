use core::num;

struct Solution;
impl Solution {
    // LC: 0ms 2.06mb
    pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
        if nums.len() == 0 {
            return vec![];
        }
        let mut pre = nums[0];
        let mut ret = vec![];

        for id in 0..nums.len() - 1 {
            if nums[id + 1] != nums[id] + 1 {
                if nums[id] == pre {
                    ret.push(format!("{}", pre));
                } else {
                    ret.push(format!("{}->{}", pre, nums[id]));
                }
                pre = nums[id + 1];
            }
        }

        if nums[nums.len() - 1] == pre {
            ret.push(format!("{}", pre));
        } else {
            ret.push(format!("{}->{}", pre, nums[nums.len() - 1]));
        }

        ret
    }

    // LC: 0ms 1.98mb
    pub fn summary_ranges1(nums: Vec<i32>) -> Vec<String> {
        if nums.len() == 0 {
            return vec![];
        }
        let mut pre = nums[0];
        let mut ret = vec![];

        nums.windows(2).for_each(|win| {
            if win[1] != win[0] + 1 {
                if win[0] == pre {
                    ret.push(format!("{}", pre));
                } else {
                    ret.push(format!("{}->{}", pre, win[0]));
                }
                pre = win[1];
            }
        });

        if let Some(&last) = nums.last() {
            if last == pre {
                ret.push(format!("{}", pre));
            } else {
                ret.push(format!("{}->{}", pre, last));
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
        assert_eq!(
            Solution::summary_ranges(vec![0, 1, 2, 4, 5, 7]),
            vec!["0->2".to_string(), "4->5".to_string(), "7".to_string()]
        );

        assert_eq!(
            Solution::summary_ranges(vec![0, 2, 3, 4, 6, 8, 9]),
            vec![
                "0".to_string(),
                "2->4".to_string(),
                "6".to_string(),
                "8->9".to_string()
            ]
        );
    }

    #[test]
    fn it_works1() {
        assert_eq!(
            Solution::summary_ranges1(vec![0, 1, 2, 4, 5, 7]),
            vec!["0->2".to_string(), "4->5".to_string(), "7".to_string()]
        );

        assert_eq!(
            Solution::summary_ranges1(vec![0, 2, 3, 4, 6, 8, 9]),
            vec![
                "0".to_string(),
                "2->4".to_string(),
                "6".to_string(),
                "8->9".to_string()
            ]
        );
    }
}
