package solutions

func ShoppingListSort(array[] string)[]string {
	for i := 0; i < len(array) - 1; i++ {
		 for j := 0; j < len(array) - i - 1; j++ {
				if (len(array[j]) > len(array[j + 1])) {
					 array[j], array[j + 1] = array[j + 1], array[j]
					}
      }
   }
   return array
}
