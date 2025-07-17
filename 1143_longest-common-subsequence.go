func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1);
    n := len(text2);
    lcq := make([][]int, m);

    for i := range lcq {
        lcq[i] = make([]int, n);
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if text1[i] == text2[j] {
                if i==0 || j==0 {
                    lcq[i][j] = 1;
                    continue;
                }
                lcq[i][j] = lcq[i-1][j-1] + 1;
            } else {
                if i==0 && j==0 {
                    lcq[i][j] = 0;
                    continue;
                }
                if j==0 {
                    lcq[i][j] = lcq[i-1][0];
                    continue;
                }
                if i==0 {
                    lcq[i][j] = lcq[0][j-1];
                    continue;
                }
                lcq[i][j] = max(lcq[i][j-1], lcq[i-1][j]);
            }
        }
    }

    return lcq[m-1][n-1];
}