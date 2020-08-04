fn main() {
    println!("Hello, world!");
}

pub struct Solution;
impl Solution {
    pub fn can_finish(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> bool {
        let mut num_pres = vec![0; num_courses as usize]; //记录每个课程依赖的课程个数,先修课程都学完了就可以学习当前课程
        let mut sufs = vec![Vec::new(); num_courses as usize]; //记录依赖这个课程的课程列表

        for pre in &prerequisites {
            num_pres[pre[0] as usize] = num_pres[pre[0] as usize] + 1;
            sufs[pre[1] as usize].push(pre[0]);
        }

        let mut bfs_courses = Vec::new();

        for i in 0..num_pres.len() {
            if num_pres[i] == 0 { //没有先修课程，就可以先进行学习
                bfs_courses.push(i)
            }
        }

        let mut learned = 0;
        while let Some(cur) = bfs_courses.pop() {
            learned = learned+1;
            for i in &sufs[cur] {
                num_pres[*i as usize] = num_pres[*i as usize] - 1;

                //如果这个课程i 所需的先修课程已经学完了
                if num_pres[*i as usize] == 0 {
                    bfs_courses.push(*i as usize);
                }
            }
        }



        learned == num_courses
    }
}
