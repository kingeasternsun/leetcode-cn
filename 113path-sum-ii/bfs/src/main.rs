fn main() {
    println!("Hello, world!");
    let mut a =vec![1,2];
    path_sum(& mut a);
    println!("{:?}",a);

    // let mut res = Vec::new();
    // res.push(a);
    // a.push(4);
    // res.push(a);
    // println!("{:?}",a);


}


pub fn path_sum(res: & mut Vec<i32>){

    res.push(3);
}

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

pub struct Solution {}

use std::cell::RefCell;
use std::collections::VecDeque;
use std::rc::Rc;
impl Solution {
    pub fn path_sum1(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        if root.is_none() {
            return res;
        }
        let mut deq = VecDeque::new();
        deq.push_back((0, Vec::new(), root.clone()));
        while !deq.is_empty() {
            if let Some((acc, mut vec, Some(node))) = deq.pop_front() {
                let acc = acc + node.borrow().val;
                vec.push(node.borrow().val);
                if node.borrow().left.is_none() && node.borrow().right.is_none() {
                    if acc == sum {
                        res.push(vec);
                    }
                } else {
                    if node.borrow().left.is_some() {
                        deq.push_back((acc, vec.clone(), node.borrow().left.clone()));
                    }
                    if node.borrow().right.is_some() {
                        deq.push_back((acc, vec.clone(), node.borrow().right.clone()));
                    }
                }
            }
        }
        res
    }

    //0ms 2.5MB
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> Vec<Vec<i32>> {

        let mut res = Vec::new();
        if root.is_none() {
            return res;
        }

        Self::dfs_sum(&root.unwrap(), sum,& mut Vec::new(),& mut res);

        res

    }

    pub fn dfs_sum(root: &Rc<RefCell<TreeNode>>,sum :i32, vec:&mut Vec<i32>, res: & mut Vec<Vec<i32>>){

        vec.push(root.borrow().val);

        if root.borrow().left.is_none() &&  root.borrow().right.is_none(){

            if root.borrow().val == sum{
                
                res.push(vec.clone());
        
            }

            vec.pop(); //记得这里要弹出
            return 
        }

        if root.borrow().left.is_some(){
            Self::dfs_sum(root.borrow().left.as_ref().unwrap(), sum-root.borrow().val, vec, res);
        }
        
        if root.borrow().right.is_some(){
            Self::dfs_sum(root.borrow().right.as_ref().unwrap(), sum-root.borrow().val, vec, res);
        }

        vec.pop();//记得这里要弹出

    }


}

// submission codes end
