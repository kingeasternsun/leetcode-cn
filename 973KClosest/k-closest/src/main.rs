fn main() {
    println!("Hello, world!");
    assert_eq!(Solution::k_closest(vec![vec![1,3],vec![-2,2]], 1),vec![vec![-2,2]]);
    assert_eq!(Solution::k_closest(vec![vec![-2,2]], 1),vec![vec![-2,2]]);
    assert_eq!(
        Solution::k_closest(vec![vec![3, 3], vec![5, -1], vec![-2, 4]], 2),
        vec![vec![-2, 4], vec![3, 3]]
    );


}

pub struct Solution;

use std::cmp::Ordering;
// #[derive(Debug)]
pub struct Item(i32, i32);
impl Eq for Item {}
impl PartialEq for Item {
    fn eq(&self, other: &Self) -> bool {
        self.0 == other.0 && self.1 == other.1
    }
}

impl Ord for Item {
    fn cmp(&self, other: &Item) -> Ordering {
        let a = self.0 * self.0 + self.1 * self.1;
        let b = other.0 * other.0 + other.1 * other.1;
        a.cmp(&b)
    }
}

impl PartialOrd for Item {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

/*
使用最大堆，每次超过k个就把最大的弹出来，所以最终剩下的就是最小的
*/

impl Solution {
    pub fn k_closest1(points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let k = k as usize;
        let mut res = vec![];
        if points.len() == k {
            return points;
        }

        let mut min_heap = std::collections::BinaryHeap::new();
        for v in &points {
            min_heap.push(Item(v[0], v[1]));
            if min_heap.len() > k {
                min_heap.pop();
                // print!("{:?}",tmp);
            }
        }

        min_heap.into_iter().for_each(|x| res.push(vec![x.0, x.1]));

        res
    }

    //利用快排
    pub fn k_closest(mut points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let k = k as usize;
        if points.len() == k {
            return points;
        }

    
        Self::quick_sort(&mut points[0..], k);



       points.into_iter().take(k).collect()
    }

    //利用快排
    pub fn quick_sort( points: &mut [Vec<i32>], k: usize)-> usize  {

        let mut beg = 0;//[0 beg]里面的都是比points[0]小于或相等的
        let mut end = 1; //(beg ~]里面的都是比points[0]大的

        while end < points.len(){

            if points[end][0]*points[end][0] + points[end][1]*points[end][1] > points[0][0]*points[0][0] +points[0][1]*points[0][1]{
                end+=1;
                continue;
            }

            points.swap(beg+1, end);
            end +=1;
            beg +=1;

        }

        // print!("{:?}\n",points);
        points.swap(0, beg);


        if beg == k-1{
            return  beg
        }

        if beg > k-1{
            return Self::quick_sort(& mut points[0..beg], k);
        }

        return Self::quick_sort(& mut points[beg+1..], k-(beg+1));
        
    }
}
