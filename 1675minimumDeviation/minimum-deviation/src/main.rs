fn main() {
    println!("Hello, world!");
}

pub struct Solution;

impl Solution {
    pub fn minimum_deviation(nums: Vec<i32>) -> i32 {
        if nums.len() <= 1 {
            return 0;
        }

        //先把奇数变为偶数
        let nums = nums
            .iter()
            .map(|&x| {
                if x & 1 == 1 {
                    return 2 * x;
                }
                return x;
            })
            .collect::<Vec<_>>();
        let mut lo = nums.iter().min().unwrap().clone(); //最小值

        let mut max_heap = std::collections::BinaryHeap::from(nums);
        let mut res = max_heap.peek().unwrap() - lo;
        if res == 0 {
            return res;
        }

        while max_heap.peek().unwrap() & 1 == 0 {
            let top = max_heap.pop().unwrap() / 2;
            max_heap.push(top);

            //计算新的差值
            lo = std::cmp::min(top, lo);
            res = std::cmp::min(res, max_heap.peek().unwrap() - lo);
        }

        res
    }

}

#[cfg(test)]
mod tests {

    use super::*;
    #[test]
    fn one() {
        assert_eq!(Solution::minimum_deviation(vec![1, 2, 3, 4]), 1);
        assert_eq!(Solution::minimum_deviation(vec![4, 1, 5, 20, 3]), 3);
        assert_eq!(Solution::minimum_deviation(vec![2, 10, 8]), 3);
    }
}
