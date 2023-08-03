struct Solution;

impl Solution {
    /*
    ┌─┬─┬─┐     ┌─┬─┐     ┌─┐
    │1│3│4│     │3│4│     │4│
    ├─┼─┼─┤     ├─┼─┤     ├─┤
    │2│6│8│     │6│8│     │8│
    └─┴─┴─┘     └─┴─┘     └─┘
         */
    // 如果是double的数组，那么这个原始数组的最小值肯定有个2倍的值也存在于双倍数组中
    // 我们把这个最小值和最小值的2倍从双倍数组中移除，剩下的数组肯定也是双倍数组
    // 算法：每次移除最小值和对应的double值
    // 使用 排序 和 HashMap
    // 注意点： 0 的 2倍 也是 0
    pub fn find_original_array(mut changed: Vec<i32>) -> Vec<i32> {
        if changed.len() & 1 == 1 {
            return vec![];
        }
        let orgin_len = changed.len() / 2;

        changed.sort_unstable();

        let mut map: std::collections::HashMap<i32, i32> = std::collections::HashMap::new();
        let mut ret = vec![];
        changed.iter().for_each(|x| {
            *map.entry(*x).or_default() += 1;
        });

        for n in changed {
            let v1 = map.get(&n).unwrap().clone();
            if v1 == 0 {
                continue;
            }

            // 移除一个n， 注意 n 可能为0，所以要放在get 2*n 之前
            map.insert(n, v1 - 1);

            let v2 = map.get(&(2 * n)).map_or(0, |v| v.clone());
            if v2 == 0 {
                break;
            }
            ret.push(n);

            map.insert(2 * n, v2 - 1);
        }

        // println!("{:?}", ret);
        if ret.len() == orgin_len {
            return ret;
        }

        return vec![];
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::find_original_array(vec![1, 3, 4, 2, 6, 8]),
            vec![1, 3, 4]
        );
        assert_eq!(
            Solution::find_original_array(vec![1, 2, 3, 4, 6, 8]),
            vec![1, 3, 4]
        );
        assert_eq!(Solution::find_original_array(vec![0, 1, 1, 0]), vec![]);
        assert_eq!(Solution::find_original_array(vec![0, 1, 2, 0]), vec![0, 1]);
        assert_eq!(Solution::find_original_array(vec![0, 1, 0, 0]), vec![]);
        assert_eq!(Solution::find_original_array(vec![0, 1]), vec![]);
        assert_eq!(Solution::find_original_array(vec![0, 0]), vec![0]);
        assert_eq!(Solution::find_original_array(vec![0, 0, 0, 0]), vec![0, 0]);
    }
}
