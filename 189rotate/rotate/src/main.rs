fn main() {
    println!("Hello, world!");
}

pub struct Solution;
impl Solution {
    pub fn rotate(nums: &mut Vec<i32>, k: i32) {
        if k ==0{
            return 
        }

        let len=nums.len();
        nums.rotate_right(k as usize % len);

        
    }
}

#[cfg(test)]
mod tests{

    use super::*;

    #[test]
    fn rotate(){
        
    }

}