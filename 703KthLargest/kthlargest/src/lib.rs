/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-13 09:22:07
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-13 23:52:54
 * @FilePath: /kthlargest/src/lib.rs
 */



/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
 use std::collections::BinaryHeap;
 use std::cmp::Reverse;
 
 struct KthLargest {
     min_heap: BinaryHeap<Reverse<i32>>,
     min_heap_len: usize,
 }
 
 impl KthLargest {
     fn new(k: i32, mut nums: Vec<i32>) -> Self {
         let mut r = KthLargest {
             min_heap: BinaryHeap::new(),
             min_heap_len: k as usize,
         };
         nums.into_iter().for_each(|n| { r.add(n); });
         r
     }
     
     pub fn add(&mut self, val: i32) -> i32 {
         if self.min_heap.len() < self.min_heap_len {
             self.min_heap.push(Reverse(val));
         } else if self.min_heap.peek().unwrap().0 < val {
             self.min_heap.pop();
             self.min_heap.push(Reverse(val));
         }
         self.min_heap.peek().unwrap().0
     }
 }
 

/**
 * Your KthLargest object will be instantiated and called as such:
 * let obj = KthLargest::new(k, nums);
 * let ret_1: i32 = obj.add(val);
 */
#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
