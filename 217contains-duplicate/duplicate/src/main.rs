fn main() {
    println!("Hello, world!");
}


pub struct Solution{}


use std::collections::HashSet;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {

        let mut hs = HashSet::new();

        for i in 0..nums.len(){
           if hs.contains(&nums[i]){
               return true
           }

           hs.insert(nums[i]);
        }

        false

    }
}