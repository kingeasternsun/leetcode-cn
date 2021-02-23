/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 16:56:39
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 17:10:04
 * @FilePath: \240\search_matrix\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let rows = matrix.len();
        if rows == 0 {
            return false;
        }

        let cols = matrix[0].len();
        if cols == 0 {
            return false;
        }

        let mut x = 0;
        let mut y = cols - 1;
        while x < rows {
            if matrix[x][y] == target {
                return true;
            }
            //比目标要大 ，向左移动
            if matrix[x][y] > target {
                if y ==0 {
                    return false
                }
                y -= 1;
            } else {
                //比目标值小，向下移动
                x += 1;
            }
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(Solution::search_matrix(vec![vec![1,4,7,11,15],vec![2,5,8,12,19],vec![3,6,9,16,22],vec![10,13,14,17,24],vec![18,21,23,26,30]],5), true);
        assert_eq!(Solution::search_matrix(vec![vec![1,4,7,11,15],vec![2,5,8,12,19],vec![3,6,9,16,22],vec![10,13,14,17,24],vec![18,21,23,26,30]],-5), false);
    }
}
