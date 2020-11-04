fn main() {
    println!("Hello, world!");
    assert_eq!(Solution::furthest_building(vec![4,2,7,6,9,14,12], 5, 1),4);
    assert_eq!(Solution::furthest_building(vec![4,12,2,7,3,18,20,3,19], 10, 2),7);
    assert_eq!(Solution::furthest_building(vec![14,3,19,3], 17, 0),3);
}

pub struct Solution;
impl Solution {
    pub fn furthest_building(heights: Vec<i32>, bricks: i32, ladders: i32) -> i32 {
        if heights.len() <= 1 {
            return 0;
        }

        let mut bricks = bricks;
        let mut min_heap = std::collections::BinaryHeap::new();

        for i in 1..heights.len(){

            let dif = heights[i]-heights[i-1];
            if dif <=0{
                continue
            }
            min_heap.push(std::cmp::Reverse(dif));
            //梯子数量不够了,把之前差值最小的换为砖块
            if min_heap.len() as i32 > ladders{
                if let Some(std::cmp::Reverse(mid_dif)) = min_heap.pop(){

                    //到目前位置 梯子都用于最高的几次跳跃了，砖块却不够
                    if mid_dif>bricks{
                        return i as i32 -1
                    }

                    bricks = bricks - mid_dif;

                } 
            }
        }

        heights.len() as i32 - 1
    }
}
