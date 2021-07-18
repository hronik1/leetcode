func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) == 0 {
        return findMedianSortedArray(nums2)
    }

    if len(nums2) == 0 {
        return findMedianSortedArray(nums1)
    }
    
    l := len(nums1)+len(nums2)
    mid := l/2
    i, j := 0, 0
    cur, prev := 0, 0
    for k := 0; k <= mid; k++ {
        if i >= len(nums1) {
            prev = cur
            cur = nums2[j]
            j++
        } else if j >= len(nums2) {
            prev = cur
            cur = nums1[i]
            i++
        } else if nums1[i] < nums2[j] {
            prev = cur
            cur = nums1[i]
            i++
        } else {
            prev = cur
            cur = nums2[j]
            j++
        }
    }
    
    fmt.Printf("mid:%d, prev:%d, cur:%d", mid, prev, cur)
    if l%2 == 1 {
        return float64(cur)
    }
    
    return float64(prev) + float64(cur-prev)/2.0
}

func findMedianSortedArray(nums []int) float64 {
    mid := len(nums)/2
    if len(nums)%2 == 1 {
        return float64(nums[mid])
    } else {
        return float64(nums[mid-1]) + float64(nums[mid]-nums[mid-1])/2.0
    }
}
