# cli shop


![clishop screnshot](/clishop-ss.png)


## Definiton:

[https://db.dai.fmph.uniba.sk/teaching/db2/projekt/](https://db.dai.fmph.uniba.sk/teaching/db2/projekt/)

[https://docs.google.com/document/d/12sy75jUickbnZCY0HU3kC7AHjLCxAdfOPmqy-dRTYis/edit#](https://docs.google.com/document/d/12sy75jUickbnZCY0HU3kC7AHjLCxAdfOPmqy-dRTYis/edit#)


## To run:
configure database connection in /db/db.go

```
go run main.go
```


## Requirements

Golang 1.18 (because of generics)

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# db

```go
import "ondrejmaksi.com/db2project/db"
```

## Index

- [func GetDatabase() *sql.DB](<#func-getdatabase>)
- [func RunScript(path string) error](<#func-runscript>)


## func GetDatabase

```go
func GetDatabase() *sql.DB
```

public function to get database instance as singleton

## func RunScript

```go
func RunScript(path string) error
```

function to run sql script



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# rdg

```go
import "ondrejmaksi.com/db2project/rdg"
```

## Index

- [func AddProductToBasket(user User, product Product)](<#func-addproducttobasket>)
- [func CreateOrder(order Order) int](<#func-createorder>)
- [func CreateOrderItem(orderItem OrderItem)](<#func-createorderitem>)
- [func CreateProduct(product Product)](<#func-createproduct>)
- [func CreateUser(user User)](<#func-createuser>)
- [func DeleteBasketItem(product Product)](<#func-deletebasketitem>)
- [func DeleteProduct(product Product)](<#func-deleteproduct>)
- [func DeleteUser(user User)](<#func-deleteuser>)
- [func EmptyBasket(userId int)](<#func-emptybasket>)
- [func ListOrderItems(orderId int)](<#func-listorderitems>)
- [func UpdateBasketItemQuantity(product Product)](<#func-updatebasketitemquantity>)
- [func UpdateProduct(product Product)](<#func-updateproduct>)
- [func UpdateUser(user User)](<#func-updateuser>)
- [type Jsonable](<#type-jsonable>)
  - [func (u *Jsonable) ToJson() string](<#func-jsonable-tojson>)
- [type Order](<#type-order>)
  - [func ListOrders() []Order](<#func-listorders>)
- [type OrderItem](<#type-orderitem>)
- [type Product](<#type-product>)
  - [func GetProducts() []Product](<#func-getproducts>)
  - [func ListProductsInBasket(userId int) []Product](<#func-listproductsinbasket>)
  - [func SearchForProductByAttribute(searchText string) ([]Product, error)](<#func-searchforproductbyattribute>)
- [type SaleStatistic](<#type-salestatistic>)
  - [func GetSaleStatistics() []SaleStatistic](<#func-getsalestatistics>)
- [type SearchStatistic](<#type-searchstatistic>)
  - [func GetSearchStatistics(year string) []SearchStatistic](<#func-getsearchstatistics>)
- [type User](<#type-user>)
  - [func GetUser(id int) []User](<#func-getuser>)
  - [func GetUsers() []User](<#func-getusers>)


## func AddProductToBasket

```go
func AddProductToBasket(user User, product Product)
```

## func CreateOrder

```go
func CreateOrder(order Order) int
```

## func CreateOrderItem

```go
func CreateOrderItem(orderItem OrderItem)
```

## func CreateProduct

```go
func CreateProduct(product Product)
```

## func CreateUser

```go
func CreateUser(user User)
```

## func DeleteBasketItem

```go
func DeleteBasketItem(product Product)
```

## func DeleteProduct

```go
func DeleteProduct(product Product)
```

## func DeleteUser

```go
func DeleteUser(user User)
```

## func EmptyBasket

```go
func EmptyBasket(userId int)
```

## func ListOrderItems

```go
func ListOrderItems(orderId int)
```

## func UpdateBasketItemQuantity

```go
func UpdateBasketItemQuantity(product Product)
```

## func UpdateProduct

```go
func UpdateProduct(product Product)
```

## func UpdateUser

```go
func UpdateUser(user User)
```

## type Jsonable

Json serialization struct

```go
type Jsonable struct {
}
```

### func \(\*Jsonable\) ToJson

```go
func (u *Jsonable) ToJson() string
```

## type Order

An order\, containing an address\, Total price and status\.

```go
type Order struct {
    Id      int
    UserId  int
    Address string
    Total   float64
    Status  string
}
```

### func ListOrders

```go
func ListOrders() []Order
```

## type OrderItem

An item of an order\, containing the product id reference\, the order id reference and the quantity\.

```go
type OrderItem struct {
    Id        int
    UserId    int
    ProductId int
    OrderId   int
    Quantity  int
}
```

## type Product

A product row with Name\, Price and Quantity

```go
type Product struct {
    Jsonable
    Id       uint
    Name     string
    Price    float64
    Quantity uint
}
```

### func GetProducts

```go
func GetProducts() []Product
```

### func ListProductsInBasket

```go
func ListProductsInBasket(userId int) []Product
```

### func SearchForProductByAttribute

```go
func SearchForProductByAttribute(searchText string) ([]Product, error)
```

Fulltext search for product by string input\. Searches in all product attributes and product name

## type SaleStatistic

Detailnos?? popisu produktu je po??et atrib??tov\, ktor?? m?? produkt v syst??me vyplnen??\. Pre ka??d?? detailnos?? popisu produktu vypo????tajte po??et predan??ch kusov produktov s danou detailnos??ou produktu\.

```go
type SaleStatistic struct {
    AttributeCount int
    TotalSold      int
}
```

### func GetSaleStatistics

```go
func GetSaleStatistics() []SaleStatistic
```

## type SearchStatistic

Pou????vate?? zad?? na vstupe po??iato??n?? rok\. Pre ka??d?? mesiac v danom a nasledovn??ch rokoch \(januar 2018\, februar 2018\, ???\) vypo????tajte\, ak?? je pomer medzi t??mi vyh??ad??vac??mi dopytmi\, ktor?? obsahuj?? n??zov produktu a t??mi\, ktor?? n??zov produktu neobsahuj??\. Vo v??stupe musia by?? mesiace\, pre ktor?? nie s?? v datab??ze d??ta\. Ak dan?? pomer nie je pre dan?? mesiac definovan??\, vo v??stupe m?? by?? NULL\.

```go
type SearchStatistic struct {
    Month time.Time
    Share sql.NullFloat64
}
```

### func GetSearchStatistics

```go
func GetSearchStatistics(year string) []SearchStatistic
```

## type User

User struct containing id\, email and basket item count

```go
type User struct {
    Id              int
    Email           string
    BasketItemCount int
}
```

### func GetUser

```go
func GetUser(id int) []User
```

### func GetUsers

```go
func GetUsers() []User
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# ts

```go
import "ondrejmaksi.com/db2project/ts"
```

## Index

- [func CancelOrder(order rdg.Order)](<#func-cancelorder>)
- [func CreateOrder(userId int, address string)](<#func-createorder>)
- [func ExpediteOrder(order rdg.Order) error](<#func-expediteorder>)
- [func PayOrder(order rdg.Order, amount float64) error](<#func-payorder>)


## func CancelOrder

```go
func CancelOrder(order rdg.Order)
```

Objedn??vky\, ktor?? e??te neboli zaplaten??\, sa daj?? zru??i??\. Zarezervovan?? tovar sa mus?? odblokova??\. Zaplaten?? ani expedovan?? objedn??vky nie je mo??n?? zru??i??\.

## func CreateOrder

```go
func CreateOrder(userId int, address string)
```

Pou????vate?? si zvol??\, ??e chce objedna?? tovar\, ktor?? m?? vlo??en?? v ko????ku\. Zad?? miesto dodania\. Vypo????ta sa celkov?? cena \(produkty \+ doprava\)\. V syst??me sa vytvor?? objedn??vka v stave vytvoren??\. Tovar sa na sklade rezervuje\. Ak nie je v??etok tovar na sklade\, objedn??vka sa stale vytvor??\, no pou????vate?? je na to upozornen??\.

## func ExpediteOrder

```go
func ExpediteOrder(order rdg.Order) error
```

Zamestnanec vyberie id objedn??vky\, ktor?? chce expedova??\. Ak je objedn??vka zaplaten?? a je dostatok tovarov na sklade\, objedn??vka sa presunie do stavu expedovan??\. Z??rove?? sa upravia skladov?? z??soby tovarov\.

## func PayOrder

```go
func PayOrder(order rdg.Order, amount float64) error
```

Po zaplaten?? sa v syst??me eviduje platba\. V pr??pade\, ??e v????ka platby nesed?? so sumou objedn??vky\, platba mus?? by?? vr??ten?? \(toto simulujte nejak??m v??pisom do konzoly\)\. Po ??spe??nej platbe sa objedn??vka pres??va do stavu zaplaten??\. V pr??pade\, ??e na sklade nie je dostato??n?? mno??stvo nejak??ho tovaru\, tento tovar sa automaticky objedn?? z ve??koskladu a dopln?? v po??adovanom po??te \(simulujte v??pisom do konzoly\)\. Dvakr??t zaplati?? t?? ist?? objedn??vku mo??n?? nie je\. Syst??m to mus?? odmietnu?? a ozn??mi?? chybu\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# ui

```go
import "ondrejmaksi.com/db2project/ui"
```

## Index

- [func Init() *tview.Application](<#func-init>)
- [func NewApp() *tview.Application](<#func-newapp>)
- [func Noop[T any](T)](<#func-noop>)
- [type Actions](<#type-actions>)
  - [func (c *Actions) AddBasketItem(product rdg.Product)](<#func-actions-addbasketitem>)
  - [func (c *Actions) CancelOrder(order rdg.Order)](<#func-actions-cancelorder>)
  - [func (c *Actions) CreateOrder(userId int, address string)](<#func-actions-createorder>)
  - [func (c *Actions) CreateProduct(product rdg.Product)](<#func-actions-createproduct>)
  - [func (c *Actions) CreateUser(user rdg.User)](<#func-actions-createuser>)
  - [func (c *Actions) DeleteBasketItem(product rdg.Product)](<#func-actions-deletebasketitem>)
  - [func (c *Actions) DeleteProduct(product rdg.Product)](<#func-actions-deleteproduct>)
  - [func (c *Actions) DeleteUser(user rdg.User)](<#func-actions-deleteuser>)
  - [func (c *Actions) ExpediteOrder(order rdg.Order)](<#func-actions-expediteorder>)
  - [func (c *Actions) GetSearchStatistics(searchText string)](<#func-actions-getsearchstatistics>)
  - [func (c *Actions) PayOrder(order rdg.Order, amount float64)](<#func-actions-payorder>)
  - [func (c *Actions) SearchByAttribute(searchText string)](<#func-actions-searchbyattribute>)
  - [func (c *Actions) ShowAddBasketItem(product rdg.Product)](<#func-actions-showaddbasketitem>)
  - [func (c *Actions) ShowBasket(user rdg.User)](<#func-actions-showbasket>)
  - [func (c *Actions) ShowCreateOrder(user rdg.User)](<#func-actions-showcreateorder>)
  - [func (c *Actions) ShowEditBasketItem(product rdg.Product)](<#func-actions-showeditbasketitem>)
  - [func (c *Actions) ShowEditProduct(product rdg.Product)](<#func-actions-showeditproduct>)
  - [func (c *Actions) ShowEditUser(user rdg.User)](<#func-actions-showedituser>)
  - [func (c *Actions) ShowNewUser()](<#func-actions-shownewuser>)
  - [func (c *Actions) ShowOrderList()](<#func-actions-showorderlist>)
  - [func (c *Actions) ShowPayOrder(order rdg.Order)](<#func-actions-showpayorder>)
  - [func (c *Actions) ShowProductList()](<#func-actions-showproductlist>)
  - [func (c *Actions) ShowSaleStatistics()](<#func-actions-showsalestatistics>)
  - [func (c *Actions) ShowUserList()](<#func-actions-showuserlist>)
  - [func (c *Actions) UpdateBasketItem(product rdg.Product)](<#func-actions-updatebasketitem>)
  - [func (c *Actions) UpdateProduct(product rdg.Product)](<#func-actions-updateproduct>)
  - [func (c *Actions) UpdateUser(user rdg.User)](<#func-actions-updateuser>)
- [type AppState](<#type-appstate>)


## func Init

```go
func Init() *tview.Application
```

Application initialization

## func NewApp

```go
func NewApp() *tview.Application
```

Application constructor\, creates main application window and components\, adds handling for state changes

## func Noop

```go
func Noop[T any](T)
```

## type Actions

Actions struct\, holds actions modifying state\, actions beginning with Show change screen content

```go
type Actions struct {
}
```

### func \(\*Actions\) AddBasketItem

```go
func (c *Actions) AddBasketItem(product rdg.Product)
```

### func \(\*Actions\) CancelOrder

```go
func (c *Actions) CancelOrder(order rdg.Order)
```

### func \(\*Actions\) CreateOrder

```go
func (c *Actions) CreateOrder(userId int, address string)
```

### func \(\*Actions\) CreateProduct

```go
func (c *Actions) CreateProduct(product rdg.Product)
```

### func \(\*Actions\) CreateUser

```go
func (c *Actions) CreateUser(user rdg.User)
```

### func \(\*Actions\) DeleteBasketItem

```go
func (c *Actions) DeleteBasketItem(product rdg.Product)
```

### func \(\*Actions\) DeleteProduct

```go
func (c *Actions) DeleteProduct(product rdg.Product)
```

### func \(\*Actions\) DeleteUser

```go
func (c *Actions) DeleteUser(user rdg.User)
```

### func \(\*Actions\) ExpediteOrder

```go
func (c *Actions) ExpediteOrder(order rdg.Order)
```

### func \(\*Actions\) GetSearchStatistics

```go
func (c *Actions) GetSearchStatistics(searchText string)
```

### func \(\*Actions\) PayOrder

```go
func (c *Actions) PayOrder(order rdg.Order, amount float64)
```

### func \(\*Actions\) SearchByAttribute

```go
func (c *Actions) SearchByAttribute(searchText string)
```

### func \(\*Actions\) ShowAddBasketItem

```go
func (c *Actions) ShowAddBasketItem(product rdg.Product)
```

### func \(\*Actions\) ShowBasket

```go
func (c *Actions) ShowBasket(user rdg.User)
```

### func \(\*Actions\) ShowCreateOrder

```go
func (c *Actions) ShowCreateOrder(user rdg.User)
```

### func \(\*Actions\) ShowEditBasketItem

```go
func (c *Actions) ShowEditBasketItem(product rdg.Product)
```

### func \(\*Actions\) ShowEditProduct

```go
func (c *Actions) ShowEditProduct(product rdg.Product)
```

### func \(\*Actions\) ShowEditUser

```go
func (c *Actions) ShowEditUser(user rdg.User)
```

### func \(\*Actions\) ShowNewUser

```go
func (c *Actions) ShowNewUser()
```

### func \(\*Actions\) ShowOrderList

```go
func (c *Actions) ShowOrderList()
```

### func \(\*Actions\) ShowPayOrder

```go
func (c *Actions) ShowPayOrder(order rdg.Order)
```

### func \(\*Actions\) ShowProductList

```go
func (c *Actions) ShowProductList()
```

### func \(\*Actions\) ShowSaleStatistics

```go
func (c *Actions) ShowSaleStatistics()
```

### func \(\*Actions\) ShowUserList

```go
func (c *Actions) ShowUserList()
```

### func \(\*Actions\) UpdateBasketItem

```go
func (c *Actions) UpdateBasketItem(product rdg.Product)
```

### func \(\*Actions\) UpdateProduct

```go
func (c *Actions) UpdateProduct(product rdg.Product)
```

### func \(\*Actions\) UpdateUser

```go
func (c *Actions) UpdateUser(user rdg.User)
```

## type AppState

Struct holding all of the application UI state

```go
type AppState struct {
    // contains filtered or unexported fields
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# cmp

```go
import "ondrejmaksi.com/db2project/ui/cmp"
```

## Index

- [func NewBasketItemForm(product rdg.Product, onSubmit func(product rdg.Product), onCancel func()) *tview.Form](<#func-newbasketitemform>)
- [func NewLog() *tview.List](<#func-newlog>)
- [func NewMainMenu(onMenuSelect func(int, string)) *tview.List](<#func-newmainmenu>)
- [func NewOrderForm(userId int, onSubmit func(userId int, address string), onCancel func()) *tview.Form](<#func-neworderform>)
- [func NewPaymentForm(order rdg.Order, onSubmit func(order rdg.Order, amount float64), onCancel func()) *tview.Form](<#func-newpaymentform>)
- [func NewProductForm(product rdg.Product, onSubmit func(product rdg.Product), onCancel func()) *tview.Form](<#func-newproductform>)
- [func NewSearchForm(onSubmit func(searchText string), onCancel func()) *tview.Form](<#func-newsearchform>)
- [func NewSearchStatisticsForm(onSubmit func(year string), onCancel func()) *tview.Form](<#func-newsearchstatisticsform>)
- [func NewStatusBar() (*tview.Flex, *tview.TextView, *tview.TextView)](<#func-newstatusbar>)
- [func NewUserForm(user rdg.User, onSubmit func(user rdg.User), onCancel func()) *tview.Form](<#func-newuserform>)
- [type Content](<#type-content>)
  - [func NewContent() *Content](<#func-newcontent>)
  - [func (c *Content) SetContent(content tview.Primitive)](<#func-content-setcontent>)
- [type OrderTable](<#type-ordertable>)
  - [func NewOrdersTable(orders []rdg.Order, onCancel func(order rdg.Order), onPay func(order rdg.Order), onExpedite func(order rdg.Order), onDone func()) *OrderTable](<#func-neworderstable>)
- [type ProductTable](<#type-producttable>)
  - [func NewProductsTable(products []rdg.Product, onEdit func(product rdg.Product), onDelete func(product rdg.Product), onAdd func(user rdg.Product), onDone func()) *ProductTable](<#func-newproductstable>)
  - [func (pt *ProductTable) SetProps(products []rdg.Product, onDelete func(user rdg.Product), onEdit func(user rdg.Product), onAdd func(user rdg.Product), onDone func())](<#func-producttable-setprops>)
  - [func (pt *ProductTable) ShowProducts(products []rdg.Product)](<#func-producttable-showproducts>)
- [type SaleStatisticTable](<#type-salestatistictable>)
  - [func NewSaleStatisticsTable(saleStatistics []rdg.SaleStatistic, onDone func()) *SaleStatisticTable](<#func-newsalestatisticstable>)
  - [func (sst *SaleStatisticTable) SetProps(saleStatistics []rdg.SaleStatistic, onDone func())](<#func-salestatistictable-setprops>)
  - [func (sst *SaleStatisticTable) ShowSaleStatistics(saleStatistics []rdg.SaleStatistic)](<#func-salestatistictable-showsalestatistics>)
- [type SearchStatisticTable](<#type-searchstatistictable>)
  - [func NewSearchStatisticsTable(searchStatistics []rdg.SearchStatistic, onDone func()) *SearchStatisticTable](<#func-newsearchstatisticstable>)
  - [func (sst *SearchStatisticTable) SetProps(searchStatistics []rdg.SearchStatistic, onDone func())](<#func-searchstatistictable-setprops>)
  - [func (sst *SearchStatisticTable) ShowSearchStatistics(searchStatistics []rdg.SearchStatistic)](<#func-searchstatistictable-showsearchstatistics>)
- [type UserTable](<#type-usertable>)
  - [func NewUsersTable(users []rdg.User, onDelete func(user rdg.User), onEdit func(user rdg.User), onBasket func(user rdg.User), onOrder func(user rdg.User), onDone func()) *UserTable](<#func-newuserstable>)
  - [func (ut *UserTable) SetProps(users []rdg.User, onDelete func(user rdg.User), onEdit func(user rdg.User), onBasket func(user rdg.User), onOrder func(user rdg.User), onDone func())](<#func-usertable-setprops>)
  - [func (ut *UserTable) ShowUsers(users []rdg.User)](<#func-usertable-showusers>)


## func NewBasketItemForm

```go
func NewBasketItemForm(product rdg.Product, onSubmit func(product rdg.Product), onCancel func()) *tview.Form
```

Creates a new basket item form\.

## func NewLog

```go
func NewLog() *tview.List
```

## func NewMainMenu

```go
func NewMainMenu(onMenuSelect func(int, string)) *tview.List
```

## func NewOrderForm

```go
func NewOrderForm(userId int, onSubmit func(userId int, address string), onCancel func()) *tview.Form
```

## func NewPaymentForm

```go
func NewPaymentForm(order rdg.Order, onSubmit func(order rdg.Order, amount float64), onCancel func()) *tview.Form
```

## func NewProductForm

```go
func NewProductForm(product rdg.Product, onSubmit func(product rdg.Product), onCancel func()) *tview.Form
```

## func NewSearchForm

```go
func NewSearchForm(onSubmit func(searchText string), onCancel func()) *tview.Form
```

## func NewSearchStatisticsForm

```go
func NewSearchStatisticsForm(onSubmit func(year string), onCancel func()) *tview.Form
```

## func NewStatusBar

```go
func NewStatusBar() (*tview.Flex, *tview.TextView, *tview.TextView)
```

## func NewUserForm

```go
func NewUserForm(user rdg.User, onSubmit func(user rdg.User), onCancel func()) *tview.Form
```

## type Content

```go
type Content struct {
    *tview.Pages
}
```

### func NewContent

```go
func NewContent() *Content
```

### func \(\*Content\) SetContent

```go
func (c *Content) SetContent(content tview.Primitive)
```

## type OrderTable

```go
type OrderTable struct {
    *tview.Table
}
```

### func NewOrdersTable

```go
func NewOrdersTable(orders []rdg.Order, onCancel func(order rdg.Order), onPay func(order rdg.Order), onExpedite func(order rdg.Order), onDone func()) *OrderTable
```

Creates a new table for displaying a list of orders

## type ProductTable

```go
type ProductTable struct {
    *tview.Table
}
```

### func NewProductsTable

```go
func NewProductsTable(products []rdg.Product, onEdit func(product rdg.Product), onDelete func(product rdg.Product), onAdd func(user rdg.Product), onDone func()) *ProductTable
```

### func \(\*ProductTable\) SetProps

```go
func (pt *ProductTable) SetProps(products []rdg.Product, onDelete func(user rdg.Product), onEdit func(user rdg.Product), onAdd func(user rdg.Product), onDone func())
```

### func \(\*ProductTable\) ShowProducts

```go
func (pt *ProductTable) ShowProducts(products []rdg.Product)
```

## type SaleStatisticTable

```go
type SaleStatisticTable struct {
    *tview.Table
}
```

### func NewSaleStatisticsTable

```go
func NewSaleStatisticsTable(saleStatistics []rdg.SaleStatistic, onDone func()) *SaleStatisticTable
```

### func \(\*SaleStatisticTable\) SetProps

```go
func (sst *SaleStatisticTable) SetProps(saleStatistics []rdg.SaleStatistic, onDone func())
```

### func \(\*SaleStatisticTable\) ShowSaleStatistics

```go
func (sst *SaleStatisticTable) ShowSaleStatistics(saleStatistics []rdg.SaleStatistic)
```

## type SearchStatisticTable

```go
type SearchStatisticTable struct {
    *tview.Table
}
```

### func NewSearchStatisticsTable

```go
func NewSearchStatisticsTable(searchStatistics []rdg.SearchStatistic, onDone func()) *SearchStatisticTable
```

### func \(\*SearchStatisticTable\) SetProps

```go
func (sst *SearchStatisticTable) SetProps(searchStatistics []rdg.SearchStatistic, onDone func())
```

### func \(\*SearchStatisticTable\) ShowSearchStatistics

```go
func (sst *SearchStatisticTable) ShowSearchStatistics(searchStatistics []rdg.SearchStatistic)
```

## type UserTable

```go
type UserTable struct {
    *tview.Table
}
```

### func NewUsersTable

```go
func NewUsersTable(users []rdg.User, onDelete func(user rdg.User), onEdit func(user rdg.User), onBasket func(user rdg.User), onOrder func(user rdg.User), onDone func()) *UserTable
```

### func \(\*UserTable\) SetProps

```go
func (ut *UserTable) SetProps(users []rdg.User, onDelete func(user rdg.User), onEdit func(user rdg.User), onBasket func(user rdg.User), onOrder func(user rdg.User), onDone func())
```

### func \(\*UserTable\) ShowUsers

```go
func (ut *UserTable) ShowUsers(users []rdg.User)
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# lib

```go
import "ondrejmaksi.com/db2project/ui/lib"
```

## Index

- [type GenericState](<#type-genericstate>)
  - [func NewGenericState[T any](value T) *GenericState[T]](<#func-newgenericstate>)
  - [func (s *GenericState[T]) AddHandler(handler func(T))](<#func-genericstatet-addhandler>)
  - [func (s *GenericState[T]) GetState() T](<#func-genericstatet-getstate>)
  - [func (s *GenericState[T]) NotifyAll()](<#func-genericstatet-notifyall>)
  - [func (s *GenericState[T]) SetState(value T)](<#func-genericstatet-setstate>)


## type GenericState

Generic single value state with Setter\, Getter and ability to add handler\, and notify all listeners

```go
type GenericState[T any] struct {
    // contains filtered or unexported fields
}
```

### func NewGenericState

```go
func NewGenericState[T any](value T) *GenericState[T]
```

### func \(\*GenericState\[T\]\) AddHandler

```go
func (s *GenericState[T]) AddHandler(handler func(T))
```

### func \(\*GenericState\[T\]\) GetState

```go
func (s *GenericState[T]) GetState() T
```

### func \(\*GenericState\[T\]\) NotifyAll

```go
func (s *GenericState[T]) NotifyAll()
```

### func \(\*GenericState\[T\]\) SetState

```go
func (s *GenericState[T]) SetState(value T)
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# state

```go
import "ondrejmaksi.com/db2project/ui/state"
```

## Index

- [type ContentState](<#type-contentstate>)
  - [func NewContentState(value string) *ContentState](<#func-newcontentstate>)
- [type FocusState](<#type-focusstate>)
  - [func NewFocusState(value string) *FocusState](<#func-newfocusstate>)
  - [func (f *FocusState) AddFocusTarget(name string, p tview.Primitive)](<#func-focusstate-addfocustarget>)
  - [func (s *FocusState) Content()](<#func-focusstate-content>)
  - [func (f *FocusState) GetPrimitives() map[string]tview.Primitive](<#func-focusstate-getprimitives>)
  - [func (s *FocusState) Menu()](<#func-focusstate-menu>)
- [type KeyHintsState](<#type-keyhintsstate>)
  - [func NewKeyHintsState(value string) *KeyHintsState](<#func-newkeyhintsstate>)
  - [func (k *KeyHintsState) Clear()](<#func-keyhintsstate-clear>)
  - [func (k *KeyHintsState) SetForAddToBasket()](<#func-keyhintsstate-setforaddtobasket>)
  - [func (k *KeyHintsState) SetForBasket()](<#func-keyhintsstate-setforbasket>)
  - [func (k *KeyHintsState) SetForOrder()](<#func-keyhintsstate-setfororder>)
  - [func (k *KeyHintsState) SetForProduct()](<#func-keyhintsstate-setforproduct>)
  - [func (k *KeyHintsState) SetForUser()](<#func-keyhintsstate-setforuser>)
- [type MessageState](<#type-messagestate>)
  - [func NewMessageState(value string) *MessageState](<#func-newmessagestate>)
  - [func (s *MessageState) Fail(message string)](<#func-messagestate-fail>)
  - [func (s *MessageState) SetMessage(message string)](<#func-messagestate-setmessage>)
  - [func (s *MessageState) Success(message string)](<#func-messagestate-success>)
- [type OrderState](<#type-orderstate>)
  - [func NewOrderState(user rdg.Order) *OrderState](<#func-neworderstate>)
- [type OrdersState](<#type-ordersstate>)
  - [func NewOrdersState(users []rdg.Order) *OrdersState](<#func-newordersstate>)
- [type ProductState](<#type-productstate>)
  - [func NewProductState(user rdg.Product) *ProductState](<#func-newproductstate>)
- [type ProductsState](<#type-productsstate>)
  - [func NewProductsState(users []rdg.Product) *ProductsState](<#func-newproductsstate>)
- [type SaleStatisticsState](<#type-salestatisticsstate>)
  - [func NewSaleStatisticsState(ss []rdg.SaleStatistic) *SaleStatisticsState](<#func-newsalestatisticsstate>)
- [type SearchStatisticsState](<#type-searchstatisticsstate>)
  - [func NewSearchStatisticsState(ss []rdg.SearchStatistic) *SearchStatisticsState](<#func-newsearchstatisticsstate>)
- [type TitleState](<#type-titlestate>)
  - [func NewTitleState(value string) *TitleState](<#func-newtitlestate>)
- [type UserState](<#type-userstate>)
  - [func NewUserState(user rdg.User) *UserState](<#func-newuserstate>)
- [type UsersState](<#type-usersstate>)
  - [func NewUsersState(users []rdg.User) *UsersState](<#func-newusersstate>)


## type ContentState

```go
type ContentState struct {
    // contains filtered or unexported fields
}
```

### func NewContentState

```go
func NewContentState(value string) *ContentState
```

## type FocusState

```go
type FocusState struct {
    // contains filtered or unexported fields
}
```

### func NewFocusState

```go
func NewFocusState(value string) *FocusState
```

### func \(\*FocusState\) AddFocusTarget

```go
func (f *FocusState) AddFocusTarget(name string, p tview.Primitive)
```

### func \(\*FocusState\) Content

```go
func (s *FocusState) Content()
```

### func \(\*FocusState\) GetPrimitives

```go
func (f *FocusState) GetPrimitives() map[string]tview.Primitive
```

### func \(\*FocusState\) Menu

```go
func (s *FocusState) Menu()
```

## type KeyHintsState

```go
type KeyHintsState struct {
    // contains filtered or unexported fields
}
```

### func NewKeyHintsState

```go
func NewKeyHintsState(value string) *KeyHintsState
```

### func \(\*KeyHintsState\) Clear

```go
func (k *KeyHintsState) Clear()
```

### func \(\*KeyHintsState\) SetForAddToBasket

```go
func (k *KeyHintsState) SetForAddToBasket()
```

### func \(\*KeyHintsState\) SetForBasket

```go
func (k *KeyHintsState) SetForBasket()
```

### func \(\*KeyHintsState\) SetForOrder

```go
func (k *KeyHintsState) SetForOrder()
```

### func \(\*KeyHintsState\) SetForProduct

```go
func (k *KeyHintsState) SetForProduct()
```

### func \(\*KeyHintsState\) SetForUser

```go
func (k *KeyHintsState) SetForUser()
```

## type MessageState

```go
type MessageState struct {
    // contains filtered or unexported fields
}
```

### func NewMessageState

```go
func NewMessageState(value string) *MessageState
```

### func \(\*MessageState\) Fail

```go
func (s *MessageState) Fail(message string)
```

### func \(\*MessageState\) SetMessage

```go
func (s *MessageState) SetMessage(message string)
```

### func \(\*MessageState\) Success

```go
func (s *MessageState) Success(message string)
```

## type OrderState

```go
type OrderState struct {
    // contains filtered or unexported fields
}
```

### func NewOrderState

```go
func NewOrderState(user rdg.Order) *OrderState
```

## type OrdersState

```go
type OrdersState struct {
    // contains filtered or unexported fields
}
```

### func NewOrdersState

```go
func NewOrdersState(users []rdg.Order) *OrdersState
```

## type ProductState

```go
type ProductState struct {
    // contains filtered or unexported fields
}
```

### func NewProductState

```go
func NewProductState(user rdg.Product) *ProductState
```

## type ProductsState

```go
type ProductsState struct {
    // contains filtered or unexported fields
}
```

### func NewProductsState

```go
func NewProductsState(users []rdg.Product) *ProductsState
```

## type SaleStatisticsState

```go
type SaleStatisticsState struct {
    // contains filtered or unexported fields
}
```

### func NewSaleStatisticsState

```go
func NewSaleStatisticsState(ss []rdg.SaleStatistic) *SaleStatisticsState
```

## type SearchStatisticsState

```go
type SearchStatisticsState struct {
    // contains filtered or unexported fields
}
```

### func NewSearchStatisticsState

```go
func NewSearchStatisticsState(ss []rdg.SearchStatistic) *SearchStatisticsState
```

## type TitleState

```go
type TitleState struct {
    // contains filtered or unexported fields
}
```

### func NewTitleState

```go
func NewTitleState(value string) *TitleState
```

## type UserState

```go
type UserState struct {
    // contains filtered or unexported fields
}
```

### func NewUserState

```go
func NewUserState(user rdg.User) *UserState
```

## type UsersState

```go
type UsersState struct {
    // contains filtered or unexported fields
}
```

### func NewUsersState

```go
func NewUsersState(users []rdg.User) *UsersState
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
