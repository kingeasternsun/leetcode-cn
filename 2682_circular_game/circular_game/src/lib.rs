struct Solution;
impl Solution {
    pub fn circular_game_losers(n: i32, k: i32) -> Vec<i32> {
        let mut iter = 1;
        let mut cur_friend = 1;
        let mut friends = vec![false; n as usize + 1];
        friends[1] = true;
        loop {
            cur_friend = (cur_friend + iter * k) % n;
            if cur_friend == 0 {
                cur_friend = n;
            }

            // 如果之前已经接到过球
            if friends[cur_friend as usize] {
                break;
            }
            friends[cur_friend as usize] = true;
            iter += 1;
            cur_friend = cur_friend;
        }

        let mut ret = vec![];
        for i in 1..friends.len() {
            if !friends[i] {
                ret.push(i as i32)
            }
        }
        return ret;
    }

    // 把id为0的friend 当id为n的friend, 需要考虑 n 为 1 的情况，没有上面简单
    pub fn circular_game_losers1(n: i32, k: i32) -> Vec<i32> {
        let n = n as usize;
        let k = k as usize;
        let mut iter = 1;
        let mut cur_friend = 1;
        let mut friends = vec![false; n];
        friends[1 % n] = true;
        loop {
            cur_friend = (cur_friend + iter * k) % n;

            // 如果之前已经接到过球
            if friends[cur_friend] {
                break;
            }
            friends[cur_friend] = true;
            iter += 1;
            cur_friend = cur_friend;
        }

        let mut ret = vec![];
        for i in 0..friends.len() {
            if !friends[i] {
                if i == 0 {
                    ret.push(n as i32);
                } else {
                    ret.push(i as i32);
                }
            }
        }
        return ret;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::circular_game_losers(5, 2), vec![4, 5]);
        assert_eq!(Solution::circular_game_losers(4, 4), vec![2, 3, 4]);
        assert_eq!(Solution::circular_game_losers(1, 1), vec![]);
        assert_eq!(Solution::circular_game_losers(2, 1), vec![]);
        assert_eq!(Solution::circular_game_losers(3, 1), vec![3]);
        assert_eq!(Solution::circular_game_losers(2, 2), vec![2]);
        assert_eq!(Solution::circular_game_losers(3, 2), vec![2]);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::circular_game_losers1(5, 2), vec![5, 4]);
        assert_eq!(Solution::circular_game_losers1(4, 4), vec![4, 2, 3]);
        assert_eq!(Solution::circular_game_losers1(1, 1), vec![]);
        assert_eq!(Solution::circular_game_losers1(2, 1), vec![]);
        assert_eq!(Solution::circular_game_losers1(3, 1), vec![3]);
        assert_eq!(Solution::circular_game_losers1(2, 2), vec![2]);
        assert_eq!(Solution::circular_game_losers1(3, 2), vec![2]);
    }
}
