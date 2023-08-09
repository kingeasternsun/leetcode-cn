struct Solution;

impl Solution {
    pub fn subtract_product_and_sum(n: i32) -> i32 {
        let mut sum = 0;
        let mut product = 1;
        let mut n = n;
        while n > 0 {
            let m = n % 10;
            n = n / 10;
            sum += m;
            product *= m;
        }
        product - sum
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::subtract_product_and_sum(234), 15);
        assert_eq!(Solution::subtract_product_and_sum(4421), 21);
        assert_eq!(Solution::subtract_product_and_sum(1), 0);
        assert_eq!(Solution::subtract_product_and_sum(2), 0);
        assert_eq!(Solution::subtract_product_and_sum(10), -1);
        assert_eq!(Solution::subtract_product_and_sum(20), -2);
    }
}
