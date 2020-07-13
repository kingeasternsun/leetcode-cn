fn main() {
    println!("Hello, world!");
}
pub struct Solution {}

use std::collections::BinaryHeap;
use std::cmp::Ordering;

#[derive(Copy,Clone)]
struct State{
    cost:f64,
    position:usize,
}

impl PartialEq for State {
    fn eq(&self,other: &State)->bool{
        self.position == other.position && (self.cost - other.cost).abs()< 1e-10
    }
}

impl Eq for State { }


impl Ord for State{
    fn cmp(&self,other: &State)->Ordering{
        other.cost.partial_cmp(&self.cost).unwrap().then_with(|| self.position.cmp(&other.position))
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &State) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Solution {

   
    pub fn max_probability(n: i32, edges: Vec<Vec<i32>>, succ_prob: Vec<f64>, start: i32, end: i32) -> f64 {
        
        pub const MAX: f64 = 1.7976931348623157e+308_f64;
        let mut dist = vec![MAX;n as usize]; //到这个点的最小值
    
        let mut visited = vec![false;n as usize];  //这个点是否处理过了

        let mut g = vec![Vec::<(usize,f64)>::new();n as usize];
        for i in 0..edges.len(){
            let edge = &edges[i];
            g[edge[0] as usize].push((edge[1] as usize,-succ_prob[i].log2()));
            g[edge[1] as usize].push((edge[0] as usize,-succ_prob[i].log2()));

        }


        let  mut   heap =BinaryHeap::new();
        heap.push(State{cost:0.0,position:start as usize});
        dist[start as usize] = 0.0;
        visited[start as usize] = true;


        while let Some(State{cost,position}) = heap.pop(){

            if position == end as usize {
                return (-cost).exp2()
            }

            visited[position]=true;

            for i in 0..g[position].len(){

                let (next_position,next_cost) = g[position][i];
                if visited[next_position]{
                    continue
                }

                if cost+next_cost > dist[next_position]{
                    continue
                } 

                dist[next_position] = cost+next_cost;
                heap.push(State{cost:cost+next_cost,position:next_position});

            }

        }
    



        0.0

    }
}