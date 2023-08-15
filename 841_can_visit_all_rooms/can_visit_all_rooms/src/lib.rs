struct Solution;
impl Solution {
    // BFS 把每个房间当做树中节点
    pub fn can_visit_all_rooms(rooms: Vec<Vec<i32>>) -> bool {
        let mut queue = std::collections::VecDeque::new();
        let mut room_map = vec![false; rooms.len()];
        queue.push_back(0);
        room_map[0] = true;
        while let Some(cur) = queue.pop_front() {
            rooms[cur].iter().for_each(|&x| {
                let x = x as usize;
                // 还没有进入过这个房间
                if !room_map[x] {
                    room_map[x] = true;
                    queue.push_back(x);
                }
            });
        }

        room_map.into_iter().all(|x| x)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::can_visit_all_rooms(vec![vec![1], vec![2], vec![3], vec![]]),
            true
        );
        assert_eq!(
            Solution::can_visit_all_rooms(vec![vec![1, 3], vec![3, 0, 1], vec![2], vec![0]]),
            false
        );
    }
}
