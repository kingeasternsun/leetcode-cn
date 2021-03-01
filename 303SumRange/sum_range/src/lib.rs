/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-01 09:30:46
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-01 09:44:45
 * @FilePath: /sum_range/src/lib.rs
 */

 struct NumArray {
    pre_sum: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        let mut pre = 0;
        let mut pre_sum: Vec<i32> = nums
            .iter()
            .map(|x| {
                pre += x;
                pre
            })
            .collect();

        NumArray { pre_sum: pre_sum }
    }

    fn sum_range(&self, i: i32, j: i32) -> i32 {
        if i > j {
            return 0;
        }
        if i == 0 {
            return self.pre_sum[j as usize];
        }

        return self.pre_sum[j as usize] - self.pre_sum[i as usize - 1];
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * let obj = NumArray::new(nums);
 * let ret_1: i32 = obj.sum_range(i, j);
 */

#[cfg(test)]
mod tests {
    use crate::NumArray;
    #[test]
    fn it_works() {
        let obj = NumArray::new(vec![-2, 0, 3, -5, 2, -1]);
        // let ret_1: i32 = obj.sum_range(i, j);
        assert_eq!(obj.sum_range(0, 2), 1);
        assert_eq!(obj.sum_range(2, 5), -1);
        assert_eq!(obj.sum_range(0, 5), -3);
    }
}
