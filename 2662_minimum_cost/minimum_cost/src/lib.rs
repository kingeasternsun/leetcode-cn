struct Solution;
use std::collections::{BinaryHeap, HashMap};

#[derive(Debug)]
struct Value(usize, usize);
impl PartialEq for Value {
    fn eq(&self, other: &Self) -> bool {
        self.1 == other.1
    }
}
impl Eq for Value {}

impl PartialOrd for Value {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Value {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other.1.cmp(&self.1)
    }
}

impl Solution {
    // https://www.bilibili.com/video/BV1es4y1w7An/
    /*
    1. 基于起始位置，终止位置，特殊路径的首尾位置， 创建为图中的点V
    2. 为这些点创建图中的边E
     */
    pub fn minimum_cost(start: Vec<i32>, target: Vec<i32>, special_roads: Vec<Vec<i32>>) -> i32 {
        let start = (start[0], start[1]);
        let target = (target[0], target[1]);
        // 构建图中点
        let mut pos_map: HashMap<(i32, i32), usize> = HashMap::new();
        pos_map.entry(start).or_insert(0);
        pos_map.entry(target).or_insert(1);

        special_roads.iter().for_each(|sr| {
            let len = pos_map.len();
            pos_map.entry((sr[0], sr[1])).or_insert(len);

            let len = pos_map.len();
            pos_map.entry((sr[2], sr[3])).or_insert(len);
        });

        // 构建图
        let mut graph = vec![vec![usize::MAX; pos_map.len()]; pos_map.len()];

        // 添加边
        pos_map.iter().for_each(|(pos, _)| {
            pos_map.iter().for_each(|(pos2, _)| {
                let from = pos_map.get(pos).unwrap().clone();
                let to = pos_map.get(pos2).unwrap().clone();

                graph[from][to] = ((pos.0 - pos2.0).abs() + (pos.1 - pos2.1).abs()) as usize;
            });
        });

        // 更新特殊边
        special_roads.iter().for_each(|sr| {
            let from = pos_map.get(&(sr[0], sr[1])).unwrap().clone();
            let to = pos_map.get(&(sr[2], sr[3])).unwrap().clone();

            graph[from][to] = graph[from][to].min(sr[4] as usize);
        });

        // Dijkstra
        let mut priority_queue = BinaryHeap::new();
        // the min distance from start to specified V
        let mut min_dist = vec![usize::MAX; pos_map.len()];
        min_dist[pos_map.get(&start).unwrap().clone()] = 0;
        let target_id = pos_map.get(&target).unwrap().clone();

        priority_queue.push(Value(0, 0));
        while let Some(Value(pos_id, pos_dist)) = priority_queue.pop() {
            if pos_dist > min_dist[pos_id] {
                continue;
            }

            if pos_id == target_id {
                return pos_dist as i32;
            }

            for (n, e) in graph[pos_id].iter().enumerate() {
                if pos_dist + e < min_dist[n] {
                    min_dist[n] = pos_dist + e;
                    priority_queue.push(Value(n, min_dist[n]));
                }
            }
        }

        -1
    }

    // fn create_e(x:i32, y:i32, map: std::collections::HashMap<(i32,i32), i32> ){
    //     len:=map.len();
    //     map.entry((x,y)).or_insert(len);
    // }
}

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
