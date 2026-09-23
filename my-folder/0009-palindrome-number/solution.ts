function isPalindrome(x: number): boolean {
  let reverse = 0;
  let copy = x;
  while(copy > 0) {
    reverse *= 10; // 0
    reverse += copy % 10; // 1
    copy = Math.floor(copy / 10);
  }

  return reverse == x;  
};
