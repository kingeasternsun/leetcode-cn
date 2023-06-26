struct Solution;
impl Solution {
    pub fn pivot_integer(n: i32) -> i32 {
        let total =(1+n)*n/2;
        let mut size = n;
        let mut left = 1;
        let mut right = n+1;
        while left < right {
            let mid = left + size/2;
            let cur_sum = (1+mid)*mid/2;
            match cur_sum.cmp(&(total+mid-cur_sum)){
                std::cmp::Ordering::Less => {left = mid+1;},
                std::cmp::Ordering::Greater =>{right = mid;},
                std::cmp::Ordering::Equal => return mid,
            }
            size = right - left;
        }

        -1

    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::pivot_integer(8), 6);
        assert_eq!(Solution::pivot_integer(1), 1);
        
        assert_eq!(Solution::pivot_integer(2), -1);
        assert_eq!(Solution::pivot_integer(3), -1);
        assert_eq!(Solution::pivot_integer(4), -1);
        assert_eq!(Solution::pivot_integer(5), -1);
    }
}
