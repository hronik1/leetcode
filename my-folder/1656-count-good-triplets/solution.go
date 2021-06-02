import "math"

func countGoodTriplets(arr []int, a int, b int, c int) int {
    count := 0
    l := len(arr)
    for i := 0; i < l-2; i++ {
        for j := i+1; j < l-1; j++ {
            for k := j+1; k < l; k++ {
                if int(math.Abs(float64(arr[j]-arr[i]))) <= a && int(math.Abs(float64(arr[k]-arr[j]))) <= b && int(math.Abs(float64(arr[k]-arr[i]))) <= c {
                    count++
                }            
            }
        } 
    }
    
    return count
}
