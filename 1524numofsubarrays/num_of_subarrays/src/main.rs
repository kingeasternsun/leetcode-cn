fn main() {
    println!("Hello, world!");
}

pub struct Solution{}
use std::mem;
impl Solution {
    pub fn num_of_subarrays(arr: Vec<i32>) -> i32 {

        let mut ans = 0;
        let mut odd = 0;
        let mut even = 0;

        for i in &arr{
            if (i&1)==0{
                even+=1;
            }else{
                even+=1;
                mem::swap(&mut even,&mut odd);
            }

            ans += odd;
            ans = ans %(1000000000+7)
        }

        ans 
    }
}