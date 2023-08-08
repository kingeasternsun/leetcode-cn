struct Solution;

use std::collections::HashMap;
impl Solution {
    pub fn sum_of_unique0(nums: Vec<i32>) -> i32 {
        let mut map: HashMap<i32, i32> = HashMap::new();
        for n in nums {
            *map.entry(n).or_default() += 1;
        }

        map.into_iter().filter(|x| x.1 == 1).map(|x| x.0).sum()
    }

    pub fn sum_of_unique(nums: Vec<i32>) -> i32 {
        let mut map: HashMap<i32, i32> = HashMap::new();
        for n in nums {
            *map.entry(n).or_default() += 1;
        }

        map.into_iter()
            .filter_map(|(key, value)| if value == 1 { Some(key) } else { None })
            .sum()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works0() {
        assert_eq!(Solution::sum_of_unique0(vec![1, 2, 3, 2]), 4);
        assert_eq!(Solution::sum_of_unique0(vec![1, 1, 1, 1, 1]), 0);
        assert_eq!(Solution::sum_of_unique0(vec![1, 2, 3, 4, 5]), 15);
        assert_eq!(Solution::sum_of_unique0(vec![1]), 1);
        assert_eq!(Solution::sum_of_unique0(vec![1, 1]), 0);
    }

    #[test]
    fn it_works() {
        assert_eq!(Solution::sum_of_unique(vec![1, 2, 3, 2]), 4);
        assert_eq!(Solution::sum_of_unique(vec![1, 1, 1, 1, 1]), 0);
        assert_eq!(Solution::sum_of_unique(vec![1, 2, 3, 4, 5]), 15);
        assert_eq!(Solution::sum_of_unique(vec![1]), 1);
        assert_eq!(Solution::sum_of_unique(vec![1, 1]), 0);
    }
}
