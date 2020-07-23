fn main() {
    println!("Hello, world!");
    Solution::min_path_sum(vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1],]);
}

pub struct Solution{}

impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let rows = grid.len();
        if rows ==0{
            return 0
        }

        let cols = grid[0].len();
        if cols ==0{
            return 0
        }

        let mut dp: Vec<Vec<i32>> = vec![vec![i32::MAX;cols];2];
        dp[1][0]=0;
        // for j in 0..cols{

        //     if j ==0{
        //         dp[0][j] =  grid[0][j]
        //     }else{
        //         dp[0][j] =  grid[0][j] + dp[0][j-1]
        //     }

            
        // }

        for i in 0..rows{
            let preID =(i+1)&1;
            let curID =(i)&1;

            for j in 0..cols{
                dp[curID][j] = dp[preID][j];
                if j >0 && dp[curID][j-1] < dp[curID][j]{
                    dp[curID][j] = dp[curID][j-1]
                }

                dp[curID][j]+=grid[i][j]
            }


        }

        dp[(rows-1)&1][cols-1]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_64() {
        assert_eq!(Solution::min_path_sum(vec![vec![2]]), 2);
        assert_eq!(
            Solution::min_path_sum(vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1],]),
            7
        );
        assert_eq!(Solution::min_path_sum(vec![vec![1, 3, 1],]), 5);
    }
}