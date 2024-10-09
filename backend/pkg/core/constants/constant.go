package constants

import "time"

const (
	GachaBoardMinimumChoose          = 1
	GachaBoardMaximumChoose          = 9
	ForgetPasswordExpiredDuration    = 15 * time.Minute
	GachaRewardLevel1                = 0
	GachaRewardLevel2                = 10000
	GachaRewardLevel3                = 20000
	GachaRewardLevel4                = 50000
	GachaRewardLevel5                = 100000
	GachaRewardLevel6                = 150000
	GachaRewardLevel7                = 200000
	GachaRewardLevel8                = 250000
	GachaRewardLevel9                = 300000
	WalletNumberStart                = 4200000000000
	MinimumPasswordLength            = 6
	MaximumPasswordLength            = 20
	ForgetPasswordTokenLength        = 6
	MinimumTopUpAmount               = 50000
	MaximumTopUpAmount               = 10000000
	MinimumTransferAmount            = 1000
	MaximumTransferAmount            = 50000000
	MaximumTransferDescriptionLength = 35
)

var (
	GachaBoard = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
)
