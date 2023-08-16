struct Solution;
impl Solution {
    // 4ms 2.2MB
    pub fn max_dist_to_closest(seats: Vec<i32>) -> i32 {
        let seat_len = seats.len();
        let mut ret = 0;
        let mut pre_seat = None;
        for (id, people) in seats.into_iter().enumerate() {
            if people == 0 {
                continue;
            }

            match pre_seat {
                None => {
                    ret = id;
                    pre_seat = Some(id);
                }

                Some(pre_id) => {
                    // 101  -> 2/1 = 1
                    // 1001 -> 3/2 = 1
                    if id >= pre_id + 2 {
                        ret = ret.max((id - pre_id) / 2);
                    }
                    pre_seat = Some(id);
                }
            }
        }

        match pre_seat {
            None => panic!("would not happen"),
            Some(pre_id) => {
                ret = ret.max(seat_len - 1 - pre_id);
            }
        }

        ret as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::max_dist_to_closest(vec![1, 0, 0, 0, 1, 0, 1]), 2);
        assert_eq!(Solution::max_dist_to_closest(vec![1, 0, 0, 0]), 3);
        assert_eq!(Solution::max_dist_to_closest(vec![0, 0, 0, 1]), 3);
        assert_eq!(Solution::max_dist_to_closest(vec![0, 1]), 1);
    }
}
