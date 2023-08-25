struct Solution;
impl Solution {
    pub fn count_servers(grid: Vec<Vec<i32>>) -> i32 {
        let mut same_row = vec![vec![]; grid.len()];
        let mut same_col = vec![vec![]; grid[0].len()];

        for row_id in 0..grid.len() {
            for col_id in 0..grid[0].len() {
                if grid[row_id][col_id] == 1 {
                    same_row[row_id].push(col_id);
                    same_col[col_id].push(row_id);
                }
            }
        }

        // println!("{:?}", same_row);
        // println!("{:?}", same_col);

        let mut map = std::collections::HashSet::new();

        for row_id in 0..grid.len() {
            if same_row[row_id].len() > 1 {
                for &col_id in same_row[row_id].iter() {
                    map.insert((row_id, col_id));
                }
            }
        }

        println!("{:?}", map);

        for col_id in 0..grid[0].len() {
            if same_col[col_id].len() > 1 {
                for &row_id in same_col[col_id].iter() {
                    map.insert((row_id, col_id));
                }
            }
        }

        map.len() as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::count_servers(vec![vec![1, 0], vec![0, 1]]), 0);
        assert_eq!(Solution::count_servers(vec![vec![1, 0], vec![1, 1]]), 3);
        assert_eq!(
            Solution::count_servers(vec![
                vec![1, 1, 0, 0],
                vec![0, 0, 1, 0],
                vec![0, 0, 1, 0],
                vec![0, 0, 0, 1]
            ]),
            4
        );
    }
}
