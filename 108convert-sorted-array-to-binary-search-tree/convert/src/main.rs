fn main() {
    println!("Hello, world!");

    Solution::sorted_array_to_bst(vec![-10, -3, 0, 5, 9]);
}
pub struct Solution {}


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
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        let end = nums.len();
        // Self::help2(&nums[..],0,end-1)
        Self::help(&nums[..],0,end as i32 -1)
    }

    //4ms 
    pub fn help(nums: &[i32], beg:i32,end:i32)->Option<Rc<RefCell<TreeNode>>> {

        if nums.is_empty(){
            return None
        }

        if beg > end{
            return None
        }

        let mid = (beg+end)/2+1 ;

        // println!("{:?} {:?} {:?}",beg,end,nums[mid as usize]);


        Some(
            Rc::new(RefCell::new(
                TreeNode{
                    val:nums[mid as usize],
                    left:Self::help(nums,beg,mid-1),
                    right:Self::help(nums,mid+1,end),
                }
            ))
        )
    }

    // 0ms
    pub fn help2(nums: &[i32])->Option<Rc<RefCell<TreeNode>>> {

        if nums.is_empty(){
            return None
        }


        Some(
            Rc::new(RefCell::new(
                TreeNode{
                    val:nums[nums.len()/2],
                    left:Self::help2(&nums[..(nums.len()/2)]),
                    right:Self::help2(&nums[(nums.len()/2+1)..]),
                }
            ))
        )

     
    }

}