import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

/**
 * description:
 * 
 *  Given "abcabcbb", the answer is "abc", which the length is 3.
    Given "bbbbb", the answer is "b", with the length of 1.
    Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 * 
 * LeetCode:LengthOfLongestSubstring3
 * author:xieydd
 * date:Wed Mar 14 22:00:21 CST 2018
 */
public class LengthOfLongestSubstring3 {

    /**
     * force
     */
    public static int lengthOfLongestSubstring1(String s) {
        int length = s.length();
        int ans = 0;
        for(int i = 0;i < length;i++) {
            for(int j = i+1;j < length;j++) {
                if(allUnique(s, i, j)) ans =  Math.max(ans,j-i);
            }
        }
        return ans;
    }

    public static boolean allUnique(String s,int start,int end) {
        Set<Character> set = new HashSet<Character>();
        for(int i = start;i < end;i++) {
            Character ch = s.charAt(i);
            if(set.contains(ch)) return false;
            set.add(ch);
        }
        return true;
    }

    /**
     * slide Windows
     */
    public static int lengthOfLongestSubstring2(String s) {
        Set<Character> set = new HashSet<Character>();
        int length = s.length();
        int ans = 0 ,i =0,j = 0;
        while(i < length && j < length) {
            if(!set.contains(s.charAt(j))) {
                set.add(s.charAt(j++));
                ans = Math.max(ans,j-i);
            }else {
                set.remove(s.charAt(i++));
            }
        }  
        return ans;
    }

    

    /**
     * sliding Window Optimized
     */
    public static int lengthOfLongestSubstring3(String s) {
        int length = s.length(),ans = 0;
        Map<Character,Integer> map = new HashMap<Character,Integer>();
        for(int i = 0, j = 0;j < length;j++){
            if(map.containsKey(s.charAt(j))) {
                i = Math.max(map.get(s.charAt(j)),i);
            }
            ans = Math.max(ans,j-i+1);
            map.put(s.charAt(j),j+1);
        }
        return ans;
    }

    public static int lengthOfLongestSubstring4(String s) {

        /**
         * array is smaller than Map.
         *  int[26] for Letters 'a' - 'z' or 'A' - 'Z'
            int[128] for ASCII
            int[256] for Extended ASCII
         */
        int[] index = new int[128];
        int length = s.length();
        int ans = 0;
        for(int i = 0,j=0;j<length;j++) {
            i = Math.max(index[s.charAt(j)],i);
            ans = Math.max(ans,j-i+1);
            index[s.charAt(j)] = j+1;
        }
        return ans;
    }
    public static void main(String[] args) {
        String s = "abacdffe";
        int ans = lengthOfLongestSubstring4(s);
        System.out.println(ans);
    }
 }