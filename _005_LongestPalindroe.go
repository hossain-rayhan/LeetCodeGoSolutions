func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    
    return longestPalindromeN2(s)
    //return longestPalindromeN3(s)
}

// Using DP, imporoved one
// Time Complexity O(N^2), so N2 added after the func name
func longestPalindromeN2(s string) string {
    n := len(s)
    
    // isPalindrome[i][j] is true if substr[i,j] is a palindrome
    isPalindrome := make([][]bool, n)
    for i := range isPalindrome {
        isPalindrome[i] = make([]bool, n)
    }
    
    // All substring of length 1 is a palindrome
    for i := 0; i < n; i++ {
        isPalindrome[i][i] = true;
    }
    
    maxLength := 1
    startIndex := 0
    
    // Check for substring of length 2
    for i := 0; i < n - 1; i++ {
        if s[i] == s[i+1] {
            isPalindrome[i][i+1] = true
            startIndex = i
            maxLength = 2
        }
    }
    
    // Check for substring of length l >= 3
    for l := 3; l <= n; l++ {
        for start := 0; start < n - l + 1; start++ {
            end := start + l - 1;
            
            //check if the stbstring is a palindrome
            if isPalindrome[start+1][end-1] && s[start] == s[end] {
                isPalindrome[start][end] = true
                if l > maxLength {
                    startIndex = start
                    maxLength = l
                }
            } 
        }
    }
    
    return s[startIndex : startIndex + maxLength]
}

// The very first thought-- easy way to think
// Time complexity O(N^3), so N3 added after the func name
func longestPalindromeN3(s string) string {
    n := len(s) 
    longSub := ""
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if isPalindrome(s, i, j) && j - i + 1 > len(longSub) {
                longSub = s[i:j+1]
            }
        }
    }
    return longSub
}

func isPalindrome(s string, start int, end int) bool{
    for ; end >= start; {
        if s[start] != s[end] {
            return false
        }
        start++
        end--
    }
    return true
}
