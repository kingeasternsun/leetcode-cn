fn main() {
    println!("Hello, world!");
    Solution::top_k_frequent(vec![1,1,1,2,2,4], 2);
}

pub struct Solution {}


use std::cmp::Ordering;
#[derive(Debug)]
pub struct Item{
    num :i32,
    cnt :i32,
}
impl Eq for Item {}
impl PartialEq for Item {
    fn eq(&self, other: &Self) -> bool {
        self.cnt == other.cnt
    }
}

impl PartialOrd for Item {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        other.cnt.partial_cmp(&self.cnt)
    }
}

impl Ord for Item {
    fn cmp(&self, other: &Item) -> Ordering {
        // self.cnt.cmp(&other.cnt)
        other.cnt.cmp(&self.cnt)
    }
}

// use std::cmp::Reverse;
use std::collections::HashMap;
use std::collections::BinaryHeap;

impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut min_heap = BinaryHeap::new();

        let mut res = Vec::new();
        let mut map = HashMap::new();
        for x in nums{
            if let Some(cnt) = map.get(&x){
                let cnt  = *cnt+1;
                map.insert(x, cnt);
            }else{
                map.insert(x, 1);
            }
        }
        println!("{:?}",map);

        for (key, val) in map {
            // println!("key: {} val: {}", key, val);
            if min_heap.len()< k as usize{
                min_heap.push(Item{num:key,cnt:val});
            }else{

                if val > min_heap.peek().unwrap().cnt{
                    min_heap.pop();
                    min_heap.push(Item{num:key,cnt:val});
                }

            }

        }


        for key in min_heap{
            res.push(key.num);
        }


        res
    }
}


// submission codes end

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_215() {
        assert_eq!(
            Solution::top_k_frequent(vec![1,1,1,2,2,4], 2),
            vec![2,1]
        );
        assert_eq!(Solution::top_k_frequent(vec![1], 1), vec![1]);
    }
}