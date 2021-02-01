pub struct Solution;
impl Solution {
    pub fn num_identical_pairs1(nums: Vec<i32>) -> i32 {
        let mut num_set = std::collections::HashMap::with_capacity(nums.len());
        for i in nums{
            let counter = num_set.entry(i).or_insert(0);
            *counter += 1;
        }

        num_set.into_iter().map(|(_k,v)|{v*(v-1)/2}).sum()
        
    }

    pub fn num_identical_pairs(nums: Vec<i32>) -> i32 {
        let mut num_set = std::collections::HashMap::with_capacity(nums.len());
        nums.iter().for_each(|i|{
            let counter = num_set.entry(i).or_insert(0);
            *counter += 1;
        });

        num_set.into_iter().map(|(_k,v)|{v*(v-1)/2}).sum()
        
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
