package main

/*
	請寫出對於method : GetMachineGame 的測試
*/
import (
	"testing"
)

/*func testMain() {
	GetMachineGame("123456")
}*/

//測試機率是否有額度在遊戲中
func testGetMachineGame(t *testing.T) {
	gotests := []struct {
		name   string
		userID string
		want   int
	}{
		//測試結果比較資料
		{name: "沒有找到會員", userID: "000", want: 0},
		{name: "BetBase欄位沒有資料", userID: "326", want: 0},
		{name: "無法轉換格式", userID: "980", want: 0},
		{name: "BetBase欄位拆解後為0", userID: "234", want: 0},
		{name: "沒有額度在機率遊戲中", userID: "521", want: 0},
		{name: "機率遊戲額度大於0", userID: "789", want: 1},
	}
	for _, tt := range gotests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMachineGame(tt.userID); got != tt.want {
				t.Errorf("GetMachineGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
