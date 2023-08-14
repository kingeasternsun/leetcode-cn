use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

struct Solution;
impl Solution {
    pub fn merge_trees(
        root1: Option<Rc<RefCell<TreeNode>>>,
        root2: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        match (root1.clone(), root2.clone()) {
            (Some(r1), Some(r2)) => Some(Rc::new(RefCell::new(TreeNode {
                val: r1.borrow().val + r2.borrow().val,
                left: Self::merge_trees(r1.borrow_mut().left.take(), r2.borrow_mut().left.take()),
                right: Self::merge_trees(
                    r1.borrow_mut().right.take(),
                    r2.borrow_mut().right.take(),
                ),
            }))),
            _ => root1.or(root2),
        }
    }
    pub fn merge_trees0(
        root1: Option<Rc<RefCell<TreeNode>>>,
        root2: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        match (root1.clone(), root2.clone()) {
            (Some(r1), Some(r2)) => {
                let r1_child_left = r1.borrow_mut().left.take();
                let r1_child_right = r1.borrow_mut().right.take();

                let r2_child_left = r2.borrow_mut().left.take();
                let r2_child_right = r2.borrow_mut().right.take();

                Some(Rc::new(RefCell::new(TreeNode {
                    val: r1.borrow().val + r2.borrow().val,
                    left: Self::merge_trees(r1_child_left, r2_child_left),
                    right: Self::merge_trees(r1_child_right, r2_child_right),
                })))
            }
            _ => root1.or(root2),
        }
    }

    pub fn merge_trees1(
        root1: Option<Rc<RefCell<TreeNode>>>,
        root2: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        match (root1.clone(), root2.clone()) {
            (Some(r1), Some(r2)) => {
                let r1_child = (r1.borrow_mut().left.take(), r1.borrow_mut().right.take());
                let r2_child = (r2.borrow_mut().left.take(), r2.borrow_mut().right.take());

                Some(Rc::new(RefCell::new(TreeNode {
                    val: r1.borrow().val + r2.borrow().val,
                    left: Self::merge_trees(r1_child.0, r2_child.0),
                    right: Self::merge_trees(r1_child.1, r2_child.1),
                })))
            }
            _ => root1.or(root2),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use leetcode_prelude::btree;

    #[test]
    fn it_works() {
        let btree = btree![1, 2, 2, null, null, 3, 3];
    }
}
