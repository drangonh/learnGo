/**
 * @Author: dragon
 * @Description:
 * @File:  hasGroupsSizeX
 * @Version: 1.0.0
 * @Date: 2020/3/27 上午11:26
 */

package algorithm

/*
给定一副牌，每张牌上都写着一个整数。
此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。
*/
func HasGroupsSizeX(deck []int) bool {
	m := make(map[int]int, len(deck))
	for _, v := range deck {
		m[v]++
	}
	g := -1
	for _, v := range m {
		if g == -1 {
			g = v
		} else {
			g = Gcd(g, v)
		}
	}
	return g >= 2
}
