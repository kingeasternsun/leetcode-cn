use std::cell::RefCell;
use std::rc::Rc;

// 定义二叉树节点
#[derive(Debug)]
struct TreeNode {
    val: i32,
    left: Option<Rc<RefCell<TreeNode>>>,
    right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    fn new(val: i32) -> Rc<RefCell<TreeNode>> {
        Rc::new(RefCell::new(TreeNode {
            val,
            left: None,
            right: None,
        }))
    }
}

struct Solution;
impl Solution {
    // 24ms 6.7mb
    pub fn good_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut pre_max = i32::MIN;
        let mut cnt = 0;
        Self::dfs(root, pre_max, &mut cnt);
        cnt
    }

    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, mut pre_max: i32, cnt: &mut i32) {
        if let Some(cur) = root {
            if cur.borrow().val >= pre_max {
                *cnt = *cnt + 1;
                pre_max = cur.borrow().val;
            }

            Self::dfs(cur.borrow().left.clone(), pre_max, cnt);
            Self::dfs(cur.borrow().right.clone(), pre_max, cnt);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {}
}
