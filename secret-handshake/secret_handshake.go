package secret

var events = [4]string{"wink", "double blink", "close your eyes", "jump"}

//Handshake return the secret sequence based on the number given
func Handshake(n uint) []string {
	reverse := n > 16
	result := []string{}
	n = n % 16
	for _, s := range events {
		if n&1 == 1 {
			if reverse {
				result = append([]string{s}, result...)
			} else {
				result = append(result, s)
			}
		}
		n >>= 1
	}
	return result
}
