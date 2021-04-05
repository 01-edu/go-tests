package solutions

func Sweetproblem(a, b, c int) int {
	if a > b {
		f := a
		a = b
		b = f
	}
	if a > c {
		f := a
		a = c
		c = f
	}
	if b > c {
		f := b
		b = c
		c = f
	}
	ans := a
	if c-b >= a {
		c -= a
	} else {
		a -= c - b
		half := a / 2
		c -= half
		b -= a - half
	}
	if b <= c {
		ans += b
	} else {
		ans += c
	}
	return ans
}
