fn main() {
    println!("Hello, world!");
    Solution::product_except_self(vec![1, 2, 3, 4]);
}
pub struct Solution {}

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        

        if nums.len()<2{
            return vec![];
        }

        let mut res = vec![1;nums.len()];
        let mut n = 1;
        for i in 0..nums.len(){
            n *=nums[i];
            res[i] = n;
        }

        n = 1;
        for i in (1..nums.len()).rev(){
            res[i]=res[i-1]*n;
            n*=nums[i];
        }
        res[0] = n;

        res

    }
}

// submission codes end

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_238() {
        assert_eq!(
            Solution::product_except_self(vec![1, 2, 3, 4]),
            vec![24, 12, 8, 6]
        );
    }
}