struct Solution;
impl Solution {
    // the beauty of iterator
    // 20ms 4.36mb
    pub fn sequence_reconstruction(nums: Vec<i32>, sequences: Vec<Vec<i32>>) -> bool {
        let mut map = std::collections::HashMap::new();
        sequences.into_iter().for_each(|seq| {
            seq.windows(2).for_each(|win| {
                map.entry(win[0])
                    .or_insert(std::collections::HashSet::new())
                    .insert(win[1]);
            });
        });

        !nums
            .windows(2)
            .any(|win| map.get(&win[0]).and_then(|set| set.get(&win[1])).is_none())
    }
    pub fn sequence_reconstruction1(nums: Vec<i32>, sequences: Vec<Vec<i32>>) -> bool {
        let mut map = std::collections::HashMap::new();
        sequences.into_iter().for_each(|seq| {
            seq.windows(2).for_each(|win| {
                map.entry(win[0])
                    .or_insert(std::collections::HashSet::new())
                    .insert(win[1]);
            });
        });

        // 写法二:
        nums.windows(2)
            .all(|win| map.get(&win[0]).map_or(false, |set| set.contains(&win[1])))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::sequence_reconstruction(vec![1, 2, 3], vec![vec![1, 2], vec![1, 3]]),
            false
        );
        assert_eq!(
            Solution::sequence_reconstruction(vec![1, 2, 3], vec![vec![1, 2]]),
            false
        );
        assert_eq!(
            Solution::sequence_reconstruction(
                vec![1, 2, 3],
                vec![vec![1, 2], vec![1, 3], vec![2, 3]]
            ),
            true
        );
    }

    #[test]
    fn it_works1() {
        assert_eq!(
            Solution::sequence_reconstruction1(vec![1, 2, 3], vec![vec![1, 2], vec![1, 3]]),
            false
        );
        assert_eq!(
            Solution::sequence_reconstruction1(vec![1, 2, 3], vec![vec![1, 2]]),
            false
        );
        assert_eq!(
            Solution::sequence_reconstruction1(
                vec![1, 2, 3],
                vec![vec![1, 2], vec![1, 3], vec![2, 3]]
            ),
            true
        );
    }
}
