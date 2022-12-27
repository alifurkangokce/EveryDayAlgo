using System;
using System.Collections.Generic;
using System.Linq;

class Result
{
    public static string fairRations(List<int> B)
    {
        int odd = 0;
        List<int> oddIndices = new List<int>();
        for(int i = 0; i < B.Count; i++){ 
            if(B[i] % 2 == 1) {
                oddIndices.Add(i);
                odd++;
            }
        }
   
        if(odd.Equals(0)) return "0";
        if(odd % 2 == 1) return "NO";
        

        int min = 0;
        for(int i = 0; i < (oddIndices.Count - 1); i += 2){
            min += (oddIndices[i + 1] - oddIndices[i]);
        }
        
  
        return $"{min * 2}";
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        //TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        int N = Convert.ToInt32(Console.ReadLine().Trim());

        List<int> B = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(BTemp => Convert.ToInt32(BTemp)).ToList();

        string result = Result.fairRations(B);

        // textWriter.WriteLine(result);

        // textWriter.Flush();
        // textWriter.Close();
        Console.WriteLine(result);
    }
}