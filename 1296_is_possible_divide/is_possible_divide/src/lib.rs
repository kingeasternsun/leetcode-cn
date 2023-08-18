struct Solution;
impl Solution {
    pub fn is_possible_divede(nums: Vec<i32>, k: i32) -> bool {
        let k = k as usize;
        // 不能整除
        if nums.len() % k > 0 {
            return false;
        }

        let mut nums = nums;
        nums.sort_unstable();

        let mut map = std::collections::HashMap::new();
        nums.iter().for_each(|&x| {
            *map.entry(x).or_insert(0) += 1;
        });

        for x in nums.into_iter() {
            let v = map.get(&x).map_or(0, |v| v.clone());
            if v == 0 {
                continue;
            }

            for k in x..x + k as i32 {
                match map.get(&k) {
                    None => return false,
                    Some(&v) => {
                        if v == 1 {
                            map.remove(&k);
                        } else {
                            map.insert(k, v - 1);
                        }
                    }
                }
            }
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::is_possible_divede(vec![1, 2, 3, 3, 4, 4, 5, 6], 4),
            true
        );
        assert_eq!(
            Solution::is_possible_divede(vec![3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11], 3),
            true
        );
        assert_eq!(
            Solution::is_possible_divede(vec![3, 3, 2, 2, 1, 1], 3),
            true
        );
        assert_eq!(Solution::is_possible_divede(vec![1, 2, 3, 4], 3), false);
    }
}
