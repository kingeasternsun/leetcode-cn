/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 16:27:37
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 16:35:46
 * @FilePath: \74\search_matrix\src\lib.rs
 */

pub struct Solution;

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let row = matrix.len();
        if row == 0 {
            return false;
        }

        let col = matrix[0].len();
        if col == 0 {
            return false;
        }

        let mut beg = 0;
        let mut end = row * col - 1;
        while beg <= end {
            let mid = (end - beg) / 2 + beg;

            if matrix[mid / col][mid % col] == target {
                return true;
            }
            if matrix[mid / col][mid % col] > target {
                if mid ==0 { //rust 二分查找的时候 一定要注意这个
                    return false
                }
                end = mid - 1
            } else {
                beg = mid + 1
            }
        }
        return false;
    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::search_matrix(vec![vec![1,3,5,7],vec![10,11,16,20],vec![23,30,34,60]], 3),true);
        assert_eq!(Solution::search_matrix(vec![vec![1,3,5,7],vec![10,11,16,20],vec![23,30,34,60]], 13),false);
        assert_eq!(Solution::search_matrix(vec![vec![1]], 13),false);
        assert_eq!(Solution::search_matrix(vec![vec![1]], -1),false);
    }
}
