pub struct Solution;
pub struct Queue {
    data: Vec<(i32, i32)>, //record the 'O' that has been fliped
    front: usize,          // record the index of front element
}
impl Queue {
    pub fn new() -> Self {
        Queue {
            data: Vec::new(),
            front: 0,
        }
    }
    fn push(&mut self, x: (i32, i32)) {
        self.data.push(x);
    }

    fn pop(&mut self) -> (i32, i32) {
        self.front += 1;
        self.data[self.front - 1]
    }

    fn len(&self) -> usize {
        self.data.len() - self.front
    }

    fn data_size(&self) -> usize {
        self.data.len()
    }
}

use std::option::Option;
impl Solution {
    fn can_flip(chessboard: &Vec<Vec<char>>, cur: (i32, i32), dir: &(i32, i32)) -> Option<i32> {
        let board_size = (chessboard.len() as i32, chessboard[0].len() as i32);
        let mut cur = (cur.0 + dir.0, cur.1 + dir.1);
        let mut n = 0; // the number of 'O'
        while cur.0 >= 0 && cur.0 < board_size.0 && cur.1 >= 0 && cur.1 < board_size.1 {
            match chessboard[cur.0 as usize][cur.1 as usize] {
                '.' => return Option::None,
                'X' => return Option::Some(n),
                _ => {
                    cur.0 += dir.0;
                    cur.1 += dir.1;
                    n += 1;
                }
            }
        }

        Option::None
    }

    pub fn flip_chess(chessboard: Vec<String>) -> i32 {
        let mut ans = 0;
        let dirs = vec![
            (0, 1),
            (0, -1),
            (1, 0),
            (-1, 0),
            (1, 1),
            (-1, 1),
            (1, -1),
            (-1, -1),
        ];
        let mut chessboard: Vec<Vec<char>> = chessboard
            .iter()
            .map(|x| x.chars().collect::<Vec<char>>())
            .collect();
        let board_size = (chessboard.len(), chessboard[0].len());
        for i in 0..board_size.0 {
            for j in 0..board_size.1 {
                if chessboard[i][j] != '.' {
                    continue;
                }

                let mut queue = Queue::new();
                chessboard[i][j] = 'X';
                queue.push((i as i32, j as i32));
                while queue.len() > 0 {
                    let cur = queue.pop();
                    for dir in &dirs {
                        if let Some(n) = Self::can_flip(&chessboard, (cur.0 as i32, cur.1 as i32), dir){
                            let mut tmp = cur;
                            for _ in 0..n {
                                tmp = (tmp.0 + dir.0, tmp.1 + dir.1);
                                chessboard[tmp.0 as usize][tmp.1 as usize] = 'X';
                                queue.push(tmp);
                            }
                        }
                    }
                }

                if queue.data_size() - 1 > ans {
                    ans = queue.data_size() - 1;
                }

                for p in queue.data {
                    chessboard[p.0 as usize][p.1 as usize] = 'O';
                }

                chessboard[i][j] = '.';
            }
        }
        ans as i32
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(
            Solution::flip_chess(vec![
                String::from("....X."),
                String::from("....X."),
                String::from("XOOO.."),
                String::from("......"),
                String::from("......"),
            ]),
            3
        );
    }

    #[test]
    fn it_works2() {
        assert_eq!(
            Solution::flip_chess(vec![
                String::from(".X."),
                String::from(".O."),
                String::from("XO."),
            ]),
            2
        );
    }
}
