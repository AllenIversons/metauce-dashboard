// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PlayGameBnbGasLink is an auto generated low-level Go binding around an user-defined struct.
type PlayGameBnbGasLink struct {
	Index    *big.Int
	Rate     *big.Int
	LinkName string
}

// PlayGameCarNFT is an auto generated low-level Go binding around an user-defined struct.
type PlayGameCarNFT struct {
	Pid         *big.Int
	Power       *big.Int
	Surplus     *big.Int
	Level       *big.Int
	TokenId     *big.Int
	CurrentTime *big.Int
	Owner       common.Address
	Used        bool
}

// PlayGameMapNFT is an auto generated low-level Go binding around an user-defined struct.
type PlayGameMapNFT struct {
	Pid         *big.Int
	Level       *big.Int
	TotalUsdt   *big.Int
	ConsumeUsdt *big.Int
	TokenId     *big.Int
	CurrentTime *big.Int
	Completed   bool
	Owner       common.Address
}

// PlayGameNFTPool is an auto generated low-level Go binding around an user-defined struct.
type PlayGameNFTPool struct {
	Index     *big.Int
	NftType   *big.Int
	Level     *big.Int
	Min       *big.Int
	Max       *big.Int
	MintPrice *big.Int
	Day       *big.Int
	NftUrl    string
	NftName   string
	MintPause bool
	Erc721    common.Address
}

// PlayGameUserIncome is an auto generated low-level Go binding around an user-defined struct.
type PlayGameUserIncome struct {
	TotalIncome  *big.Int
	TotalLocking *big.Int
	WithdrawAt   *big.Int
	Submits      *big.Int
	Settled      bool
}

// PlayGameUserIncomeDetail is an auto generated low-level Go binding around an user-defined struct.
type PlayGameUserIncomeDetail struct {
	TotalIncome *big.Int
	TokenId     *big.Int
	Level       *big.Int
	CurrentTime *big.Int
}

// PlayGameWithdrawInfo is an auto generated low-level Go binding around an user-defined struct.
type PlayGameWithdrawInfo struct {
	Index      *big.Int
	WithdrawAt *big.Int
	Amount     *big.Int
	Payfree    *big.Int
}

// PlayMetaData contains all meta data concerning the Play contract.
var PlayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PLAYNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allNFTPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nftUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nftName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"mintPause\",\"type\":\"bool\"},{\"internalType\":\"contractIQLF\",\"name\":\"erc721\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minResourceCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burnRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_carRate\",\"type\":\"uint256\"}],\"name\":\"setTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGameParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_minResourceCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burnRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_carRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawfree\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawReduce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"addPro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"updatePro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"}],\"name\":\"_findNFTLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_linkName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"addBnbGasLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"updateBnbGasLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nftName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nftUrl\",\"type\":\"string\"},{\"internalType\":\"contractIQLF\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_day\",\"type\":\"uint256\"}],\"name\":\"addNFTPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_nftName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nftUrl\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_mintPause\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_day\",\"type\":\"uint256\"},{\"internalType\":\"contractIQLF\",\"name\":\"_erc721\",\"type\":\"address\"}],\"name\":\"updateNFTPoolInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_free\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reduce\",\"type\":\"uint256\"}],\"name\":\"setWithdrawfree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bnbgas\",\"type\":\"uint256\"}],\"name\":\"setMiningInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_erc20Token\",\"type\":\"address\"}],\"name\":\"setErc20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCODE_IQLF\",\"name\":\"_codeIqlf\",\"type\":\"address\"}],\"name\":\"setCodeIQLF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserCarNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"surplus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"}],\"internalType\":\"structPlayGame.CarNFT[]\",\"name\":\"_carNFT\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getAllBnbGasLink\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"linkName\",\"type\":\"string\"}],\"internalType\":\"structPlayGame.BnbGasLink\",\"name\":\"_link\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserMapNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"consumeUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPlayGame.MapNFT[]\",\"name\":\"_mapNFT\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserIncome\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalIncome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalLocking\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submits\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"internalType\":\"structPlayGame.UserIncome\",\"name\":\"income\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserIncomeDetail\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalIncome\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"internalType\":\"structPlayGame.UserIncomeDetail[]\",\"name\":\"detail\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserWithdraws\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payfree\",\"type\":\"uint256\"}],\"internalType\":\"structPlayGame.WithdrawInfo[]\",\"name\":\"detail\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mapindex\",\"type\":\"uint256\"}],\"name\":\"getCarFromMap\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexs\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_carindex\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_mapindex\",\"type\":\"uint256\"}],\"name\":\"addCarToMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_carindex\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_mapindex\",\"type\":\"uint256\"}],\"name\":\"delCarToMap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"sellNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"canalSellNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"buyNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getUsersellNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"power\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"surplus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"}],\"internalType\":\"structPlayGame.CarNFT\",\"name\":\"_car\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"consumeUsdt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structPlayGame.MapNFT\",\"name\":\"_map\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"withdrawPayFree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_typeNFT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mapIndex\",\"type\":\"uint256\"}],\"name\":\"allowMiningResource\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_carIndex\",\"type\":\"uint256[]\"}],\"name\":\"checkCarUnlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"}],\"name\":\"loginInviteCodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_code\",\"type\":\"string\"}],\"name\":\"registerInviteCodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_carIndex\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_mapIndex\",\"type\":\"uint256\"}],\"name\":\"miningResource\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_carIndex\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_subMethod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoveryCar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_carIndex\",\"type\":\"uint256[]\"}],\"name\":\"getRecoveryCarBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_typeNFT\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subMethod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_typeNFT\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintOne\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"erc721Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_typeNFT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recvAddress\",\"type\":\"address\"}],\"name\":\"giveNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"_findNFTPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nftUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nftName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"mintPause\",\"type\":\"bool\"},{\"internalType\":\"contractIQLF\",\"name\":\"erc721\",\"type\":\"address\"}],\"internalType\":\"structPlayGame.NFTPool\",\"name\":\"_nftpool\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"randCar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"withdrawFree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PlayABI is the input ABI used to generate the binding from.
// Deprecated: Use PlayMetaData.ABI instead.
var PlayABI = PlayMetaData.ABI

// Play is an auto generated Go binding around an Ethereum contract.
type Play struct {
	PlayCaller     // Read-only binding to the contract
	PlayTransactor // Write-only binding to the contract
	PlayFilterer   // Log filterer for contract events
}

// PlayCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlaySession struct {
	Contract     *Play             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlayCallerSession struct {
	Contract *PlayCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PlayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlayTransactorSession struct {
	Contract     *PlayTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlayRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlayRaw struct {
	Contract *Play // Generic contract binding to access the raw methods on
}

// PlayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlayCallerRaw struct {
	Contract *PlayCaller // Generic read-only contract binding to access the raw methods on
}

// PlayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlayTransactorRaw struct {
	Contract *PlayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlay creates a new instance of Play, bound to a specific deployed contract.
func NewPlay(address common.Address, backend bind.ContractBackend) (*Play, error) {
	contract, err := bindPlay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Play{PlayCaller: PlayCaller{contract: contract}, PlayTransactor: PlayTransactor{contract: contract}, PlayFilterer: PlayFilterer{contract: contract}}, nil
}

// NewPlayCaller creates a new read-only instance of Play, bound to a specific deployed contract.
func NewPlayCaller(address common.Address, caller bind.ContractCaller) (*PlayCaller, error) {
	contract, err := bindPlay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlayCaller{contract: contract}, nil
}

// NewPlayTransactor creates a new write-only instance of Play, bound to a specific deployed contract.
func NewPlayTransactor(address common.Address, transactor bind.ContractTransactor) (*PlayTransactor, error) {
	contract, err := bindPlay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlayTransactor{contract: contract}, nil
}

// NewPlayFilterer creates a new log filterer instance of Play, bound to a specific deployed contract.
func NewPlayFilterer(address common.Address, filterer bind.ContractFilterer) (*PlayFilterer, error) {
	contract, err := bindPlay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlayFilterer{contract: contract}, nil
}

// bindPlay binds a generic wrapper to an already deployed contract.
func bindPlay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Play *PlayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Play.Contract.PlayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Play *PlayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Play.Contract.PlayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Play *PlayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Play.Contract.PlayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Play *PlayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Play.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Play *PlayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Play.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Play *PlayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Play.Contract.contract.Transact(opts, method, params...)
}

