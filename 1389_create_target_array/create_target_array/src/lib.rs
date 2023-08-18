struct Solution;
impl Solution {
    // 0ms 2.1mb
    pub fn create_target_array(nums: Vec<i32>, index: Vec<i32>) -> Vec<i32> {
        let mut target = vec![0; nums.len()];
        let origin_len = nums.len();
        index.into_iter().zip(nums.into_iter()).for_each(|(id, n)| {
            target.insert(id as usize, n);
        });

        target.truncate(origin_len);
        target
    }

    pub fn create_target_array1(nums: Vec<i32>, index: Vec<i32>) -> Vec<i32> {
        let mut target = vec![];
        index.into_iter().zip(nums.into_iter()).for_each(|(id, n)| {
            target.insert(id as usize, n);
        });
        target
    }

    // 0ms 2.1mb
    pub fn create_target_array2(nums: Vec<i32>, index: Vec<i32>) -> Vec<i32> {
        let mut target = vec![-1; nums.len()];
        let len = nums.len();
        index.into_iter().zip(nums.into_iter()).for_each(|(id, n)| {
            if target[id as usize] == -1 {
                target[id as usize] = n;
            } else {
                target.copy_within(id as usize..len - 1, id as usize + 1);
                target[id as usize] = n;
            }
        });
        target
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::create_target_array(vec![0, 1, 2, 3, 4], vec![0, 1, 2, 2, 1]),
            vec![0, 4, 1, 3, 2]
        );

        assert_eq!(
            Solution::create_target_array(vec![1, 2, 3, 4, 0], vec![0, 1, 2, 3, 0]),
            vec![0, 1, 2, 3, 4]
        );
        assert_eq!(Solution::create_target_array(vec![1], vec![0]), vec![1]);
    }

    #[test]
    fn it_works1() {
        assert_eq!(
            Solution::create_target_array1(vec![0, 1, 2, 3, 4], vec![0, 1, 2, 2, 1]),
            vec![0, 4, 1, 3, 2]
        );

        assert_eq!(
            Solution::create_target_array1(vec![1, 2, 3, 4, 0], vec![0, 1, 2, 3, 0]),
            vec![0, 1, 2, 3, 4]
        );
        assert_eq!(Solution::create_target_array1(vec![1], vec![0]), vec![1]);
    }

    #[test]
    fn it_works2() {
        assert_eq!(
            Solution::create_target_array2(vec![0, 1, 2, 3, 4], vec![0, 1, 2, 2, 1]),
            vec![0, 4, 1, 3, 2]
        );

        assert_eq!(
            Solution::create_target_array2(vec![1, 2, 3, 4, 0], vec![0, 1, 2, 3, 0]),
            vec![0, 1, 2, 3, 4]
        );
        assert_eq!(Solution::create_target_array2(vec![1], vec![0]), vec![1]);
    }
}
