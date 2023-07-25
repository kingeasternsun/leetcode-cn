struct Solution;
impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
        let mut res = 0;

        let mut pre_max_id = 0;
        let mut sum = 0;
        for i in 1..height.len() {
            if height[i] < height[pre_max_id] {
                sum += height[i];
                continue;
            }

            res = res + (i as i32 - pre_max_id as i32 - 1) * height[pre_max_id] - sum;
            sum = 0;
            pre_max_id = i;
        }

        if pre_max_id >= height.len() - 1 {
            return res;
        }

        let pre_max = height[pre_max_id];
        let mut sum = 0; // IMPORTE
        pre_max_id = height.len() - 1;

        for i in (0..height.len() - 1).rev() {
            if height[i] < height[pre_max_id] {
                sum += height[i];
                continue;
            }
            res = res + (pre_max_id as i32 - i as i32 - 1) * height[pre_max_id] - sum;
            sum = 0;
            pre_max_id = i;

            if height[i] == pre_max {
                break;
            }
        }

        res
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::trap(vec![0,1,0,2,1,0,1,3,2,1,2,1]), 6);
        assert_eq!(Solution::trap(vec![4,2,0,3,2,5]), 9);
        assert_eq!(Solution::trap(vec![1]), 0);
        assert_eq!(Solution::trap(vec![1,1]), 0);
        assert_eq!(Solution::trap(vec![1,1,1]), 0);
        assert_eq!(Solution::trap(vec![0]), 0);
        assert_eq!(Solution::trap(vec![0,0]), 0);
        assert_eq!(Solution::trap(vec![0,0,0]), 0);
    }
}
