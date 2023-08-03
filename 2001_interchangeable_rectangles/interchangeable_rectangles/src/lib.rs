struct Solution;
impl Solution {
    pub fn interchangeable_rectangles(rectangles: Vec<Vec<i32>>) -> i64 {
        let mut map = std::collections::HashMap::new();
        rectangles.iter().for_each(|x| {
            let k = Self::tidy(x);
            let a: &mut i32 = map.entry(k).or_insert(0);
            *a = *a + 1;
        });
        map.iter()
            .map(|(_, v)| (*v as i64) * (*v as i64 - 1) / 2)
            .sum()
    }

    // iterator's beauty
    pub fn interchangeable_rectangles2(rectangles: Vec<Vec<i32>>) -> i64 {
        let mut map: std::collections::HashMap<(i32, i32), i32> = std::collections::HashMap::new();
        rectangles.iter().map(Self::tidy).for_each(|k| {
            *map.entry(k).or_default() += 1;
        });
        map.iter()
            .map(|(_, v)| (*v as i64) * (*v as i64 - 1) / 2)
            .sum()
    }

    fn tidy(x: &Vec<i32>) -> (i32, i32) {
        let lcm = Self::lcm(x[0], x[1]);
        (x[0] / lcm, x[1] / lcm)
    }
    fn lcm(a: i32, b: i32) -> i32 {
        if a == b || a % b == 0 {
            return b;
        }

        if a < b {
            return Self::lcm(b, a);
        }
        return Self::lcm(b, a % b);
    }
}

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn lcm_works() {
        assert_eq!(Solution::lcm(4, 8), 4);
        assert_eq!(Solution::lcm(1, 8), 1);
        assert_eq!(Solution::lcm(4, 12), 4);
        assert_eq!(Solution::lcm(3, 9), 3);
    }

    #[test]
    fn interchangeable_rectangles_works() {
        assert_eq!(
            Solution::interchangeable_rectangles(vec![
                vec![4, 8],
                vec![3, 6],
                vec![10, 20],
                vec![15, 30]
            ]),
            6
        );
        assert_eq!(
            Solution::interchangeable_rectangles(vec![
                vec![4, 8],
                vec![3, 6],
                vec![10, 30],
                vec![15, 30]
            ]),
            3
        );
        assert_eq!(
            Solution::interchangeable_rectangles(vec![vec![4, 5], vec![7, 8]]),
            0
        );
    }
    fn interchangeable2_rectangles_works() {
        assert_eq!(
            Solution::interchangeable_rectangles2(vec![
                vec![4, 8],
                vec![3, 6],
                vec![10, 20],
                vec![15, 30]
            ]),
            6
        );
        assert_eq!(
            Solution::interchangeable_rectangles2(vec![
                vec![4, 8],
                vec![3, 6],
                vec![10, 30],
                vec![15, 30]
            ]),
            3
        );
        assert_eq!(
            Solution::interchangeable_rectangles2(vec![vec![4, 5], vec![7, 8]]),
            0
        );
    }
}
