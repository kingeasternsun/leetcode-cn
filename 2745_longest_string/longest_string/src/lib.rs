struct Solution;

impl Solution {
    // 由于 AB 自己可以和自己无限拼接，然后AB 既可以拼AA，也可以拼BB
    // 所以 先考虑 AA，BB 最多可以拼多少
    //  如果 AA 和 BB 个数一样，那么AA，BB 都可以拼完: 2(x+y)
    //  如果 AA 和 BB 个数不一样，那么 可以拼 AABB...BBAA 或 BBAA...AABB, AA或BB只能多一个: 2 * x.min(y) + 1
    // 然后 AB 有多少拼多少
    // 0ms 1.85mb
    pub fn longest_string(x: i32, y: i32, z: i32) -> i32 {
        if x == y {
            return (x + y + z) << 1;
        }

        return ((x.min(y) << 1) + 1 + z) << 1;
    }

    // 4ms 2.16mb
    pub fn longest_string1(x: i32, y: i32, z: i32) -> i32 {
        if x == y {
            return 2 * (x + y + z);
        }

        return 2 * (2 * x.min(y) + 1 + z);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::longest_string(2, 2, 1), 10);
        assert_eq!(Solution::longest_string(2, 5, 1), 12);
        assert_eq!(Solution::longest_string(3, 2, 2), 14);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::longest_string1(2, 2, 1), 10);
        assert_eq!(Solution::longest_string1(2, 5, 1), 12);
        assert_eq!(Solution::longest_string1(3, 2, 2), 14);
    }
}
