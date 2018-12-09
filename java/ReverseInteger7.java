/**
 * description: Given a 32-bit signed integer, reverse digits of an integer.
Input: -123
Output: -321

Input: 120
Output: 21
 * 
 * LeetCode:ReverseInteger7
 * author:xieydd
 * date:Mon Mar 26 21:39:48 CST 2018
 */
class ReverseInteger7 {
    public static int reverse(int x) {
        int result = 0;
        while (x != 0)
        {
            int tail = x % 10;
            int newResult = result * 10 + tail;
            if ((newResult - tail) / 10 != result)
            { return 0; }
            result = newResult;
            x = x / 10;
        }
        return result;
    }

    public static void main(String[] args) {
        int x = 1124534645;
        int result = reverse(x);
        System.out.println(result);
    }
}