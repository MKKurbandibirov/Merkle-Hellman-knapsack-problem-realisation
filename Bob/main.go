package main

import (
	"cursach/bob"
	"fmt"
	"log"
	"math/big"
	"net"
	"strconv"
	"strings"
)

// func createBasis(publicKey []*big.Int, elem *big.Int, length int) [][]*big.Int {
// 	Basis := make([][]*big.Int, length)
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < length; j++ {
// 			tmp := big.NewInt(0)
// 			if j == length - 1 {
// 				if i == j {
// 					tmp.Set(elem)
// 				} else {
// 					tmp.Set(big.NewInt(0).Neg(publicKey[i]))
// 				}
// 			} else if i == j{
// 				tmp.Set(big.NewInt(1))
// 			}
// 			Basis[i] = append(Basis[i], tmp)
// 		}
// 	}
// 	return Basis
// }

// func scalarProduct(a, b []*big.Float, length int) *big.Float{
// 	var res *big.Float = big.NewFloat(0)
// 	for i := 0; i < length; i++ {
// 		res.Add(res, big.NewFloat(0).Mul(a[i], b[i]))
// 	}
// 	return res
// }

// func gramSchmidt(Basis [][]*big.Int, length int) ([][]*big.Float, [][]*big.Float){
// 	var rbs [][]*big.Float
// 	var mus [][]*big.Float

// 	for i := 0; i < length; i++ {
// 		tmp_rb := make([]*big.Float, length)
// 		for j := 0; j < length; j++ {
// 			tmp_rb[j] = big.NewFloat(0).Set(big.NewFloat(0).SetInt(Basis[i][j]))
// 		}
// 		var tmp_mu []*big.Float = make([]*big.Float, 0)

// 		for _, r := range rbs {
// 			coef := big.NewFloat(0).Quo(scalarProduct(r, tmp_rb, length), scalarProduct(r, r, length))
// 			for k, r_k := range r {
// 				tmp_rb[k].Sub(tmp_rb[k], big.NewFloat(0).Mul(coef, r_k))
// 			}
// 			tmp_mu = append(tmp_mu, coef)
// 		}
// 		rbs = append(rbs, tmp_rb)
// 		mus = append(mus, tmp_mu)
// 	}
// 	return rbs, mus
// }

// func copyMatrix(Matrix [][]*big.Int, length int) [][]*big.Int {
// 	var tmp [][]*big.Int = make([][]*big.Int, length)
// 	for i := 0; i < length; i++ {
// 		tmp[i] = make([]*big.Int, length)
// 		for j := 0; j < length; j++ {
// 			tmp[i][j] = big.NewInt(0).Set(Matrix[i][j])
// 		}
// 	}
// 	return tmp
// }

// func findBadInd(gsh, mus [][]*big.Float, delta float64, length int) int {
// 	for i := 0; i < length - 1; i++ {
// 		lhs := new(big.Float).Mul(scalarProduct(gsh[i], gsh[i], length), new(big.Float).SetFloat64(delta))
// 		rhs := new(big.Float).Add(new(big.Float).Mul(scalarProduct(gsh[i], gsh[i], length),
// 			new(big.Float).Mul(mus[i + 1][i], mus[i + 1][i])), scalarProduct(gsh[i + 1], gsh[i + 1], length))
// 		if rhs.Cmp(lhs) == 1 {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func deltaLLLSwap(lll_b [][]*big.Int, gsh, mus [][]*big.Float, delta float64, length int) bool {
// 	idx := findBadInd(gsh, mus, delta, length)
// 	if idx == -1 {
// 		return true
// 	}
// 	lll_b[idx], lll_b[idx + 1] = lll_b[idx + 1], lll_b[idx]
// 	return false
// }

// func deltaLLLReduce(Basis [][]*big.Int, mus [][]*big.Float, length int) ([][]*big.Int, [][]*big.Float) {
// 	for i := 1; i < length; i++ {
// 		for j := i - 1; j >= 0; j-- {
// 			c, _ := big.NewFloat(0).Add(mus[i][j], new(big.Float).SetFloat64(0.5*float64(mus[i][j].Sign()))).Int(nil)
// 			for ind, num := range Basis[j] {
// 				Basis[i][ind].Sub(Basis[i][ind], new(big.Int).Mul(c, num))
// 			}
// 			mus[i][j].Sub(mus[i][j], new(big.Float).SetInt(c))
// 			for k := 0; k < j; k++ {
// 				mus[i][k].Sub(mus[i][k], new(big.Float).Mul(new(big.Float).SetInt(c), mus[j][k]))
// 			}
// 		}
// 	}
// 	return Basis, mus
// }

// func deltaLLL(Basis [][]*big.Int, delta float64, length int) [][]*big.Int {
// 	alright := false
// 	var lll_b [][]*big.Int = copyMatrix(Basis, length)

// 	for alright {
// 		gsh, mus := gramSchmidt(lll_b, length)
// 		lll_b, mus = deltaLLLReduce(lll_b, mus, length)
// 		alright = deltaLLLSwap(lll_b, gsh, mus, delta, length)
// 	}
// 	return lll_b
// }

// func cracking(publicKey, CryptedMsg []*big.Int, Alice *alice.T_Alice) {
// 	var delta float64 = 0.99
// 	var length = Alice.KeyLen + 1
// 	for _, elem := range CryptedMsg {
// 		Basis := createBasis(publicKey, elem, length)
// 		probable_answer := deltaLLL(Basis, delta, length)
// 		for i, j := 0, len(probable_answer[0])-1; i < j; i, j = i+1, j-1 {
// 			probable_answer[0][i], probable_answer[0][j] = probable_answer[0][j], probable_answer[0][i]
// 		}
// 		for i := 0; i < length; i++ {
// 			fmt.Print(probable_answer[0][i].String(), " ")
// 		}
// 		fmt.Println()
// 	}
// }

// type Listener int

// type Reply struct {
// 	Data string
// }

// func (l *Listener) PrintDecryptedMsg(msg string, reply *Reply) error {
// 	*reply = Reply{msg}
// 	fmt.Println(reply.Data)
// 	return nil
// }

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4045")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Отправка ключа
	var Bob *bob.T_Bob = bob.CreateBob()
	var b strings.Builder
	for i := 0; i < Bob.KeyLen; i++ {
		b.WriteString(Bob.PublicKey[i].String())
		b.WriteByte(' ')
	}
	conn.Write([]byte(b.String()))

	buf := make([]byte, 8192)
	n, err := conn.Read(buf)
	if n == 0 || err != nil {
		log.Print(err)
	}
	var tmp []string = strings.Split(string(buf[:n]), " ")
	var MsgLen int = len(tmp) - 1
	Len, _ := strconv.ParseInt(tmp[MsgLen], 10, 64)
	Bob.MsgLen = int(Len)
	tmp = tmp[:MsgLen]
	fmt.Println(tmp)

	var CryptedMsg []*big.Int = make([]*big.Int, MsgLen)
	for i := 0; i < MsgLen; i++ {
		CryptedMsg[i] = new(big.Int)
		CryptedMsg[i].SetString(tmp[i], 10)
	}
	Bob.CryptedMsg = CryptedMsg
	decryptedMsg := Bob.Encrypting()
	fmt.Println(decryptedMsg)
}