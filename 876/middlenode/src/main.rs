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

pub struct Solution ;
impl Solution {
    pub fn middle_node(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {

        if head.is_none(){
            return head
        }

        let mut one = &head;//一次跳一步
        let mut two = one.as_ref().unwrap().next.as_ref();//技巧 一次跳两步

        while let Some(cur)=two{ //因为tow 肯定比one块，所以只需要判断two 就可以了
            one = &(one.as_ref().unwrap().next);
            two = cur.next.as_ref();
            if let Some(next)=two{
                two = next.next.as_ref()
            }

        }

        one.clone()
        
    }
}

fn main() {
    println!("Hello, world!");
}



