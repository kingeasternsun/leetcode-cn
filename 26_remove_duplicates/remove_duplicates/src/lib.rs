struct Solution;
impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        nums.dedup();
        nums.len() as i32

    }

    pub fn remove_duplicates1(nums: &mut Vec<i32>) -> i32 {
       let mut left = 0;
       let mut right = 0;
       while right < nums.len(){
        if nums[left] == nums[right]{
            right+=1;
        }else{
            left +=1;
            nums[left] = nums[right];
            right+=1;
        }

       }
       left as i32 +1

    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
    }
}
