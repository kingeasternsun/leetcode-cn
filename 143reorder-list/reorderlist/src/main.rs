fn main() {
    println!("Hello, world!");
    let a = Box::new(ListNode { val: 1, next: None });
    let mut b = a.clone();
    b.as_mut().val = 3;
    print!("{:?}", a);
    print!("{:?}", b);
}
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

pub struct Solution;

impl Solution {
    pub fn reorder_list(head: &mut Option<Box<ListNode>>) {
        let mut list = Vec::new();
        let mut beg = head.clone();
        // let mut i = 0;
        while let Some(mut cur) = beg {
            let next = cur.next.take();
            list.push(cur);
            beg = next;
            // i = i + 1;
        }

        let cnt = list.len();
        if cnt < 3 {
            //少于3个不需要进行调整
            return;
        }

        
        let mut end = None;
        if cnt & 1 == 1 {
            end = Some(list[cnt / 2].clone());
        }

        // let mut beg = None;
        for i in (cnt + 1) / 2..cnt {
            let mut one = list[cnt - 1 - i].clone();
            let mut two = list[i].clone();

            if end.is_some() {
                two.next.replace(end.unwrap());
            }
            one.next.replace(two);

            end = Some(one);
        }
        head.replace(end.unwrap());
        
    }
}
