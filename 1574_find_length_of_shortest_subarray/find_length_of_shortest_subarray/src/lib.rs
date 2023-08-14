use std::cmp::Ordering;

struct Solution;
impl Solution {
    pub fn find_length_of_shortest_subarray(array: Vec<i32>) -> i32 {
        if array.len() <= 1 {
            return 0;
        }

        // 从0开始的最长非递减子数组
        let mut left_i = 0;
        while left_i < array.len() - 1 && array[left_i] <= array[left_i + 1] {
            left_i += 1;
        }

        if left_i == array.len() - 1 {
            return 0;
        }

        // 以 array-len()-1 结尾的最长非递增子数组
        let mut right_i = array.len() - 1;
        while array[right_i - 1] <= array[right_i] {
            // NOTE: right_i 不会为1
            right_i -= 1;
        }

        let mut ret = array.len() - 1;
        for i in 0..left_i {
            let r = array[right_i..]
                .binary_search_by(|x| {
                    if *x < array[i] {
                        std::cmp::Ordering::Less
                    } else {
                        std::cmp::Ordering::Greater
                    }
                })
                .unwrap_or_else(|j| j);
            ret = ret.min(right_i + r - i - 1);
        }

        ret as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_length_of_shortest_subarray(vec![1, 2, 3, 10, 4, 2, 3, 5]),
            3
        );

        assert_eq!(
            Solution::find_length_of_shortest_subarray(vec![5, 4, 3, 2, 1]),
            4
        );

        assert_eq!(Solution::find_length_of_shortest_subarray(vec![1, 2, 3]), 0);
    }
}
