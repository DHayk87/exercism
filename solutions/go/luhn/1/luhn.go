package luhn

func Valid(id string) bool {
	sum := 0
	numCount := 0
	shouldDouble := true
	for i := len(id) - 1; i >= 0; i-- {
		if id[i] == byte(' ')  {
			continue
		}else if  id[i] < '0' || id[i] > '9'{
			return false
		} else if shouldDouble {
			sum += int(id[i] - '0')
		} else {
			current := int(id[i]-'0') * 2
			if current > 9 {
				current -= 9
			}
			sum += current
		}
		numCount++
		shouldDouble = !shouldDouble
	}

	return  sum % 10 == 0 && numCount > 1
}
