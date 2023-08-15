struct Solution;
impl Solution {
    // 0ms 2.05mb
    pub fn flip_and_invert_image(image: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut image = image;
        image.iter_mut().for_each(|row| {
            let (mut left, mut right) = (0, row.len() - 1);
            while left < right {
                row.swap(left, right);
                left += 1;
                right -= 1;
            }
            for i in 0..row.len() {
                row[i] = 1 - row[i];
            }
        });

        image
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::flip_and_invert_image(vec![vec![1, 1, 0], vec![1, 0, 1], vec![0, 0, 0]]),
            vec![vec![1, 0, 0], vec![0, 1, 0], vec![1, 1, 1]]
        );

        assert_eq!(
            Solution::flip_and_invert_image(vec![
                vec![1, 1, 0, 0],
                vec![1, 0, 0, 1],
                vec![0, 1, 1, 1],
                vec![1, 0, 1, 0]
            ]),
            vec![
                vec![1, 1, 0, 0],
                vec![0, 1, 1, 0],
                vec![0, 0, 0, 1],
                vec![1, 0, 1, 0]
            ]
        );

        assert_eq!(
            Solution::flip_and_invert_image(vec![vec![0, 1], vec![1, 0]]),
            vec![vec![0, 1], vec![1, 0]]
        );

        assert_eq!(
            Solution::flip_and_invert_image(vec![vec![1]]),
            vec![vec![0]]
        );
    }
}
