struct Solution;
impl Solution {
    // BFS 12ms 2.3MB
    pub fn color_border(grid: Vec<Vec<i32>>, row: i32, col: i32, color: i32) -> Vec<Vec<i32>> {
        let dirs = vec![(-1, 0), (1, 0), (0, 1), (0, -1)];
        let mut grid = grid;
        // let (row, col): (usize, usize) = (row as usize, col as usize);
        let (row_len, col_len) = (grid.len() as i32, grid[0].len() as i32);
        let mut queue = std::collections::VecDeque::new();
        let mut point_set = std::collections::HashSet::new();
        // 不能在遍历的时候就涂色，不然会影响后续的边界判断，所以要先保存边界点
        let mut color_queue = vec![];
        queue.push_back((row, col));
        while let Some((cur_row, cur_col)) = queue.pop_front() {
            let cur_color = grid[cur_row as usize][cur_col as usize];
            let mut is_boudry = false;
            if cur_row == 0
                || cur_row == row_len as i32 - 1
                || cur_col == 0
                || cur_col == col_len as i32 - 1
            {
                is_boudry = true;
            }

            for dir in dirs.iter() {
                let new_point = (cur_row + dir.0, cur_col + dir.1);
                if 0 <= new_point.0
                    && new_point.0 < row_len as i32
                    && 0 <= new_point.1
                    && new_point.1 < col_len as i32
                {
                    // 周围颜色不一样，当前也是边界
                    if grid[new_point.0 as usize][new_point.1 as usize] != cur_color {
                        is_boudry = true;
                        // 周围颜色一样，且没有访问过
                    } else if !point_set.contains(&new_point) {
                        queue.push_back(new_point);
                        point_set.insert(new_point);
                    }
                }
            }

            if is_boudry {
                color_queue.push((cur_row, cur_col));
            }
        }

        for point in color_queue {
            grid[point.0 as usize][point.1 as usize] = color;
        }

        grid
    }

    // BFS 4ms 2.2MB
    pub fn color_border1(grid: Vec<Vec<i32>>, row: i32, col: i32, color: i32) -> Vec<Vec<i32>> {
        let mut grid = grid;
        if grid.len() == 0 {
            return grid;
        }
        let grid_size = (grid.len(), grid[0].len());
        let direct: Vec<i32> = vec![0, 1, 0, -1, 0];
        let mut borders = vec![]; // 记录边界
        let mut visited = vec![vec![false; grid_size.1]; grid_size.0]; // 记录哪些已经访问过
        let mut queue = std::collections::VecDeque::new(); // 记录将要访问的节点

        queue.push_back((row as usize, col as usize));
        visited[row as usize][col as usize] = true;

        while let Some((cur_x, cur_y)) = queue.pop_front() {
            let mut is_border = false;
            for i in 0..4 {
                let n_x = cur_x as i32 + direct[i];
                let n_y = cur_y as i32 + direct[i + 1];
                // 出界了,当前节点就是边界
                if n_x < 0
                    || n_x >= grid_size.0 as i32
                    || n_y < 0
                    || n_y >= grid_size.1 as i32
                    || grid[n_x as usize][n_y as usize] != grid[row as usize][col as usize]
                {
                    is_border = true;
                } else {
                    // 相连 ，并且没有访问过
                    if !visited[n_x as usize][n_y as usize] {
                        queue.push_back((n_x as usize, n_y as usize));
                        visited[n_x as usize][n_y as usize] = true;
                    }
                }
            }

            if is_border {
                borders.push((cur_x, cur_y));
            }
        }

        for (x, y) in borders {
            grid[x][y] = color;
        }

        grid
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            /*
            1 1
            1 2
             */
            Solution::color_border(vec![vec![1, 1], vec![1, 2]], 0, 0, 3),
            vec![vec![3, 3], vec![3, 2]]
        );
        assert_eq!(
            Solution::color_border(vec![vec![1, 2, 2], vec![2, 3, 2]], 0, 1, 3),
            vec![vec![1, 3, 3], vec![2, 3, 3]]
        );
        assert_eq!(
            Solution::color_border(vec![vec![1, 1, 1], vec![1, 1, 1], vec![1, 1, 1]], 1, 1, 2),
            vec![vec![2, 2, 2], vec![2, 1, 2], vec![2, 2, 2]]
        )
    }

    #[test]
    fn it_works1() {
        assert_eq!(
            /*
            1 1
            1 2
             */
            Solution::color_border1(vec![vec![1, 1], vec![1, 2]], 0, 0, 3),
            vec![vec![3, 3], vec![3, 2]]
        );
        assert_eq!(
            Solution::color_border1(vec![vec![1, 2, 2], vec![2, 3, 2]], 0, 1, 3),
            vec![vec![1, 3, 3], vec![2, 3, 3]]
        );
        assert_eq!(
            Solution::color_border1(vec![vec![1, 1, 1], vec![1, 1, 1], vec![1, 1, 1]], 1, 1, 2),
            vec![vec![2, 2, 2], vec![2, 1, 2], vec![2, 2, 2]]
        )
    }
}
