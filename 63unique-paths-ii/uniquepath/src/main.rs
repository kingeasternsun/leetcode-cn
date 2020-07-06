fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {

       let rows = obstacle_grid.len();
       if rows == 0 {
           return 0
       }

       let cols = obstacle_grid[0].len();
       if cols == 0 {
           return 0
       }

       if obstacle_grid[0][0] == 1 || obstacle_grid[rows as usize -1][cols as usize -1]==1{
           return 0
       }
       

       let mut dp = vec![vec![0;cols];2];
       dp[1][0] = 1;
       for  row in 0..rows{

            let pre_id  = (row+1)&1 as usize;
        for  col in 0..cols{
            if obstacle_grid[row][col] == 1{
                dp[row&1][col]=0
            }else if col == 0{
                dp[row&1][col] = dp[pre_id][col] 
            }else{
                dp[row&1][col] = dp[pre_id][col] +  dp[row&1][col-1]
            }

        }
       }
     

       dp[(rows-1)&1][cols - 1]

    }
}