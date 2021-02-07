/*
 * @Description: 665
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-07 10:08:19
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-07 14:05:54
 * @FilePath: \665checkPossibility\check_possibility\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn check_possibility(nums: Vec<i32>) -> bool {
        let mut incre_nums = Vec::new();
        let nums_len = nums.len();
        for n in nums {
            // println!("{:?}",incre_nums);
            let idx = incre_nums.binary_search(&(n + 1)).unwrap_or_else(|x| x);
            if idx == incre_nums.len() {
                incre_nums.push(n);
            } else {
                //如果找到了 要在处理下 找到最左边界
                let mut j = idx;
                loop {
                    if j == 0 {
                        break;
                    }

                    if incre_nums[j - 1] != incre_nums[j] {
                        break;
                    }
                    j -= 1;
                }

                incre_nums[j] = n;
            }

            // println!("after {:?}",incre_nums);
        }

        return incre_nums.len() >= nums_len - 1;
    }


}
//4 ,5, 1,2,3

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::check_possibility(vec![4, 2, 3]), true);
        assert_eq!(Solution::check_possibility(vec![4, 2, 1]), false);
        assert_eq!(Solution::check_possibility(vec![2, 3, 3, 2, 2]), false);
    }
}
