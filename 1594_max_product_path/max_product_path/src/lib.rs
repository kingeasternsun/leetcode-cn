struct Solution;
impl Solution {
    // 由于在矩阵中，只能向右或向下移动，所以每个位置只能从左边或上边到达
    // 由于乘以负数后最大积会变为最小值，最小积会变为最大积，所以要同时记录最大积和最小积
    // 0ms 1.98mb
    pub fn max_product_path(grid: Vec<Vec<i32>>) -> i32 {
        let mut dp = vec![vec![(0_i64, 0_i64); grid[0].len()]; 2];
        dp[0][0] = (grid[0][0] as i64, grid[0][0] as i64);
        for col in 1..grid[0].len() {
            dp[0][col] = Self::mul(dp[0][col - 1], grid[0][col] as i64);
        }

        for row in 1..grid.len() {
            dp[row & 1][0] = Self::mul(dp[(row - 1) & 1][0], grid[row][0] as i64);
            for col in 1..grid[0].len() {
                let left = Self::mul(dp[row & 1][col - 1], grid[row][col] as i64);
                let up = Self::mul(dp[(row - 1) & 1][col], grid[row][col] as i64);
                dp[row & 1][col] = (left.0.min(up.0), left.1.max(up.1));
            }
        }

        let ret = (dp[(grid.len() - 1) & 1][grid[0].len() - 1].1 % 1000000007) as i32;
        if ret < 0 {
            return -1;
        }
        ret
    }

    fn mul(a: (i64, i64), b: i64) -> (i64, i64) {
        if b > 0 {
            return (a.0 * b, a.1 * b);
        }
        return (a.1 * b, a.0 * b);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::max_product_path(vec![vec![-1, -2, -3], vec![-2, -3, -3], vec![-3, -3, -2]]),
            -1
        );
        assert_eq!(
            Solution::max_product_path(vec![vec![1, -2, 1], vec![1, -2, 1], vec![3, -4, 1]]),
            8
        );
        assert_eq!(Solution::max_product_path(vec![vec![1, 3], vec![0, -4]]), 0);
    }
}
