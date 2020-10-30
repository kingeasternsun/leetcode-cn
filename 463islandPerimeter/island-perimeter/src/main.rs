fn main() {
    // println!("Hello, world!");
    let grid = vec![vec![0,1,0,0],
    vec![1,1,1,0],
    vec![0,1,0,0],
    vec![1,1,0,0]];
    
    assert_eq!(16,Solution::island_perimeter(grid));
}

pub struct Solution;
impl Solution {
    pub fn island_perimeter(grid: Vec<Vec<i32>>) -> i32 {
        let rows = grid.len();
        if grid.get(0).is_none() {
            return 0;
        }

        let cols = grid.get(0).unwrap().len();

        let dirs: Vec<i32> = vec![0, 1, 0, -1, 0];

        let mut sum = 0;

        for i in 0..rows {
            for j in 0..cols {
                if grid[i][j] == 0 {
                    continue;
                }

                //对每一个方向，如果相邻的是水或边界，就加1
                dirs.windows(2).for_each(|dir| {
                    let new_x = (i as i32 + dir[0]) as usize;
                    let new_y = (j as i32 + dir[1]) as usize;

                    if let Some(row) = grid.get(new_x) {
                        if let Some(next) = row.get(new_y) {
                            if *next == 0 {
                                sum += 1;
                            }
                        } else {
                            sum += 1;
                        }
                    } else {
                        sum += 1;
                    }
                });
            }
        }

        sum
    }
}
