fn main() {
    println!("Hello, world!");

    let word_list = vec![
        String::from("hot"),
        String::from("dot"),
        String::from("dog"),
        String::from("lot"),
        String::from("log"),
        String::from("cog"),
    ];
    assert_eq!(
        Solution::ladder_length(String::from("hit"), String::from("cog"), word_list),
        5
    );

    
    let word_list2 = vec![
        String::from("hot"),
        String::from("dot"),
        String::from("dog"),
        String::from("lot"),
        String::from("log"),
    ];
    assert_eq!(
        Solution::ladder_length(String::from("hit"), String::from("cog"), word_list2),
        0
    );

    let word_list3 = vec![
        String::from("hot"),
        String::from("dog"),
        String::from("cog"),
        String::from("pot"),
        String::from("dot"),
    ];
    assert_eq!(
        Solution::ladder_length(String::from("hot"), String::from("dog"), word_list3),
        3
    );

    let word_list4 = vec![
        String::from("a"),
        String::from("b"),
        String::from("c"),
    ];
    assert_eq!(
        Solution::ladder_length(String::from("a"), String::from("c"), word_list4),
        2
    );

    
}
pub struct Solution;


impl Solution {
    /*
    1. BFS
    2. 从 begin_word 开始 220ms  2.7MB
    */
    pub fn ladder_length_dfs(begin_word: String, end_word: String, word_list: Vec<String>) -> i32 {
        let mut exist = false;
        for v in &word_list {
            if *v == end_word {
                exist = true;

                break;
            }
        }

        if !exist {
            return 0;
        }

        let mut word = std::collections::HashSet::new();

        let mut queue = std::collections::VecDeque::new();
        queue.push_back(begin_word.clone());
        word.insert(begin_word.clone());


        let mut step = 1;

        while queue.len() > 0 {
            let num = queue.len();

            for i in 0..num {
                if let Some(node) = queue.pop_front() {
                    for x in &word_list {
                        //如果词典里的这个单词和当前节点的单词 差一个字符,并且之前没有访问过
                        if Self::one_char_diff(&node, x) && (!word.contains(x)) {

                            if *x == end_word{
                                return  step+1
                            }
                           
                            queue.push_back(x.clone());
                            word.insert(x.clone());

                        }
                    }
                }
            }

            step += 1;
        }

        0
    }

    //判断两个字符是否只差一个字符不一样
    fn one_char_diff(from: &String, to: &String) -> bool {
        if from.len() != to.len() {
            return false;
        }

        let mut cnt = 0;
        for i in 0..from.len() {
            if from.as_bytes()[i] != to.as_bytes()[i] {
                cnt += 1;
                if cnt == 2 {
                    return false;
                }
            }
        }

        return cnt == 1;
    }

    //双边BFS 20ms 2.5MB
    pub fn ladder_length(begin_word: String, end_word: String, word_list: Vec<String>) -> i32 {

        let mut exist = false;
        let mut end_id = word_list.len() ;
        for i in 0..word_list.len() {
            if word_list[i] == end_word {
                exist = true;
                end_id = i;
                break;
            }
        }

        if !exist {
            return 0;
        }



        let mut flags = vec![0;word_list.len()];

        //两个队列分别用于两边BFS
        let mut first_queue = std::collections::VecDeque::new();
        first_queue.push_back(&begin_word);

        let mut second_queue = std::collections::VecDeque::new();
        second_queue.push_back(&end_word);
        flags[end_id] = 2; //这里要注意，end_word已经加入了右边的队列，所以要赋值2

        let mut cur_queue;
        let mut flag;
        let mut step = 1;
        while first_queue.len()>0 && second_queue.len()>0 {

            if first_queue.len()>second_queue.len(){
                cur_queue = & mut second_queue;
                flag = 2;
            }else{
                //如果大小相同，就从左边开始 
                cur_queue = & mut first_queue;
                flag = 1;
            }

            let num = cur_queue.len();

            for _ in 0..num{

                if let Some(node) = cur_queue.pop_front(){


                    for i in 0..word_list.len() {
                        //如果词典里的这个单词和当前节点的单词 差一个字符,并且这个方向没有访问过
                        if Self::one_char_diff(&node, &word_list[i]) && (flags[i] != flag) {

                            flags[i]+=flag;
                            if flags[i] ==3{
                                return step+1;
                            }
                            cur_queue.push_back(&word_list[i]);

                        }
                    }


                }
            }
            step+=1;

        }



        0
    }
}
