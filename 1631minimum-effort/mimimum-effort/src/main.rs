fn main() {
    // println!("Hello, world!");
    let heights1 = vec![vec![1, 2, 2], vec![3, 8, 2], vec![5, 3, 5]];
    print!("{}\n", Solution::minimum_effort_path(heights1));

    let heights2 = vec![vec![1, 2, 3], vec![3, 8, 4], vec![5, 3, 5]];
    print!("{}\n", Solution::minimum_effort_path(heights2));

    let heights3 = vec![
        vec![1, 2, 1, 1, 1],
        vec![1, 2, 1, 2, 1],
        vec![1, 2, 1, 2, 1],
        vec![1, 2, 1, 2, 1],
        vec![1, 1, 1, 2, 1],
    ];
    print!("{}\n", Solution::minimum_effort_path(heights3));

    let heights4 = vec![vec![1, 10, 6, 7, 9, 10, 4, 9]];
    print!("{}\n", Solution::minimum_effort_path(heights4));
    let heights5 = vec![vec![1, 1, 9, 9]];
    print!("{}\n", Solution::minimum_effort_path(heights5));
}

/*
LeetCode 1102. 得分最高的路径（优先队列BFS/极大极小化 二分查找）
LeetCode 410. 分割数组的最大值（极小极大化 二分查找）
LeetCode 774. 最小化去加油站的最大距离（极小极大化 二分查找）
LeetCode 875. 爱吃香蕉的珂珂（二分查找）
LeetCode LCP 12. 小张刷题计划（二分查找）
LeetCode 1011. 在 D 天内送达包裹的能力（二分查找）
LeetCode 1062. 最长重复子串（二分查找）
LeetCode 5438. 制作 m 束花所需的最少天数（二分查找）
LeetCode 5489. 两球之间的磁力（极小极大化 二分查找）

作者：kobe24o
链接：https://leetcode-cn.com/problems/path-with-minimum-effort/solution/dfs-er-fen-cha-zhao-mo-ban-ti-by-kobe24o/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
pub struct Solution;

use std::cmp::Ordering;
use std::collections::BinaryHeap;
#[derive(Debug)]
pub struct Item(usize, usize, i32);

impl Eq for Item {}
impl PartialEq for Item {
    fn eq(&self, other: &Self) -> bool {
        self.2 == other.2
    }
}

impl PartialOrd for Item {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        other.2.partial_cmp(&self.2)
    }
}

impl Ord for Item {
    fn cmp(&self, other: &Item) -> Ordering {
        // self.cnt.cmp(&other.cnt)
        other.2.cmp(&self.2)
    }
}
impl Solution {
    //timeout
    pub fn minimum_effort_path_timeout(heights: Vec<Vec<i32>>) -> i32 {
        let rows = heights.len();
        if rows == 0 {
            return 0;
        }
        let cols = heights[0].len();
        if cols == 0 {
            return 0;
        }

        //记录到各个点的体力
        let mut efforts = vec![vec![std::i32::MAX; cols]; rows];

        let mut stack = Vec::new();
        stack.push((0, 0));
        efforts[0][0] = 0;

        let dir: Vec<i32> = vec![0, 1, 0, -1, 0];
        while let Some((x, y)) = stack.pop() {
            //当前位置

            //对每一方向
            dir.windows(2).for_each(|dir| {
                let dx = dir[0];
                let dy = dir[1];

                if let Some(row) = efforts.get(x + dx as usize) {
                    if let Some(v) = row.get(y + dy as usize) {
                        //当前体力 和 从当前到(x+dy,y+dy)高度差的绝对值
                        let tmph = heights[x][y] - heights[x + dx as usize][y + dy as usize];
                        let new_effort = std::cmp::max(efforts[x][y], tmph.abs());

                        if new_effort < *v {
                            // efforts[x + dx as usize][y + dy as usize] = new_effort;
                            *efforts
                                .get_mut(x + dx as usize)
                                .unwrap()
                                .get_mut(y + dy as usize)
                                .unwrap() = new_effort;
                            stack.push((x + dx as usize, y + dy as usize));
                        }
                    }
                }
            });
        }

        efforts[rows - 1][cols - 1]
    }

    //32ms 2.3MB
    pub fn minimum_effort_path_dijkstra(heights: Vec<Vec<i32>>) -> i32 {
        let rows = heights.len();
        if rows == 0 {
            return 0;
        }
        let cols = heights[0].len();
        if cols == 0 {
            return 0;
        }

        //记录到各个点的体力
        let mut efforts = vec![vec![std::i32::MAX; cols]; rows];
        let mut seen = vec![vec![0; cols]; rows];

        let mut heap = BinaryHeap::new();
        let dir: Vec<i32> = vec![0, 1, 0, -1, 0];

        heap.push(Item(0, 0, 0));
        efforts[0][0] = 0;

        while let Some(cur) = heap.pop() {
            if cur.0 == rows - 1 && cur.1 == cols - 1 {
                return cur.2;
            }

            //已经处理过了. 因为相同点可能会加入多次
            if 1 == seen[cur.0][cur.1] {
                continue;
            }
            seen[cur.0][cur.1] = 1;
            //对每一方向
            dir.windows(2).for_each(|dir| {
                let dx = dir[0];
                let dy = dir[1];

                if let Some(row) = efforts.get(cur.0 + dx as usize) {
                    if let Some(v) = row.get(cur.1 + dy as usize) {
                        //当前体力 和 从当前到(x+dy,y+dy)高度差的绝对值
                        let tmph = heights[cur.0][cur.1]
                            - heights[cur.0 + dx as usize][cur.1 + dy as usize];
                        let new_effort = std::cmp::max(cur.2, tmph.abs());

                        if new_effort < *v {
                            // efforts[x + dx as usize][y + dy as usize] = new_effort;
                            *efforts
                                .get_mut(cur.0 + dx as usize)
                                .unwrap()
                                .get_mut(cur.1 + dy as usize)
                                .unwrap() = new_effort;

                            heap.push(Item(cur.0 + dx as usize, cur.1 + dy as usize, new_effort));
                        }
                    }
                }
            });
        }

        0
    }

    /*
    1. 记录节点之间的长度，
    2. 从小到大排序
    3. 依次插入
    52ms 3MB
    */
    pub fn minimum_effort_path_union(heights: Vec<Vec<i32>>) -> i32 {
        let rows = heights.len();
        if rows == 0 {
            return 0;
        }
        let cols = heights[0].len();
        if cols == 0 {
            return 0;
        }

        let mut dh = Vec::new();
        for x in 0..rows {
            for y in 0..cols {
                if let Some(row) = heights.get(x + 1) {
                    if let Some(left) = row.get(y) {
                        dh.push(Item(
                            x * cols + y,
                            (x + 1) * cols + y,
                            (heights[x][y] - left).abs(),
                        ));
                    }
                }

                if let Some(row) = heights.get(x) {
                    if let Some(down) = row.get(y + 1) {
                        dh.push(Item(
                            x * cols + y,
                            x * cols + y + 1,
                            (heights[x][y] - down).abs(),
                        ));
                    }
                }
            }
        }
        // print!("{:?}\n", dh);

        dh.sort_by(|a, b| a.2.cmp(&b.2));
        // print!("{:?}\n", dh);

        let mut seen = vec![0; rows * cols]; //记录每个点的上级点
        for i in 0..rows * cols {
            seen[i] = i;
        }

        // let mut res = vec![0; rows * cols]; //记录每个集合当前的最大数值

        //从小到大依次加进去
        for edge in &dh {
            //如果两个点都已经加入过路径了
            let p0 = Self::getparent(&seen, edge.0);
            let p1 = Self::getparent(&seen, edge.1);
            if p0 == p1 {
                continue;
            }

            //p1 链接到 p0上
            seen[p1] = p0;

            //加入当前变后，首尾相同了，因为边时从小到达插入的，所以当前的肯定就是最大值
            if Self::getparent(&seen, 0) == Self::getparent(&seen, rows * cols - 1) {
                return edge.2;
            }
        }

        0
    }

    pub fn getparent(parent: &Vec<usize>, x: usize) -> usize {
        let mut x = x;
        while x != parent[x] {
            x = parent[x];
        }
        return x;
    }

    //二分法，转为求一个数值 threshold，判断从[0,0] 到 [m-1,n-1] 是否存在一条路径上的体力都小于这个threshold 
    // 88ms 2.3MB
    pub fn minimum_effort_path(heights: Vec<Vec<i32>>) -> i32 {
        let rows = heights.len();
        if rows == 0 {
            return 0;
        }
        let cols = heights[0].len();
        if cols == 0 {
            return 0;
        }

        let mut left = 0;
        let mut right = Self::get_max_effort(&heights);

        while left < right{

            let mid = left + (right-left)/2;
            // print!("{:?} ",mid);
            if Self::exist_path(&heights, mid, rows, cols){
                right = mid;
            }else{
                left = mid+1;
            }
        }

        left
    }

    //dfs，判断从[0,0] 到 [m-1,n-1] 是否存在一条路径上的体力都小于这个threshold
    fn exist_path(heights: & Vec<Vec<i32>>, threashold: i32, rows: usize, cols: usize) -> bool {
        let mut seen = vec![vec![false; cols]; rows];
        let mut queue = std::collections::VecDeque::new();
        let dir: Vec<i32> = vec![0, 1, 0, -1, 0];

        queue.push_back((0, 0));
        seen[0][0] = true;

        while let Some((x, y)) = queue.pop_front() {
            if x == rows - 1 && y == cols - 1 {
                return true;
            }

            //对于每一个方向,是否在满足体力的条件下跳过去
            dir.windows(2).for_each(|dir| {
                let new_x = (x as i32 + dir[0]) as usize;
                let new_y = (y as i32 + dir[1]) as usize;
                if let Some(row) = heights.get(new_x) {
                    if let Some(next_height) = row.get(new_y) {
                        //体力超过满足条件 并且没有访问过
                        if (heights[x][y] - next_height).abs() <= threashold
                            && (!seen[new_x][new_y])
                        {
                            queue.push_back((new_x, new_y));
                            seen[new_x][new_y] = true;
                        }
                    }
                }
            });
        }

        false
    }

    //获取最大的体力
    fn get_max_effort(heights:& Vec<Vec<i32>>) -> i32 {
        let rows = heights.len();
        if rows == 0 {
            return 0;
        }
        let cols = heights[0].len();
        if cols == 0 {
            return 0;
        }

        let mut max_effort = 0;
        for x in 0..rows {
            for y in 0..cols {
                if let Some(row) = heights.get(x + 1) {
                    if let Some(left) = row.get(y) {
                        max_effort = std::cmp::max(max_effort, (heights[x][y] - left).abs())
                    }
                }

                if let Some(row) = heights.get(x) {
                    if let Some(down) = row.get(y + 1) {
                        max_effort = std::cmp::max(max_effort, (heights[x][y] - down).abs())
                    }
                }
            }
        }

        max_effort
    }
}
