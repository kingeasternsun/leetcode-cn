struct Solution;
impl Solution {
    // BFS: 4ms 2.07mb, 编程要点，注意 point_map 要是一个全局的去重map，执行过bfs的就会跳过
    pub fn max_area_of_island(grid: Vec<Vec<i32>>) -> i32 {
        let mut ret = 0;
        let mut point_map = std::collections::HashSet::new();
        for row_id in 0..grid.len() {
            for col_id in 0..grid[0].len() {
                if grid[row_id][col_id] == 1 && !point_map.contains(&(row_id as i32, col_id as i32))
                {
                    ret = ret.max(Self::bfs(
                        &grid,
                        (row_id as i32, col_id as i32),
                        &mut point_map,
                    ));
                }
            }
        }

        ret as i32
    }

    fn bfs(
        grid: &Vec<Vec<i32>>,
        start: (i32, i32),
        point_map: &mut std::collections::HashSet<(i32, i32)>,
    ) -> usize {
        let dirs = vec![(0, -1), (0, 1), (1, 0), (-1, 0)];
        let mut queue = std::collections::VecDeque::new();

        queue.push_back(start);
        point_map.insert(start);

        let mut ret = 0;
        while let Some(cur) = queue.pop_front() {
            ret += 1;
            for dir in dirs.iter() {
                let new_point = (cur.0 + dir.0, cur.1 + dir.1);
                if 0 <= new_point.0
                    && new_point.0 < grid.len() as i32
                    && 0 <= new_point.1
                    && new_point.1 < grid[0].len() as i32
                    && grid[new_point.0 as usize][new_point.1 as usize] == 1
                    && !point_map.contains(&new_point)
                {
                    queue.push_back(new_point);
                    point_map.insert(new_point);
                }
            }
        }

        ret
    }

    // BFS 0ms 2.26mb
    pub fn max_area_of_island1(grid: Vec<Vec<i32>>) -> i32 {
        let mut ret = 0;
        let mut point_map = vec![vec![false; grid[0].len()]; grid.len()];
        for row_id in 0..grid.len() {
            for col_id in 0..grid[0].len() {
                if !point_map[row_id][col_id] && grid[row_id][col_id] == 1 {
                    ret = ret.max(Self::bfs1(
                        &grid,
                        (row_id as i32, col_id as i32),
                        &mut point_map,
                    ));
                }
            }
        }

        ret as i32
    }

    fn bfs1(grid: &Vec<Vec<i32>>, start: (i32, i32), point_map: &mut Vec<Vec<bool>>) -> usize {
        let dirs = vec![(0, -1), (0, 1), (1, 0), (-1, 0)];
        let mut queue = std::collections::VecDeque::new();

        queue.push_back(start);
        point_map[start.0 as usize][start.1 as usize] = true;

        let mut ret = 0;
        while let Some(cur) = queue.pop_front() {
            ret += 1;
            for dir in dirs.iter() {
                let new_point = (cur.0 + dir.0, cur.1 + dir.1);
                if 0 <= new_point.0
                    && new_point.0 < grid.len() as i32
                    && 0 <= new_point.1
                    && new_point.1 < grid[0].len() as i32
                    && grid[new_point.0 as usize][new_point.1 as usize] == 1
                    && !point_map[new_point.0 as usize][new_point.1 as usize]
                {
                    queue.push_back(new_point);
                    point_map[new_point.0 as usize][new_point.1 as usize] = true;
                }
            }
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let grid = vec![
            vec![0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0],
        ];

        assert_eq!(Solution::max_area_of_island(grid), 6);
    }

    #[test]
    fn it_works1() {
        let grid = vec![
            vec![0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
            vec![0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
            vec![0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0],
        ];

        assert_eq!(Solution::max_area_of_island1(grid), 6);
    }
}
