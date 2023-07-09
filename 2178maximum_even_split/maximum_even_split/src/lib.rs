struct Solution;
impl Solution {
    // 贪心，每次取比前一个偶数大的最小偶数，假设M(n,sum) 表示拆分出来元素都大于n，互不相同，和为sum的元素个数
    // 假设 a<b , 则 M(0,sum) 可以选取一个a， 得到 1 + M(a, sum-a); 或选取一个b，得到 1 + M(b,sum-b)
    // 由于 a<b, 所以 sum-a > sum-b, 所以 M(a,sum-a) >= M(a, sum-b) >= M(b, sum-b)
    // 40ms 3.1MB
    pub fn maximum_even_split(final_sum: i64) -> Vec<i64> {
        if final_sum & 1 == 1 {
            return vec![];
        }

        let mut beg = 2;
        let mut left_sum = final_sum;
        let mut ans = vec![];
        while beg <= left_sum {
            left_sum -= beg;
            ans.push(beg);
            beg += 2;
        }
        if let Some(s) = ans.last_mut() {
            *s += left_sum;
        }
        ans
    }

    // 80ms 3.2MB 预先分配内存 
    pub fn maximum_even_split1(final_sum: i64) -> Vec<i64> {
        if final_sum & 1 == 1 {
            return vec![];
        }

        let mut beg = 2;
        let mut left_sum = final_sum;
        let mut ans = Vec::with_capacity((final_sum as f64).sqrt() as usize);
        while beg <= left_sum {
            left_sum -= beg;
            ans.push(beg);
            beg += 2;
        }
        if let Some(s) = ans.last_mut() {
            *s += left_sum;
        }
        ans
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::maximum_even_split(12), vec![2, 4, 6]);
        assert_eq!(Solution::maximum_even_split(7), vec![]);
        assert_eq!(Solution::maximum_even_split(28), vec![2, 4, 6, 16]);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::maximum_even_split1(12), vec![2, 4, 6]);
        assert_eq!(Solution::maximum_even_split1(7), vec![]);
        assert_eq!(Solution::maximum_even_split1(28), vec![2, 4, 6, 16]);
    }
}
