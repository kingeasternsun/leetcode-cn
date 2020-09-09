fn main() {
    println!("Hello, world!");
    let result = Solution::combination_sum(vec![2,3,5],8);
    println!("{:?}",result);
}

//如果数组 candidates 中的数字没有重复元素，而且每个元素只能选一个，那么解法如下
pub struct Solution1;
impl Solution1 {
    pub fn combination_sum(mut candidates:  Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        candidates.sort();
        let mut result = Vec::new();
        let mut tmp = Vec::new();
        Self::dfs(& candidates,target,& mut result,& mut tmp,None,0);

        result


    }

    fn dfs(candidates:  & Vec<i32>, target: i32, result :& mut Vec<Vec<i32>>, tmp : & mut Vec<i32>,tmp_id : Option<usize>,tmp_sum :i32){

        if tmp_sum == target{
            result.push(tmp.clone());
            return 
        }

        let mut beg = 0;
        if let Some(end) = tmp_id{
            beg = end+1
        }

        for i in beg..candidates.len(){

            let new_sum = tmp_sum + candidates[i];
            tmp.push(candidates[i]);
            Self::dfs(candidates,target,result,tmp,Some(i),new_sum);
            tmp.pop();

            if new_sum >= target{
                return 
            }

        }


    }
}

pub struct Solution;
impl Solution {
    pub fn combination_sum(mut candidates:  Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        candidates.sort();
        let mut result = Vec::new();
        let mut tmp = Vec::new();
        Self::dfs(& candidates,target,& mut result,& mut tmp,0);

        result


    }

    /*
     target 表示距离目标还差多少 
     tmp 表示当前已经装入了哪些
     tmp_id 表示 下一层级 从第几个开始
    */

    fn dfs(candidates:  & Vec<i32>, target: i32, result :& mut Vec<Vec<i32>>, tmp : & mut Vec<i32>,tmp_id : usize){

        if 0 == target{
            result.push(tmp.clone());
            return 
        }
        if candidates[tmp_id] > target{
            return 
        }


        for i in (tmp_id..candidates.len()).take_while(|&x| candidates[x]<=target){

            tmp.push(candidates[i]);
            Self::dfs(candidates,target-candidates[i],result,tmp,i);
            tmp.pop();

        }


    }
}
