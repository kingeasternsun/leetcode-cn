struct Solution;
impl Solution {
    pub fn ways(pizza: Vec<String>, k: i32) -> i32 {
        let mut ret = 0_u64;
        let col_num = pizza[0].as_bytes().len();
        let row_num = pizza.len();
        let mut cache = vec![vec![-1; col_num]; pizza.len()];
        let mut sum = vec![vec![0; col_num]; pizza.len()];

        // 初始化最后一行
        for col in (0..col_num).rev() {
            sum[row_num - 1][col] = sum[row_num - 1].get(col + 1).unwrap_or(&0).clone();
            if pizza[row_num - 1].as_bytes()[col] == b'A' {
                sum[row_num - 1][col] += 1;
            }
        }

        // 初始化最后一列
        for row in (0..row_num - 1).rev() {
            sum[row][col_num - 1] = sum[row + 1][col_num - 1];
            if pizza[row].as_bytes()[col_num - 1] == b'A' {
                sum[row][col_num - 1] += 1;
            }
        }

        for row_id in (0..row_num - 1).rev() {
            for col_id in (0..col_num - 1).rev() {
                sum[row_id][col_id] =
                    sum[row_id - 1][col_id] + sum[row_id][col_id - 1] - sum[row_id - 1][col_id - 1];
                if pizza[row_id].as_bytes()[col_id] == b'A' {
                    sum[row_id][col_id] += 1;
                }
            }
        }

        (ret % 1000000007) as i32
    }

    fn dp(row:usize,col:usize, k: i32, sum:& Vec<Vec<i32>>, cache:& mut Vec<Vec<i32>>)->i32{
        if cache[row][col] >=0{
            return cache[row][col];
        }

        if k == 1{
            cache[row][col] = 1;
            return 1;
        }

        

        0
    }
}
pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
