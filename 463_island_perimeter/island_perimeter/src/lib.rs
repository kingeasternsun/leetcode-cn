struct Solution;
impl Solution {
    // BFS 岛屿边界
    pub fn island_perimeter(grid: Vec<Vec<i32>>) -> i32 {
        let root = Self::get_root(&grid);
        if root.is_none() {
            return 0;
        }

        let dirs = vec![(0, -1), (0, 1), (-1, 0), (1, 0)];
        let mut queue = std::collections::VecDeque::new();
        let mut point_set = std::collections::HashSet::new();
        let mut boundry = 0;
        queue.push_back(root.unwrap());
        point_set.insert(root.unwrap());
        while let Some(cur) = queue.pop_front() {
            for dir in dirs.iter() {
                let new_point = (cur.0 + dir.0, cur.1 + dir.1);
                // 方向 dir 是边界
                if new_point.0 < 0
                    || new_point.0 >= grid.len() as i32
                    || new_point.1 < 0
                    || new_point.1 >= grid[0].len() as i32
                    || grid[new_point.0 as usize][new_point.1 as usize] == 0
                {
                    boundry += 1;
                // 方向 dir 的相邻网格还没有访问过
                } else if !point_set.contains(&new_point) {
                    queue.push_back(new_point);
                    point_set.insert(new_point);
                }
            }
        }

        boundry
    }

    fn get_root(grid: &Vec<Vec<i32>>) -> Option<(i32, i32)> {
        for row_id in 0..grid.len() {
            for col_id in 0..grid[0].len() {
                if grid[row_id][col_id] == 1 {
                    return Some((row_id as i32, col_id as i32));
                }
            }
        }
        None
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::island_perimeter(vec![
                vec![0, 1, 0, 0],
                vec![1, 1, 1, 0],
                vec![0, 1, 0, 0],
                vec![1, 1, 0, 0]
            ]),
            16
        );

        // assert_eq!(Solution::island_perimeter(vec![vec![1]]), 4);

        // assert_eq!(Solution::island_perimeter(vec![vec![1, 0]]), 4);
    }
}
