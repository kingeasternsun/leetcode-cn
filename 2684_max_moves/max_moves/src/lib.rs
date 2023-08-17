struct Solution;
impl Solution {
    // 跟 最大乘积 类似，利用dp解决
    /*

        根据题目，可以知道 dp[row][col] 只能由 dp[row-1][col-1],dp[row][col-1],dp[row+1][col-1]得到,
        按照列做dp
    ┌──────────────┐
    │(row-1,col-1) │
    ├──────────────┼──────────────┐
    │(row  ,col-1) │(row  ,col  ) │
    ├──────────────┼──────────────┘
    │(row+1,col-1) │
    └──────────────┘

        从上又可以推断出，
        第0列中只有 dp[0][0]可达
        第1列中只有 dp[1][0]和dp[1][1]可达
        第2列中只有 dp[2][0]到dp[2][2]可达
        依次类推
        第i列中只有 dp[i][0]到dp[i][i]可达
         */
    pub fn max_moves(grid: Vec<Vec<i32>>) -> i32 {
        let mut ret = 0;
        let mut dp = vec![vec![0; 2]; grid.len()];
        for col in 1..grid[0].len() {
            let pre_col = (col - 1) & 1;
            if dp[0][pre_col] >= 0 && grid[0][col] > grid[0][col - 1] {
                dp[0][col & 1] = dp[0][pre_col] + 1;
            } else {
                dp[0][col & 1] = -1;
            }

            for row in 1..grid.len() {
                dp[row][col & 1] = -1;
                if dp[row - 1][pre_col] >= 0 && grid[row][col] > grid[row - 1][col - 1] {
                    dp[row][col & 1] = dp[row][col & 1].max(dp[row - 1][pre_col] + 1);
                }

                if dp[row][pre_col] >= 0 && grid[row][col] > grid[row][col - 1] {
                    dp[row][col & 1] = dp[row][col & 1].max(dp[row][pre_col] + 1);
                }

                if (row + 1) <= grid.len() - 1
                    && dp[row + 1][pre_col] >= 0
                    && grid[row][col] > grid[row + 1][col - 1]
                {
                    dp[row][col & 1] = dp[row][col & 1].max(dp[row + 1][pre_col] + 1);
                }

                ret = ret.max(dp[row][col & 1]);
            }
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::max_moves(vec![
                vec![2, 4, 3, 5],
                vec![5, 4, 9, 3],
                vec![3, 4, 2, 11],
                vec![10, 9, 13, 15]
            ]),
            3
        );

        assert_eq!(
            Solution::max_moves(vec![vec![3, 2, 4], vec![2, 1, 9], vec![1, 1, 7]]),
            0
        );
    }
}
