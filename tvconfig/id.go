package chainid

import "math/big"

const ChainID = 6359

const GasLimit = 1000000000000

const EnforceTips = true

const Period = 2
const Epoch = 30000

const YearBlocks = 15768000

const StartTime = 1756696906

const Reward = 200e18

const CoinBase = "0x0000000000000000000000000000000000111111"

const FeesCoinBase = "0x0000000000000000000000000000000000111111"

func CalcReward(number uint64) *big.Int {
	//year := number / YearBlocks
	//
	//blockReward := uint64(Reward)
	//for i := uint64(0); i < year; i++ {
	//	blockReward = blockReward * 9 / 10
	//}
	return new(big.Int).SetUint64(0)
}

const Addr1 = "0x2cdFC885C5604474892C891E4f24Abd4b871948a"
var Pri1 = ""

const Addr2 = "0xeEe35C2Ad0c1f50856605dB9337eE605DAd0b41A"
var Pri2 = ""
