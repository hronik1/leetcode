func merge(nums1 []int, m int, nums2 []int, n int)  {
    for writei, readi, readj := len(nums1)-1, m-1, n-1; writei >= 0; writei-- {
        if readj < 0 || (readi >= 0 && nums1[readi] > nums2[readj]) {
            nums1[writei] = nums1[readi]
            readi--
        } else {
            nums1[writei] = nums2[readj]
            readj--
        }
    }
    // add remainder of 
}