// FindNFTLevel is a free data retrieval call binding the contract method 0x4dc6a7da.
//
// Solidity: function _findNFTLevel(uint256 _nftType) view returns(uint256)
func (_Play *PlayCaller) FindNFTLevel(opts *bind.CallOpts, _nftType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "_findNFTLevel", _nftType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindNFTLevel is a free data retrieval call binding the contract method 0x4dc6a7da.
//
// Solidity: function _findNFTLevel(uint256 _nftType) view returns(uint256)
func (_Play *PlaySession) FindNFTLevel(_nftType *big.Int) (*big.Int, error) {
	return _Play.Contract.FindNFTLevel(&_Play.CallOpts, _nftType)
}

// FindNFTLevel is a free data retrieval call binding the contract method 0x4dc6a7da.
//
// Solidity: function _findNFTLevel(uint256 _nftType) view returns(uint256)
func (_Play *PlayCallerSession) FindNFTLevel(_nftType *big.Int) (*big.Int, error) {
	return _Play.Contract.FindNFTLevel(&_Play.CallOpts, _nftType)
}

// FindNFTPool is a free data retrieval call binding the contract method 0x02f441b6.
//
// Solidity: function _findNFTPool(uint256 _nftType, uint256 _level) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,bool,address) _nftpool)
func (_Play *PlayCaller) FindNFTPool(opts *bind.CallOpts, _nftType *big.Int, _level *big.Int) (PlayGameNFTPool, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "_findNFTPool", _nftType, _level)

	if err != nil {
		return *new(PlayGameNFTPool), err
	}

	out0 := *abi.ConvertType(out[0], new(PlayGameNFTPool)).(*PlayGameNFTPool)

	return out0, err

}

// FindNFTPool is a free data retrieval call binding the contract method 0x02f441b6.
//
// Solidity: function _findNFTPool(uint256 _nftType, uint256 _level) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,bool,address) _nftpool)
func (_Play *PlaySession) FindNFTPool(_nftType *big.Int, _level *big.Int) (PlayGameNFTPool, error) {
	return _Play.Contract.FindNFTPool(&_Play.CallOpts, _nftType, _level)
}

// FindNFTPool is a free data retrieval call binding the contract method 0x02f441b6.
//
// Solidity: function _findNFTPool(uint256 _nftType, uint256 _level) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,string,string,bool,address) _nftpool)
func (_Play *PlayCallerSession) FindNFTPool(_nftType *big.Int, _level *big.Int) (PlayGameNFTPool, error) {
	return _Play.Contract.FindNFTPool(&_Play.CallOpts, _nftType, _level)
}

