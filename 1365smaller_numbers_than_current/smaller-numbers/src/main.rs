fn main() {
    println!("Hello, world!");
}



pub struct Solution ;
use std::collections::HashMap;
impl Solution {
    pub fn smaller_numbers_than_current(nums: Vec<i32>) -> Vec<i32> {

        if nums.len()==0{
            return nums
        }

        if nums.len()==1{
            return vec!{0}
        }

        let mut nums2 = nums.clone();
        nums2.sort_unstable();

        let mut map = HashMap::new();
        map.insert(nums2[0],0);

  
        for i in 1..nums2.len(){

            if nums2[i]!=nums2[i-1]{
                map.insert(nums2[i],i as i32);
            }

        }

        let mut res = Vec::new();
        for i in &nums{
            res.push(map[i]);
        }


        res

    }
}