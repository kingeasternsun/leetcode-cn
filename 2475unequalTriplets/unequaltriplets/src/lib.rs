
pub fn unequal_triplets(nums: Vec<i32>) -> i32 {
    let n = nums.len();
    if n < 3 {
        return 0;
    }

    let mut kv = std::collections::HashMap::new();
    for n in nums {
        let count = kv.entry(n).or_insert(0);
        *count += 1;
    }

    
    let mut left = 0;
    let mut ret  = 0;
    for (_,v) in kv{
        ret  = ret + left * v * (n - left - v);
        left += v;
    }

    ret as i32

}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(unequal_triplets(vec![4,4,2,4,3]), 3);
        assert_eq!(unequal_triplets(vec![1,1,1]), 0);
    }
}
