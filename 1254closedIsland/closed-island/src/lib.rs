
struct Solution;
impl Solution {
    pub fn closed_island(grid: Vec<Vec<i32>>) -> i32 {
        if grid.len() == 0 {
            return 0;
        }

        let mut num = 0;
        let mut grid = grid;
        for i in 0..grid.len() {
            for j in 0..grid[i].len() {
                if grid[i][j] > 0 {
                    continue;
                }

                num = num + Self::bfs((i as i32, j as i32), &mut grid)
            }
        }
        num
    }

    pub fn bfs(p: (i32, i32), grid: &mut Vec<Vec<i32>>) -> i32 {
        let mut boundry = false;
        let mut points = vec![p];
        while let Some(p) = points.pop() {
            if p.0 == 0
                || p.1 == 0
                || p.0 == grid.len() as i32 - 1
                || p.1 == grid[0].len() as i32 - 1
            {
                boundry = true;
            }
            grid[p.0 as usize][p.1 as usize] = 2;

            let dirs = vec![0, 1, 0, -1, 0];
            for i in 0..4 {
                let new_p = (p.0 + dirs[i], p.1 + dirs[i + 1]);
                if new_p.0 < 0
                    || new_p.1 < 0
                    || new_p.0 > grid.len() as i32 - 1
                    || new_p.1 > grid[0].len() as i32 - 1
                {
                    continue;
                }

                if grid[new_p.0 as usize][new_p.1 as usize] > 0 {
                    continue;
                }

                points.push(new_p);
            }
        }

        if boundry {
            return 0;
        }
        return 1;
    }
    pub fn bfs2(p: (i32, i32), grid: &mut Vec<Vec<i32>>) -> i32 {
        let mut boundry = false;
        let mut points = vec![p];
        loop {
            let mut new_points: Vec<(i32, i32)> = vec![];
            for p in points {
                if p.0 == 0
                    || p.1 == 0
                    || p.0 == grid.len() as i32 - 1
                    || p.1 == grid[0].len() as i32 - 1
                {
                    boundry = true;
                }
                grid[p.0 as usize][p.1 as usize] = 2;

                let dirs = vec![0, 1, 0, -1, 0];
                for i in 0..4 {
                    let new_p = (p.0 + dirs[i], p.1 + dirs[i + 1]);
                    if new_p.0 < 0
                        || new_p.1 < 0
                        || new_p.0 > grid.len() as i32 - 1
                        || new_p.1 > grid[0].len() as i32 - 1
                    {
                        continue;
                    }

                    if grid[new_p.0 as usize][new_p.1 as usize] > 0 {
                        continue;
                    }

                    new_points.push(new_p);
                }
            }

            if new_points.len() == 0 {
                break;
            }
            points = new_points;
        }
        if boundry {
            return 0;
        }
        return 1;
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        let grid1 = vec![vec![1,1,1,1,1,1,1,0],vec![1,0,0,0,0,1,1,0],vec![1,0,1,0,1,1,1,0],vec![1,0,0,0,0,1,0,1],vec![1,1,1,1,1,1,1,0]];
        assert_eq!(Solution::closed_island(grid1),2);
        let grid2 = vec![vec![0,0,1,0,0],vec![0,1,0,1,0],vec![0,1,1,1,0]];
        assert_eq!(Solution::closed_island(grid2),2);
    }
}
