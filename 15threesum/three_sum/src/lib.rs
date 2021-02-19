/*
 * @Description: 15
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-19 16:57:35
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-19 18:04:15
 * @FilePath: \15threesum\three_sum\src\lib.rs
 */
pub struct Solution;

impl Solution {
    pub fn three_sum_hash(nums: Vec<i32>) -> Vec<Vec<i32>> {

        let mut res = Vec::new();
        let mut num_cnt = std::collections::HashMap::new();
        for i in &nums {
            let e = num_cnt.entry(i).or_insert(0);
            *e += 1;
        }

        let mut nums2 = nums.clone();
        nums2.sort_unstable();

        for i in 0..nums2.len() {
            if i > 0 && nums2[i] == nums2[i - 1] {
                continue;
            }

            for j in i + 1..nums2.len() {
                if j > i + 1 && nums2[j] == nums2[j - 1] {
                    continue;
                }

                let third = 0 - nums2[i] - nums2[j];
                if third< nums2[j]{
                    continue;
                }
 

                if let Some(v) = num_cnt.get(&third) {

                    let mut cnt = 1;
                    if nums2[i] == third {
                        cnt += 1;
                    }
                    if nums2[j] == third {
                        cnt += 1;
                    }
                    if *v >= cnt {
                        res.push(vec![nums2[i], nums2[j], third])
                    };
                }
            }
        }

        res
    }

    /*
     双指针
    */
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {

        let mut res = Vec::new();
        let mut nums = nums;
        nums.sort_unstable();

        for i in 0..nums.len() {
            if i > 0 && nums[i] == nums[i - 1] { //去重
                continue;
            }

            //开始滑窗查询 
            let mut beg = i+1;
            let mut end = nums.len()-1;

            while beg < end{

                if nums[beg]+nums[end] < -nums[i]{
                    beg+=1;
                    continue;
                }

                if nums[beg]+nums[end] > -nums[i]{
                    end-=1;
                    continue;
                }


                res.push(vec![nums[i],nums[beg],nums[end]]);
                while beg<end && nums[beg]==nums[beg+1]{ //去重
                    beg+=1
                }
                beg+=1;

                while beg<end && nums[end]==nums[end-1]{ //去重
                    end-=1;
                }
                end-=1;

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
        assert_eq!(Solution::three_sum_hash(vec![-1,0,1,2,-1,-4]), vec![vec![-1,-1,2],vec![-1,0,1]]);
        assert_eq!(Solution::three_sum_hash(vec![0,0,0,0]), vec![vec![0,0,0]]);
        assert_eq!(Solution::three_sum_hash(vec![0,0,0]), vec![vec![0,0,0]]);

        assert_eq!(Solution::three_sum(vec![-1,0,1,2,-1,-4]), vec![vec![-1,-1,2],vec![-1,0,1]]);
        assert_eq!(Solution::three_sum(vec![0,0,0,0]), vec![vec![0,0,0]]);
        assert_eq!(Solution::three_sum(vec![0,0,0]), vec![vec![0,0,0]]);
    }
}
