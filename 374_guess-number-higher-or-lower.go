/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

import "math"

func guessNumber(n int) int {    
    mid := int(math.Ceil(float64(n) / 2));
    diff := mid

    for true {
        res := guess(mid)
        if res == 0 {
            return mid
        }
        
        diff = int(math.Ceil(float64(diff) / 2))

        if res == -1 {
            mid = mid - diff
        } 
        if res == 1 {
            mid = mid + diff
        }
    }

    return mid
}
