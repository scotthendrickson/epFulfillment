package fulfillment

import (
	"net/http"
	"net/url"
)

//Inventories consists of the warehouse where the product is stored and relevant information regarding that product
type Inventories struct {
	Warehouse   *Warehouse `json:"warehouse,omitempty"`
	WarehouseID string     `json:"warehouse_id,omitempty"`
	Product     *Product   `json:"product,omitempty"`
	ProductID   string     `json:"product_id,omitempty"`
	Quantity    int64      `json:"quantity,omitempty"`
}

//InventoriesList is a list of inventories
type InventoriesList struct {
	Inventories []*Inventories `json:"inventories,omitempty"`
}

//ListInventories can be used to retrieve all inventories on your products or only specific products
//This one can take three arguements, a list of product id's, an Includes list, and List Options.
//If a list of product id's is not supplied it will return all products and their quantity of units in
//each warehouse. Any product id's supplied in the list will limit the returned information to just those products.
//The includes list can consist of only two string options, either "product" or "warehouse". If you put "product" in the list it
//will return the whole product object and quantity of units in each warehouse. For "warehouse" if you include it in the list it
//will return the full warehouse object but if left out it will only return the warehouse id.
//List options can be used to set per_page,page, limit, and offset
//Example calls:
// client := epFulfillment.New("YOUR-API-KEY")
//This will return a information on a single product, the warehouse id's, product id, and quantity of units in each warehouse
//inventories, err := client.ListInventories([]string{"PRODUCT-ID"}, []string{}, nil)
//
//This will return a information on a single product, the warehouse id's, the product object, and quantity of units in each warehouse
//inventories, err := client.ListInventories([]string{"PRODUCT-ID"}, []string{"product"}, nil)
//
//This will return a information on a single product, the each warehouse object, product id, and quantity of units in each warehouse
//inventories, err := client.ListInventories([]string{"PRODUCT-ID"}, []string{"warehouse"}, nil)
//
//This will return a information on a single product, the each warehouse object, the product object, and quantity of units in each warehouse
//inventories, err := client.ListInventories([]string{"PRODUCT-ID"}, []string{"product", "warehouse"}, nil)
func (c *Client) ListInventories(productIds []string, includes []string, opt *ListOptions) (il InventoriesList, err error) {
	baseURL, err := url.Parse("inventories")

	params := url.Values{}

	if productIds != nil {
		for i := range productIds {
			params.Add("product_ids[]", productIds[i])
		}
	}

	if includes != nil {
		for i := range includes {
			if includes[i] == "product" || includes[i] == "warehouse" {
				params.Add("includes[]", includes[i])
			}
		}
	}
	baseURL.RawQuery = params.Encode()

	err = c.do(nil, http.MethodGet, baseURL.String(), &opt, &il)
	return
}
