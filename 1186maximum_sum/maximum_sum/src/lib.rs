struct Solution;
impl Solution {
    // DP ，由于以第i个元素的“删除一次得到子数组最大和” 只跟 以第i-1个元素的“删除一次得到子数组最大和”和当前元素i有关，
    // 所以在DP中只需要记录前一个就可以了，这里就可以直接利用rust中迭代器的fold
    // 利用元祖的第三个元素记录最大值
    pub fn maximum_sum1(arr: Vec<i32>) -> i32 {
        let ret = arr.iter().skip(1).fold((arr[0],0,arr[0]), |pre, cur| {
            let mut cur = (pre.0.max(0) + cur, pre.0.max(pre.1 + cur),pre.2);
            cur.2 = cur.0.max(cur.1).max(cur.2);
            cur
        });
        ret.2
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::maximum_sum1(vec![1, -2, 0, 3]), 4);
        assert_eq!(Solution::maximum_sum1(vec![1, -2, -2, 3]), 3);
        assert_eq!(Solution::maximum_sum1(vec![-1, -1, -1, -1]), -1);
        assert_eq!(Solution::maximum_sum1(vec![2, 1, -2, -5, -2]), 3);
    }
}
