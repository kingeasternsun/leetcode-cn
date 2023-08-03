struct Solution;
use std::collections::HashMap;
impl Solution {
    // https://www.bilibili.com/video/BV1es4y1w7An/
    /*
    1. 基于起始位置，终止位置，特殊路径的首尾位置， 创建为图中的点V
    2. 为这些点创建图中的边E
     */
    pub fn minimum_cost(start: Vec<i32>, target: Vec<i32>, special_roads: Vec<Vec<i32>>) -> i32 {
        // 构建图中点
        let mut pos_map: HashMap<(i32, i32), i32> = HashMap::new();
        pos_map.entry((start[0], start[1])).or_insert(0);
        pos_map.entry((target[0], target[1])).or_insert(1);

        special_roads.iter().for_each(|sr| {
            let len = pos_map.len();
            pos_map.entry((sr[0], sr[1])).or_insert(len as i32);

            let len = pos_map.len();
            pos_map.entry((sr[2], sr[3])).or_insert(len as i32);
        });

        // 构建图
        let mut graph = vec![vec![2*pos_map.len(); pos_map.len()]; pos_map.len()];

        // 添加边
        pos_map.iter().for_each(|(pos, _)| {
            pos_map.iter().for_each(|(pos2, _)| {
                let from = pos_map.get(pos).unwrap().clone() as usize;
                let to = pos_map.get(pos2).unwrap().clone() as usize;

                graph[from][to] = ((pos.0 - pos2.0).abs() + (pos.1 - pos2.1).abs()) as usize;
            });
        });

        // 特殊边
        special_roads.iter().for_each(|sr| {
            let from = pos_map.get(&(sr[0], sr[1])).unwrap().clone() as usize;
            let to = pos_map.get(&(sr[2], sr[3])).unwrap().clone() as usize;

            graph[from][to] = graph[from][to].min(sr[4] as usize);
        });

        0
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
