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

fn main() {
    let a =Some(
        Box::new(ListNode{
            val:1,
            next:Some(Box::new(ListNode{
                val:2,
                next:Some(Box::new(ListNode{
                    val:1,
                    next:None,
                }))
            })),
        })
    );
    println!("Hello, world!");
    print!("{}",Solution::is_palindrome(a));
}

pub struct Solution;

impl Solution {
    pub fn is_palindrome(head: Option<Box<ListNode>>) -> bool {

        if head.is_none()|| head.as_ref().unwrap().next.is_none(){
            return true
        }

        let (mut one,cnt) = Self::find_mid(head.as_ref());

        if cnt &1 == 1{
            one = one.unwrap().next.as_ref();
        }

        //此时 one所在就是mid
        let end = cnt/2-1;
        let (res,_) = Self::help(head.as_ref(),one,0,end);


        res
    }

    pub fn find_mid<'a>(head:Option< & 'a Box<ListNode>>)-> (Option< & 'a Box<ListNode>>,i32){
        //找到中间节点
        let mut one = head;
        let mut two = head.as_ref().unwrap().next.as_ref();
        let mut cnt = 1;


        //因为 two 肯定在one后面 所以只需要判断two的终止条件就可以了
        while let Some(t) = two{
            one = one.unwrap().next.as_ref();
            two = t.next.as_ref();
            cnt = cnt +1;
            if let Some(t2) = two{
                two = t2.next.as_ref();
                cnt = cnt +1;
            }
        }

        (one,cnt)
    }

    pub fn help<'a>(left :Option< & Box<ListNode>>,mid :Option< & 'a Box<ListNode>>,  pos : i32,end :i32)->(bool,Option< & 'a Box<ListNode>>){

        if pos == end{
            return (left.unwrap().val==mid.unwrap().val, mid.unwrap().next.as_ref())
        }


        //让left 通过不断的调用  进入到中间的节点 ，然后在出栈的时候返回响应右边的节点
        let (result,right) = Self::help(left.unwrap().next.as_ref(),mid, pos+1,end);
        if result==false{
            return (false,right)
        }

        //判断当前节点和 对称右边节点的节点 next 是否相同 ,同时把next靠右边的返回
        return (left.unwrap().val==right.unwrap().val,right.unwrap().next.as_ref())
    }
}




