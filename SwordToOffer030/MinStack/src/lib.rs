/*
 * @Description: 
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-03-05 10:36:30
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-03-05 11:18:04
 * @FilePath: \SwordToOffer030\MinStack\src\lib.rs
 */


struct MinStack {
    stack:Vec<i32>,
    min_vec:Vec<i32>, //因为是从右边删除，所以要保持右边最小

}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {

    /** initialize your data structure here. */
    fn new() -> Self {
        return  MinStack{
            stack:Vec::new(),
            min_vec:Vec::new(),
        };
    }
    
    fn push(& mut self, x: i32) {
        self.stack.push(x);

        //维护右边最小
        if let Some(v) = self.min_vec.last(){
            if v >= &x{
                self.min_vec.push(x);
            }
        }else{
            self.min_vec.push(x);
        }
    }
    
    fn pop(& mut self) {
        if let Some(v) = self.stack.pop(){
            if &v <= self.min_vec.last().unwrap(){
                self.min_vec.pop();
            }
        }
        return 
        
    }
    
    fn top(&self) -> i32 {
        return *self.stack.last().unwrap();
    }
    
    fn min(&self) -> i32 {
        return *self.min_vec.last().unwrap();
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * let obj = MinStack::new();
 * obj.push(x);
 * obj.pop();
 * let ret_3: i32 = obj.top();
 * let ret_4: i32 = obj.min();
 */

#[cfg(test)]
mod tests {

    use crate::MinStack;
    #[test]
    fn it_works() {
        let mut obj = MinStack::new();
        obj.push(-2);
        obj.push(0);
        obj.push(-3);

        assert_eq!(obj.top(),-3);
        assert_eq!(obj.min(),-3);
        obj.pop();
        assert_eq!(obj.top(),0);
        assert_eq!(obj.min(),-2);
    }
}
