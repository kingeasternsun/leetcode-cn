/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-05-02 20:40:37
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-05-02 20:59:24
 * @FilePath: /least-bricks/src/lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn least_bricks(wall: Vec<Vec<i32>>) -> i32 {
        let mut len_map = std::collections::HashMap::new();
        for row in &wall{
            let mut tmp_len = 0;
            for i in 0..row.len()-1{
                tmp_len += row[i];
                let cnt =  len_map.entry(tmp_len).or_insert(0);
                *cnt+=1;
            }
        }

        let mut max_cnt = 0;
        for val in len_map.values(){
            if *val > max_cnt{
                max_cnt = *val;
            }
        }

        wall.len() as i32 - max_cnt

    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::least_bricks(vec![vec![1,2,2,1],vec![3,1,2],
            vec![1,3,2],vec![2,4],vec![3,1,2],vec![1,3,1,1]]), 2);
        assert_eq!(Solution::least_bricks(vec![vec![1],vec![1],vec![1]]),3);
        
    }
}
