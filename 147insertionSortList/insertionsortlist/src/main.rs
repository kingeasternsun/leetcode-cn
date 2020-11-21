fn main() {
    println!("Hello, world!");
}
pub struct Solution;
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


//  跟编译器斗了一天，最后放弃了，然后 参考了链接：https://leetcode-cn.com/problems/insertion-sort-list/solution/rustlian-biao-by-mengsuenyan-8/
// 这里面关键而且奇妙的地方 在于添加了一个额外的 anchor ， 然后 anchor ，node， head 就分开了，而且还能很好的处理比第一个节点要小的情况
impl Solution {

  pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut anchor = Box::new(ListNode::new(0));//这里面的val无所谓
    let mut cur = head;
    while let Some(mut node) = cur{

      // 每次都要取出当前节点，和下一个节点断开
      cur = node.next.take();

      let mut tail = & mut anchor;
      while tail.next.is_some() && tail.next.as_ref().unwrap().val < node.val{
        tail = tail.next.as_mut().unwrap();
      }


     let half2 = tail.next.take();
     node.next = half2;
     tail.next = Some(node);

      

    }


    anchor.next


}

}
