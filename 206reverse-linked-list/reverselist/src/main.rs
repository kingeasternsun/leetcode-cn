fn main() {
    println!("Hello, world!");
    let b = Some(Box::new(
        ListNode{
            val:1,
            next:Some(Box::new(
               ListNode{
                val:2,
                next:None,
               }
            )),
        },
    ));

    let c = Solution_recur1::reverse_list(b);
    println!("{:?}",c);
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
pub struct Solution_recur1;
impl Solution_recur1 {
    //递归1
    pub fn recur(cur: Option<Box<ListNode>>, pre: Box<ListNode>) -> Option<Box<ListNode>> {
        let mut cur = cur.unwrap();
        if cur.next.is_none() {
            cur.next.replace(pre);//这里忘记 坑了好久
            return Some(cur);
        }

        let mut next = cur.next.replace(pre); //当前节点指向前一个节点，同时返回当前节点旧的下一个节点

        let end = Self::recur(next.take(), cur);
        return end;
    }

    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }

        let mut head = head.unwrap();
        if head.next.is_none() {
            return Some(head);
        }

        let next = head.next.take();
        return Self::recur(next, head);
    }
      
}


// // 递归
// impl Solution {

//     //递归2
//     pub fn recur2<'a>(
//         cur: Option<&'a Box<ListNode>>,
//         // pre: Option<&'a Box<ListNode>>,
//     ) -> Option<&'a Box<ListNode>> { //返回最右边的节点
//         // let mut cur = cur.unwrap();
//         if cur.unwrap().next.is_none() {
//             return cur;
//         }

//         // let mut next = cur.next.replace(pre); //当前节点指向前一个节点，同时返回当前节点旧的下一个节点
//         let next = cur.unwrap().next.as_ref();

//         let end = Self::recur2(next);
//         next.unwrap().next.replace(*cur.unwrap());
//         return end;
//     }

//     pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
//         if head.is_none() {
//             return None;
//         }

//         return Some(*Self::recur2(head.as_ref()).unwrap());
//     }
      
// }






 // 迭代 利用replace的 最简洁的写法
pub struct Solution_iter;
impl Solution_iter {
    pub fn Solution_iter(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {

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

//迭代例外一种
pub struct Solution_iter1;
impl Solution_iter1 {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {

        let mut pre = None ;
        let mut cur = head;
        while let Some(mut node) = cur {

            if pre.is_none(){
                cur = node.next.take();
            }else{
                cur = node.next.replace(pre.unwrap());
            }

            pre = Some(node); //将当前节点变为前一个节点
        }

        pre
    }
}


