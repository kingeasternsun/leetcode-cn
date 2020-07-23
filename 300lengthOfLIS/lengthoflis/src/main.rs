fn main() {
    println!("Hello, world!");
    Solution::length_of_lis(vec![10,9,2,5,3,7,101,18]);
}

pub struct Solution{}

// learn from https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-she-ji-fang-fa-zhi-pai-you-xi-jia/
impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        
        let mut tops = Vec::new();

        for i in 0..nums.len(){

            let mut start = 0;
            let mut end = tops.len();

            while start < end {
                let mid = (start+end)>>1;
                if tops[mid] < nums[i]{
                    start = mid+1;
                }else{
                    end = mid;
                }
            }

            // can not found value in top that large or equal nums[i]
            if start == tops.len(){
                tops .push(nums[i]);
            }else{
                tops[start] = nums[i];
            }

        }

        tops.len() as i32

    }
}