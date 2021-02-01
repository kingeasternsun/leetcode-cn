pub struct Solution;
impl Solution {
    pub fn fair_candy_swap(a: Vec<i32>, b: Vec<i32>) -> Vec<i32> {

        let mut res = vec![0;2];
        let sum_a :i32  = a.iter().sum() ;
        let sum_b :i32  = b.iter().sum() ;

        let a_map :std::collections::HashSet<i32> = a.into_iter().collect();
        let delta = (sum_a-sum_b)/2;

        for b1 in &b{

            if a_map.get(&(b1+delta)).is_some(){
                res[0]=delta+b1;
                res[1]=*b1;
                return res 
            }
        }

        res

    }

}

#[cfg(test)]
mod tests {
    use crate::Solution;

    #[test]
    fn it_works() {
        assert_eq!(Solution::fair_candy_swap(vec![1,1], vec![2,2]),vec![1,2]);
        assert_eq!(Solution::fair_candy_swap(vec![2,4], vec![1,2,5]),vec![4,5]);
        assert_eq!(2 + 2, 4);
    }
}
