struct Solution;
impl Solution {
    pub fn ways(pizza: Vec<String>, k: i32) -> i32 {
        let k = k as usize;
        let col_num = pizza[0].as_bytes().len();
        let row_num = pizza.len();
        let mut cache = vec![vec![vec![-1; k + 1]; col_num]; pizza.len()];
        let mut sum = vec![vec![0; col_num]; pizza.len()];

        // 由于每次留下的是下面的，或 右边的披萨，所以求到右下角的二维前缀和
        for row_id in (0..row_num).rev() {
            for col_id in (0..col_num).rev() {
                if row_id == row_num - 1 && col_id == col_num - 1 {
                    sum[row_id][col_id] = 0_usize;
                } else if row_id == row_num - 1 {
                    // 最后一行
                    sum[row_id][col_id] = sum[row_id][col_id + 1];
                } else if col_id == col_num - 1 {
                    // 最后一列
                    sum[row_id][col_id] = sum[row_id + 1][col_id];
                } else {
                    sum[row_id][col_id] = sum[row_id + 1][col_id] + sum[row_id][col_id + 1]
                        - sum[row_id + 1][col_id + 1];
                }

                if pizza[row_id].as_bytes()[col_id] == b'A' {
                    sum[row_id][col_id] += 1;
                }
            }
        }

        let ret = Self::dp(0, 0, k, col_num, row_num, &sum, &mut cache);
        (ret % 1000000007) as i32
    }

    fn dp(
        row: usize,
        col: usize,
        k: usize,
        col_num: usize,
        row_num: usize,
        sum: &Vec<Vec<usize>>,
        cache: &mut Vec<Vec<Vec<i32>>>,
    ) -> i32 {
        if cache[row][col][k] >= 0 {
            return cache[row][col][k];
        }

        // 剩余苹果数目还不到 k 个,不够 k个人分
        if sum[row][col] < k {
            cache[row][col][k] = 0;
            return 0;
        }

        if k == 1 {
            cache[row][col][k] = 1;
            return 1;
        }

        let mut total = 0;
        // 尝试横着切
        for row_id in row + 1..row_num {
            if sum[row][col] - sum[row_id][col] > 0 {
                let ret = Self::dp(row_id, col, k - 1, col_num, row_num, sum, cache);
                if ret == 0 {
                    // 已经不够分了，再往下切就更不够分了
                    break;
                }
                total += ret;
            }
        }

        // 尝试竖着切
        for col_id in col + 1..col_num {
            if sum[row][col] - sum[row][col + 1] > 0 {
                let ret = Self::dp(row, col_id, k - 1, col_num, row_num, sum, cache);
                if ret == 0 {
                    // 已经不够分了，再往右切就更不够分了
                    break;
                }
                total += ret;
            }
        }

        cache[row][col][k] = total;
        return total;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::ways(
                vec!["A..".to_string(), "AAA".to_string(), "...".to_string()],
                3
            ),
            3
        );
        assert_eq!(
            Solution::ways(
                vec!["A..".to_string(), "AA.".to_string(), "...".to_string()],
                3
            ),
            1
        );
        assert_eq!(
            Solution::ways(
                vec!["A..".to_string(), "A..".to_string(), "...".to_string()],
                1
            ),
            1
        );

        assert_eq!(
            Solution::ways(
                vec![
                    ".A..A".to_string(),
                    "A.A..".to_string(),
                    "A.AA.".to_string(),
                    "AAAA.".to_string(),
                    "A.AA.".to_string()
                ],
                5
            ),
            153
        );
    }
}
