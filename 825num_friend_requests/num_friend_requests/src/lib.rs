mod lt825a {
    // 4ms 2.1MB
    struct Solution;
    impl Solution {
        //  only ages[y] > 0.5*ages[x]+7  and  ages[y] <= ages[x] , send friend
        pub fn num_friend_requests(ages: Vec<i32>) -> i32 {
            let mut age_cnt = vec![0; 121];
            for age in &ages {
                age_cnt[*age as usize] += 1;
            }

            //  求前缀累加和
            for i in 1..age_cnt.len() {
                age_cnt[i] = age_cnt[i - 1] + age_cnt[i]
            }

            let mut res = 0;
            for age in ages {
                let lower = (age >> 1) + 7;
                // no friend match
                if lower >= age {
                    continue;
                }

                // remember to exclude yourself
                res += age_cnt[age as usize] - age_cnt[lower as usize] - 1;
            }

            res
        }
    }
    #[cfg(test)]
    mod tests {
        use crate::lt825a::Solution;

        #[test]
        fn it_works() {
            assert_eq!(Solution::num_friend_requests(vec![16, 16]), 2);
            assert_eq!(Solution::num_friend_requests(vec![16, 17, 18]), 2);
            assert_eq!(
                Solution::num_friend_requests(vec![20, 30, 100, 110, 120]),
                3
            );
        }
    }
}

mod lt825b {
    //4ms 2.2MB
    struct Solution;
    impl Solution {
        //  only ages[y] > 0.5*ages[x]+7  and  ages[y] <= ages[x] , send friend
        pub fn num_friend_requests(ages: Vec<i32>) -> i32 {
            let mut age_cnt = vec![0; 121];
            // count
            ages.iter().for_each(|age| age_cnt[*age as usize] += 1);
            //  求前缀累加和
            for i in 1..age_cnt.len() {
                age_cnt[i] = age_cnt[i - 1] + age_cnt[i]
            }

            let mut res = 0;
            ages.into_iter()
                .filter(|&age| age > (age >> 1) + 7)
                .for_each(|age| {
                    res += age_cnt[age as usize] - age_cnt[(age >> 1) as usize + 7] - 1
                });
            res
        }
    }
    #[cfg(test)]
    mod tests {
        use crate::lt825b::Solution;

        #[test]
        fn it_works() {
            assert_eq!(Solution::num_friend_requests(vec![16, 16]), 2);
            assert_eq!(Solution::num_friend_requests(vec![16, 17, 18]), 2);
            assert_eq!(
                Solution::num_friend_requests(vec![20, 30, 100, 110, 120]),
                3
            );
        }
    }
}
