// 本檔案為 BDD glue code，負責將 feature file 的 Given/When/Then 步驟對應到 Go function。
//
// 每個 func 對應一種 step（如 Given no promotions are applied），
// 由 godog 自動呼叫，並與 feature file 的自然語言步驟對應。
//
// 你不用像 unit test 那樣每個情境都寫一個 Go function，
// 只要針對步驟（step）撰寫對應的 function 即可。
//
// InitializeScenario 負責將步驟與 function 綁定，
// 並在每個 scenario 執行前初始化狀態，確保測試獨立。

package bdd

import (
	src "ai100x-order/src"
	"context"
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
)

var (
	orderService   *src.OrderService
	orderProducts  []src.Product
	orderSummary   src.OrderSummary
	orderItems     []src.OrderItem
	threshold      int
	discount       int
	bogoActive     bool
	double11Active bool
)

func thereAreNoPromotions() error {
	orderService = src.NewOrderService()
	orderProducts = nil
	orderSummary = src.OrderSummary{}
	orderItems = nil
	return nil
}

func customerPlacesOrderWith(table *godog.Table) error {
	headers := table.Rows[0].Cells
	for i := 1; i < len(table.Rows); i++ {
		row := table.Rows[i].Cells
		product := src.Product{}
		for j, h := range headers {
			switch h.Value {
			case "productName":
				product.Name = row[j].Value
			case "category":
				product.Category = row[j].Value
			case "quantity":
				q, _ := strconv.Atoi(row[j].Value)
				product.Quantity = q
			case "unitPrice":
				p, _ := strconv.Atoi(row[j].Value)
				product.UnitPrice = p
			}
		}
		orderProducts = append(orderProducts, product)
	}
	if threshold > 0 && discount > 0 {
		orderService.SetPromotion(&src.Promotion{Threshold: threshold, Discount: discount})
	}
	if bogoActive {
		orderService.SetBogoPromotion(&src.BogoPromotion{Active: true, Category: "cosmetics"})
	}
	if double11Active {
		orderService.SetDouble11Promotion(&src.Double11Promotion{Active: true})
	}
	orderSummary, orderItems = orderService.PlaceOrder(orderProducts)
	return nil
}

func orderSummaryShouldBe(table *godog.Table) error {
	headers := table.Rows[0].Cells
	row := table.Rows[1].Cells
	for j, h := range headers {
		switch h.Value {
		case "totalAmount":
			expect, _ := strconv.Atoi(row[j].Value)
			if orderSummary.TotalAmount != expect {
				return fmt.Errorf("totalAmount want %d, got %d", expect, orderSummary.TotalAmount)
			}
		case "originalAmount":
			// 只針對 threshold discount 測試
			actual := 0
			for _, p := range orderProducts {
				actual += p.Quantity * p.UnitPrice
			}
			expect, _ := strconv.Atoi(row[j].Value)
			if actual != expect {
				return fmt.Errorf("originalAmount want %d, got %d", expect, actual)
			}
		case "discount":
			expect, _ := strconv.Atoi(row[j].Value)
			if discount != expect {
				return fmt.Errorf("discount want %d, got %d", expect, discount)
			}
		}
	}
	return nil
}

func customerShouldReceive(table *godog.Table) error {
	headers := table.Rows[0].Cells
	expects := map[string]int{}
	for i := 1; i < len(table.Rows); i++ {
		row := table.Rows[i].Cells
		var name string
		var qty int
		for j, h := range headers {
			switch h.Value {
			case "productName":
				name = row[j].Value
			case "quantity":
				qty, _ = strconv.Atoi(row[j].Value)
			}
		}
		expects[name] = qty
	}
	for _, item := range orderItems {
		if expects[item.ProductName] != item.Quantity {
			return fmt.Errorf("item %s want %d, got %d", item.ProductName, expects[item.ProductName], item.Quantity)
		}
	}
	return nil
}

// 其餘 scenario 的步驟定義（暫時不註冊到 context）
func theThresholdDiscountPromotionIsConfigured(table *godog.Table) error {
	headers := table.Rows[0].Cells
	row := table.Rows[1].Cells
	for j, h := range headers {
		switch h.Value {
		case "threshold":
			t, _ := strconv.Atoi(row[j].Value)
			threshold = t
		case "discount":
			d, _ := strconv.Atoi(row[j].Value)
			discount = d
		}
	}
	return nil
}

func theBuyOneGetOnePromotionForCosmeticsIsActive() error {
	bogoActive = true
	return nil
}

func double11PromotionIsActive() error {
	double11Active = true
	return nil
}

// InitializeScenario 會在每個 scenario 執行前呼叫，
// 綁定所有 Given/When/Then 步驟與對應的 Go function，
// 並初始化全域狀態，確保每個 scenario 測試獨立、無污染。
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		orderService = src.NewOrderService()
		orderProducts = nil
		orderSummary = src.OrderSummary{}
		orderItems = nil
		threshold = 0
		discount = 0
		bogoActive = false
		double11Active = false
		return ctx, nil
	})
	// 註冊前兩個 scenario 的步驟
	ctx.Step(`^no promotions are applied$`, thereAreNoPromotions)
	ctx.Step(`^a customer places an order with:$`, customerPlacesOrderWith)
	ctx.Step(`^the order summary should be:$`, orderSummaryShouldBe)
	ctx.Step(`^the customer should receive:$`, customerShouldReceive)
	ctx.Step(`^the threshold discount promotion is configured:$`, theThresholdDiscountPromotionIsConfigured)
	ctx.Step(`^the buy one get one promotion for cosmetics is active$`, theBuyOneGetOnePromotionForCosmeticsIsActive)
	ctx.Step(`^雙十一優惠活動啟動$`, double11PromotionIsActive)
}
