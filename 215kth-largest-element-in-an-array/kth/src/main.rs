fn main() {
    println!("Hello, world!");
    Solution::find_kth_largest(vec![3, 2, 1, 5, 6, 4], 2);
    // Solution::find_kth_largest(vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4);
    // let mut a =vec![3, 2, 1, 5, 6, 4];
    // let a1 = & mut a[..];
    // a1.swap(0,0);

    // for i in 1..2{
    //     println!("{}",i);
    // }

    
}
pub struct Solution {}
use std::cmp::Reverse;
use std::collections::BinaryHeap;
impl Solution {
    pub fn find_kth_largest1(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        let end = nums.len();
        let exp_id = end-k as usize ;

        return Self::find_k(&mut nums[..], 0,  end-1,exp_id)
    }

    pub fn plus_one(x: Option<usize>)->usize{
        match x{
            None => 0,
            Some(i)=>i+1,
        }
    }

    pub fn partion(nums : &mut[i32], left :usize,end:usize)->usize{
        let x = nums[end];
        let mut i :Option<usize>= None;
        if left > 0{
            i = Some(left-1)
        }

        for j in left..end{

            if nums[j] <= x{
                let new_i = Self::plus_one(i);
                nums.swap(new_i, j);
                i = Some(new_i);
            }

        }

        let new_i = Self::plus_one(i);
        nums.swap(new_i, end);
        // println!("{:?} {:?}",nums,new_i);
        new_i
        
    }
    pub fn find_k(nums : &mut[i32], left :usize,end:usize, exp_id :usize)->i32{
        // println!("before {:?}  {:?} {:?} {:?}",nums,left,end,exp_id);
        let par_id = Self::partion(nums, left, end);
        // println!("{:?} {:?} {:?} {:?} {:?}",nums,par_id,left,end,exp_id);
        if par_id == exp_id{
            return nums[par_id]
        }else if par_id < exp_id{
            return Self::find_k(nums, par_id+1, end, exp_id)

        }else{
            return Self::find_k(nums, left, par_id-1, exp_id)
        }
    }
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        

        let k = k as usize;
        let mut heap = BinaryHeap::with_capacity(k);
        for num in nums{

            if heap.len() < k  {
                heap.push(Reverse(num));
            }else if heap.peek().unwrap().0 < num{
                heap.pop();
                heap.push(Reverse(num));
            }

        }
        heap.pop().unwrap().0
    }

}

// submission codes end

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_215() {
        assert_eq!(
            Solution::find_kth_largest(vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4),
            4
        );
        assert_eq!(Solution::find_kth_largest(vec![3, 2, 1, 5, 6, 4], 2), 5);
    }
}