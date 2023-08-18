struct Solution;
impl Solution {
    pub fn is_possible_divede(nums: Vec<i32>, k: i32) -> bool {
        let k = k as usize;
        // 不能整除
        if nums.len() % k > 0 {
            return false;
        }

        let iter_cnt = nums.len() / k;

        // 求最大，最小值
        let min = nums.iter().min().map_or(0, |x| x.clone()) as usize;
        let max = nums.iter().max().map_or(0, |x| x.clone()) as usize;
        // 逆向推理，假设可以拆分，那么极端情况下 nums中是一个完全顺子，最大值和最小值的差为
        if max - min >= nums.len() {
            return false;
        }

        let mut num_map = vec![0; max - min + 1];
        nums.into_iter().for_each(|x| {
            num_map[x as usize - min] += 1;
        });

        println!("{:?}", num_map);
        let mut begin_id = 0;
        for iter in 0..iter_cnt {
            for i in 0..k {
                if num_map[begin_id + i] == 0 {
                    return false;
                }
                num_map[begin_id + i] -= 1;
            }

            let mut next_id = begin_id;
            while next_id < begin_id + k {
                if num_map[next_id] > 0 {
                    break;
                }
                next_id += 1;
            }

            begin_id = next_id;

            println!("{:?} {}", num_map, begin_id);
        }
        true
        // begin_id == num_map.len()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        // assert_eq!(
        //     Solution::is_possible_divede(vec![1, 2, 3, 3, 4, 4, 5, 6], 4),
        //     true
        // );
        assert_eq!(
            Solution::is_possible_divede(vec![3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11], 3),
            true
        );
        // assert_eq!(
        //     Solution::is_possible_divede(vec![3, 3, 2, 2, 1, 1], 3),
        //     true
        // );
        // assert_eq!(Solution::is_possible_divede(vec![1, 2, 3, 4], 3), false);
    }
}
