struct Solution;

impl Solution {
    // 0ms 2.26mb
    pub fn majority_element(nums: Vec<i32>) -> Vec<i32> {
        let mut first = 0;
        let mut first_cnt = 0;
        let mut second = 0;
        let mut second_cnt = 0;

        for &n in nums.iter() {
            if first_cnt > 0 && first == n {
                first_cnt += 1;
            } else if second_cnt > 0 && second == n {
                second_cnt += 1;
            } else if first_cnt == 0 {
                // 注意： 这两个判断必须放在上面两个if判断的后面，否则对于 [1,1]就会出错
                first_cnt += 1;
                first = n;
            } else if second_cnt == 0 {
                second_cnt += 1;
                second = n;
            } else {
                // 一起毁灭吧
                first_cnt -= 1;
                second_cnt -= 1;
            }
        }

        let mut ret = vec![];

        if first_cnt > 0 {
            if nums.iter().filter(|&x| *x == first).count() > nums.len() / 3 {
                ret.push(first);
            }
        }

        if second_cnt > 0 {
            if nums.iter().filter(|&x| *x == second).count() > nums.len() / 3 {
                ret.push(second);
            }
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::majority_element(vec![3, 2, 3]), vec![3]);
        assert_eq!(Solution::majority_element(vec![1]), vec![1]);
        assert_eq!(Solution::majority_element(vec![1, 2]), vec![1, 2]);
    }
}
