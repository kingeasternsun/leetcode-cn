mod lt475 {
    struct Solution;
    impl Solution {
        pub fn find_radius(houses: Vec<i32>, heaters: Vec<i32>) -> i32 {
            let mut houses = houses;
            houses.sort();
            let mut heaters = heaters;
            heaters.sort();
            let mut beg = 0;
            let mut end = houses[houses.len() - 1].max(heaters[heaters.len() - 1]);

            let mut mid = 0;
            let mut res = 0;
            while beg <= end {
                mid = (end - beg) / 2 + beg;
                if Self::radius(&houses, &heaters, mid) {
                    res = mid;
                    end = mid - 1;
                } else {
                    beg = mid + 1;
                }
            }

            res
        }

        fn radius(houses: &Vec<i32>, heaters: &Vec<i32>, r: i32) -> bool {
            let mut hi = 0;
            let mut hh = 0;

            while hi < houses.len() && hh < heaters.len() {
                if houses[hi] >= heaters[hh] - r && houses[hi] <= heaters[hh] + r {
                    hi += 1;
                    continue;
                }

                if houses[hi] < heaters[hh] - r {
                    return false;
                }
                hh += 1;
            }
            if hi == houses.len() {
                return true;
            }
            return false;
        }
    }

    #[cfg(test)]
    mod tests {
        use super::*;

        #[test]
        fn test_name() {
            assert_eq!(Solution::find_radius(vec![1, 5], vec![10]), 9);
            assert_eq!(Solution::find_radius(vec![1, 2, 3, 4], vec![1, 4]), 1);
            assert_eq!(Solution::find_radius(vec![1, 2, 3], vec![2]), 1);
            assert_eq!(Solution::find_radius(vec![1, 5], vec![2]), 3);
        }
    }
}

mod lt475_b {
    struct Solution;
    impl Solution {
        pub fn find_radius(houses: Vec<i32>, heaters: Vec<i32>) -> i32 {
            let mut heaters = heaters;
            heaters.sort_unstable();
            let mut res = 0;
            for h in &houses {
                let i = heaters.binary_search(h);
                let tmp = match i {
                    Ok(i) => 0,
                    Err(i) => {
                        let mut dis = i32::MAX;
                        if i < heaters.len() && heaters[i] - *h < dis {
                            dis = heaters[i] - *h;
                        }
                        if i > 0 && *h - heaters[i - 1] < dis {
                            dis = *h - heaters[i - 1];
                        }
                        dis
                    }
                };
                if tmp > res {
                    res = tmp;
                }
            }

            res
        }
    }

    #[cfg(test)]
    mod tests {
        use super::*;

        #[test]
        fn test_name() {
            assert_eq!(Solution::find_radius(vec![1, 5], vec![10]), 9);
            assert_eq!(Solution::find_radius(vec![1, 2, 3, 4], vec![1, 4]), 1);
            assert_eq!(Solution::find_radius(vec![1, 2, 3], vec![2]), 1);
            assert_eq!(Solution::find_radius(vec![1, 5], vec![2]), 3);
        }
    }
}
