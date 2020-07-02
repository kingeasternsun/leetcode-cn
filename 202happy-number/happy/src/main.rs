fn main() {
    println!("Hello, world!");
    Solution::is_happy(19);
}
pub struct Solution {}

use std::collections::HashSet;
impl Solution {
    pub fn is_happy(n: i32) -> bool {
        
        let mut set = HashSet::new();
        let mut n = n;
        while n!=1{


            if set.contains(&n){
                return false
            }

            set.insert(n);
            n = Self::get_sum(n);
        }

        true
    }

    pub fn get_sum(n:i32)->i32{
        let mut sum = 0;
        let mut n = n;
        while n >0 {
            sum = sum + (n%10)*(n%10);
            n = n/10;
        }
        sum 

    }
}