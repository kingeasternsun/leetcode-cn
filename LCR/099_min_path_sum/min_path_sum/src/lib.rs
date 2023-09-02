struct Solution;
impl Solution {
    // 0ms 2.24mb DP
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {

        let mut dp = vec![vec![0;grid[0].len()];2];
        for row_id in 0..grid.len(){
            let dp_pre_row_id = (row_id+1)&1;
            let dp_row_id = (row_id)&1;
            for col_id in 0..grid[row_id].len(){
                

              dp[dp_row_id][col_id]= match (row_id, col_id){
                (0,0) => grid[0][0],
                (0,_) => dp[dp_row_id][col_id-1]+grid[row_id][col_id],
                (_,0) => dp[dp_pre_row_id][col_id] + grid[row_id][col_id],
                _=>dp[dp_row_id][col_id-1].min(dp[dp_pre_row_id][col_id]) + grid[row_id][col_id],
               }
            }
        }

        dp[(grid.len()-1)&1][grid[0].len()-1]
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let grid = vec![vec![1,3,1],vec![1,5,1],vec![4,2,1]];
        assert_eq!(Solution::min_path_sum(grid), 7);

        let grid = vec![vec![1,2,3],vec![4,5,6]];
         assert_eq!(Solution::min_path_sum(grid), 12);
    }
}
