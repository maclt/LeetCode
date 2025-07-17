func tribonacci(n int) int {
    var mems []int = []int{0,1,1}
    var mod int

    switch n {
        case 0:
          return mems[0]
        case 1:
          return mems[1]
        case 2:
          return mems[2]
    }

    for i:=3; i<=n; i++ {
        mod = i%3
        mems[mod] = mems[0] + mems[1] + mems[2]
    }

    return mems[mod]
}

// not fast enough!!
// func tribonacci(n int) int {
//     switch n {
//         case 0:
//           return 0 
//         case 1:
//           return 1 
//         case 2:
//           return 1
//     }
//     return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
// }