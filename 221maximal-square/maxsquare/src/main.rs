fn main() {
    println!("Hello, world!");

    // [["1","1","1","1","1","1","1","1"],["1","1","1","1","1","1","1","0"],["1","1","1","1","1","1","1","0"],["1","1","1","1","1","0","0","0"],["0","1","1","1","1","0","0","0"]]

   let s =  Solution::maximal_square(vec![
        vec!['1','1','1','1','1','1','1','1'],
        vec!['1','1','1','1','1','1','1','0'],
        vec!['1','1','1','1','1','1','1','0'],
        vec!['1','1','1','1','1','0','0','0'],
        vec!['0','1','1','1','1','0','0','0']
    ]);

    println!("{}",s);
}

pub struct Solution {}

impl Solution {

    //相比上面的方法，这里由于 dp 使用复用行，所以为0的记得要重置为0  4ms 4.8MB
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {

        if matrix.is_empty() || matrix[0].is_empty(){
            return 0;
        }

        let (rows,cols) = (matrix.len(),matrix[0].len());

        let mut dp  = vec![vec![0;cols];2];
        let mut max_len = 0;

        for row in 0..rows{
           let  pre_id = (row+1)&1;
            for col in 0..cols{

                if matrix[row][col]=='0'{
                    dp[row&1][col] = 0; //这里一定要清除变为0 
                    continue;
                }

                if row == 0 || col == 0 {
                    dp[row&1][col]=1;

                    if max_len ==0 {
                        max_len = 1;
                    }

                    continue
                }

                if dp[row&1][col-1] == dp[pre_id][col]{

                    //对角线
                    if dp[pre_id][col-1] >= dp[row&1][col-1]{
                        dp[row&1][col]= dp[row&1][col-1] +1
                    }else{
                        dp[row&1][col]= dp[row&1][col-1]
                    }

                }else{
                    dp[row&1][col] = i32::min(dp[row&1][col-1],dp[pre_id][col])+1
                }

                //更新 max_len
                if dp[row&1][col]>max_len{
                    max_len = dp[row&1][col]
                }


            }
        }



        max_len*max_len
    }
}




#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_221() {
        // assert_eq!(
        //     Solution::maximal_square(vec![
        //         vec!['1', '0', '1', '0', '0'],
        //         vec!['1', '0', '1', '1', '1'],
        //         vec!['1', '1', '1', '1', '1'],
        //         vec!['1', '0', '0', '1', '0'],
        //     ]),
        //     4
        // );
        // // [["0","0","0","1"],["1","1","0","1"],["1","1","1","1"],["0","1","1","1"],["0","1","1","1"]]

        // assert_eq!(
        //     Solution::maximal_square(vec![
        //         vec!['0', '0', '0', '0'],
        //         vec!['1', '1', '0', '1'],
        //         vec!['1', '1', '1', '1'],
        //         vec!['0', '1', '1', '1'],
        //         vec!['0', '1', '1', '1']
        //     ]),
        //     9
        // );

        assert_eq!(
            Solution::maximal_square(vec![
                vec!['1','1','1','1','1','1','1','1'],
                vec!['1','1','1','1','1','1','1','0'],
                vec!['1','1','1','1','1','1','1','0'],
                vec!['1','1','1','1','1','0','0','0'],
                vec!['0','1','1','1','1','0','0','0']
            ]),
            16
        )


    }
}