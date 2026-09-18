package prime

// x以下の素数のリストを返す
func SieveOfEratosthenes(x int) []int {
	// 2未満に素数なし
	if x < 2 {
		return nil
	}

	// 0からxまでのインデックスを確保
	isPrime := make([]bool, x+1)
	for i := 2; i <= x; i++ {
		isPrime[i] = true
	}

	// sqrt(x)まで判定すると、それより大きい合成数も小さい素因数を持つ
	for i := 2; i*i <= x; i++ {
		if isPrime[i] {
			for j := i * i; j <= x; j += i { // jがiの倍数になる
				isPrime[j] = false
			}
		}
	}

	// 素数リストの構築
	var primes []int
	for i := 2; i <= x; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}