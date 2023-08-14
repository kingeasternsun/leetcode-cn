struct Solution;
impl Solution {
    pub fn num_special(mat: Vec<Vec<i32>>) -> i32 {
        let mut ret = 0;
        for i in 0..mat.len() {
            let mut num = None;
            for j in 0..mat[0].len() {
                if mat[i][j] == 1 {
                    if num.is_none() {
                        num.replace(j);
                    } else {
                        // 重复出现 1，置为 None
                        let _ = num.take();
                        break;
                    }
                }
            }
            // 在 i 行中，只有 mat[i][j] 为1
            if let Some(j) = num {
                // 判断 j 列是否只有 mat[i][j] 为 1
                let mut exist = false;
                for k in 0..mat.len() {
                    if k != i && mat[k][j] == 1 {
                        exist = true;
                    }
                }

                if !exist {
                    ret += 1;
                }
            }
        }

        ret
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::num_special(vec![vec![1, 0, 0], vec![0, 0, 1], vec![1, 0, 0]]),
            1
        );
        assert_eq!(
            Solution::num_special(vec![vec![1, 0, 0], vec![0, 1, 0], vec![0, 0, 1]]),
            3
        );
        assert_eq!(
            Solution::num_special(vec![
                vec![0, 0, 0, 1],
                vec![1, 0, 0, 0],
                vec![0, 1, 1, 0],
                vec![0, 0, 0, 0]
            ]),
            2
        );
    }
}
