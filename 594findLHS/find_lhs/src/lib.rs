pub struct Solution;
impl Solution {
    pub fn find_lhs(nums: Vec<i32>) -> i32 {
        // use hash to get each item count
        let mut h = std::collections::BTreeMap::new();
        for v in nums {
            *h.entry(v).or_insert(0) += 1;
        }

        let mut res = 0;
        for (k1, v1) in h.iter() {
            if let Some(v2) = h.get(&(k1 + 1)) {
                if v2 + v1 > res {
                    res = v2 + v1;
                }
            }
        }

        res
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::find_lhs(vec![1, 3, 2, 2, 5, 2, 3, 7]), 5);
        assert_eq!(Solution::find_lhs(vec![1, 2, 3, 4]), 2);
        assert_eq!(Solution::find_lhs(vec![1, 1, 1, 1]), 0);
    }
}
