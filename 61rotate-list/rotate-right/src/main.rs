fn main() {
    // println!("Hello, world!");
    let a = Some(Box::new(ListNode {
        val: 1,
        next: Some(Box::new(ListNode {
            val: 2,
            next: Some(Box::new(ListNode {
                val: 3,
                next: Some(Box::new(ListNode {
                    val: 4,
                    next: Some(Box::new(ListNode { val: 5, next: None })),
                })),
            })),
        })),
    }));

    // scan(a.clone());

    // scan(&a);

    Solution::rotate_right(a.clone(), 2);

    // Solution::rotate_right(a, 2);
}

fn scan(head: &Option<Box<ListNode>>) {
    let mut node = head;
    while let Some(cur) = node {
        print!("{}", cur.val);
        node = &cur.next;
    }
    println!("----------");
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
    pub fn rotate_right1(mut head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        if k == 0 || head.is_none() {
            return head;
        }

        let mut ref_node = head.as_ref();
        let mut len = 0;
        while let Some(node) = ref_node {
            len += 1;
            ref_node = node.next.as_ref();
        }

        let k = len - (k % len); // 注意要先取余
        if len == 1 || k == len {
            return head;
        }

        // 找到分割的位置
        let mut ref_node = head.as_mut();
        let mut i = 0;
        let mut new_head: Option<Box<ListNode>> = None;

        while let Some(node) = ref_node {
            if i == k - 1 {
                new_head = node.next.take();
                break;
            } else {
                ref_node = node.next.as_mut();
            }
            i += 1;
        }

        let mut ref_node = new_head.as_mut();
        // 重新拼接两段链表
        while let Some(node) = ref_node {
            if node.next.is_none() {
                // 将前面一段链表接到后面的一段链表里
                node.next.replace(head.unwrap());
                // node.next =head;
                break;
            }
            ref_node = node.next.as_mut();
        }

        new_head
        //   Some(*remain.unwrap())
    }

    //TODO BUG 
    pub fn rotate_right(mut head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        if k == 0 || head.is_none() {
            return head;
        }

        let mut ref_node = head.as_ref();
        let mut i = 0;
        while let Some(cur) = ref_node {
            if i == k {
                break;
            }

            i += 1;
            ref_node = cur.next.as_ref();
        }

        // 链条长度 小于等于 k
        if ref_node.is_none() {
            //这里 i 就是链表的长度
            if i == 1 {
                return head;
            }

            let k = i % k;
            if k == 0 {
                return head;
            }


            //重新走一下
            ref_node = head.as_ref();
            while let Some(cur) = ref_node {
                if i == k {
                    break;
                }
    
                i += 1;
                ref_node = cur.next.as_ref();
            }

        }


        let head_ref =  head.clone();
        let mut head_next = head_ref.unwrap().next;
        let mut one = head_next.as_mut();
        // let mut one = head_next.as_mut();
        let mut ref_next = ref_node.unwrap().clone().next;
        let mut ref_node = ref_next.as_mut();
        let mut new_head: Option<Box<ListNode>> = None;

        while let Some(cur) = ref_node{

            if cur.next.is_none(){
                if let Some(cur1) = one{
                    new_head = cur1.next.take();
                }
                cur.next.replace(head.unwrap());
                break
            }else{
                
                if let Some(cur1) = one{
                    one = cur1.next.as_mut();
                }

                ref_node = cur.next.as_mut();
            }

        }



        new_head
    }
}
