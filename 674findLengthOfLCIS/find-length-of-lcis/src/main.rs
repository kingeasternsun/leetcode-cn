fn main() {
    println!("Hello, world!");
}


pub struct Solution;

impl Solution {
    pub fn find_length_of_lcis(nums: Vec<i32>) -> i32 {

        let mut res = 0;
        let mut tmp_len = 0;

        for i in 0..nums.len(){
            if i ==0{
                res = 1;
                tmp_len  = 1;
                continue 
            }

            if &nums[i]>&nums[i-1]{
                tmp_len+=1;

            }else{
                if tmp_len > res {
                    res = tmp_len
                }

                tmp_len = 1
            }
        }

        return std::cmp::max(tmp_len,res)

    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn find() {
        assert_eq!(Solution::find_length_of_lcis(vec![2,2,2]), 1);
        assert_eq!(Solution::find_length_of_lcis(vec![1,3,5]), 3);
    }
}