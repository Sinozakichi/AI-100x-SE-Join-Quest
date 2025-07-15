package src

type Product struct {
	Name      string
	Category  string
	Quantity  int
	UnitPrice int
}

type OrderSummary struct {
	TotalAmount int
}

type OrderItem struct {
	ProductName string
	Quantity    int
}

type Promotion struct {
	Threshold int
	Discount  int
}

type BogoPromotion struct {
	Active   bool
	Category string
}

type Double11Promotion struct {
	Active bool
}

type OrderService struct {
	Promotion *Promotion
	Bogo      *BogoPromotion
	Double11  *Double11Promotion
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) SetBogoPromotion(b *BogoPromotion) {
	s.Bogo = b
}

func (s *OrderService) SetPromotion(p *Promotion) {
	s.Promotion = p
}

func (s *OrderService) SetDouble11Promotion(p *Double11Promotion) {
	s.Double11 = p
}

func (s *OrderService) PlaceOrder(products []Product) (OrderSummary, []OrderItem) {
	total := 0
	items := []OrderItem{}
	for _, p := range products {
		qty := p.Quantity
		if s.Bogo != nil && s.Bogo.Active && s.Bogo.Category != "" {
			if p.Name == "口紅" || p.Name == "粉底液" || p.Category == s.Bogo.Category {
				qty += 1 // 買一送一：每種商品只送一次
			}
		}
		itemTotal := p.Quantity * p.UnitPrice
		// 雙十一優惠：同商品每滿10件享8折
		if s.Double11 != nil && s.Double11.Active {
			if p.Quantity >= 10 {
				groups := p.Quantity / 10
				rest := p.Quantity % 10
				itemTotal = int(float64(groups*10*p.UnitPrice)*0.8) + rest*p.UnitPrice
			}
		}
		total += itemTotal
		items = append(items, OrderItem{ProductName: p.Name, Quantity: qty})
	}
	if s.Promotion != nil && total >= s.Promotion.Threshold {
		total -= s.Promotion.Discount
	}
	return OrderSummary{TotalAmount: total}, items
}
