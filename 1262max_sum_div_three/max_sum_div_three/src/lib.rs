struct Solution;
impl Solution {
    pub fn max_sum_div_three1(nums: Vec<i32>) -> i32 {
        let mut sum = 0;
        let mut mod1 = vec![]; // record the minest two values of mod 1
        let mut mod2 = vec![]; // record the maxest two values of mod 1
        nums.iter().for_each(|&x| {
            sum += x;
            match x % 3 {
                1 => {
                    mod1.push(x);
                    if mod1.len() == 3 {
                        mod1.sort();
                        mod1.pop();
                    }
                }
                2 => {
                    mod2.push(x);
                    if mod2.len() == 3 {
                        mod2.sort();
                        mod2.pop();
                    }
                }
                _ => {}
            }
        });

        if sum % 3 == 1 {
            let r1 = mod1.first();
            let r2 = mod2.chunks_exact(2).next().map(|x| x[0] + x[1]);

            return match (r1, r2) {
                (Some(n1), Some(n2)) => (sum - n1).max(sum - n2),
                (None, Some(n2)) => sum - n2,
                (Some(n1), None) => sum - n1,
                _ => 0,
            };
        }

        if sum % 3 == 2 {
            let r1: Option<i32> = mod1.chunks_exact(2).next().map(|x| x[0] + x[1]);
            let r2 = mod2.first();

            return match (r1, r2) {
                (Some(n1), Some(n2)) => (sum - n1).max(sum - n2),
                (None, Some(n2)) => sum - n2,
                (Some(n1), None) => sum - n1,
                _ => 0,
            };
        }

        sum
    }

    pub fn max_sum_div_three(nums: Vec<i32>) -> i32 {
        // note: the init i32::MIN is import
        let ret = nums.iter().fold([0, i32::MIN, i32::MIN], |sum, &item| {
            let mut tmp = [0, 0, 0];
            for i in 0..3 {
                let newid = ((i + item) % 3) as usize;
                tmp[newid] = sum[newid].max(sum[i as usize] + item);
            }
            tmp
        });
        ret[0]
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        use crate::Solution;
        assert_eq!(Solution::max_sum_div_three1(vec![3, 6, 5, 1, 8]), 18);
        assert_eq!(Solution::max_sum_div_three1(vec![3, 6, 5, 1, 9]), 24);
        assert_eq!(Solution::max_sum_div_three1(vec![4]), 0);
        assert_eq!(Solution::max_sum_div_three1(vec![3]), 3);
    }
    #[test]
    fn dp_works() {
        use crate::Solution;
        assert_eq!(Solution::max_sum_div_three(vec![3, 6, 5, 1, 8]), 18);
        assert_eq!(Solution::max_sum_div_three(vec![3, 6, 5, 1, 9]), 24);
        assert_eq!(Solution::max_sum_div_three(vec![4]), 0);
        assert_eq!(Solution::max_sum_div_three(vec![3]), 3);
    }
}
