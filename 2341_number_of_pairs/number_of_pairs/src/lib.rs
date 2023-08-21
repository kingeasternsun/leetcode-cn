struct Solution;
impl Solution {
    // 0ms 1.9mb
    pub fn number_of_pairs(nums: Vec<i32>) -> Vec<i32> {
        let mut set = std::collections::HashSet::new();
        let mut pair_num = 0;
        nums.iter().for_each(|x| {
            if set.contains(x) {
                set.remove(x);
                pair_num += 1;
            } else {
                set.insert(x.clone());
            }
        });

        vec![pair_num, set.len() as i32]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::number_of_pairs(vec![1, 3, 2, 1, 3, 2, 2]),
            vec![3, 1]
        );

        assert_eq!(Solution::number_of_pairs(vec![1, 1]), vec![1, 0]);

        assert_eq!(Solution::number_of_pairs(vec![1]), vec![0, 1]);
    }
}
