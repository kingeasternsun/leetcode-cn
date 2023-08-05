struct Solution;

use std::cmp::Ordering::{Greater, Less};
use std::collections::HashMap;
impl Solution {
    // 376ms 6.5mb
    pub fn max_taxi_earnings(n: i32, mut rides: Vec<Vec<i32>>) -> i64 {
        // 首先根据每个乘客的终点位置排序
        rides.sort_unstable_by(|a, b| a[1].cmp(&b[1]));
        let mut cache: HashMap<usize, i64> = HashMap::new();

        // 计算到每个乘客的最大收益
        (0..rides.len())
            .map(|r| Self::dp(r, &rides, &mut cache))
            .max()
            .unwrap_or_default()
    }

    // 计算到cur乘客的最大收益， 等于这个乘客起始点之前可以接单的乘客的最大收益的最大值加上当前乘客收益
    // cur乘客起始点之前可以接的乘客，他们的终点需要要在 cur 乘客的起始点或起始点之前，就用二分法查询，这也是为什么上面要基于乘客的终点排序
    fn dp(cur: usize, rides: &Vec<Vec<i32>>, cache: &mut HashMap<usize, i64>) -> i64 {
        if let Some(v) = cache.get(&cur) {
            return *v;
        }

        let cur_start = rides[cur][0];
        let partion_point = rides
            .binary_search_by(|x| if x[1] <= cur_start { Less } else { Greater })
            .unwrap_or_else(|x| x);

        let mut max = 0;
        let mut max_start = 0;
        for id in (0..partion_point).rev() {
            // 如果这个乘客在cur前面的一个乘客的前面，显然不是最优
            if rides[id][1] <= max_start {
                break;
            }

            max = max.max(Self::dp(id, rides, cache));
            max_start = max_start.max(rides[id][0]);
        }

        max = max + (rides[cur][1] - rides[cur][0] + rides[cur][2]) as i64;
        cache.insert(cur, max);
        return max;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::max_taxi_earnings(5, vec![vec![2, 5, 4], vec![1, 5, 1]]),
            7
        );
        assert_eq!(
            Solution::max_taxi_earnings(
                20,
                vec![
                    vec![1, 6, 1],
                    vec![3, 10, 2],
                    vec![10, 12, 3],
                    vec![11, 12, 2],
                    vec![12, 15, 2],
                    vec![13, 18, 1]
                ]
            ),
            20
        );
    }
}
