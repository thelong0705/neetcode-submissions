func isAnagram(s string, t string) bool {
   if len(s) != len(t) {
	return false
   }
   countS, countT := map[byte]int{}, map[byte]int{}
   
   for i, _ := range s {
	countS[s[i]]++
	countT[t[i]]++
   }

   for k, _ := range countS {
	if countS[k] != countT[k] {
		return false
	} 
   }

   return true
}
