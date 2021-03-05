/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-05 10:14:33
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-05 10:32:09
 * @FilePath: \SwordToOffer009\CQueue\src\lib.rs
 */


struct CQueue {

    in_stack:Vec<i32>,
    out_stack:Vec<i32>,

}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl CQueue {

    fn new() -> Self {
        return CQueue{
            in_stack:Vec::new(),
            out_stack:Vec::new(),
        }
    }
    
    fn append_tail(& mut self, value: i32) {
        self.in_stack.push(value)
    }
    
    fn delete_head(&mut self) -> i32 {
        if let Some(x) = self.out_stack.pop(){
            return x;
        }

        while let Some(x) = self.in_stack.pop(){
            self.out_stack.push(x);
        }
       return  self.out_stack.pop().unwrap_or(-1);
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * let obj = CQueue::new();
 * obj.append_tail(value);
 * let ret_2: i32 = obj.delete_head();
 */
#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
