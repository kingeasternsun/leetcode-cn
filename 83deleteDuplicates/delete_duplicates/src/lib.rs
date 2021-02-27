/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-25 22:01:22
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-27 14:13:29
 * @FilePath: /83deleteDuplicates/delete_duplicates/src/lib.rs
 */
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}
impl Solution {
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    

    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
