fn main() {
    println!("Hello, world!");
}

pub struct Solution;

impl Solution {
    pub fn minimum_time_required(jobs: Vec<i32>, k: i32) -> i32 {

        if jobs.len()==0{
            return 0;
        }

        let mut beg = jobs.iter().max().unwrap().clone();
        let mut end:i32 = jobs.iter().sum();

   

        let k = k as usize;

        if jobs.len()==k{
            return beg;
        }

        //二分法查找
        let mut mid = 0;
        while beg < end {
            mid = (beg + end)>>1;
            let mut work_times = vec![0;k ];
            if Self::dfs(0, &jobs[..], mid, k , &mut work_times[..]){
                end = mid;
            }else{
                beg = mid+1;
            }
            
        }

        return beg

    }

    //判断 是否可以让每个工人的总工作量不大于 mid
    fn dfs(curjob_id: usize, jobs: &[i32], mid: i32, k: usize, work_times: &mut [i32]) -> bool {

        if curjob_id == jobs.len(){
            return true;
        }
        //尝试把当前工作分配给工人work做，看是否可以满足条件
        for worker in 0..k {
            if work_times[worker] + jobs[curjob_id] <= mid {
                work_times[worker] += jobs[curjob_id];
                if Self::dfs(curjob_id + 1, jobs, mid, k, work_times) {
                    return true;
                }

                work_times[worker] -= jobs[curjob_id];
            }

            //如果当前工人的没有任何工作，表示不管怎么给这个工人安排工作都无法满足 mid 
            if work_times[worker] == 0 {
                return false;
            }
        }

        false
    }
}


#[cfg(test)]
mod tests{
    use super::*;
    #[test]
    fn one(){
        assert_eq!(Solution::minimum_time_required(vec![1,2,4,7,8],2),11);
    }
    
}

