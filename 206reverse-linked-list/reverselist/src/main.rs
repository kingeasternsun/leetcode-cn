fn main() {
    println!("Hello, world!");
}

pub struct Solution;

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

// 递归
impl Solution {


    pub fn iterator( cur:  Option<Box<ListNode>>, pre:Box<ListNode>) -> Option<Box<ListNode>>  {

        let mut cur = cur.unwrap();
        if cur.next.is_none(){
            return Some(cur)
        }
        
        let mut next = cur.next.replace(pre);//当前节点指向前一个节点，同时返回当前节点旧的下一个节点

        let end = Self::iterator(next.take(), cur);
        return end

    }

    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {

        if head.is_none(){
            return None
        }

        let mut head = head.unwrap();
        if head.next.is_none(){
            return Some(head)
        }

        return Self::iterator(head.next.take(), head)

    }
}




/* 迭代1
impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut pre = None;
        let mut cur = head;
        while let Some(mut node) = cur {
            cur = node.next.take(); //先将下一个节点保存起来
            node.next = pre; //将当前节点指向前一个节点
            pre = Some(node); //将当前节点变为前一个节点
        }

        pre
    }
}


 // 迭代2
impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {

        if head.is_none(){
            return None
        }
    

        let mut pre = head.unwrap() ;
        let mut cur = pre.next.take();
        while let Some(mut node) = cur {
            cur = node.next.replace(pre); //先将下一个节点指向前一个节点，同事保存之前的下一节点
            pre = node; //将当前节点变为前一个节点
        }

        Some(pre)
    }
}

*/