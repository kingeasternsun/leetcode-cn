struct Solution;
impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        let mut jin: i32 = 1;
        let mut ret = Vec::with_capacity(digits.len());
        for d in digits.into_iter().rev(){
            ret.push((d+jin)%10);
            jin = (d+jin)/10;
        }
        if jin >0{
            ret.push(jin);
        }

        let mut left = 0;
        let mut right = ret.len()-1;
        while left < right{
            ret.swap(left, right);
            left+=1;
            right-=1;
        }
        ret
    }

    pub fn plus_one1(digits: Vec<i32>) -> Vec<i32> {
        let mut jin: i32 = 1;
        let mut ret = Vec::with_capacity(digits.len());
        for d in digits.into_iter().rev(){
            ret.push((d+jin)%10);
            jin = (d+jin)/10;
        }
        if jin >0{
            ret.push(jin);
        }

        ret.reverse();
        ret
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::plus_one(vec![1,2,3]), vec![1,2,4]);
        assert_eq!(Solution::plus_one(vec![4,3,2,1]), vec![4,3,2,2]);
        assert_eq!(Solution::plus_one(vec![9]), vec![1,0]);
        assert_eq!(Solution::plus_one(vec![9,9]), vec![1,0,0]);
        assert_eq!(Solution::plus_one(vec![0]), vec![1]);
    }

    #[test]
    fn it_works1() {
        assert_eq!(Solution::plus_one1(vec![1,2,3]), vec![1,2,4]);
        assert_eq!(Solution::plus_one1(vec![4,3,2,1]), vec![4,3,2,2]);
        assert_eq!(Solution::plus_one1(vec![9]), vec![1,0]);
        assert_eq!(Solution::plus_one1(vec![9,9]), vec![1,0,0]);
        assert_eq!(Solution::plus_one1(vec![0]), vec![1]);
    }
}
