struct Solution;
impl Solution {
    // 把要求 group大小 一样的人聚集在一起，再根据group大小 分组
    pub fn group_the_people(group_sizes: Vec<i32>) -> Vec<Vec<i32>> {
        let mut group_map = vec![vec![]; 501];
        group_sizes.iter().enumerate().for_each(|x| {
            group_map[x.1.clone() as usize].push(x.0 as i32);
        });

        let mut ret = vec![];
        for (g_size, v) in group_map.iter().enumerate() {
            if g_size == 0 {
                continue;
            }
            for sub_group in v.chunks(g_size) {
                ret.push(sub_group.to_vec());
            }
        }

        ret
    }

    // 优化版本1, 当 人数已经可以构成一组后就加入到 ret 中，而不是等全部分类完成后再分组
    pub fn group_the_people1(group_sizes: Vec<i32>) -> Vec<Vec<i32>> {
        let mut group_map = vec![vec![]; 501];
        let mut ret = vec![];
        group_sizes.iter().enumerate().for_each(|(id, &g_size)| {
            group_map[g_size as usize].push(id as i32);
            if group_map[g_size as usize].len() == g_size as usize {
                ret.push(group_map[g_size as usize].clone());
                group_map[g_size as usize].clear();
            }
        });

        ret
    }

    // 优化版本2, 通过 std::mem::replace 零拷贝 优化
    pub fn group_the_people2(group_sizes: Vec<i32>) -> Vec<Vec<i32>> {
        let mut group_map = vec![vec![]; 501];
        let mut ret = vec![];
        group_sizes.iter().enumerate().for_each(|(id, &g_size)| {
            group_map[g_size as usize].push(id as i32);
            if group_map[g_size as usize].len() == g_size as usize {
                ret.push(std::mem::replace(&mut group_map[g_size as usize], vec![]));
            }
        });

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::group_the_people(vec![3, 3, 3, 3, 3, 1, 3]),
            vec![vec![5], vec![0, 1, 2], vec![3, 4, 6]]
        );

        assert_eq!(
            Solution::group_the_people(vec![2, 1, 3, 3, 3, 2]),
            vec![vec![1], vec![0, 5], vec![2, 3, 4]]
        );
    }

    #[test]
    fn it_works1() {
        assert_eq!(
            Solution::group_the_people1(vec![3, 3, 3, 3, 3, 1, 3]),
            vec![vec![0, 1, 2], vec![5], vec![3, 4, 6]]
        );

        assert_eq!(
            Solution::group_the_people1(vec![2, 1, 3, 3, 3, 2]),
            vec![vec![1], vec![2, 3, 4], vec![0, 5]]
        );
    }

    #[test]
    fn it_works2() {
        assert_eq!(
            Solution::group_the_people2(vec![3, 3, 3, 3, 3, 1, 3]),
            vec![vec![0, 1, 2], vec![5], vec![3, 4, 6]]
        );

        assert_eq!(
            Solution::group_the_people2(vec![2, 1, 3, 3, 3, 2]),
            vec![vec![1], vec![2, 3, 4], vec![0, 5]]
        );
    }
}
