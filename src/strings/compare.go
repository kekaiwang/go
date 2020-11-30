// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

// Compare 返回一个按字典顺序比较两个字符串的整数。
// 如果 a == b，结果将为 0；如果 a <b，结果将为-1；如果 a> b，结果将为 +1。
//
// 仅包括与包字节对称的比较。 使用内置的字符串比较运算符==，<，>等通常更清晰，总是更快。
func Compare(a, b string) int {
	// NOTE(rsc): This function does NOT call the runtime cmpstring function,
	// because we do not want to provide any performance justification for
	// using strings.Compare. Basically no one should use strings.Compare.
	// As the comment above says, it is here only for symmetry with package bytes.
	// If performance is important, the compiler should be changed to recognize
	// the pattern so that all code doing three-way comparisons, not just code
	// using strings.Compare, can benefit.
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
