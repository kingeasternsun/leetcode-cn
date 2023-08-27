struct Solution;

impl Solution {
    // 4ms 2.83
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        if intervals.len() == 0 {
            return ret;
        }
        let mut intervals = intervals;
        intervals.sort_unstable_by_key(|x| x[0]);
        let mut pre = (intervals[0][0], intervals[0][1]);
        for inter in intervals {
            if inter[0] <= pre.1 {
                pre.1 = pre.1.max(inter[1]);
            } else {
                ret.push(vec![pre.0, pre.1]);
                pre = (inter[0], inter[1]);
            }
        }

        ret.push(vec![pre.0, pre.1]);

        ret
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::merge(vec![vec![1, 3], vec![2, 6], vec![8, 10], vec![15, 18]]),
            vec![vec![1, 6], vec![8, 10], vec![15, 18]]
        );

        assert_eq!(
            Solution::merge(vec![vec![1, 4], vec![4, 5]]),
            vec![vec![1, 5]]
        );
    }
}
