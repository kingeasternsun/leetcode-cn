struct Solution;

impl Solution {
    // 统计'1'的个数，然后看是否可以被3整除
    // 如果可以被3整除，假设得到m
    // 要把字符串分割为包含相同'1'个数的三个子字符串，只能在 m 个'1'和 m+1 个'1'之间切割，以及 2m,2m+1 个'1'之间切割
    // 那么根据第 m,m+1 个'1'所在的索引，以及 2m,2m+1 个'1'所在的索引就可以算出来
    // 4ms 2.17mb
    pub fn num_ways(s: String) -> i32 {
        let bytes = s.as_bytes();
        let num1 = bytes.iter().filter(|&x| *x == b'1').count();
        if num1 == 0 {
            let ret = (bytes.len() - 1) * (bytes.len() - 2) / 2 % 1000000007;
            return ret as i32;
        }

        if num1 % 3 > 0 {
            return 0;
        }

        // println!("{}", num1);

        let m = num1 / 3;
        let mut first = (0, 0);
        let mut second = (0, 0);
        let mut cnt = 0;
        for i in 0..bytes.len() {
            if bytes[i] == b'1' {
                cnt += 1;
                if cnt == m {
                    first.0 = i
                }

                if cnt == m + 1 {
                    first.1 = i
                }
                if cnt == 2 * m {
                    second.0 = i
                }
                if cnt == 2 * m + 1 {
                    second.1 = i
                }
            }
        }

        // println!("{:?}{:?}", first, second);

        let ret = (first.1 - first.0) * (second.1 - second.0) % 1000000007;
        return ret as i32;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::num_ways(String::from("10101")), 4);
        assert_eq!(Solution::num_ways(String::from("1001")), 0);
        assert_eq!(Solution::num_ways(String::from("0000")), 3);
        assert_eq!(Solution::num_ways(String::from("100100010100110")), 12);
    }
}
