/**
 * @Author: dragon
 * @Description:
 * @File:  gcd
 * @Version: 1.0.0
 * @Date: 2020/3/27 上午11:26
 */

package algorithm

/**
 * 计算两个数最大公约数
 * 原理：欧几里德算法又称辗转相除法, 该算法用于计算两个整数的最大公约数. 定理如下:
 * gcd(a,b) = gcd(b,a mod b)
 */
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
