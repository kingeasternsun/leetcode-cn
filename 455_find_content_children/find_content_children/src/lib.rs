struct Solution;
impl Solution {
    // 8ms 2.17mb 贪心算法，先满足胃口最小的孩子
    pub fn find_content_children(g: Vec<i32>, s: Vec<i32>) -> i32 {
        let mut g = g;
        let mut s = s;

        g.sort_unstable();
        s.sort_unstable();

        let mut ret = 0;
        let mut i = 0;
        let mut j = 0;
        while i < g.len() && j < s.len() {
            if g[i] <= s[j] {
                ret += 1;
                i += 1;
                j += 1;
            } else {
                j += 1;
            }
        }
        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_content_children(vec![1, 2, 3], vec![1, 1]),
            1
        );
        assert_eq!(
            Solution::find_content_children(vec![1, 2], vec![1, 2, 3]),
            2
        );
    }
}
