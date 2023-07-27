struct Solution;
impl Solution {
    pub fn delete_greatest_value(mut grid: Vec<Vec<i32>>) -> i32 {
        let mut sum = 0;
        let col = grid[0].len();
        grid.iter_mut().for_each(|x| x.sort_unstable());
        for i in 0..col {
            sum += grid.iter().map(|x| x[i]).max().unwrap_or(0);
        }

        sum
    }

    // the beautify of iterator
    pub fn delete_greatest_value2(mut grid: Vec<Vec<i32>>) -> i32 {
        grid.iter_mut().for_each(|x| x.sort_unstable());
        (0..grid[0].len()).map(|i| grid.iter().map(|x| x[i]).max().unwrap_or(0)).sum()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut grid1 = vec![vec![1, 2, 4], vec![3, 3, 1]];
        let mut grid2 = vec![vec![1, 2, 4]];
        let mut grid3 = vec![vec![10]];
        let mut grid4 = vec![vec![1], vec![3]];
        assert_eq!(Solution::delete_greatest_value(grid1), 8);
        assert_eq!(Solution::delete_greatest_value(grid2), 7);
        assert_eq!(Solution::delete_greatest_value(grid3), 10);
        assert_eq!(Solution::delete_greatest_value(grid4), 3);
    }

    #[test]
    fn it_works2() {
        let mut grid1 = vec![vec![1, 2, 4], vec![3, 3, 1]];
        let mut grid2 = vec![vec![1, 2, 4]];
        let mut grid3 = vec![vec![10]];
        let mut grid4 = vec![vec![1], vec![3]];
        assert_eq!(Solution::delete_greatest_value2(grid1), 8);
        assert_eq!(Solution::delete_greatest_value2(grid2), 7);
        assert_eq!(Solution::delete_greatest_value2(grid3), 10);
        assert_eq!(Solution::delete_greatest_value2(grid4), 3);
    }
}
