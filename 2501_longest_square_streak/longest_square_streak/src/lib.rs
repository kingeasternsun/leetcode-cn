use core::num;

struct Solution;
impl Solution {
    pub fn longest_square_streak(mut nums: Vec<i32>) -> i32 {
        let mut set  = std::collections::BTreeSet::new();
        let mut ret = -1;
        nums.sort_unstable();
        for n in &nums{

            if set.contains(n){
                continue;
            }

            let mut tmp = *n;
            let mut len = 0;
            while let Ok(_) = nums.binary_search(&tmp) {
                set.insert(tmp);
                len += 1;
                if let Some(newTmp) = tmp.checked_mul(tmp){
                    tmp = newTmp;
                }else{
                    break;
                }

            }

            ret = ret.max(len);

        }

        if ret <= 1{
            return -1
        }
        return ret

    }
}
pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        
        assert_eq!(Solution::longest_square_streak(vec![4,3,6,16,8,2]), 3);
        assert_eq!(Solution::longest_square_streak(vec![2,3,5,6,7]), -1);
        assert_eq!(Solution::longest_square_streak(vec![2,3,4,6,7]), 2);
    }
}
