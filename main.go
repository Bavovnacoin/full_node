package main

import "bavovnacoin/testing/singleFunctionTesting"

func main() {
	// a, _ := new(big.Int).SetString("00ffff0000000000000000000000000000000000", 16)
	// fmt.Printf("%x\n", blockchain.TargetToBits(a))
	var ct singleFunctionTesting.CryptoTest
	ct.Sha1TestAmmount = 5
	ct.Launch("SHA1")

	// node.Launch()
	//synchronization.GetCheckpHashes(1)

}
