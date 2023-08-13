struct Solution;

impl Solution {
    // 首先 可以证明，窃取至少K个 和 只能窃取K个 的窃取能力是一样的。
    // 假设窃取  K+n 个房屋，那么我们可以把n个房屋去掉，留下的K个房屋包含最大值的话，就跟窃取 K个 的结果是一样的
    pub fn min_capability(nums: Vec<i32>, k: i32) -> i32 {
        let mut min_cap = i32::MAX;
        Self::dp(&nums[0..], k as usize, 0, &mut min_cap);
        min_cap
    }

    // 超时
    fn dp(nums: &[i32], k: usize, max_cap: i32, min_cap: &mut i32) {
        if k == 1 {
            match nums.iter().min() {
                None => return,
                Some(&v) => {
                    // 从剩下的nums中挑一个最小的值，然后和 max_cap 比较
                    // 为什么是选 nums 中的最小值呢，因为是求最小能力，所以在最后一个可以窃取的房屋里，肯定要选最少的
                    *min_cap = v.max(max_cap).min(*min_cap);
                    return;
                }
            }
        }

        if (nums.len() + 1) / 2 < k {
            return;
        }

        // steal nums[0]
        Self::dp(&nums[2..], k - 1, max_cap.max(nums[0]), min_cap);
        // not steal nums[0]
        Self::dp(&nums[1..], k, max_cap, min_cap);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::min_capability(vec![2, 3, 5, 9], 2), 5);
        assert_eq!(Solution::min_capability(vec![2, 7, 9, 3, 1], 2), 2);
    }
}
