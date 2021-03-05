struct MyQueue {
    in_stack: Vec<i32>,
    out_stack: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyQueue {
    /** Initialize your data structure here. */
    fn new() -> Self {
        return MyQueue {
            in_stack: Vec::new(),
            out_stack: Vec::new(),
        };
    }

    /** Push element x to the back of queue. */
    fn push(& mut self, x: i32) {
        self.in_stack.push(x);
    }

    /** Removes the element from in front of queue and returns that element. */
    fn pop(& mut self) -> i32 {
        if let Some(x) = self.out_stack.pop() {
            return x;
        }

        while let Some(x) = self.in_stack.pop() {
            self.out_stack.push(x);
        }

        return self.out_stack.pop().unwrap_or(0);
    }

    /** Get the front element. */
    fn peek(& mut self) -> i32 {
        if let Some(x) = self.out_stack.last() {
            return *x;
        }

        while let Some(x) = self.in_stack.pop() {
            self.out_stack.push(x);
        }

        if let Some(x) = self.out_stack.last() {
            return *x;
        }
        return  0;
    }

    /** Returns whether the queue is empty. */
    fn empty(&self) -> bool {
        return self.in_stack.is_empty() && self.out_stack.is_empty();
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * let obj = MyQueue::new();
 * obj.push(x);
 * let ret_2: i32 = obj.pop();
 * let ret_3: i32 = obj.peek();
 * let ret_4: bool = obj.empty();
 */

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
