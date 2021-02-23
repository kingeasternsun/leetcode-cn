/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-23 09:32:32
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-23 09:59:05
 * @FilePath: \1052maxSatisfied\max_satisfied\src\lib.rs
 */

pub struct Solution;

impl Solution {
    pub fn max_satisfied(customers: Vec<i32>, grumpy: Vec<i32>, x: i32) -> i32 {
        let x = x as usize;
        let total_sum = customers
            .iter()
            .zip(grumpy.iter())
            .fold(0, |acc, x| acc + x.0 * (1 - x.1));


        //满意提升度
        let mut tmp = customers
        .iter()
        .zip(grumpy.iter())
        .take(x)
        .fold(0, |acc, x| acc + x.0 * x.1);

        let mut max_delta = tmp;
        for i in x..customers.len(){
            if grumpy[i]==1{
                tmp+=customers[i];
            }

            if grumpy[i-x]==1{
                tmp-=customers[i-x];
            }

            if tmp>max_delta{
                max_delta = tmp;
            }
        }

        total_sum+max_delta
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::max_satisfied(vec![1, 0, 1, 2, 1, 1, 7, 5], vec![0, 1, 0, 1, 0, 1, 0, 1], 3),16);
        assert_eq!(Solution::max_satisfied(vec![4, 10, 10], vec![1,1,0], 2),24);
        assert_eq!(Solution::max_satisfied(vec![2], vec![0], 1),2);
        assert_eq!(Solution::max_satisfied(vec![2], vec![1], 1),2);
        assert_eq!(Solution::max_satisfied(vec![2,3], vec![1,0], 1),5);
    }
}
