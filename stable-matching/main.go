package main

import "fmt"

type people map[string]info
type info struct {
	orderOfPreference []string
	matched           string
}

func main() {
	mList := people{
		"mA": info{
			orderOfPreference: []string{"fA", "fB", "fC", "fD"},
		},
		"mB": info{
			orderOfPreference: []string{"fC", "fB", "fA", "fD"},
		},
		"mC": info{
			orderOfPreference: []string{"fA", "fB", "fD", "fC"},
		},
		"mD": info{
			orderOfPreference: []string{"fC", "fA", "fD", "fB"},
		},
	}
	fList := people{
		"fA": info{
			orderOfPreference: []string{"mA", "mB", "mC", "mD"},
		},
		"fB": info{
			orderOfPreference: []string{"mB", "mA", "mD", "mC"},
		},
		"fC": info{
			orderOfPreference: []string{"mB", "mC", "mA", "mD"},
		},
		"fD": info{
			orderOfPreference: []string{"mA", "mD", "mC", "mB"},
		},
	}

	matching := stableMatching(mList, fList)
	fmt.Println(matching)
}

func matchedList(mList people) map[string]string {
	matching := make(map[string]string, len(mList))
	for name, info := range mList {
		matching[name] = info.matched
	}
	return matching
}

func contract(order []string, matched string, proposer string) bool {
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

func noMatched(p people) string {
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
		// 独身男性
		mName := noMatched(mList)
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
			// その子の中で自分の方が順位が上ならマッチング
			if contract(fInfo.orderOfPreference, fInfo.matched, mName) {
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
