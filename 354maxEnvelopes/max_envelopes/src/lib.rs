/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-04 09:44:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-04 10:09:29
 * @FilePath: \max_envelopes\src\lib.rs
 */
pub struct Solution;

use std::cmp::Ordering;
impl Solution {
    pub fn max_envelopes(envelopes: Vec<Vec<i32>>) -> i32 {
        if envelopes.len() == 0 {
            return 0;
        }

        let mut envelopes = envelopes;

        //先排序，按照 宽w进行排序。如果宽w相同，就不能互相套，所以要把高h大的放前面
        envelopes.sort_unstable_by(|a, b| {
            if a[0].cmp(&b[0]) == Ordering::Equal {
                b[1].cmp(&a[1])
            } else {
                a[0].cmp(&b[0])
            }
        });


        //接下来转为求 高h 的最长递增子序列
        let mut hs = Vec::new();
        for e in &envelopes{
            let idx =  hs.binary_search(&e[1]).unwrap_or_else(|x|x);
            if idx == hs.len(){
                hs.push(e[1]);
            }else{
                hs[idx] = e[1];
            }
            
        }

        hs.len() as i32

    }
}
#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::max_envelopes(vec![vec![5,4],vec![6,4],vec![6,7],vec![2,3]]),3);
    }
}
