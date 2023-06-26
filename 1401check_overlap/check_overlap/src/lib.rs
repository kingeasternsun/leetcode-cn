struct Solution;
impl Solution {
    pub fn check_overlap(
        radius: i32,
        x_center: i32,
        y_center: i32,
        x1: i32,
        y1: i32,
        x2: i32,
        y2: i32,
    ) -> bool {
        let r = radius.pow(2);
        if Self::distantce(x1, y1, x_center, y_center) < r {
            return true;
        }

        if Self::distantce(x2, y2, x_center, y_center) <= r {
            return true;
        }

        if Self::distantce(x1, y2, x_center, y_center) <= r {
            return true;
        }

        if Self::distantce(x2, y1, x_center, y_center) <= r {
            return true;
        }

        let max_x = x1.max(x2);
        let min_x = x1.min(x2);
        let max_y = y1.max(y2);
        let min_y = y1.min(y2);

        if x_center <= min_x {
            let mut dis = (x_center - min_x).pow(2);
            if y_center < min_y {
                dis = Self::distantce(min_x, min_y, x_center, y_center)
            } else if y_center > max_y {
                dis = Self::distantce(min_x, max_y, x_center, y_center)
            }

            if dis <= r {
                return true;
            }

            return false;
        }

        if x_center >= max_x {
            let mut dis = (x_center - max_x).pow(2);
            if y_center < min_y {
                dis = Self::distantce(max_x, min_y, x_center, y_center)
            } else if y_center > max_y {
                dis = Self::distantce(max_x, max_y, x_center, y_center)
            }

            if dis <= r {
                return true;
            }

            return false;
        }

        if y_center < min_y {
            if (y_center - min_y).pow(2) <= r {
                return true;
            }

            return false;
        }

        if y_center > max_y {
            if (y_center - max_y).pow(2) <= r {
                return true;
            }

            return false;
        }

        // the center is inside of the rectangle
        true
    }

    fn distantce(x: i32, y: i32, x_center: i32, y_center: i32) -> i32 {
        (x - x_center).pow(2) + (y - y_center).pow(2)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(Solution::check_overlap(1, 0, 0, 1, -1, 3, 1), true);
        assert_eq!(Solution::check_overlap(1, 1, 1, 1, -3, 2, -1), false);
        assert_eq!(Solution::check_overlap(1, 0, 0, -1, 0, 0, 1), true);
    }
}
