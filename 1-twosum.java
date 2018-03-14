/**
 * description:LeetCode:1_twosum
 * author:xieydd
 * date:Tue Mar 13 20:16:50 CST 2018
 */
class TwoSum{
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
        //int[] result = twoSum(nums,target);
        System.out.println(result);
    }
}
