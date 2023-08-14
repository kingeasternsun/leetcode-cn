struct Solution;

use std::cmp::Reverse;
impl Solution {
    // 这道题目中，关键点在于移除气球而不是更改气球的颜色，等价于把连续重复颜色的气球进行去重
    // 移除气球时，肯定选移除时间最小的， 利用堆排序 获取 连续重复颜色的气球中移除时间最小的
    // 需要注意：最后一个气球也是重复的情况
    // LC:36ms 3.16mb
    pub fn min_cost(colors: String, needed_time: Vec<i32>) -> i32 {
        let mut ret = 0;
        let mut min_heap = std::collections::BinaryHeap::new();
        let bytes = colors.as_bytes();
        for i in 0..bytes.len() {
            if i == bytes.len() - 1 || bytes[i] != bytes[i + 1] {
                if min_heap.len() > 0 {
                    min_heap.push(Reverse(needed_time[i]));

                    // 此时，min_heap 里面都是重复颜色的气球
                    // 每次从堆顶移除的就是移除时间最小的
                    while min_heap.len() > 1 {
                        ret += min_heap.pop().unwrap().0;
                    }

                    // 这里最后要清除掉，不然后面会判断错误
                    min_heap.clear()
                }
            } else {
                // i == bytes.len()-1 && bytes[i]==bytes[i+1]
                min_heap.push(Reverse(needed_time[i]));
            }
        }

        ret
    }

    // 对于连续的重复颜色的气球，只需要将移除时间最大的气球之外的气球移除即可
    // 利用冒泡法，求出连续重复颜色的气球中除了移除时间最大的气球之外的气球移除时间之和
    pub fn min_cost1(colors: String, needed_time: Vec<i32>) -> i32 {
        let mut ret = 0;
        let bytes = colors.as_bytes();
        let mut needed_time = needed_time;
        for i in 0..bytes.len() - 1 {
            if bytes[i] == bytes[i + 1] {
                if needed_time[i] > needed_time[i + 1] {
                    needed_time.swap(i, i + 1);
                }
                ret += needed_time[i];
            }
        }

        ret
    }

    // the beauty of iterator
    pub fn min_cost2(colors: String, needed_time: Vec<i32>) -> i32 {
        let bytes = colors.as_bytes();
        let mut needed_time = needed_time;
        bytes.windows(2).enumerate().fold(0, |acc, (i, win)| {
            if win[0] == win[1] {
                if needed_time[i] > needed_time[i + 1] {
                    needed_time.swap(i, i + 1);
                }
                acc + needed_time[i]
            } else {
                acc
            }
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::min_cost(String::from("abaac"), vec![1, 2, 3, 4, 5]),
            3
        );
        assert_eq!(Solution::min_cost(String::from("abc"), vec![1, 2, 3]), 0);
        assert_eq!(
            Solution::min_cost(String::from("aabaa"), vec![1, 2, 3, 4, 1]),
            2
        );
        assert_eq!(Solution::min_cost(String::from("aaa"), vec![1, 2, 3]), 3);
        assert_eq!(Solution::min_cost(String::from("aaa"), vec![1, 3, 2]), 3);
        assert_eq!(Solution::min_cost(String::from("aaa"), vec![2, 1, 3]), 3);
        assert_eq!(Solution::min_cost(String::from("aaa"), vec![2, 3, 1]), 3);
        assert_eq!(Solution::min_cost(String::from("aaa"), vec![3, 2, 1]), 3);
        assert_eq!(Solution::min_cost(String::from("aaa"), vec![3, 1, 2]), 3);

        assert_eq!(
            Solution::min_cost(String::from("aaab"), vec![1, 2, 3, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost(String::from("aaab"), vec![1, 3, 2, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost(String::from("aaab"), vec![2, 1, 3, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost(String::from("aaab"), vec![2, 3, 1, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost(String::from("aaab"), vec![3, 2, 1, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost(String::from("aaab"), vec![3, 1, 2, 3]),
            3
        );
    }

    #[test]
    fn it_works1() {
        assert_eq!(
            Solution::min_cost1(String::from("abaac"), vec![1, 2, 3, 4, 5]),
            3
        );
        assert_eq!(Solution::min_cost1(String::from("abc"), vec![1, 2, 3]), 0);
        assert_eq!(
            Solution::min_cost1(String::from("aabaa"), vec![1, 2, 3, 4, 1]),
            2
        );
        assert_eq!(Solution::min_cost1(String::from("aaa"), vec![1, 2, 3]), 3);
        assert_eq!(Solution::min_cost1(String::from("aaa"), vec![1, 3, 2]), 3);
        assert_eq!(Solution::min_cost1(String::from("aaa"), vec![2, 1, 3]), 3);
        assert_eq!(Solution::min_cost1(String::from("aaa"), vec![2, 3, 1]), 3);
        assert_eq!(Solution::min_cost1(String::from("aaa"), vec![3, 2, 1]), 3);
        assert_eq!(Solution::min_cost1(String::from("aaa"), vec![3, 1, 2]), 3);

        assert_eq!(
            Solution::min_cost1(String::from("aaab"), vec![1, 2, 3, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost1(String::from("aaab"), vec![1, 3, 2, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost1(String::from("aaab"), vec![2, 1, 3, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost1(String::from("aaab"), vec![2, 3, 1, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost1(String::from("aaab"), vec![3, 2, 1, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost1(String::from("aaab"), vec![3, 1, 2, 3]),
            3
        );
    }

    #[test]
    fn it_works2() {
        assert_eq!(
            Solution::min_cost2(String::from("abaac"), vec![1, 2, 3, 4, 5]),
            3
        );
        assert_eq!(Solution::min_cost2(String::from("abc"), vec![1, 2, 3]), 0);
        assert_eq!(
            Solution::min_cost2(String::from("aabaa"), vec![1, 2, 3, 4, 1]),
            2
        );
        assert_eq!(Solution::min_cost2(String::from("aaa"), vec![1, 2, 3]), 3);
        assert_eq!(Solution::min_cost2(String::from("aaa"), vec![1, 3, 2]), 3);
        assert_eq!(Solution::min_cost2(String::from("aaa"), vec![2, 1, 3]), 3);
        assert_eq!(Solution::min_cost2(String::from("aaa"), vec![2, 3, 1]), 3);
        assert_eq!(Solution::min_cost2(String::from("aaa"), vec![3, 2, 1]), 3);
        assert_eq!(Solution::min_cost2(String::from("aaa"), vec![3, 1, 2]), 3);

        assert_eq!(
            Solution::min_cost2(String::from("aaab"), vec![1, 2, 3, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost2(String::from("aaab"), vec![1, 3, 2, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost2(String::from("aaab"), vec![2, 1, 3, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost2(String::from("aaab"), vec![2, 3, 1, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost2(String::from("aaab"), vec![3, 2, 1, 3]),
            3
        );
        assert_eq!(
            Solution::min_cost2(String::from("aaab"), vec![3, 1, 2, 3]),
            3
        );
    }
}