// AllNFTPool is a free data retrieval call binding the contract method 0xe5894422.
//
// Solidity: function allNFTPool(uint256 ) view returns(uint256 index, uint256 nftType, uint256 level, uint256 min, uint256 max, uint256 mintPrice, uint256 day, string nftUrl, string nftName, bool mintPause, address erc721)
func (_Play *PlayCaller) AllNFTPool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index     *big.Int
	NftType   *big.Int
	Level     *big.Int
	Min       *big.Int
	Max       *big.Int
	MintPrice *big.Int
	Day       *big.Int
	NftUrl    string
	NftName   string
	MintPause bool
	Erc721    common.Address
}, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "allNFTPool", arg0)

	outstruct := new(struct {
		Index     *big.Int
		NftType   *big.Int
		Level     *big.Int
		Min       *big.Int
		Max       *big.Int
		MintPrice *big.Int
		Day       *big.Int
		NftUrl    string
		NftName   string
		MintPause bool
		Erc721    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftType = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Min = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Max = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MintPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Day = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.NftUrl = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.NftName = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.MintPause = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.Erc721 = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AllNFTPool is a free data retrieval call binding the contract method 0xe5894422.
//
// Solidity: function allNFTPool(uint256 ) view returns(uint256 index, uint256 nftType, uint256 level, uint256 min, uint256 max, uint256 mintPrice, uint256 day, string nftUrl, string nftName, bool mintPause, address erc721)
func (_Play *PlaySession) AllNFTPool(arg0 *big.Int) (struct {
	Index     *big.Int
	NftType   *big.Int
	Level     *big.Int
	Min       *big.Int
	Max       *big.Int
	MintPrice *big.Int
	Day       *big.Int
	NftUrl    string
	NftName   string
	MintPause bool
	Erc721    common.Address
}, error) {
	return _Play.Contract.AllNFTPool(&_Play.CallOpts, arg0)
}

// AllNFTPool is a free data retrieval call binding the contract method 0xe5894422.
//
// Solidity: function allNFTPool(uint256 ) view returns(uint256 index, uint256 nftType, uint256 level, uint256 min, uint256 max, uint256 mintPrice, uint256 day, string nftUrl, string nftName, bool mintPause, address erc721)
func (_Play *PlayCallerSession) AllNFTPool(arg0 *big.Int) (struct {
	Index     *big.Int
	NftType   *big.Int
	Level     *big.Int
	Min       *big.Int
	Max       *big.Int
	MintPrice *big.Int
	Day       *big.Int
	NftUrl    string
	NftName   string
	MintPause bool
	Erc721    common.Address
}, error) {
	return _Play.Contract.AllNFTPool(&_Play.CallOpts, arg0)
}

// AllowMiningResource is a free data retrieval call binding the contract method 0x0633aeed.
//
// Solidity: function allowMiningResource(address _address, uint256 _mapIndex) view returns(bool)
func (_Play *PlayCaller) AllowMiningResource(opts *bind.CallOpts, _address common.Address, _mapIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "allowMiningResource", _address, _mapIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowMiningResource is a free data retrieval call binding the contract method 0x0633aeed.
//
// Solidity: function allowMiningResource(address _address, uint256 _mapIndex) view returns(bool)
func (_Play *PlaySession) AllowMiningResource(_address common.Address, _mapIndex *big.Int) (bool, error) {
	return _Play.Contract.AllowMiningResource(&_Play.CallOpts, _address, _mapIndex)
}

// AllowMiningResource is a free data retrieval call binding the contract method 0x0633aeed.
//
// Solidity: function allowMiningResource(address _address, uint256 _mapIndex) view returns(bool)
func (_Play *PlayCallerSession) AllowMiningResource(_address common.Address, _mapIndex *big.Int) (bool, error) {
	return _Play.Contract.AllowMiningResource(&_Play.CallOpts, _address, _mapIndex)
}

// CheckCarUnlock is a free data retrieval call binding the contract method 0x21cce05c.
//
// Solidity: function checkCarUnlock(address _address, uint256[] _carIndex) view returns(bool)
func (_Play *PlayCaller) CheckCarUnlock(opts *bind.CallOpts, _address common.Address, _carIndex []*big.Int) (bool, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "checkCarUnlock", _address, _carIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCarUnlock is a free data retrieval call binding the contract method 0x21cce05c.
//
// Solidity: function checkCarUnlock(address _address, uint256[] _carIndex) view returns(bool)
func (_Play *PlaySession) CheckCarUnlock(_address common.Address, _carIndex []*big.Int) (bool, error) {
	return _Play.Contract.CheckCarUnlock(&_Play.CallOpts, _address, _carIndex)
}

// CheckCarUnlock is a free data retrieval call binding the contract method 0x21cce05c.
//
// Solidity: function checkCarUnlock(address _address, uint256[] _carIndex) view returns(bool)
func (_Play *PlayCallerSession) CheckCarUnlock(_address common.Address, _carIndex []*big.Int) (bool, error) {
	return _Play.Contract.CheckCarUnlock(&_Play.CallOpts, _address, _carIndex)
}

// GetAllBnbGasLink is a free data retrieval call binding the contract method 0x7adba392.
//
// Solidity: function getAllBnbGasLink(uint256 _index) view returns((uint256,uint256,string) _link)
func (_Play *PlayCaller) GetAllBnbGasLink(opts *bind.CallOpts, _index *big.Int) (PlayGameBnbGasLink, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getAllBnbGasLink", _index)

	if err != nil {
		return *new(PlayGameBnbGasLink), err
	}

	out0 := *abi.ConvertType(out[0], new(PlayGameBnbGasLink)).(*PlayGameBnbGasLink)

	return out0, err

}

// GetAllBnbGasLink is a free data retrieval call binding the contract method 0x7adba392.
//
// Solidity: function getAllBnbGasLink(uint256 _index) view returns((uint256,uint256,string) _link)
func (_Play *PlaySession) GetAllBnbGasLink(_index *big.Int) (PlayGameBnbGasLink, error) {
	return _Play.Contract.GetAllBnbGasLink(&_Play.CallOpts, _index)
}

// GetAllBnbGasLink is a free data retrieval call binding the contract method 0x7adba392.
//
// Solidity: function getAllBnbGasLink(uint256 _index) view returns((uint256,uint256,string) _link)
func (_Play *PlayCallerSession) GetAllBnbGasLink(_index *big.Int) (PlayGameBnbGasLink, error) {
	return _Play.Contract.GetAllBnbGasLink(&_Play.CallOpts, _index)
}

// GetCarFromMap is a free data retrieval call binding the contract method 0x2739c866.
//
// Solidity: function getCarFromMap(address _address, uint256 _mapindex) view returns(uint256[] _indexs)
func (_Play *PlayCaller) GetCarFromMap(opts *bind.CallOpts, _address common.Address, _mapindex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getCarFromMap", _address, _mapindex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCarFromMap is a free data retrieval call binding the contract method 0x2739c866.
//
// Solidity: function getCarFromMap(address _address, uint256 _mapindex) view returns(uint256[] _indexs)
func (_Play *PlaySession) GetCarFromMap(_address common.Address, _mapindex *big.Int) ([]*big.Int, error) {
	return _Play.Contract.GetCarFromMap(&_Play.CallOpts, _address, _mapindex)
}

// GetCarFromMap is a free data retrieval call binding the contract method 0x2739c866.
//
// Solidity: function getCarFromMap(address _address, uint256 _mapindex) view returns(uint256[] _indexs)
func (_Play *PlayCallerSession) GetCarFromMap(_address common.Address, _mapindex *big.Int) ([]*big.Int, error) {
	return _Play.Contract.GetCarFromMap(&_Play.CallOpts, _address, _mapindex)
}

// GetGameParams is a free data retrieval call binding the contract method 0x7586a386.
//
// Solidity: function getGameParams() view returns(uint256 _minResourceCount, uint256 _burnRate, uint256 _carRate, uint256 _withdrawfree, uint256 _withdrawReduce)
func (_Play *PlayCaller) GetGameParams(opts *bind.CallOpts) (struct {
	MinResourceCount *big.Int
	BurnRate         *big.Int
	CarRate          *big.Int
	Withdrawfree     *big.Int
	WithdrawReduce   *big.Int
}, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getGameParams")

	outstruct := new(struct {
		MinResourceCount *big.Int
		BurnRate         *big.Int
		CarRate          *big.Int
		Withdrawfree     *big.Int
		WithdrawReduce   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinResourceCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BurnRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CarRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Withdrawfree = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.WithdrawReduce = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGameParams is a free data retrieval call binding the contract method 0x7586a386.
//
// Solidity: function getGameParams() view returns(uint256 _minResourceCount, uint256 _burnRate, uint256 _carRate, uint256 _withdrawfree, uint256 _withdrawReduce)
func (_Play *PlaySession) GetGameParams() (struct {
	MinResourceCount *big.Int
	BurnRate         *big.Int
	CarRate          *big.Int
	Withdrawfree     *big.Int
	WithdrawReduce   *big.Int
}, error) {
	return _Play.Contract.GetGameParams(&_Play.CallOpts)
}

// GetGameParams is a free data retrieval call binding the contract method 0x7586a386.
//
// Solidity: function getGameParams() view returns(uint256 _minResourceCount, uint256 _burnRate, uint256 _carRate, uint256 _withdrawfree, uint256 _withdrawReduce)
func (_Play *PlayCallerSession) GetGameParams() (struct {
	MinResourceCount *big.Int
	BurnRate         *big.Int
	CarRate          *big.Int
	Withdrawfree     *big.Int
	WithdrawReduce   *big.Int
}, error) {
	return _Play.Contract.GetGameParams(&_Play.CallOpts)
}

// GetRecoveryCarBalance is a free data retrieval call binding the contract method 0x19c97ebe.
//
// Solidity: function getRecoveryCarBalance(address _address, uint256[] _carIndex) view returns(uint256)
func (_Play *PlayCaller) GetRecoveryCarBalance(opts *bind.CallOpts, _address common.Address, _carIndex []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getRecoveryCarBalance", _address, _carIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRecoveryCarBalance is a free data retrieval call binding the contract method 0x19c97ebe.
//
// Solidity: function getRecoveryCarBalance(address _address, uint256[] _carIndex) view returns(uint256)
func (_Play *PlaySession) GetRecoveryCarBalance(_address common.Address, _carIndex []*big.Int) (*big.Int, error) {
	return _Play.Contract.GetRecoveryCarBalance(&_Play.CallOpts, _address, _carIndex)
}

// GetRecoveryCarBalance is a free data retrieval call binding the contract method 0x19c97ebe.
//
// Solidity: function getRecoveryCarBalance(address _address, uint256[] _carIndex) view returns(uint256)
func (_Play *PlayCallerSession) GetRecoveryCarBalance(_address common.Address, _carIndex []*big.Int) (*big.Int, error) {
	return _Play.Contract.GetRecoveryCarBalance(&_Play.CallOpts, _address, _carIndex)
}

// GetUserCarNFT is a free data retrieval call binding the contract method 0x589c4c19.
//
// Solidity: function getUserCarNFT(address _address) view returns((uint256,uint256,uint256,uint256,uint256,uint256,address,bool)[] _carNFT)
func (_Play *PlayCaller) GetUserCarNFT(opts *bind.CallOpts, _address common.Address) ([]PlayGameCarNFT, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getUserCarNFT", _address)

	if err != nil {
		return *new([]PlayGameCarNFT), err
	}

	out0 := *abi.ConvertType(out[0], new([]PlayGameCarNFT)).(*[]PlayGameCarNFT)

	return out0, err

}

// GetUserCarNFT is a free data retrieval call binding the contract method 0x589c4c19.
//
// Solidity: function getUserCarNFT(address _address) view returns((uint256,uint256,uint256,uint256,uint256,uint256,address,bool)[] _carNFT)
func (_Play *PlaySession) GetUserCarNFT(_address common.Address) ([]PlayGameCarNFT, error) {
	return _Play.Contract.GetUserCarNFT(&_Play.CallOpts, _address)
}

// GetUserCarNFT is a free data retrieval call binding the contract method 0x589c4c19.
//
// Solidity: function getUserCarNFT(address _address) view returns((uint256,uint256,uint256,uint256,uint256,uint256,address,bool)[] _carNFT)
func (_Play *PlayCallerSession) GetUserCarNFT(_address common.Address) ([]PlayGameCarNFT, error) {
	return _Play.Contract.GetUserCarNFT(&_Play.CallOpts, _address)
}

// GetUserIncome is a free data retrieval call binding the contract method 0x7f081b16.
//
// Solidity: function getUserIncome(address _address) view returns((uint256,uint256,uint256,uint256,bool) income)
func (_Play *PlayCaller) GetUserIncome(opts *bind.CallOpts, _address common.Address) (PlayGameUserIncome, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getUserIncome", _address)

	if err != nil {
		return *new(PlayGameUserIncome), err
	}

	out0 := *abi.ConvertType(out[0], new(PlayGameUserIncome)).(*PlayGameUserIncome)

	return out0, err

}

// GetUserIncome is a free data retrieval call binding the contract method 0x7f081b16.
//
// Solidity: function getUserIncome(address _address) view returns((uint256,uint256,uint256,uint256,bool) income)
func (_Play *PlaySession) GetUserIncome(_address common.Address) (PlayGameUserIncome, error) {
	return _Play.Contract.GetUserIncome(&_Play.CallOpts, _address)
}

// GetUserIncome is a free data retrieval call binding the contract method 0x7f081b16.
//
// Solidity: function getUserIncome(address _address) view returns((uint256,uint256,uint256,uint256,bool) income)
func (_Play *PlayCallerSession) GetUserIncome(_address common.Address) (PlayGameUserIncome, error) {
	return _Play.Contract.GetUserIncome(&_Play.CallOpts, _address)
}

// GetUserIncomeDetail is a free data retrieval call binding the contract method 0x245d36db.
//
// Solidity: function getUserIncomeDetail(address _address) view returns((uint256,uint256,uint256,uint256)[] detail)
func (_Play *PlayCaller) GetUserIncomeDetail(opts *bind.CallOpts, _address common.Address) ([]PlayGameUserIncomeDetail, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getUserIncomeDetail", _address)

	if err != nil {
		return *new([]PlayGameUserIncomeDetail), err
	}

	out0 := *abi.ConvertType(out[0], new([]PlayGameUserIncomeDetail)).(*[]PlayGameUserIncomeDetail)

	return out0, err

}

// GetUserIncomeDetail is a free data retrieval call binding the contract method 0x245d36db.
//
// Solidity: function getUserIncomeDetail(address _address) view returns((uint256,uint256,uint256,uint256)[] detail)
func (_Play *PlaySession) GetUserIncomeDetail(_address common.Address) ([]PlayGameUserIncomeDetail, error) {
	return _Play.Contract.GetUserIncomeDetail(&_Play.CallOpts, _address)
}

// GetUserIncomeDetail is a free data retrieval call binding the contract method 0x245d36db.
//
// Solidity: function getUserIncomeDetail(address _address) view returns((uint256,uint256,uint256,uint256)[] detail)
func (_Play *PlayCallerSession) GetUserIncomeDetail(_address common.Address) ([]PlayGameUserIncomeDetail, error) {
	return _Play.Contract.GetUserIncomeDetail(&_Play.CallOpts, _address)
}

// GetUserMapNFT is a free data retrieval call binding the contract method 0xb4138b11.
//
// Solidity: function getUserMapNFT(address _address) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool,address)[] _mapNFT)
func (_Play *PlayCaller) GetUserMapNFT(opts *bind.CallOpts, _address common.Address) ([]PlayGameMapNFT, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getUserMapNFT", _address)

	if err != nil {
		return *new([]PlayGameMapNFT), err
	}

	out0 := *abi.ConvertType(out[0], new([]PlayGameMapNFT)).(*[]PlayGameMapNFT)

	return out0, err

}

// GetUserMapNFT is a free data retrieval call binding the contract method 0xb4138b11.
//
// Solidity: function getUserMapNFT(address _address) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool,address)[] _mapNFT)
func (_Play *PlaySession) GetUserMapNFT(_address common.Address) ([]PlayGameMapNFT, error) {
	return _Play.Contract.GetUserMapNFT(&_Play.CallOpts, _address)
}

// GetUserMapNFT is a free data retrieval call binding the contract method 0xb4138b11.
//
// Solidity: function getUserMapNFT(address _address) view returns((uint256,uint256,uint256,uint256,uint256,uint256,bool,address)[] _mapNFT)
func (_Play *PlayCallerSession) GetUserMapNFT(_address common.Address) ([]PlayGameMapNFT, error) {
	return _Play.Contract.GetUserMapNFT(&_Play.CallOpts, _address)
}

// GetUserWithdraws is a free data retrieval call binding the contract method 0x515d223a.
//
// Solidity: function getUserWithdraws(address _address) view returns((uint256,uint256,uint256,uint256)[] detail)
func (_Play *PlayCaller) GetUserWithdraws(opts *bind.CallOpts, _address common.Address) ([]PlayGameWithdrawInfo, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getUserWithdraws", _address)

	if err != nil {
		return *new([]PlayGameWithdrawInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PlayGameWithdrawInfo)).(*[]PlayGameWithdrawInfo)

	return out0, err

}

// GetUserWithdraws is a free data retrieval call binding the contract method 0x515d223a.
//
// Solidity: function getUserWithdraws(address _address) view returns((uint256,uint256,uint256,uint256)[] detail)
func (_Play *PlaySession) GetUserWithdraws(_address common.Address) ([]PlayGameWithdrawInfo, error) {
	return _Play.Contract.GetUserWithdraws(&_Play.CallOpts, _address)
}

// GetUserWithdraws is a free data retrieval call binding the contract method 0x515d223a.
//
// Solidity: function getUserWithdraws(address _address) view returns((uint256,uint256,uint256,uint256)[] detail)
func (_Play *PlayCallerSession) GetUserWithdraws(_address common.Address) ([]PlayGameWithdrawInfo, error) {
	return _Play.Contract.GetUserWithdraws(&_Play.CallOpts, _address)
}

// GetUsersellNFT is a free data retrieval call binding the contract method 0x1f64cc1f.
//
// Solidity: function getUsersellNFT(address from, uint256 _tokenId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,address,bool) _car, (uint256,uint256,uint256,uint256,uint256,uint256,bool,address) _map)
func (_Play *PlayCaller) GetUsersellNFT(opts *bind.CallOpts, from common.Address, _tokenId *big.Int) (struct {
	Car PlayGameCarNFT
	Map PlayGameMapNFT
}, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "getUsersellNFT", from, _tokenId)

	outstruct := new(struct {
		Car PlayGameCarNFT
		Map PlayGameMapNFT
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Car = *abi.ConvertType(out[0], new(PlayGameCarNFT)).(*PlayGameCarNFT)
	outstruct.Map = *abi.ConvertType(out[1], new(PlayGameMapNFT)).(*PlayGameMapNFT)

	return *outstruct, err

}

// GetUsersellNFT is a free data retrieval call binding the contract method 0x1f64cc1f.
//
// Solidity: function getUsersellNFT(address from, uint256 _tokenId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,address,bool) _car, (uint256,uint256,uint256,uint256,uint256,uint256,bool,address) _map)
func (_Play *PlaySession) GetUsersellNFT(from common.Address, _tokenId *big.Int) (struct {
	Car PlayGameCarNFT
	Map PlayGameMapNFT
}, error) {
	return _Play.Contract.GetUsersellNFT(&_Play.CallOpts, from, _tokenId)
}

// GetUsersellNFT is a free data retrieval call binding the contract method 0x1f64cc1f.
//
// Solidity: function getUsersellNFT(address from, uint256 _tokenId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,address,bool) _car, (uint256,uint256,uint256,uint256,uint256,uint256,bool,address) _map)
func (_Play *PlayCallerSession) GetUsersellNFT(from common.Address, _tokenId *big.Int) (struct {
	Car PlayGameCarNFT
	Map PlayGameMapNFT
}, error) {
	return _Play.Contract.GetUsersellNFT(&_Play.CallOpts, from, _tokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Play *PlayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Play *PlaySession) Owner() (common.Address, error) {
	return _Play.Contract.Owner(&_Play.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Play *PlayCallerSession) Owner() (common.Address, error) {
	return _Play.Contract.Owner(&_Play.CallOpts)
}

// RandCar is a free data retrieval call binding the contract method 0x4ccfb146.
//
// Solidity: function randCar(uint256 _min, uint256 _max, uint256 _level) view returns(uint256 result)
func (_Play *PlayCaller) RandCar(opts *bind.CallOpts, _min *big.Int, _max *big.Int, _level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "randCar", _min, _max, _level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandCar is a free data retrieval call binding the contract method 0x4ccfb146.
//
// Solidity: function randCar(uint256 _min, uint256 _max, uint256 _level) view returns(uint256 result)
func (_Play *PlaySession) RandCar(_min *big.Int, _max *big.Int, _level *big.Int) (*big.Int, error) {
	return _Play.Contract.RandCar(&_Play.CallOpts, _min, _max, _level)
}

// RandCar is a free data retrieval call binding the contract method 0x4ccfb146.
//
// Solidity: function randCar(uint256 _min, uint256 _max, uint256 _level) view returns(uint256 result)
func (_Play *PlayCallerSession) RandCar(_min *big.Int, _max *big.Int, _level *big.Int) (*big.Int, error) {
	return _Play.Contract.RandCar(&_Play.CallOpts, _min, _max, _level)
}

// WithdrawPayFree is a free data retrieval call binding the contract method 0xb423fe37.
//
// Solidity: function withdrawPayFree(address _address) view returns(uint256)
func (_Play *PlayCaller) WithdrawPayFree(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Play.contract.Call(opts, &out, "withdrawPayFree", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawPayFree is a free data retrieval call binding the contract method 0xb423fe37.
//
// Solidity: function withdrawPayFree(address _address) view returns(uint256)
func (_Play *PlaySession) WithdrawPayFree(_address common.Address) (*big.Int, error) {
	return _Play.Contract.WithdrawPayFree(&_Play.CallOpts, _address)
}

// WithdrawPayFree is a free data retrieval call binding the contract method 0xb423fe37.
//
// Solidity: function withdrawPayFree(address _address) view returns(uint256)
func (_Play *PlayCallerSession) WithdrawPayFree(_address common.Address) (*big.Int, error) {
	return _Play.Contract.WithdrawPayFree(&_Play.CallOpts, _address)
}

// AddBnbGasLink is a paid mutator transaction binding the contract method 0xee145903.
//
// Solidity: function addBnbGasLink(string _linkName, uint256 _rate) returns()
func (_Play *PlayTransactor) AddBnbGasLink(opts *bind.TransactOpts, _linkName string, _rate *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "addBnbGasLink", _linkName, _rate)
}

// AddBnbGasLink is a paid mutator transaction binding the contract method 0xee145903.
//
// Solidity: function addBnbGasLink(string _linkName, uint256 _rate) returns()
func (_Play *PlaySession) AddBnbGasLink(_linkName string, _rate *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddBnbGasLink(&_Play.TransactOpts, _linkName, _rate)
}

// AddBnbGasLink is a paid mutator transaction binding the contract method 0xee145903.
//
// Solidity: function addBnbGasLink(string _linkName, uint256 _rate) returns()
func (_Play *PlayTransactorSession) AddBnbGasLink(_linkName string, _rate *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddBnbGasLink(&_Play.TransactOpts, _linkName, _rate)
}

// AddCarToMap is a paid mutator transaction binding the contract method 0xdad8195e.
//
// Solidity: function addCarToMap(uint256[] _carindex, uint256 _mapindex) returns(uint256)
func (_Play *PlayTransactor) AddCarToMap(opts *bind.TransactOpts, _carindex []*big.Int, _mapindex *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "addCarToMap", _carindex, _mapindex)
}

// AddCarToMap is a paid mutator transaction binding the contract method 0xdad8195e.
//
// Solidity: function addCarToMap(uint256[] _carindex, uint256 _mapindex) returns(uint256)
func (_Play *PlaySession) AddCarToMap(_carindex []*big.Int, _mapindex *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddCarToMap(&_Play.TransactOpts, _carindex, _mapindex)
}

// AddCarToMap is a paid mutator transaction binding the contract method 0xdad8195e.
//
// Solidity: function addCarToMap(uint256[] _carindex, uint256 _mapindex) returns(uint256)
func (_Play *PlayTransactorSession) AddCarToMap(_carindex []*big.Int, _mapindex *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddCarToMap(&_Play.TransactOpts, _carindex, _mapindex)
}

// AddNFTPool is a paid mutator transaction binding the contract method 0x2905ddd9.
//
// Solidity: function addNFTPool(string _nftName, string _nftUrl, address _erc721, uint256 _price, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max, uint256 _day) returns()
func (_Play *PlayTransactor) AddNFTPool(opts *bind.TransactOpts, _nftName string, _nftUrl string, _erc721 common.Address, _price *big.Int, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int, _day *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "addNFTPool", _nftName, _nftUrl, _erc721, _price, _nftType, _level, _min, _max, _day)
}

// AddNFTPool is a paid mutator transaction binding the contract method 0x2905ddd9.
//
// Solidity: function addNFTPool(string _nftName, string _nftUrl, address _erc721, uint256 _price, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max, uint256 _day) returns()
func (_Play *PlaySession) AddNFTPool(_nftName string, _nftUrl string, _erc721 common.Address, _price *big.Int, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int, _day *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddNFTPool(&_Play.TransactOpts, _nftName, _nftUrl, _erc721, _price, _nftType, _level, _min, _max, _day)
}

// AddNFTPool is a paid mutator transaction binding the contract method 0x2905ddd9.
//
// Solidity: function addNFTPool(string _nftName, string _nftUrl, address _erc721, uint256 _price, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max, uint256 _day) returns()
func (_Play *PlayTransactorSession) AddNFTPool(_nftName string, _nftUrl string, _erc721 common.Address, _price *big.Int, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int, _day *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddNFTPool(&_Play.TransactOpts, _nftName, _nftUrl, _erc721, _price, _nftType, _level, _min, _max, _day)
}

// AddPro is a paid mutator transaction binding the contract method 0xb92af962.
//
// Solidity: function addPro(uint256 _nftType, uint256 _level, uint256 _min, uint256 _max) returns()
func (_Play *PlayTransactor) AddPro(opts *bind.TransactOpts, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "addPro", _nftType, _level, _min, _max)
}

// AddPro is a paid mutator transaction binding the contract method 0xb92af962.
//
// Solidity: function addPro(uint256 _nftType, uint256 _level, uint256 _min, uint256 _max) returns()
func (_Play *PlaySession) AddPro(_nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddPro(&_Play.TransactOpts, _nftType, _level, _min, _max)
}

// AddPro is a paid mutator transaction binding the contract method 0xb92af962.
//
// Solidity: function addPro(uint256 _nftType, uint256 _level, uint256 _min, uint256 _max) returns()
func (_Play *PlayTransactorSession) AddPro(_nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Play.Contract.AddPro(&_Play.TransactOpts, _nftType, _level, _min, _max)
}

// Burn is a paid mutator transaction binding the contract method 0x8df2c58a.
//
// Solidity: function burn(uint256 _typeNFT, uint256 _tokenId, uint256 _level, uint256 _index) returns()
func (_Play *PlayTransactor) Burn(opts *bind.TransactOpts, _typeNFT *big.Int, _tokenId *big.Int, _level *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "burn", _typeNFT, _tokenId, _level, _index)
}

// Burn is a paid mutator transaction binding the contract method 0x8df2c58a.
//
// Solidity: function burn(uint256 _typeNFT, uint256 _tokenId, uint256 _level, uint256 _index) returns()
func (_Play *PlaySession) Burn(_typeNFT *big.Int, _tokenId *big.Int, _level *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _Play.Contract.Burn(&_Play.TransactOpts, _typeNFT, _tokenId, _level, _index)
}

// Burn is a paid mutator transaction binding the contract method 0x8df2c58a.
//
// Solidity: function burn(uint256 _typeNFT, uint256 _tokenId, uint256 _level, uint256 _index) returns()
func (_Play *PlayTransactorSession) Burn(_typeNFT *big.Int, _tokenId *big.Int, _level *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _Play.Contract.Burn(&_Play.TransactOpts, _typeNFT, _tokenId, _level, _index)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xc916311d.
//
// Solidity: function buyNFT(uint256 _nftType, address from, address to, uint256 _tokenId) returns()
func (_Play *PlayTransactor) BuyNFT(opts *bind.TransactOpts, _nftType *big.Int, from common.Address, to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "buyNFT", _nftType, from, to, _tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xc916311d.
//
// Solidity: function buyNFT(uint256 _nftType, address from, address to, uint256 _tokenId) returns()
func (_Play *PlaySession) BuyNFT(_nftType *big.Int, from common.Address, to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Play.Contract.BuyNFT(&_Play.TransactOpts, _nftType, from, to, _tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0xc916311d.
//
// Solidity: function buyNFT(uint256 _nftType, address from, address to, uint256 _tokenId) returns()
func (_Play *PlayTransactorSession) BuyNFT(_nftType *big.Int, from common.Address, to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Play.Contract.BuyNFT(&_Play.TransactOpts, _nftType, from, to, _tokenId)
}

// CanalSellNFT is a paid mutator transaction binding the contract method 0x9effbdcb.
//
// Solidity: function canalSellNFT(uint256 _nftType, uint256 _tokenId, address _address) returns()
func (_Play *PlayTransactor) CanalSellNFT(opts *bind.TransactOpts, _nftType *big.Int, _tokenId *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "canalSellNFT", _nftType, _tokenId, _address)
}

// CanalSellNFT is a paid mutator transaction binding the contract method 0x9effbdcb.
//
// Solidity: function canalSellNFT(uint256 _nftType, uint256 _tokenId, address _address) returns()
func (_Play *PlaySession) CanalSellNFT(_nftType *big.Int, _tokenId *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Play.Contract.CanalSellNFT(&_Play.TransactOpts, _nftType, _tokenId, _address)
}

// CanalSellNFT is a paid mutator transaction binding the contract method 0x9effbdcb.
//
// Solidity: function canalSellNFT(uint256 _nftType, uint256 _tokenId, address _address) returns()
func (_Play *PlayTransactorSession) CanalSellNFT(_nftType *big.Int, _tokenId *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Play.Contract.CanalSellNFT(&_Play.TransactOpts, _nftType, _tokenId, _address)
}

// DelCarToMap is a paid mutator transaction binding the contract method 0xefc3590a.
//
// Solidity: function delCarToMap(uint256[] _carindex, uint256 _mapindex) returns()
func (_Play *PlayTransactor) DelCarToMap(opts *bind.TransactOpts, _carindex []*big.Int, _mapindex *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "delCarToMap", _carindex, _mapindex)
}

// DelCarToMap is a paid mutator transaction binding the contract method 0xefc3590a.
//
// Solidity: function delCarToMap(uint256[] _carindex, uint256 _mapindex) returns()
func (_Play *PlaySession) DelCarToMap(_carindex []*big.Int, _mapindex *big.Int) (*types.Transaction, error) {
	return _Play.Contract.DelCarToMap(&_Play.TransactOpts, _carindex, _mapindex)
}

// DelCarToMap is a paid mutator transaction binding the contract method 0xefc3590a.
//
// Solidity: function delCarToMap(uint256[] _carindex, uint256 _mapindex) returns()
func (_Play *PlayTransactorSession) DelCarToMap(_carindex []*big.Int, _mapindex *big.Int) (*types.Transaction, error) {
	return _Play.Contract.DelCarToMap(&_Play.TransactOpts, _carindex, _mapindex)
}

// GiveNFT is a paid mutator transaction binding the contract method 0x00c18a6b.
//
// Solidity: function giveNFT(address erc721Token, uint256 _typeNFT, uint256 _index, uint256 _tokenId, address _recvAddress) payable returns()
func (_Play *PlayTransactor) GiveNFT(opts *bind.TransactOpts, erc721Token common.Address, _typeNFT *big.Int, _index *big.Int, _tokenId *big.Int, _recvAddress common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "giveNFT", erc721Token, _typeNFT, _index, _tokenId, _recvAddress)
}

// GiveNFT is a paid mutator transaction binding the contract method 0x00c18a6b.
//
// Solidity: function giveNFT(address erc721Token, uint256 _typeNFT, uint256 _index, uint256 _tokenId, address _recvAddress) payable returns()
func (_Play *PlaySession) GiveNFT(erc721Token common.Address, _typeNFT *big.Int, _index *big.Int, _tokenId *big.Int, _recvAddress common.Address) (*types.Transaction, error) {
	return _Play.Contract.GiveNFT(&_Play.TransactOpts, erc721Token, _typeNFT, _index, _tokenId, _recvAddress)
}

// GiveNFT is a paid mutator transaction binding the contract method 0x00c18a6b.
//
// Solidity: function giveNFT(address erc721Token, uint256 _typeNFT, uint256 _index, uint256 _tokenId, address _recvAddress) payable returns()
func (_Play *PlayTransactorSession) GiveNFT(erc721Token common.Address, _typeNFT *big.Int, _index *big.Int, _tokenId *big.Int, _recvAddress common.Address) (*types.Transaction, error) {
	return _Play.Contract.GiveNFT(&_Play.TransactOpts, erc721Token, _typeNFT, _index, _tokenId, _recvAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Play *PlayTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Play *PlaySession) Initialize() (*types.Transaction, error) {
	return _Play.Contract.Initialize(&_Play.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Play *PlayTransactorSession) Initialize() (*types.Transaction, error) {
	return _Play.Contract.Initialize(&_Play.TransactOpts)
}

// LoginInviteCodes is a paid mutator transaction binding the contract method 0x7902651a.
//
// Solidity: function loginInviteCodes(string _code) returns(uint256, uint256)
func (_Play *PlayTransactor) LoginInviteCodes(opts *bind.TransactOpts, _code string) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "loginInviteCodes", _code)
}

// LoginInviteCodes is a paid mutator transaction binding the contract method 0x7902651a.
//
// Solidity: function loginInviteCodes(string _code) returns(uint256, uint256)
func (_Play *PlaySession) LoginInviteCodes(_code string) (*types.Transaction, error) {
	return _Play.Contract.LoginInviteCodes(&_Play.TransactOpts, _code)
}

// LoginInviteCodes is a paid mutator transaction binding the contract method 0x7902651a.
//
// Solidity: function loginInviteCodes(string _code) returns(uint256, uint256)
func (_Play *PlayTransactorSession) LoginInviteCodes(_code string) (*types.Transaction, error) {
	return _Play.Contract.LoginInviteCodes(&_Play.TransactOpts, _code)
}

// MiningResource is a paid mutator transaction binding the contract method 0x56845b10.
//
// Solidity: function miningResource(uint256[] _carIndex, uint256 _mapIndex) returns(uint256)
func (_Play *PlayTransactor) MiningResource(opts *bind.TransactOpts, _carIndex []*big.Int, _mapIndex *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "miningResource", _carIndex, _mapIndex)
}

// MiningResource is a paid mutator transaction binding the contract method 0x56845b10.
//
// Solidity: function miningResource(uint256[] _carIndex, uint256 _mapIndex) returns(uint256)
func (_Play *PlaySession) MiningResource(_carIndex []*big.Int, _mapIndex *big.Int) (*types.Transaction, error) {
	return _Play.Contract.MiningResource(&_Play.TransactOpts, _carIndex, _mapIndex)
}

// MiningResource is a paid mutator transaction binding the contract method 0x56845b10.
//
// Solidity: function miningResource(uint256[] _carIndex, uint256 _mapIndex) returns(uint256)
func (_Play *PlayTransactorSession) MiningResource(_carIndex []*big.Int, _mapIndex *big.Int) (*types.Transaction, error) {
	return _Play.Contract.MiningResource(&_Play.TransactOpts, _carIndex, _mapIndex)
}

// Mint is a paid mutator transaction binding the contract method 0xbf2c8a94.
//
// Solidity: function mint(uint256 _typeNFT, address to, uint256 _subMethod, uint256 _balance) returns()
func (_Play *PlayTransactor) Mint(opts *bind.TransactOpts, _typeNFT *big.Int, to common.Address, _subMethod *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "mint", _typeNFT, to, _subMethod, _balance)
}

// Mint is a paid mutator transaction binding the contract method 0xbf2c8a94.
//
// Solidity: function mint(uint256 _typeNFT, address to, uint256 _subMethod, uint256 _balance) returns()
func (_Play *PlaySession) Mint(_typeNFT *big.Int, to common.Address, _subMethod *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Play.Contract.Mint(&_Play.TransactOpts, _typeNFT, to, _subMethod, _balance)
}

// Mint is a paid mutator transaction binding the contract method 0xbf2c8a94.
//
// Solidity: function mint(uint256 _typeNFT, address to, uint256 _subMethod, uint256 _balance) returns()
func (_Play *PlayTransactorSession) Mint(_typeNFT *big.Int, to common.Address, _subMethod *big.Int, _balance *big.Int) (*types.Transaction, error) {
	return _Play.Contract.Mint(&_Play.TransactOpts, _typeNFT, to, _subMethod, _balance)
}

// MintOne is a paid mutator transaction binding the contract method 0xbd6056f7.
//
// Solidity: function mintOne(uint256 _typeNFT, address to) returns(uint256)
func (_Play *PlayTransactor) MintOne(opts *bind.TransactOpts, _typeNFT *big.Int, to common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "mintOne", _typeNFT, to)
}

// MintOne is a paid mutator transaction binding the contract method 0xbd6056f7.
//
// Solidity: function mintOne(uint256 _typeNFT, address to) returns(uint256)
func (_Play *PlaySession) MintOne(_typeNFT *big.Int, to common.Address) (*types.Transaction, error) {
	return _Play.Contract.MintOne(&_Play.TransactOpts, _typeNFT, to)
}

// MintOne is a paid mutator transaction binding the contract method 0xbd6056f7.
//
// Solidity: function mintOne(uint256 _typeNFT, address to) returns(uint256)
func (_Play *PlayTransactorSession) MintOne(_typeNFT *big.Int, to common.Address) (*types.Transaction, error) {
	return _Play.Contract.MintOne(&_Play.TransactOpts, _typeNFT, to)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Play *PlayTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Play *PlaySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Play.Contract.OnERC721Received(&_Play.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Play *PlayTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Play.Contract.OnERC721Received(&_Play.TransactOpts, arg0, arg1, arg2, arg3)
}

// RecoveryCar is a paid mutator transaction binding the contract method 0x11a04d6d.
//
// Solidity: function recoveryCar(uint256[] _carIndex, uint256 _subMethod, uint256 _amount) returns()
func (_Play *PlayTransactor) RecoveryCar(opts *bind.TransactOpts, _carIndex []*big.Int, _subMethod *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "recoveryCar", _carIndex, _subMethod, _amount)
}

// RecoveryCar is a paid mutator transaction binding the contract method 0x11a04d6d.
//
// Solidity: function recoveryCar(uint256[] _carIndex, uint256 _subMethod, uint256 _amount) returns()
func (_Play *PlaySession) RecoveryCar(_carIndex []*big.Int, _subMethod *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Play.Contract.RecoveryCar(&_Play.TransactOpts, _carIndex, _subMethod, _amount)
}

// RecoveryCar is a paid mutator transaction binding the contract method 0x11a04d6d.
//
// Solidity: function recoveryCar(uint256[] _carIndex, uint256 _subMethod, uint256 _amount) returns()
func (_Play *PlayTransactorSession) RecoveryCar(_carIndex []*big.Int, _subMethod *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Play.Contract.RecoveryCar(&_Play.TransactOpts, _carIndex, _subMethod, _amount)
}

// RegisterInviteCodes is a paid mutator transaction binding the contract method 0xf86be3b6.
//
// Solidity: function registerInviteCodes(string _code) returns()
func (_Play *PlayTransactor) RegisterInviteCodes(opts *bind.TransactOpts, _code string) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "registerInviteCodes", _code)
}

// RegisterInviteCodes is a paid mutator transaction binding the contract method 0xf86be3b6.
//
// Solidity: function registerInviteCodes(string _code) returns()
func (_Play *PlaySession) RegisterInviteCodes(_code string) (*types.Transaction, error) {
	return _Play.Contract.RegisterInviteCodes(&_Play.TransactOpts, _code)
}

// RegisterInviteCodes is a paid mutator transaction binding the contract method 0xf86be3b6.
//
// Solidity: function registerInviteCodes(string _code) returns()
func (_Play *PlayTransactorSession) RegisterInviteCodes(_code string) (*types.Transaction, error) {
	return _Play.Contract.RegisterInviteCodes(&_Play.TransactOpts, _code)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Play *PlayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Play *PlaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Play.Contract.RenounceOwnership(&_Play.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Play *PlayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Play.Contract.RenounceOwnership(&_Play.TransactOpts)
}

// SellNFT is a paid mutator transaction binding the contract method 0x7c237eea.
//
// Solidity: function sellNFT(uint256 _nftType, uint256 _tokenId, address _address) returns()
func (_Play *PlayTransactor) SellNFT(opts *bind.TransactOpts, _nftType *big.Int, _tokenId *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "sellNFT", _nftType, _tokenId, _address)
}

// SellNFT is a paid mutator transaction binding the contract method 0x7c237eea.
//
// Solidity: function sellNFT(uint256 _nftType, uint256 _tokenId, address _address) returns()
func (_Play *PlaySession) SellNFT(_nftType *big.Int, _tokenId *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Play.Contract.SellNFT(&_Play.TransactOpts, _nftType, _tokenId, _address)
}

// SellNFT is a paid mutator transaction binding the contract method 0x7c237eea.
//
// Solidity: function sellNFT(uint256 _nftType, uint256 _tokenId, address _address) returns()
func (_Play *PlayTransactorSession) SellNFT(_nftType *big.Int, _tokenId *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Play.Contract.SellNFT(&_Play.TransactOpts, _nftType, _tokenId, _address)
}

// SetCodeIQLF is a paid mutator transaction binding the contract method 0x03aa4d91.
//
// Solidity: function setCodeIQLF(address _codeIqlf) returns()
func (_Play *PlayTransactor) SetCodeIQLF(opts *bind.TransactOpts, _codeIqlf common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "setCodeIQLF", _codeIqlf)
}

// SetCodeIQLF is a paid mutator transaction binding the contract method 0x03aa4d91.
//
// Solidity: function setCodeIQLF(address _codeIqlf) returns()
func (_Play *PlaySession) SetCodeIQLF(_codeIqlf common.Address) (*types.Transaction, error) {
	return _Play.Contract.SetCodeIQLF(&_Play.TransactOpts, _codeIqlf)
}

// SetCodeIQLF is a paid mutator transaction binding the contract method 0x03aa4d91.
//
// Solidity: function setCodeIQLF(address _codeIqlf) returns()
func (_Play *PlayTransactorSession) SetCodeIQLF(_codeIqlf common.Address) (*types.Transaction, error) {
	return _Play.Contract.SetCodeIQLF(&_Play.TransactOpts, _codeIqlf)
}

// SetErc20Token is a paid mutator transaction binding the contract method 0x26e5cb69.
//
// Solidity: function setErc20Token(address _erc20Token) returns()
func (_Play *PlayTransactor) SetErc20Token(opts *bind.TransactOpts, _erc20Token common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "setErc20Token", _erc20Token)
}

// SetErc20Token is a paid mutator transaction binding the contract method 0x26e5cb69.
//
// Solidity: function setErc20Token(address _erc20Token) returns()
func (_Play *PlaySession) SetErc20Token(_erc20Token common.Address) (*types.Transaction, error) {
	return _Play.Contract.SetErc20Token(&_Play.TransactOpts, _erc20Token)
}

// SetErc20Token is a paid mutator transaction binding the contract method 0x26e5cb69.
//
// Solidity: function setErc20Token(address _erc20Token) returns()
func (_Play *PlayTransactorSession) SetErc20Token(_erc20Token common.Address) (*types.Transaction, error) {
	return _Play.Contract.SetErc20Token(&_Play.TransactOpts, _erc20Token)
}

// SetMiningInterval is a paid mutator transaction binding the contract method 0xfe11e77a.
//
// Solidity: function setMiningInterval(uint256 _interval, uint256 _bnbgas) returns()
func (_Play *PlayTransactor) SetMiningInterval(opts *bind.TransactOpts, _interval *big.Int, _bnbgas *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "setMiningInterval", _interval, _bnbgas)
}

// SetMiningInterval is a paid mutator transaction binding the contract method 0xfe11e77a.
//
// Solidity: function setMiningInterval(uint256 _interval, uint256 _bnbgas) returns()
func (_Play *PlaySession) SetMiningInterval(_interval *big.Int, _bnbgas *big.Int) (*types.Transaction, error) {
	return _Play.Contract.SetMiningInterval(&_Play.TransactOpts, _interval, _bnbgas)
}

// SetMiningInterval is a paid mutator transaction binding the contract method 0xfe11e77a.
//
// Solidity: function setMiningInterval(uint256 _interval, uint256 _bnbgas) returns()
func (_Play *PlayTransactorSession) SetMiningInterval(_interval *big.Int, _bnbgas *big.Int) (*types.Transaction, error) {
	return _Play.Contract.SetMiningInterval(&_Play.TransactOpts, _interval, _bnbgas)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0x8590208d.
//
// Solidity: function setTotalSupply(uint256 _totalSupply, uint256 _minResourceCount, uint256 _burnRate, uint256 _carRate) returns()
func (_Play *PlayTransactor) SetTotalSupply(opts *bind.TransactOpts, _totalSupply *big.Int, _minResourceCount *big.Int, _burnRate *big.Int, _carRate *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "setTotalSupply", _totalSupply, _minResourceCount, _burnRate, _carRate)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0x8590208d.
//
// Solidity: function setTotalSupply(uint256 _totalSupply, uint256 _minResourceCount, uint256 _burnRate, uint256 _carRate) returns()
func (_Play *PlaySession) SetTotalSupply(_totalSupply *big.Int, _minResourceCount *big.Int, _burnRate *big.Int, _carRate *big.Int) (*types.Transaction, error) {
	return _Play.Contract.SetTotalSupply(&_Play.TransactOpts, _totalSupply, _minResourceCount, _burnRate, _carRate)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0x8590208d.
//
// Solidity: function setTotalSupply(uint256 _totalSupply, uint256 _minResourceCount, uint256 _burnRate, uint256 _carRate) returns()
func (_Play *PlayTransactorSession) SetTotalSupply(_totalSupply *big.Int, _minResourceCount *big.Int, _burnRate *big.Int, _carRate *big.Int) (*types.Transaction, error) {
	return _Play.Contract.SetTotalSupply(&_Play.TransactOpts, _totalSupply, _minResourceCount, _burnRate, _carRate)
}

// SetWithdrawfree is a paid mutator transaction binding the contract method 0xabc22ea0.
//
// Solidity: function setWithdrawfree(uint256 _free, uint256 _reduce) returns()
func (_Play *PlayTransactor) SetWithdrawfree(opts *bind.TransactOpts, _free *big.Int, _reduce *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "setWithdrawfree", _free, _reduce)
}

// SetWithdrawfree is a paid mutator transaction binding the contract method 0xabc22ea0.
//
// Solidity: function setWithdrawfree(uint256 _free, uint256 _reduce) returns()
func (_Play *PlaySession) SetWithdrawfree(_free *big.Int, _reduce *big.Int) (*types.Transaction, error) {
	return _Play.Contract.SetWithdrawfree(&_Play.TransactOpts, _free, _reduce)
}

// SetWithdrawfree is a paid mutator transaction binding the contract method 0xabc22ea0.
//
// Solidity: function setWithdrawfree(uint256 _free, uint256 _reduce) returns()
func (_Play *PlayTransactorSession) SetWithdrawfree(_free *big.Int, _reduce *big.Int) (*types.Transaction, error) {
	return _Play.Contract.SetWithdrawfree(&_Play.TransactOpts, _free, _reduce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Play *PlayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Play *PlaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Play.Contract.TransferOwnership(&_Play.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Play *PlayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Play.Contract.TransferOwnership(&_Play.TransactOpts, newOwner)
}

// UpdateBnbGasLink is a paid mutator transaction binding the contract method 0x4a240caf.
//
// Solidity: function updateBnbGasLink(uint256 _index, uint256 _rate) returns()
func (_Play *PlayTransactor) UpdateBnbGasLink(opts *bind.TransactOpts, _index *big.Int, _rate *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "updateBnbGasLink", _index, _rate)
}

// UpdateBnbGasLink is a paid mutator transaction binding the contract method 0x4a240caf.
//
// Solidity: function updateBnbGasLink(uint256 _index, uint256 _rate) returns()
func (_Play *PlaySession) UpdateBnbGasLink(_index *big.Int, _rate *big.Int) (*types.Transaction, error) {
	return _Play.Contract.UpdateBnbGasLink(&_Play.TransactOpts, _index, _rate)
}

// UpdateBnbGasLink is a paid mutator transaction binding the contract method 0x4a240caf.
//
// Solidity: function updateBnbGasLink(uint256 _index, uint256 _rate) returns()
func (_Play *PlayTransactorSession) UpdateBnbGasLink(_index *big.Int, _rate *big.Int) (*types.Transaction, error) {
	return _Play.Contract.UpdateBnbGasLink(&_Play.TransactOpts, _index, _rate)
}

// UpdateNFTPoolInfo is a paid mutator transaction binding the contract method 0x30700e82.
//
// Solidity: function updateNFTPoolInfo(uint256 _index, string _nftName, string _nftUrl, uint256 _mintPrice, bool _mintPause, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max, uint256 _day, address _erc721) returns()
func (_Play *PlayTransactor) UpdateNFTPoolInfo(opts *bind.TransactOpts, _index *big.Int, _nftName string, _nftUrl string, _mintPrice *big.Int, _mintPause bool, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int, _day *big.Int, _erc721 common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "updateNFTPoolInfo", _index, _nftName, _nftUrl, _mintPrice, _mintPause, _nftType, _level, _min, _max, _day, _erc721)
}

// UpdateNFTPoolInfo is a paid mutator transaction binding the contract method 0x30700e82.
//
// Solidity: function updateNFTPoolInfo(uint256 _index, string _nftName, string _nftUrl, uint256 _mintPrice, bool _mintPause, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max, uint256 _day, address _erc721) returns()
func (_Play *PlaySession) UpdateNFTPoolInfo(_index *big.Int, _nftName string, _nftUrl string, _mintPrice *big.Int, _mintPause bool, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int, _day *big.Int, _erc721 common.Address) (*types.Transaction, error) {
	return _Play.Contract.UpdateNFTPoolInfo(&_Play.TransactOpts, _index, _nftName, _nftUrl, _mintPrice, _mintPause, _nftType, _level, _min, _max, _day, _erc721)
}

// UpdateNFTPoolInfo is a paid mutator transaction binding the contract method 0x30700e82.
//
// Solidity: function updateNFTPoolInfo(uint256 _index, string _nftName, string _nftUrl, uint256 _mintPrice, bool _mintPause, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max, uint256 _day, address _erc721) returns()
func (_Play *PlayTransactorSession) UpdateNFTPoolInfo(_index *big.Int, _nftName string, _nftUrl string, _mintPrice *big.Int, _mintPause bool, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int, _day *big.Int, _erc721 common.Address) (*types.Transaction, error) {
	return _Play.Contract.UpdateNFTPoolInfo(&_Play.TransactOpts, _index, _nftName, _nftUrl, _mintPrice, _mintPause, _nftType, _level, _min, _max, _day, _erc721)
}

// UpdatePro is a paid mutator transaction binding the contract method 0x4bbcb66e.
//
// Solidity: function updatePro(uint256 _index, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max) returns()
func (_Play *PlayTransactor) UpdatePro(opts *bind.TransactOpts, _index *big.Int, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "updatePro", _index, _nftType, _level, _min, _max)
}

// UpdatePro is a paid mutator transaction binding the contract method 0x4bbcb66e.
//
// Solidity: function updatePro(uint256 _index, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max) returns()
func (_Play *PlaySession) UpdatePro(_index *big.Int, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Play.Contract.UpdatePro(&_Play.TransactOpts, _index, _nftType, _level, _min, _max)
}

// UpdatePro is a paid mutator transaction binding the contract method 0x4bbcb66e.
//
// Solidity: function updatePro(uint256 _index, uint256 _nftType, uint256 _level, uint256 _min, uint256 _max) returns()
func (_Play *PlayTransactorSession) UpdatePro(_index *big.Int, _nftType *big.Int, _level *big.Int, _min *big.Int, _max *big.Int) (*types.Transaction, error) {
	return _Play.Contract.UpdatePro(&_Play.TransactOpts, _index, _nftType, _level, _min, _max)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address _receiver, address _token) returns()
func (_Play *PlayTransactor) Withdraw(opts *bind.TransactOpts, _receiver common.Address, _token common.Address) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "withdraw", _receiver, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address _receiver, address _token) returns()
func (_Play *PlaySession) Withdraw(_receiver common.Address, _token common.Address) (*types.Transaction, error) {
	return _Play.Contract.Withdraw(&_Play.TransactOpts, _receiver, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(address _receiver, address _token) returns()
func (_Play *PlayTransactorSession) Withdraw(_receiver common.Address, _token common.Address) (*types.Transaction, error) {
	return _Play.Contract.Withdraw(&_Play.TransactOpts, _receiver, _token)
}

// WithdrawFree is a paid mutator transaction binding the contract method 0x3bf7284b.
//
// Solidity: function withdrawFree(address _receiver, uint256 _balance) returns()
func (_Play *PlayTransactor) WithdrawFree(opts *bind.TransactOpts, _receiver common.Address, _balance *big.Int) (*types.Transaction, error) {
	return _Play.contract.Transact(opts, "withdrawFree", _receiver, _balance)
}

// WithdrawFree is a paid mutator transaction binding the contract method 0x3bf7284b.
//
// Solidity: function withdrawFree(address _receiver, uint256 _balance) returns()
func (_Play *PlaySession) WithdrawFree(_receiver common.Address, _balance *big.Int) (*types.Transaction, error) {
	return _Play.Contract.WithdrawFree(&_Play.TransactOpts, _receiver, _balance)
}

// WithdrawFree is a paid mutator transaction binding the contract method 0x3bf7284b.
//
// Solidity: function withdrawFree(address _receiver, uint256 _balance) returns()
func (_Play *PlayTransactorSession) WithdrawFree(_receiver common.Address, _balance *big.Int) (*types.Transaction, error) {
	return _Play.Contract.WithdrawFree(&_Play.TransactOpts, _receiver, _balance)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Play *PlayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Play.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Play *PlaySession) Receive() (*types.Transaction, error) {
	return _Play.Contract.Receive(&_Play.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Play *PlayTransactorSession) Receive() (*types.Transaction, error) {
	return _Play.Contract.Receive(&_Play.TransactOpts)
}

// PlayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Play contract.
type PlayOwnershipTransferredIterator struct {
	Event *PlayOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlayOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlayOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlayOwnershipTransferred represents a OwnershipTransferred event raised by the Play contract.
type PlayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Play *PlayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Play.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlayOwnershipTransferredIterator{contract: _Play.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Play *PlayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Play.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlayOwnershipTransferred)
				if err := _Play.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Play *PlayFilterer) ParseOwnershipTransferred(log types.Log) (*PlayOwnershipTransferred, error) {
	event := new(PlayOwnershipTransferred)
	if err := _Play.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlayPLAYNFTIterator is returned from FilterPLAYNFT and is used to iterate over the raw logs and unpacked data for PLAYNFT events raised by the Play contract.
type PlayPLAYNFTIterator struct {
	Event *PlayPLAYNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlayPLAYNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlayPLAYNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlayPLAYNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlayPLAYNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlayPLAYNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlayPLAYNFT represents a PLAYNFT event raised by the Play contract.
type PlayPLAYNFT struct {
	Power   *big.Int
	Level   *big.Int
	TokenId *big.Int
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPLAYNFT is a free log retrieval operation binding the contract event 0x4317e9d8b46117e4e5cf4c041b87acdc477326074d912d2add0d9074731edc45.
//
// Solidity: event PLAYNFT(uint256 power, uint256 level, uint256 tokenId, address indexed user)
func (_Play *PlayFilterer) FilterPLAYNFT(opts *bind.FilterOpts, user []common.Address) (*PlayPLAYNFTIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Play.contract.FilterLogs(opts, "PLAYNFT", userRule)
	if err != nil {
		return nil, err
	}
	return &PlayPLAYNFTIterator{contract: _Play.contract, event: "PLAYNFT", logs: logs, sub: sub}, nil
}

// WatchPLAYNFT is a free log subscription operation binding the contract event 0x4317e9d8b46117e4e5cf4c041b87acdc477326074d912d2add0d9074731edc45.
//
// Solidity: event PLAYNFT(uint256 power, uint256 level, uint256 tokenId, address indexed user)
func (_Play *PlayFilterer) WatchPLAYNFT(opts *bind.WatchOpts, sink chan<- *PlayPLAYNFT, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Play.contract.WatchLogs(opts, "PLAYNFT", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlayPLAYNFT)
				if err := _Play.contract.UnpackLog(event, "PLAYNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePLAYNFT is a log parse operation binding the contract event 0x4317e9d8b46117e4e5cf4c041b87acdc477326074d912d2add0d9074731edc45.
//
// Solidity: event PLAYNFT(uint256 power, uint256 level, uint256 tokenId, address indexed user)
func (_Play *PlayFilterer) ParsePLAYNFT(log types.Log) (*PlayPLAYNFT, error) {
	event := new(PlayPLAYNFT)
	if err := _Play.contract.UnpackLog(event, "PLAYNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlayWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Play contract.
type PlayWithdrawIterator struct {
	Event *PlayWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlayWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlayWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlayWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlayWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlayWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlayWithdraw represents a Withdraw event raised by the Play contract.
type PlayWithdraw struct {
	Index  *big.Int
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 index, address indexed user, uint256 amount)
func (_Play *PlayFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*PlayWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Play.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PlayWithdrawIterator{contract: _Play.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 index, address indexed user, uint256 amount)
func (_Play *PlayFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PlayWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Play.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlayWithdraw)
				if err := _Play.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x9da6493a92039daf47d1f2d7a782299c5994c6323eb1e972f69c432089ec52bf.
//
// Solidity: event Withdraw(uint256 index, address indexed user, uint256 amount)
func (_Play *PlayFilterer) ParseWithdraw(log types.Log) (*PlayWithdraw, error) {
	event := new(PlayWithdraw)
	if err := _Play.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
