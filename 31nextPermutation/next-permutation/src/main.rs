fn main() {
    println!("Hello, world!");

    let mut input = vec![vec![1, 1, 3, 3, 2, 3, 7, 9, 9, 8, 8, 3], vec![1, 2, 3],vec![3, 2, 1],vec![1, 1, 5],vec![1, 3, 2],vec![5, 4, 7, 5, 3, 2]];
    let ans = vec![vec![1, 1, 3, 3, 2, 3, 8, 3, 7, 8, 9, 9], vec![1, 3, 2], vec![1, 2, 3], vec![1, 5, 1], vec![2, 1, 3],vec![5, 5, 2, 3, 4, 7]];

    for i in 0..input.len() {
        Solution::next_permutation(&mut input[i]);
        assert_eq!(input[i], ans[i]);
    }

    // Solution::next_permutation(&mut vec![1, 3, 2])
}




pub struct Solution;

impl Solution {
    pub fn next_permutation(nums: &mut Vec<i32>) {
        if nums.len() <= 1 {
            return;
        }

        let len = nums.len();
        if nums[len - 2] < nums[len - 1] {
            nums.swap(len - 1, len - 2);
            return;
        }

        let mut i = len - 2;
        // let mut j = len - 1;
        let mut beg = None;

        loop {
            if nums[i] < nums[i + 1] {
                beg = Some(i);
                break;
            }
            if i == 0 {
                break;
            }
            i -= 1;
        }

        // the nums all desending
        if beg.is_none() {
            nums.reverse();
            return;
        }

        let j = Self::bisearch(nums, i + 1, nums[i]); // find the j that it is the lowest nums[j] > nums[i]
        nums.swap(i, j); // exchange the value i j

        nums[(i+1)..].reverse();
    }

 
    fn bisearch(nums: &mut Vec<i32>, beg: usize, val: i32) -> usize {
        let mut beg = beg;
        let mut end = nums.len()-1;

        while beg < end {
            if beg == end - 1 {
                if nums[end ] > val {
                    return end ;
                }

                return beg;
            }

            let mid = (beg + end) / 2;
            if nums[mid] > val {
                beg = mid
            } else {
                end = mid - 1
            }
        }

        beg
    }
}
