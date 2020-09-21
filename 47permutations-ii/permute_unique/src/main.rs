fn main() {
    println!("Hello, world!");
    Solution::permute_unique(vec![1,2,3]);
}

pub struct Solution;

use std::collections::HashSet;
impl Solution {
    pub fn permute_unique(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result = Vec::new();
        let mut tmp = Vec::new();
        let mut nums = nums;
        Self::dfs(& mut nums[..],& mut tmp,& mut result);

        result
    }

    /*
    将每个元素和第一个元素替换，然后再对后面的部分递归
    考虑重复，使用map
    */
    pub fn dfs(nums: &mut [i32], tmp: &mut Vec<i32>, result: &mut Vec<Vec<i32>>) {
        if nums.is_empty() {
             println!("pass {:?} ",tmp);
            result.push(tmp.to_vec());
            return;
        }

        let mut set: HashSet<i32> = HashSet::with_capacity(nums.len());

        for i in 0..nums.len() {
            //如果作为一个元素尝试过
            if set.contains(&nums[i]) {
                // println!("pass {:?} {}",nums,i);
                continue;
            }

            set.insert(nums[i]);

            if i > 0 {
                nums.swap(i, 0);
            }

            tmp.push(nums[0]);
            Self::dfs(& mut nums[1..],tmp,result);
            if i > 0 {
                nums.swap(i, 0);// 这里一定要记得恢复
            }
            tmp.pop();

        }
    }
}
