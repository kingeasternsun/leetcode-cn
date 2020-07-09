fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {
    pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {

        if nums.is_empty() {
            return vec![];
        }

        let mut pre = 0_usize;
        let mut res = Vec::new();
        for i in 1..nums.len(){
            if nums[i]==nums[i-1]+1{
                continue
            }

            Self::push_range(& mut res,&nums,pre,i-1);
            pre = i;
        }

        Self::push_range(& mut res,&nums,pre,nums.len()-1);

        res 
    }

    pub fn push_range(res:& mut Vec<String>,nums: &Vec<i32>,i: usize,j: usize){

        if i == j{
            res.push(nums[i].to_string())
        }else{
            res.push(format!("{}->{}",nums[i],nums[j]))
        }
    }
}

