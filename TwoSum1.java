import java.util.HashMap;
import java.util.Map;


/**
 * description:
 * 
 *  Given nums = [2, 7, 11, 15], target = 9,
    Because nums[0] + nums[1] = 2 + 7 = 9,
    return [0, 1].
 * 
 * LeetCode:TwoSum1
 * author:xieydd
 * date:Tue Mar 13 20:16:50 CST 2018
 */
class TwoSum1{

    /**
     * force
     */
    public static int[]  twoSum_1(int[] nums,int target) {
        for(int i = 0;i<nums.length;i++) {
            for(int j=i+1;j<nums.length;j++) {
                if(nums[j] == target-nums[i]) {
                    return new int[] {i,j};
                }
            }
        }
        throw new IllegalArgumentException("No two sum solution");
    }
    /**
     * reverse
     */
    public static int[] twoSum_2(int[] nums,int target){
        for(int i =0;i<nums.length;i++) {
            Map<Integer,Integer> hm = new HashMap<Integer,Integer>();
            hm.put(nums[i],i);

            for(int j=0;j<nums.length;j++) {
                int complement = target - nums[j];
                if(hm.containsKey(complement) && hm.get(complement) != j) {
                    return new int[] {j,hm.get(complement)};
                }
            }
        }
        throw new IllegalArgumentException("No two sum solution");
    }

    public static int[] twoSum_3(int[] nums,int target){
        Map<Integer,Integer> hm = new HashMap<Integer,Integer>();
        for(int i =0;i<nums.length;i++) {
            int complement = target - nums[i];
            if(hm.containsKey(complement)) {
                return new int[] {hm.get(complement),i};
            }
            hm.put(nums[i],i);
        }throw new IllegalArgumentException("No two sum solution");
    }
    

    
    public static void main(String[] args) {
        int[] nums = {3,2,4};
        int target = 6;
        int[] result = twoSum_3(nums,target);
        System.out.println(result.toString());
    }
}
