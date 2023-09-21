func merge(nums1 []int, m int, nums2 []int, n int)  {
    writeI := len(nums1) - 1
    readI := m - 1
    j := n - 1
    
    for readI >= 0 && j >= 0 {
        if nums2[j] > nums1[readI] {
            nums1[writeI] = nums2[j]
            j -= 1
        } else {
            nums1[writeI] = nums1[readI]
            readI -= 1
        }
        
        writeI -= 1
    }
    
    for j >= 0 {
        nums1[writeI] = nums2[j]
        j -= 1
        writeI -= 1
    }
}
