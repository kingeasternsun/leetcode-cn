struct Solution;
impl Solution {
    // 8ms 2.67mb
    pub fn all_paths_source_target(graph: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        let mut path = vec![];
        Self::dfs(&graph, 0, graph.len() as i32, &mut path, &mut ret);
        ret
    }

    // id 表示即将访问的节点， path中表示已经路过的节点列表
    fn dfs(graph: &Vec<Vec<i32>>, id: i32, n: i32, path: &mut Vec<i32>, ret: &mut Vec<Vec<i32>>) {
        path.push(id);
        if id == n - 1 {
            ret.push(path.clone());
            path.pop();
            return;
        }

        for &neigh in graph[id as usize].iter() {
            Self::dfs(graph, neigh, n, path, ret);
        }

        path.pop();
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::all_paths_source_target(vec![vec![1, 2], vec![3], vec![3], vec![]]),
            vec![vec![0, 1, 3], vec![0, 2, 3]]
        );

        assert_eq!(
            Solution::all_paths_source_target(vec![
                vec![4, 3, 1],
                vec![3, 2, 4],
                vec![3],
                vec![4],
                vec![]
            ]),
            vec![
                vec![0, 4],
                vec![0, 3, 4],
                vec![0, 1, 3, 4],
                vec![0, 1, 2, 3, 4],
                vec![0, 1, 4]
            ]
        );

        assert_eq!(
            Solution::all_paths_source_target(vec![vec![1, 2, 3], vec![2], vec![3], vec![]]),
            vec![vec![0, 1, 2, 3], vec![0, 2, 3], vec![0, 3]]
        );
    }
}
