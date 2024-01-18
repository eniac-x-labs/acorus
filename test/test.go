package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type A interface {
	StartUnpack()
}

type B struct {
	b string
}

func (b *B) StartUnpack() {
	fmt.Println("B: Start" + b.b)
	// B 的 Start 方法逻辑
}

type C struct {
	// C 的字段
	c string
}

func (c *C) StartUnpack() {
	fmt.Println("C: Start" + c.c)
	c.ss()
	// C 的 Start 方法逻辑
}

func (c *C) ss() {
	fmt.Println("ss")
}

func main() {
	//objMap := make(map[string]A)
	//
	//b := &B{"123"}
	//c := &C{"456"}
	//
	//objMap["B"] = b
	//objMap["C"] = c
	//
	//for _, obj := range objMap {
	//	obj.StartUnpack()
	//}
	eventLog := "0xf89b94dac17f958d2ee523a2206206994597c13d831ec7f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa000000000000000000000000056eddb7aa87536c09ccc2793473599fd21a8b17fa0000000000000000000000000d7e1cde6e2cd5c762197211e8306ba1b93af12f2a0000000000000000000000000000000000000000000000000000000001113e0ba"
	b, err := hexutil.Decode(eventLog)
	if err != nil {
		fmt.Println(err)
	}

	var result types.Log
	if err := rlp.DecodeBytes(b, &result); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.BlockNumber)

	headerstr := "0xf90242f9023da08aa14011256323e181a81c725e5ece9999d880583dce39dca2533b0f370715d6a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d4934794de1e78c6ea6ca1ff6e7ced4da1dd936609051964a0c9a5867e19953ee5e8a0f67c087a2e210f60de5c61e0023bb63b6d9f452cff33a072d717409b86bfa582bf314f3794a47ac90cdc8ea799ef94194c01b049ac0eada087d73dcc3dd42521e94fe31c67385565c02d905f87da2c383a0e3947b664c409b90100f46d4244e595498fe092a3aab318417da9c96a80cc14d85241ab395ebfa15922dfe5df2a9030a6197cb34af0b91c4b0e2b918e0bbe82b4be9d23c2ee8a3ef14e60e057ccf81a58b97cab5afdcabda17cb225a56807441c1b4aade55d886699b2fa695e82d60e87ab61398be3189238f7803dda2b48fc7f9db2464d153c1969545e8ac35d9184582989c91caf4a80f832552f74b7aba7b3f9728dae57e2999e0caae5df66f953f7363afccad3547c455e461f9474d723af036b89d12cb3830a7591a1a59688590e93455168679177fef92def035ab0db479ddf25ce3ac2a2a7ec4cbeea08085e993f0e57452efb5275372442fd62601d71e7c3addf5e26b2641f808401176a1c8401c9c38084010cbb7e846523b8ff99d883010c02846765746888676f312e32302e37856c696e7578a039c7d2343fc44cf309435e16a6e871745457ed9c2cdc477e770044556b3fed948800000000000000008501f71a1d62a0cdfab9e9605f440d18a1945bfe9b05d668e3a0e27ceb49de5610ab06d13789e2c0c0"
	h, _ := hexutil.Decode(headerstr)
	fmt.Println(string(h))

	var header types.Header

	rlp.DecodeBytes(h, &header)
	json, _ := header.MarshalJSON()
	fmt.Println(string(json))
	//hexutil.Encode()
	fmt.Println(header.Number)
}
