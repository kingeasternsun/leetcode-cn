pub struct Solution;
impl Solution {
    pub fn restore_array(adjacent_pairs: Vec<Vec<i32>>) -> Vec<i32> {
        let mut res = vec![];

        let mut g = std::collections::HashMap::new();
        for v in &adjacent_pairs {
            let mut ad = g.entry(v[0]).or_insert(vec![]);
            ad.push(v[1]);

            let mut ad1 = g.entry(v[1]).or_insert(vec![]);
            ad1.push(v[0]);
        }

        let mut visited = std::collections::HashSet::new(); //记录已经访问了哪些

        for (k, v) in &g {
            if v.len() == 1 {
                visited.insert(k);
                res.push(*k);
                break;
            }
        }

        loop {
            let mut new_add = 0;
            let end = res[res.len() - 1];
            for ad in g.get(&end).unwrap() {
                if visited.get(ad).is_none() {
                    //没有访问郭
                    visited.insert(ad);
                    res.push(*ad);
                    new_add += 1;
                }
            }

            if new_add == 0 {
                break;
            }
        }

        res
    }
}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        let adjacentPairs = vec![vec![2,1],vec![3,4],vec![3,2]];
        assert_eq!(Solution::restore_array(adjacentPairs),vec![1,2,3,4]);
    }
}
