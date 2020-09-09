fn main() {
    println!("Hello, world!");
    // let result = Solution::combination_sum2(vec![10,1,2,7,6,1,5],8);
    let result = Solution::combination_sum2(vec![2,5,2,1,2],5);
    println!("{:?}",result);
}
pub struct Solution;
impl Solution {
    pub fn combination_sum2(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut result = Vec::new();
        let mut tmp = Vec::new();
        candidates.sort();
        Self::dfs(& candidates,target,0,& mut tmp,& mut result,false);
        result
    }

    fn dfs(
        candidates: &Vec<i32>,
        target: i32,
        next_id: usize,
        tmp: &mut Vec<i32>,
        result: &mut Vec<Vec<i32>>,
        flag: bool,
    ) {
        if target == 0 {
            result.push(tmp.clone());
            return;
        }

        if next_id == candidates.len(){
            return ;
        }


        if candidates[next_id] > target {
            return;
        }

 
        // for i in (next_id..candidates.len()).take_while(|&x| candidates[x] <=target){

        //     if next_id >0 && candidates[next_id]==candidates[next_id-1]{
        //         //第next_id不选
        //         Self::dfs(candidates,target,)

        //     }

        // }

        // 重复的元素，之前相同的没有选，那么当前也不能选，避免重复组合
        if next_id > 0 && candidates[next_id] == candidates[next_id - 1] && flag == false {
            //next_id 位置上的number 不放入 tmp
            Self::dfs(candidates, target, next_id + 1, tmp, result, false);
        } else {
            //next_id 位置上的number 不放入 tmp
            Self::dfs(candidates, target, next_id + 1, tmp, result, false);

            //   next_id 位置上的number 放入 tmp
            tmp.push(candidates[next_id]);
            Self::dfs(
                candidates,
                target - candidates[next_id],
                next_id + 1,
                tmp,
                result,
                true,
            );
            tmp.pop();
        }
    }
}
