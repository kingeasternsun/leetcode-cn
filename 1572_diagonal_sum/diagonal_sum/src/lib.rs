struct Solution;

impl Solution {
    pub fn diagonal_sum(mat: Vec<Vec<i32>>) -> i32 {
        let mut sum = 0;
        for i in 0..mat.len() {
            sum += mat[i][i];
            let j = mat.len() - 1 - i;
            if i != j {
                sum += mat[i][j]
            }
        }

        sum
    }

    // the beauty of iterator
    pub fn diagonal_sum2(mat: Vec<Vec<i32>>) -> i32 {
        mat.iter().enumerate().fold(0, |mut acc, (i, row)| {
            acc += row[i];
            if mat.len() - i - 1 != i {
                acc += row[mat.len() - i - 1];
            }
            acc
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::diagonal_sum(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]]),
            25
        );
        assert_eq!(
            Solution::diagonal_sum(vec![
                vec![1, 1, 1, 1],
                vec![1, 1, 1, 1],
                vec![1, 1, 1, 1],
                vec![1, 1, 1, 1]
            ]),
            8
        );
        assert_eq!(Solution::diagonal_sum(vec![vec![5]]), 5);
    }

    #[test]
    fn it_works2() {
        assert_eq!(
            Solution::diagonal_sum2(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]]),
            25
        );
        assert_eq!(
            Solution::diagonal_sum2(vec![
                vec![1, 1, 1, 1],
                vec![1, 1, 1, 1],
                vec![1, 1, 1, 1],
                vec![1, 1, 1, 1]
            ]),
            8
        );
        assert_eq!(Solution::diagonal_sum2(vec![vec![5]]), 5);
    }
}
