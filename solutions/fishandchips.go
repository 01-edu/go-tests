package solutions

import "strconv"

func FishAndChips (n int) string {
  if (n % 2 == 0) && (n % 3 == 0) {
    return "fish and chips\n"
  } else if (n % 2 == 0) && !(n % 3 == 0) {
     return "fish\n"
  } else if (n % 3 == 0)  && !(n % 2 == 0) {
     return "chips\n"
	} else if !(n < 0) || !((n%2 == 0) && (n%3 == 0)) {
		 return "error: number is negative or non divisible"
	}
	 return strconv.Itoa(n)
}
