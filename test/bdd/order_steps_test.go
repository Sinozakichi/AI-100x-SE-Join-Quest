// 本檔案為 BDD 測試入口，使用 godog 套件執行 cucumber feature 驗收測試。
//
// 與傳統 unit test（如 testify）不同：
// - 傳統 unit test 會在 *_test.go 檔案裡針對每個功能寫一個一個一個 TestXXX(t *testing.T) function，
//   測試資料、流程、預期結果都寫在 Go code 裡。
// - godog 採用 BDD 流程，所有測試案例（scenarios）都寫在 .feature 檔（Gherkin/Cucumber 語法），
//   Go code 只需撰寫 step 對應的 function（step definitions），不需為每個情境寫一個 Go function。
//
// TestMain 是 godog 的啟動點，負責：
// - 指定 feature 檔案路徑
// - 註冊 glue code（step definitions）
// - 執行所有 feature/scenario
//
// godog 會自動解析 feature file，把每個 scenario 轉成一個個測試步驟，
// 並呼叫你在 glue code 裡註冊的 Go function（step definitions）。
//
// 你不用像 unit test 那樣每個 case 都寫一個 Go function，
// 只要對應 step（Given/When/Then）即可。

package bdd

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
)

// TestMain 是 godog 測試的唯一入口，會自動讀取 feature file，
// 並根據 glue code 執行所有 BDD 驗收情境。
func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                "order",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"../../features/order.feature"},
		},
	}.Run()
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
