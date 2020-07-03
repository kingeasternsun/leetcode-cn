fn main() {
    println!("Hello, world!");
}
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

pub struct Solution {}
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
    //4ms
    pub fn remove_elements1( head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {

        let mut dummy = Some(Box::new(ListNode::new(0)));
        let mut next = dummy.as_mut();
        let mut cur = head;

        while let Some(mut inner) = cur{
            cur = inner.next.take();
            if inner.val != val{
                next.as_mut().unwrap().next = Some(inner);
                next = next.unwrap().next.as_mut();
            }

        }


        dummy.unwrap().next
    }

    //TODO 
    pub fn remove_elements( head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {

        // let mut dummy = Some(Box::new(ListNode::new(0)));
        // // let mut cur = dummy.as_mut();
        // // let t = cur.as_mut();
        // dummy.as_mut().as_mut().unwrap().next = head;
        let mut cur = Some(head.unwrap());
        // cur = cur.next;
        // // cur = cur.unwrap().next;
        // head.unwrap().next = cur;

        // head

        while let Some( inner) = cur{

            let next = inner.next;
            if inner.val != val{
                head.as_mut().unwrap().next = next;
            }
        }


        None
 
    }

    // 找到不等于val的下一个节点 
    pub fn remove( mut cur: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {

        while let Some( inner) = cur{
            cur = inner.next;
            if inner.val != val{
                return cur
            }
        }

        cur

    }

}