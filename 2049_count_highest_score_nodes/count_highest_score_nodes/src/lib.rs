struct Solution;
impl Solution {
    pub fn count_highest_score_nodes(parents: Vec<i32>) -> i32 {
        // let mut ret;
        let total_num = parents.len();

        let mut childs: Vec<Vec<usize>> = vec![Vec::new(); total_num];
        parents
            .iter()
            .enumerate()
            .filter(|x| *x.1 >= 0)
            .for_each(|(c, p)| {
                childs[*p as usize].push(c);
            });

        // 利用递归，求每个子树的大小和得分
        // NOTE: 乘积有可能溢出，所以要用u64
        let mut sub_tree_score = vec![0_u64; total_num];
        Self::dp(0, total_num, &childs, &mut sub_tree_score);

        let max_sub_tree_score = sub_tree_score.iter().max().unwrap_or(&0);
        sub_tree_score
            .iter()
            .filter(|&x| *x == *max_sub_tree_score)
            .count() as i32
    }

    // 这个不需要记忆化，因为每个子树只会被使用一次
    fn dp(
        n: usize,
        total_num: usize,
        childs: &Vec<Vec<usize>>,
        sub_tree_score: &mut Vec<u64>,
    ) -> i32 {
        // this is leaf 叶子节点
        if childs[n].len() == 0 {
            sub_tree_score[n] = total_num as u64 - 1;
            return 1;
        }

        let sub_trees = childs[n]
            .iter()
            .map(|&x| Self::dp(x, total_num, childs, sub_tree_score))
            .collect::<Vec<i32>>();

        let sub_tree_num = sub_trees.iter().sum::<i32>() + 1;
        sub_tree_score[n] = sub_trees.iter().map(|&x| x as u64).product::<u64>();
        if sub_tree_num != total_num as i32 {
            sub_tree_score[n] *= total_num as u64 - sub_tree_num as u64;
        }
        sub_tree_num
    }

    pub fn count_highest_score_nodes2(parents: Vec<i32>) -> i32 {
        // let mut ret;
        let total_num = parents.len();

        let mut childs: Vec<Vec<usize>> = vec![Vec::new(); total_num];
        parents
            .iter()
            .enumerate()
            .filter(|x| *x.1 >= 0)
            .for_each(|(c, p)| {
                childs[*p as usize].push(c);
            });

        // 利用递归，求每个子树的大小和得分
        // NOTE: 乘积有可能溢出，所以要用u64
        let mut sub_tree_score = (0_u64,0);
        Self::dp2(0, total_num, &childs, &mut sub_tree_score);

        sub_tree_score.1
    }
    // 这个不需要记忆化，因为每个子树只会被使用一次
    fn dp2(
        n: usize,
        total_num: usize,
        childs: &Vec<Vec<usize>>,
        max_sub_tree_score: &mut (u64, i32),
    ) -> i32 {
        let mut sub_tree_score = 0;
        let mut sub_tree_num = 0;
        // this is leaf 叶子节点
        if childs[n].len() == 0 {
            sub_tree_score = total_num as u64 - 1;
            sub_tree_num = 1;
        } else {
            let sub_trees = childs[n]
                .iter()
                .map(|&x| Self::dp2(x, total_num, childs, max_sub_tree_score))
                .collect::<Vec<i32>>();

            sub_tree_num = sub_trees.iter().sum::<i32>() + 1;
            sub_tree_score = sub_trees.iter().map(|&x| x as u64).product::<u64>();
            if sub_tree_num != total_num as i32 {
                sub_tree_score *= total_num as u64 - sub_tree_num as u64;
            }
        }

        if sub_tree_score > max_sub_tree_score.0 {
            *max_sub_tree_score = (sub_tree_score, 1)
        } else if sub_tree_score == max_sub_tree_score.0 {
            max_sub_tree_score.1 += 1;
        }

        sub_tree_num
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::count_highest_score_nodes(vec![-1, 2, 0, 2, 0]), 3);
        assert_eq!(Solution::count_highest_score_nodes(vec![-1, 2, 0]), 2);
    }

    #[test]
    fn it_works2() {
        assert_eq!(Solution::count_highest_score_nodes2(vec![-1, 2, 0, 2, 0]), 3);
        assert_eq!(Solution::count_highest_score_nodes2(vec![-1, 2, 0]), 2);
    }
}
