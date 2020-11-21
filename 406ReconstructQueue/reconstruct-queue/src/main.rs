fn main() {
    println!("Hello, world!");
    
let people = vec![
    // vec![],
    vec![7,0], vec![4,4], vec![7,1], vec![5,0], vec![6,1], vec![5,2]
];

println!("{:?}",Solution::reconstruct_queue(people));
    
}

pub struct Solution;

impl Solution {
    pub fn reconstruct_queue(people: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        
        if people.len()==0{
            return people
        }

        let mut people = people;
        let mut res = Vec::new();
        people.sort_by(|a,b|if a[0]==b[0]{ a[1].cmp(&b[1]) }else{b[0].cmp(&a[0])});
      

        // println!("{:?}",people);
        people.into_iter().for_each(|p| res.insert(p[1]as usize, p));

        res

    }
}