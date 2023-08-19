pub struct Solution;
impl Solution {
    // BFS 92ms  2.36mb
    pub fn largest_area(grid: Vec<String>) -> i32 {
        let mut ret = 0;
        let mut point_map = vec![vec![false; grid[0].len()]; grid.len()];
        for row_id in 0..grid.len() {
            for col_id in 0..grid[0].as_bytes().len() {
                if grid[row_id].as_bytes()[col_id] != b'0' && !point_map[row_id][col_id] {
                    let tmp = Self::bfs(&grid, &mut point_map, (row_id as i32, col_id as i32));
                    if !tmp.1 {
                        ret = ret.max(tmp.0);
                    }
                }
            }
        }
        ret
    }
    fn bfs(
        grid: &Vec<String>,
        point_map: &mut Vec<Vec<bool>>,
        start_point: (i32, i32),
    ) -> (i32, bool) {
        let target = grid[start_point.0 as usize].as_bytes()[start_point.1 as usize];
        let dirs = vec![(-1, 0), (1, 0), (0, 1), (0, -1)];
        let mut queue = std::collections::VecDeque::new();
        queue.push_back(start_point);
        point_map[start_point.0 as usize][start_point.1 as usize] = true;

        let mut ret = 0;
        let mut boudary = false;
        while let Some(cur) = queue.pop_front() {
            ret += 1;
            for dir in dirs.iter() {
                let new_point = (cur.0 + dir.0, cur.1 + dir.1);
                // 判断是否挨着走廊
                if new_point.0 < 0
                    || new_point.0 >= grid.len() as i32
                    || new_point.1 < 0
                    || new_point.1 >= grid[0].len() as i32
                    || grid[new_point.0 as usize].as_bytes()[new_point.1 as usize] == b'0'
                {
                    boudary = true;
                    // 判断是否连通且未访问过
                } else if !point_map[new_point.0 as usize][new_point.1 as usize]
                    && grid[new_point.0 as usize].as_bytes()[new_point.1 as usize] == target
                {
                    queue.push_back(new_point);
                    point_map[new_point.0 as usize][new_point.1 as usize] = true;
                }
            }
        }
        (ret, boudary)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::largest_area(vec![
                "110".to_string(),
                "231".to_string(),
                "221".to_string()
            ]),
            1
        );
        assert_eq!(
            Solution::largest_area(vec![
                "11111100000".to_string(),
                "21243101111".to_string(),
                "21224101221".to_string(),
                "11111101111".to_string()
            ]),
            3
        );
    }
}
