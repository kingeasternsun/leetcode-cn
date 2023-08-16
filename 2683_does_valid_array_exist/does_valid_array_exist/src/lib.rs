struct Solution;
impl Solution {
    // 找第一个突破口 然后逐个反向逆推
    pub fn does_valid_array_exist(derived: Vec<i32>) -> bool {
        // derived[i] = origin[i] ^ origin[i+1] 所以  origin[i+1] = derived[i] ^ origin[i]
        // origin[0] 无非为0 或 1， 分别尝试反向逆推即可
        let mut origin = vec![0; derived.len()];
        if Self::check_origin(&derived, &mut origin) {
            return true;
        }

        origin[0] = 1;
        return Self::check_origin(&derived, &mut origin);
    }

    fn check_origin(derived: &Vec<i32>, origin: &mut Vec<i32>) -> bool {
        for i in 1..derived.len() {
            origin[i] = derived[i - 1] ^ origin[i - 1];
        }

        // 判断 derived[n-1] = origin[n-1] ^ origin[0]  是否满足
        derived[derived.len() - 1] == (origin[derived.len() - 1] ^ origin[0])
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        // 0,1,1 ->  1,0,1
        assert_eq!(Solution::does_valid_array_exist(vec![1, 1, 0]), true);
        assert_eq!(Solution::does_valid_array_exist(vec![1, 0, 1]), true);
        assert_eq!(Solution::does_valid_array_exist(vec![1, 1]), true);
        assert_eq!(Solution::does_valid_array_exist(vec![1, 0]), false);
    }
}
