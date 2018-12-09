/**
 * description:
 * 
 *  PAYPALISHIRING is written in a zigzag pattern given number of rows like this
 *  P   A   H   N
    A P L S I I G
    Y   I   R
    convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR"
 * 
 * 
 * zigzag pattern
 * /*n=numRows
Δ=2n-2    1                           2n-1                         4n-3
Δ=        2                     2n-2  2n                    4n-4   4n-2
Δ=        3               2n-3        2n+1              4n-5       .
Δ=        .           .               .               .            .
Δ=        .       n+2                 .           3n               .
Δ=        n-1 n+1                     3n-3    3n-1                 5n-5
Δ=2n-2    n                           3n-2                         5n-4

that’s the zigzag pattern the question asked!
Be careful with nR=1 && nR=2
 * 
 * 
 * LeetCode:ZigZagConversion6
 * author:xieydd
 * date:Thu Mar 15 21:44:41 CST 2018
 */
class ZigZagConversion6 {
    public static String convert(String s,int rows) {
        StringBuilder[] sb = new StringBuilder[rows];
        char[] c = s.toCharArray();
        int len = c.length;

        for(int i = 0;i < sb.length;i++) sb[i] = new StringBuilder();

        int i = 0;
        while(i < len){
            for(int idx = 0; idx < rows && i < len;idx++) {
                sb[idx].append(c[i++]);
            }

            for(int idx = rows -2;idx >= 1 && i <len;idx--) {
                sb[idx].append(c[i++]);
            }
        }


        for(int idx = 1;idx < sb.length;idx++) {
            sb[0].append(sb[idx]);
        }
        return sb[0].toString();
    }

    /**
     * easy to understand
     * from:https://segmentfault.com/a/1190000005751185
     */

    public static String convert2(String s, int numRows) {
        if (numRows == 1)
            return s;
        StringBuilder sb = new StringBuilder();
        int magic = numRows * 2 - 2;
        int initialDistance = magic;
        for (int i = 0; i < numRows; i++) {
            fill(sb, i, initialDistance, magic, s);
            initialDistance = initialDistance - 2;
        }
        return sb.toString();
    }


    public static void fill(StringBuilder sb, int start, int initialDistance, int magic, String s) {
        while (start < s.length()) {
            if (initialDistance == 0)    
                initialDistance = magic;
            sb.append(s.charAt(start));
            start = start + initialDistance;  
            initialDistance = magic - initialDistance;
            
        }
    }

    public static void main(String[] args) {
        String s = "PAYPALISHIRING";
        int rows = 3;
        String result = convert2(s, rows);
        System.out.println(result);
    }
}