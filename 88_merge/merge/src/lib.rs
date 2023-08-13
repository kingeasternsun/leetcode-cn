struct Solution;
impl Solution {
    // 双指针： p1 指向 nums1， p2 指向 nums2, p0指向要插入的位置，都 从后往前
    // 如果 *p1 >= *p2, 则把 *p2 塞入 p0
    // 否则，则把 *p1 塞入 p0
    // 可以证明，无论是 p1 的塞进去，还是 p2 的塞进去，nums1中的剩余空闲空间可以塞入 nums2 中剩下的元素

    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut p1 = m - 1;
        let mut p2 = n - 1;
        let mut p0 = (m + n) - 1;

        while p1 >= 0 && p2 >= 0 {
            if nums1[p1 as usize] >= nums2[p2 as usize] {
                nums1[p0 as usize] = nums1[p1 as usize];
                p1 -= 1;
            } else {
                nums1[p0 as usize] = nums2[p2 as usize];
                p2 -= 1;
            }

            p0 -= 1;
        }

        if p2 >= 0 {
            // NOTE: 注意 src 和 dst 的长度必须一样
            nums1[0..(p2 as usize + 1)].copy_from_slice(&nums2[0..(p2 as usize + 1)]);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let mut nums1 = vec![1, 2, 3, 0, 0, 0];
        let mut nums2 = vec![2, 5, 6];
        Solution::merge(&mut nums1, 3, &mut nums2, 3);
        assert_eq!(nums1, vec![1, 2, 2, 3, 5, 6]);

        let mut nums1 = vec![1];
        let mut nums2 = vec![];
        Solution::merge(&mut nums1, 1, &mut nums2, 0);
        assert_eq!(nums1, vec![1]);

        let mut nums1 = vec![0];
        let mut nums2 = vec![1];
        Solution::merge(&mut nums1, 0, &mut nums2, 1);
        assert_eq!(nums1, vec![1]);
    }
}
