struct Solution;
impl Solution {
    // 36ms 3.83mb
    pub fn maximum_sum(nums: Vec<i32>) -> i32 {
        let mut map = std::collections::HashMap::new();
        nums.into_iter().for_each(|n| {
            let entry = map.entry(Self::get_sum(n)).or_insert(vec![]);
            entry.push(n);
            if entry.len() > 2 {
                entry.sort_unstable_by_key(|k| -k);
                entry.pop();
            }
        });

        let mut ret = -1;
        for (_, v) in map {
            if v.len() == 2 {
                ret = ret.max(v[0] + v[1]);
            }
        }

        ret
    }
    fn get_sum(mut n: i32) -> i32 {
        let mut ret = 0;
        while n > 0 {
            ret += n % 10;
            n = n / 10;
        }
        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::maximum_sum(vec![18, 43, 36, 13, 7]), 54);
        assert_eq!(Solution::maximum_sum(vec![10, 12, 19, 14]), -1);
    }
}
