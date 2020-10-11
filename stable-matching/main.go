package main

import "fmt"

type people map[string]info
type info struct {
	orderOfPreference []string
	matched           string
}

func main() {
	mList := people{
		"a": info{
			orderOfPreference: []string{"1", "2", "3", "4"},
		},
		"b": info{
			orderOfPreference: []string{"3", "2", "1", "4"},
		},
		"c": info{
			orderOfPreference: []string{"1", "2", "4", "3"},
		},
		"d": info{
			orderOfPreference: []string{"3", "1", "4", "2"},
		},
	}
	fList := people{
		"1": info{
			orderOfPreference: []string{"a", "b", "c", "d"},
		},
		"2": info{
			orderOfPreference: []string{"b", "a", "d", "c"},
		},
		"3": info{
			orderOfPreference: []string{"b", "c", "a", "d"},
		},
		"4": info{
			orderOfPreference: []string{"a", "d", "c", "b"},
		},
	}

	matching := stableMatching(mList, fList)
	fmt.Println(matching)
}

func list2matchedList(mList people) map[string]string {
	matching := make(map[string]string, len(mList))
	for name, info := range mList {
		matching[name] = info.matched
	}
	return matching
}

func tryProposal(order []string, matched string, proposer string) bool {
	for _, name := range order {
		if name == matched {
			return false
		}
		if name == proposer {
			return true
		}
	}
	return false
}

func rmExpectaion(orderOfPreference []string, rejecter string) []string {
	for i, p := range orderOfPreference {
		if p == rejecter {
			copy(orderOfPreference[i:], orderOfPreference[i+1:])
			orderOfPreference[len(orderOfPreference)-1] = ""
			orderOfPreference = orderOfPreference[:len(orderOfPreference)-1]
		}
	}
	return orderOfPreference
}

func selectNotMatched(p people) string {
	for name, info := range p {
		if info.matched == "" {
			return name
		}
	}
	return ""
}

func stableMatching(mList people, fList people) map[string]string {
	var matchNum int
	for matchNum != len(mList) {
		// 独身男性を選択
		mName := selectNotMatched(mList)
		mInfo := mList[mName]

		// ターゲットの女性
		fName := mInfo.orderOfPreference[0]
		fInfo, _ := fList[fName]

		// ターゲットの女性に相手がいなかったらマッチング
		if fInfo.matched == "" {
			fInfo.matched = mName
			mInfo.matched = fName

			matchNum++
		} else {
			// ターゲットの女性に相手がいるなら、
			// その女性の中で自分の方が順位が上ならマッチング
			// すでにいた相手とは関係破棄
			if tryProposal(fInfo.orderOfPreference, fInfo.matched, mName) {
				rejectedInfo := mList[fInfo.matched]
				rejectedInfo.matched = ""
				mList[fInfo.matched] = rejectedInfo

				fInfo.matched = mName
				mInfo.matched = fName
			}
		}

		// 一度ターゲットにした女性はリストから消す
		mInfo.orderOfPreference = rmExpectaion(mInfo.orderOfPreference, fName)

		// List更新
		mList[mName] = mInfo
		fList[fName] = fInfo
	}

	return matchedList(mList)
}
