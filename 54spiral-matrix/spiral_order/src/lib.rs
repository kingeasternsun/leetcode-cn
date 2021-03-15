/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-15 09:56:21
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-15 11:24:59
 * @FilePath: /54spiral-matrix/spiral_order/src/lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {

        let mut res = Vec::new();
        if matrix.len()==0{
            return res;
        }

        let rows = matrix.len();
        let cols = matrix[0].len();
        let iter_num = std::cmp::min(rows, cols)/2; //迭代次数  如果是奇数的话 最后一次就是一个行或列 单独考虑

        for iter in 0..iter_num{

            //上面遍历
            for col_id in iter..(cols-iter-1){
                res.push(matrix[iter][col_id]);
            }
            //右边遍历
            for row_id in iter..(rows-iter-1){
                res.push(matrix[row_id][cols-iter-1]);
            }
            //下方遍历
            for col_id in (iter+1..cols-iter).rev(){
                res.push(matrix[rows-iter-1][col_id]);
            }
            //左边遍历
            for row_id in (iter+1..rows-iter).rev(){
                res.push(matrix[row_id][iter]);
            }
        }

        //考虑剩下单行的情况 
        if rows <=cols && (rows&1>0){
            for col_id in iter_num..cols-iter_num{
                res.push(matrix[iter_num][col_id]);
            }
        }

        //考虑剩下单列的情况
        if cols < rows && (cols&1>0){
            for row_id in iter_num..rows-iter_num{
                res.push(matrix[row_id][iter_num]);
            }
        }

        res

    }


}
#[cfg(test)]
mod tests {
    use crate::Solution;
    #[test]
    fn it_works() {
        assert_eq!(Solution::spiral_order(vec![vec![1,2,3],vec![4,5,6],vec![7,8,9]]),vec![1,2,3,6,9,8,7,4,5]);
        assert_eq!(Solution::spiral_order(vec![vec![1,2,3,4],vec![5,6,7,8],vec![9,10,11,12]]),vec![1,2,3,4,8,12,11,10,9,5,6,7]);
        assert_eq!(Solution::spiral_order(vec![vec![1,2,3],vec![4,5,6],vec![7,8,9],vec![10,11,12]]),vec![1,2,3,6,9,12,11,10,7,4,5,8]);
      
    }
}

