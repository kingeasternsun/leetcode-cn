struct Solution;
use std::collections::HashMap;
impl Solution {
    // beautify of iterator
    pub fn minimum_time(n: i32, relations: Vec<Vec<i32>>, time: Vec<i32>) -> i32 {
        let n = n as usize;
        let mut pre_classes: Vec<Vec<usize>> = vec![Vec::new(); n + 1];
        for rt in relations {
            pre_classes[rt[1] as usize].push(rt[0] as usize);
        }

        let mut memo = HashMap::new();
        (1..n + 1)
            .into_iter()
            .map(|x| Self::dp(x, &pre_classes, &mut memo, &time))
            .max()
            .unwrap()
    }

    fn dp(
        i: usize,
        pre_classes: &Vec<Vec<usize>>,
        memo: &mut HashMap<usize, i32>,
        time: &Vec<i32>,
    ) -> i32 {
        match memo.get(&i) {
            Some(ret) => return *ret,
            None => {
                let ret = pre_classes[i as usize]
                    .iter()
                    .map(|pre| Self::dp(*pre, pre_classes, memo, time))
                    .max()
                    .unwrap_or(0)
                    + time[i - 1];
                memo.insert(i, ret);
                ret
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::minimum_time(3, vec![vec![1, 3], vec![2, 3]], vec![3, 2, 5]),
            8
        );
        assert_eq!(
            Solution::minimum_time(
                5,
                vec![vec![1, 5], vec![2, 5], vec![3, 5], vec![3, 4], vec![4, 5]],
                vec![1,2,3,4,5]
            ),
            12
        );
        assert_eq!(
            Solution::minimum_time(
                5,
                vec![],
                vec![1,2,3,4,5]
            ),
            5
        );
    }
}
