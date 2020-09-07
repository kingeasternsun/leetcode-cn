fn main() {
    println!("Hello, world!");
}


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
// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
  pub val: i32,
  pub left: Option<Rc<RefCell<TreeNode>>>,
  pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    TreeNode {
      val,
      left: None,
      right: None
    }
  }
}

pub struct Solution;

use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn sorted_list_to_bst(head: Option<Box<ListNode>>) -> Option<Rc<RefCell<TreeNode>>> {

        let mut vec = Vec::new();
        let mut cur = head;
        
        while let Some(node) = cur{
            vec.push(node.val);
            cur = node.next;
        }

        return Self::helper(&vec[..])
    }

    pub fn helper(num : &[i32])-> Option<Rc<RefCell<TreeNode>>>{

        if num.is_empty(){
            return None;
        }

        Some(
            Rc::new(RefCell::new(TreeNode{
                val:num[num.len()/2],
                left:Self::helper(&num[0..num.len()/2]),
                right:Self::helper(&num[num.len()/2+1..]),
            }))
        )
    }
}