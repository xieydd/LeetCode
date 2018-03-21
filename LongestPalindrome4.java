/**
 * description:
 * 
 *  Input: "babad"
    Output: "bab"
    Note: "aba" is also a valid answer.
 * 
 * LeetCode:LongestPalindrome4
 * author:xieydd
 * date:Thu Mar 15 21:44:41 CST 2018
 */
class LongestPalindrome4{

    /**
     * Expand Around Center but get wrong ans
     */
    public static String longestPalindrome1(String s) {
        int start = 0,end = 0;
        for(int i =0;i<s.length();i++) {
            int len1 = expandAroundCenter(s,i,i);
            int len2 = expandAroundCenter(s,i,i+1);
            int len = Math.max(len1,len2);
            if(len > end -start) {
                start = i - (len -1) /2;
                end = i + len/2;
            }
        }
        return s.substring(start,end+1);
    }

    private static int expandAroundCenter(String s,int left,int right){
        while(left >= 0 && right <s.length() && s.charAt(left) == s.charAt(right) ) {
            left --;
            right ++;
        }
        return right - left - 1;
    }


    /**
     * faster and simple to understand
     */
    private static int lo,maxLen;
    public static String longestPalindrome2(String s) {
        int len = s.length();
        if(len < 2) {
            return s;
        }

        for(int i =0;i < len-1;i++) {
            extendPalindrome(s, i, i);
            extendPalindrome(s, i, i+1);
        }
        return s.substring(lo,lo+maxLen);
    }

    private static void extendPalindrome(String s,int j,int k) {
        while(j >= 0 && k < s.length() && s.charAt(j) == s.charAt(k)) {
            j--;
            k++;
        }
        if(maxLen < k-j-1) {
            lo = j+1;
            maxLen = k-j-1;
        }
    }



    public static void main(String[] args) {
        String s = "babad";
        String result = longestPalindrome2(s);
        System.out.println(result);
    }
}