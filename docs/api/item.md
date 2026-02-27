# Item

An item is a thing that your company buys, sells, or re-sells, such as products and services. An item is shown as a line on an invoice or other sales form. The Item.Type attribute, which specifies how the item is used, has one of the following values: - strong: Inventory - definition: Used in transactions to track merchandise that your business purchases, stocks, and re-sells as inventory. QuickBooks tracks the current number of inventory items in stock, cost of goods sold, and the asset value of the inventory after the purchase and sale of every item. - strong: Group Used as a container for a bundle of items with a count for each item. For example, a Gift Basket with 2 apples, 5 pencils and 1 stack of paper. The bundle is the Gift Basket, the bundle items are apples, pencils and paper. Note: Creating them via the QuickBooks Online API is not supported. • Bundles cannot contain other bundles. • Bundles cannot contain categories. • An item can be listed more than once with same or different quantities. • Bundles can be added to transactions. - strong: Service - definition: Used in transactions to track services that you charge on the purchase. For example, specialized labor, consulting hours, and professional fees. - strong: NonInventory - definition: Used for goods you buy but don’t track, like office supplies. Used in transactions for goods and materials for a specific job that you charge back to the customer and don't track yourself. In addition to the above, QuickBooks companies supports item categories to define item hierarchies. Use Item.Type set to Category to create hierarchies. Of note: • Non-category items used as parent items and used for things the company sells cannot be freely mixed. • An app can now clearly distinguish between things the company sells and categories used to build a hierarchy to organize them. • Categories do not have a price, income account, or expense accounts. • Items—-the things the company sells—-cannot have children. That is, if your items are organized into a hierarchy, items can only be at the leaf level of the hierarchy. • Categories are only available on companies that have enabled Categories. Test the CompanyInfo.NameValue.Name.ItemCategoriesFeature flag: • true— categories are enabled • false— categories are not enabled.

## Attributes

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| Id |  | Required for update | IdType   Unique Identifier for an Intuit entity (object). Required for the update operation. |
| ItemCategoryType |  | Required |  |
| Name |  | Required | Name of the item. This value is unique. |
| SyncToken | String | Required for update | Version number of the entity. Required for the update operation. |
| InvStartDate | Date | Conditionally required | Date of opening balance for the inventory transaction. For read operations, the date returned in this field is always the originally provided inventory start date. |
| Type |  | Conditionally required | minorVersion: specified. String   Classification that specifies the use of this item. See the description at the top of the Item entity page for details about supported item types. |
| When |  | Optional | querying Item objects with minor versions earlier than 4 specified, NonInventory types are returned as type Service. |
| QtyOnHand | Decimal | Conditionally required | Current quantity of the Inventory items available for sale. Not used for Service or NonInventory type items.Required for Inventory type items. |
| AssetAccountRef | ReferenceType | Conditionally required | Reference to the Inventory Asset account that tracks the current value of the inventory. |
| Sku |  | Optional | The stock keeping unit (SKU) for this Item. This is a company-defined identifier for an item or product used in tracking inventory. |
| SalesTaxIncluded | Boolean | Optional | True if the sales tax is included in the item amount, and therefore is not calculated for the transaction. |
| TrackQtyOnHand | Boolean | Optional | True if there is quantity on hand to be tracked. Once this value is true, it cannot be updated to false. Applicable for items of type Inventory. |
| SalesTaxCodeRef | ReferenceType | Optional | Reference to the sales tax code for the Sales item. Applicable to Service and Sales item types only. |
| ClassRef | ReferenceType | Optional | Reference to the Class for the item. Query the Class name list resource to determine the appropriate object for this reference. Use Class.Id and Class. |
| Source | String | Optional | The Source type of the transactions created by QuickBooks Commerce. Valid values include: QBCommerce |
| PurchaseTaxIncluded | Boolean | Optional | True if the purchase tax is included in the item amount, and therefore is not calculated for the transaction. |
| Description |  | Optional | Description of the item. |
| AbatementRate |  | Optional |  |
| SubItem | Boolean | Optional | If true, this is a sub item. If false or null, this is a top-level item. |
| Taxable |  | Optional |  |
| UQCDisplayText |  | Optional | of 25 chars |
| ReorderPoint | Decimal | Optional | The minimum quantity of a particular inventory item that you need to restock at any given time. The ReorderPoint value cannot be set to null for sparse updates(sparse=true). |
| PurchaseDesc |  | Optional | 1000 chars String Purchase description for the item. |
| MetaData | ModificationMetaData | Read-only | Descriptive information about the entity. The MetaData values are set by Data Services and are  for all applications. |
| PrefVendorRef | ReferenceType | Optional | Reference to the preferred vendor of this item. Query the Vendor name list resource to determine the appropriate object for this reference. Use Vendor.Id and Vendor. |
| UQCId |  | Optional |  |
| ReverseChargeRate |  | Optional |  |
| PurchaseTaxCodeRef | ReferenceType | Optional | Reference to the purchase tax code for the item. Applicable to Service, Other Charge, and Product (Non-Inventory) item types. |
| ServiceType |  | Optional |  |
| PurchaseCost |  | Optional | of 99999999999 Decimal Amount paid when buying or ordering the item, as expressed in the home currency. |
| ParentRef | ReferenceType | Optional | The immediate parent of the sub item in the hierarchical Item:SubItem list. If SubItem is true, then ParenRef is required. If SubItem is true, then ParenRef is required. |
| UnitPrice |  | Optional | of 99999999999 Decimal  Corresponds to the Price/Rate column on the QuickBooks Online UI to specify either unit price, a discount, or a tax rate for item. |
| FullyQualifiedName | String | Read-only | Fully qualified name of the entity. The fully qualified name prepends the topmost parent, followed by each sub element separated by colons. Takes the form of Item:SubItem. |
| ExpenseAccountRef | ReferenceType | Optional | Reference to the expense account used to pay the vendor for this item. Must be an account with account type of Cost of Goods Sold. |
| Level | Integer | Read-only | Specifies the level of the hierarchy in which the entity is located. Zero specifies the top level of the hierarchy; anything above will be the next level with respect to the parent. |
| IncomeAccountRef |  | Optional | * Conditionally Required ReferenceType Reference to the posting account, that is, the account that records the proceeds from the sale of this item. |
| TaxClassificationRef | ReferenceType | Optional | Tax classification segregates different items into different classifications and the tax classification is one of the key parameters to determine appropriate tax on transactions involving items. |

## Sample Object

```json
{
  "Item": {
    "FullyQualifiedName": "Trees",
    "domain": "QBO",
    "Name": "Trees",
    "SyncToken": "0",
    "sparse": false,
    "Active": true,
    "Type": "Category",
    "Id": "29",
    "MetaData": {
      "CreateTime": "2015-10-06T08:50:34-07:00",
      "LastUpdatedTime": "2015-10-06T08:50:34-07:00"
    }
  },
  "time": "2015-10-06T08:50:34.863-07:00"
}
```

## Operations

### Create an item

- **Method**: POST
- **URL**: `/v3/company/{realmID}/item`

**Request Body**:
```json
{
  "TrackQtyOnHand": true,
  "Name": "Garden Supplies",
  "QtyOnHand": 10,
  "IncomeAccountRef": {
    "name": "Sales of Product Income",
    "value": "79"
  },
  "AssetAccountRef": {
    "name": "Inventory Asset",
    "value": "81"
  },
  "InvStartDate": "2015-01-01",
  "Type": "Inventory",
  "ExpenseAccountRef": {
    "name": "Cost of Goods Sold",
    "value": "80"
  }
}
```

**Response**:
```json
{
  "Item": {
    "FullyQualifiedName": "Garden Supplies",
    "domain": "QBO",
    "Id": "19",
    "Name": "Garden Supplies",
    "TrackQtyOnHand": true,
    "UnitPrice": 0,
    "PurchaseCost": 0,
    "QtyOnHand": 10,
    "IncomeAccountRef": {
      "name": "Sales of Product Income",
      "value": "79"
    },
    "AssetAccountRef": {
      "name": "Inventory Asset",
      "value": "81"
    },
    "Taxable": false,
    "sparse": false,
    "Active": true,
    "SyncToken": "0",
    "InvStartDate": "2015-01-01",
    "Type": "Inventory",
    "ExpenseAccountRef": {
      "name": "Cost of Goods Sold",
      "value": "80"
    },
    "MetaData": {
      "CreateTime": "2015-12-09T11:12:39-08:00",
      "LastUpdatedTime": "2015-12-09T11:12:41-08:00"
    }
  },
  "time": "2015-12-09T11:12:39.748-08:00"
}
```

### Create a category

- **Method**: POST
- **URL**: `/v3/company/{realmID}/item?minorversion=4`

**Request Body**:
```json
{
  "SubItem": true,
  "Type": "Category",
  "Name": "Cedar",
  "ParentRef": {
    "name": "Trees",
    "value": "29"
  }
}
```

**Response**:
```json
{
  "Item": {
    "FullyQualifiedName": "Trees:Cedar",
    "domain": "QBO",
    "Name": "Cedar",
    "Level": 1,
    "sparse": false,
    "SubItem": true,
    "ParentRef": {
      "name": "Trees",
      "value": "29"
    },
    "Active": true,
    "SyncToken": "0",
    "Type": "Category",
    "Id": "30",
    "MetaData": {
      "CreateTime": "2015-10-06T10:50:42-07:00",
      "LastUpdatedTime": "2015-10-06T10:50:42-07:00"
    }
  },
  "time": "2015-10-06T10:50:42.707-07:00"
}
```

### Query a bundle

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}&minorversion=4`

**Sample Query**: `select * from Item where Type='Group'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Item": [
      {
        "Sku": "234",
        "FullyQualifiedName": "Deluxe Fountain",
        "domain": "QBO",
        "Name": "Deluxe Fountain",
        "TrackQtyOnHand": false,
        "Type": "Group",
        "PurchaseCost": 0,
        "Taxable": false,
        "ItemGroupDetail": {
          "ItemGroupLine": [
            {
              "Qty": 1,
              "ItemRef": {
                "type": "Inventory",
                "name": "Pump",
                "value": "11"
              }
            },
            {
              "Qty": 1,
              "ItemRef": {
                "type": "Inventory",
                "name": "Rock Fountain",
                "value": "5"
              }
            },
            {
              "Qty": 2,
              "ItemRef": {
                "type": "Service",
                "name": "Lighting",
                "value": "8"
              }
            },
            {
              "Qty": 4,
              "ItemRef": {
                "type": "Service",
                "name": "Installation",
                "value": "7"
              }
            }
          ]
        },
        "sparse": false,
        "Active": true,
        "PrintGroupedItems": true,
        "SyncToken": "1",
        "UnitPrice": 0,
        "Id": "49",
        "MetaData": {
          "CreateTime": "2016-06-23T10:51:32-07:00",
          "LastUpdatedTime": "2016-06-23T10:52:20-07:00"
        }
      }
    ],
    "maxResults": 1
  },
  "time": "2016-06-08T13:59:00.697-07:00"
}
```

### Query a category

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}&minorversion=4`

**Sample Query**: `select * from Item where Type='Category'`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Item": [
      {
        "FullyQualifiedName": "Flowers",
        "domain": "QBO",
        "Name": "Flowers",
        "SyncToken": "1",
        "sparse": false,
        "Active": true,
        "Type": "Category",
        "Id": "20",
        "MetaData": {
          "CreateTime": "2015-09-16T13:03:07-07:00",
          "LastUpdatedTime": "2015-09-16T13:40:51-07:00"
        }
      },
      {
        "FullyQualifiedName": "Flowers :Daises",
        "domain": "QBO",
        "Name": "Daises",
        "Level": 1,
        "sparse": false,
        "SubItem": true,
        "ParentRef": {
          "name": "Flowers",
          "value": "20"
        },
        "Active": true,
        "SyncToken": "0",
        "Type": "Category",
        "Id": "22",
        "MetaData": {
          "CreateTime": "2015-09-16T13:16:41-07:00",
          "LastUpdatedTime": "2015-09-16T13:31:46-07:00"
        }
      },
      {
        "FullyQualifiedName": "Flowers :Roses",
        "domain": "QBO",
        "Name": "Roses",
        "Level": 1,
        "sparse": false,
        "SubItem": true,
        "ParentRef": {
          "name": "Flowers",
          "value": "20"
        },
        "Active": true,
        "SyncToken": "0",
        "Type": "Category",
        "Id": "21",
        "MetaData": {
          "CreateTime": "2015-09-16T13:14:11-07:00",
          "LastUpdatedTime": "2015-09-16T13:14:11-07:00"
        }
      },
      {
        "FullyQualifiedName": "Garden Supplies",
        "domain": "QBO",
        "Name": "Garden Supplies",
        "SyncToken": "0",
        "sparse": false,
        "Active": true,
        "Type": "Category",
        "Id": "19",
        "MetaData": {
          "CreateTime": "2015-09-16T13:02:07-07:00",
          "LastUpdatedTime": "2015-09-16T13:02:07-07:00"
        }
      },
      {
        "FullyQualifiedName": "Organic Fir",
        "domain": "QBO",
        "Name": "Organic Fir",
        "SyncToken": "2",
        "sparse": false,
        "Active": true,
        "Type": "Category",
        "Id": "34",
        "MetaData": {
          "CreateTime": "2015-10-07T12:43:54-07:00",
          "LastUpdatedTime": "2015-10-07T12:48:23-07:00"
        }
      },
      {
        "FullyQualifiedName": "Organic Trees",
        "domain": "QBO",
        "Name": "Organic Trees",
        "SyncToken": "2",
        "sparse": false,
        "Active": true,
        "Type": "Category",
        "Id": "29",
        "MetaData": {
          "CreateTime": "2015-10-06T08:50:34-07:00",
          "LastUpdatedTime": "2015-10-07T12:48:23-07:00"
        }
      },
      {
        "FullyQualifiedName": "Organic Trees:Cedar",
        "domain": "QBO",
        "Name": "Cedar",
        "Level": 1,
        "sparse": false,
        "SubItem": true,
        "ParentRef": {
          "name": "Organic Trees",
          "value": "29"
        },
        "Active": true,
        "SyncToken": "0",
        "Type": "Category",
        "Id": "30",
        "MetaData": {
          "CreateTime": "2015-10-06T10:50:42-07:00",
          "LastUpdatedTime": "2015-10-07T12:38:03-07:00"
        }
      },
      {
        "FullyQualifiedName": "Organic Trees:Fig",
        "domain": "QBO",
        "Name": "Fig",
        "Level": 1,
        "sparse": false,
        "SubItem": true,
        "ParentRef": {
          "name": "Organic Trees",
          "value": "29"
        },
        "Active": true,
        "SyncToken": "0",
        "Type": "Category",
        "Id": "31",
        "MetaData": {
          "CreateTime": "2015-10-06T11:07:23-07:00",
          "LastUpdatedTime": "2015-10-07T12:38:03-07:00"
        }
      }
    ],
    "maxResults": 8
  },
  "time": "2015-10-08T13:59:00.697-07:00"
}
```

### Query an item

- **Method**: GET
- **URL**: `/v3/company/{realmID}/query?query={selectStatement}`

**Sample Query**: `select * from Item maxresults 2`

**Response**:
```json
{
  "QueryResponse": {
    "startPosition": 1,
    "Item": [
      {
        "FullyQualifiedName": "Concrete",
        "domain": "QBO",
        "Name": "Concrete",
        "TrackQtyOnHand": false,
        "Type": "Service",
        "PurchaseCost": 0,
        "IncomeAccountRef": {
          "name": "Landscaping Services :Job Materials:Fountains and Garden Lighting",
          "value": "48"
        },
        "Taxable": true,
        "MetaData": {
          "CreateTime": "2014-09-16T10:36:03-07:00",
          "LastUpdatedTime": "2014-09-19T12:47:47-07:00"
        },
        "sparse": false,
        "Active": true,
        "SyncToken": "1",
        "UnitPrice": 0,
        "Id": "3",
        "Description": "Concrete for fountain installation"
      },
      {
        "FullyQualifiedName": "Design",
        "domain": "QBO",
        "Name": "Design",
        "TrackQtyOnHand": false,
        "Type": "Service",
        "PurchaseCost": 0,
        "IncomeAccountRef": {
          "name": "Design income",
          "value": "82"
        },
        "Taxable": false,
        "MetaData": {
          "CreateTime": "2014-09-16T10:41:38-07:00",
          "LastUpdatedTime": "2015-04-17T14:31:10-07:00"
        },
        "sparse": false,
        "Active": true,
        "SyncToken": "1",
        "UnitPrice": 75,
        "Id": "4",
        "Description": "Custom Design"
      }
    ],
    "maxResults": 2
  },
  "time": "2015-04-22T11:04:34.194-07:00"
}
```

### Read a bundle

- **Method**: GET
- **URL**: `/v3/company/{realmID}/item/{itemId}`

Retrieves the details of a item bundle object that has been previously created.

**Response**:
```json
{
  "Item": {
    "Sku": "234",
    "FullyQualifiedName": "Deluxe Fountain",
    "domain": "QBO",
    "Name": "Deluxe Fountain",
    "TrackQtyOnHand": false,
    "Type": "Group",
    "PurchaseCost": 0,
    "Taxable": false,
    "ItemGroupDetail": {
      "ItemGroupLine": [
        {
          "Qty": 1,
          "ItemRef": {
            "type": "Inventory",
            "name": "Pump",
            "value": "11"
          }
        },
        {
          "Qty": 1,
          "ItemRef": {
            "type": "Inventory",
            "name": "Rock Fountain",
            "value": "5"
          }
        },
        {
          "Qty": 2,
          "ItemRef": {
            "type": "Service",
            "name": "Lighting",
            "value": "8"
          }
        },
        {
          "Qty": 4,
          "ItemRef": {
            "type": "Service",
            "name": "Installation",
            "value": "7"
          }
        }
      ]
    },
    "sparse": false,
    "Active": true,
    "PrintGroupedItems": true,
    "SyncToken": "1",
    "UnitPrice": 0,
    "Id": "49",
    "MetaData": {
      "CreateTime": "2016-06-23T10:51:32-07:00",
      "LastUpdatedTime": "2016-06-23T10:52:20-07:00"
    }
  },
  "time": "2016-06-23T15:14:21.695-07:00"
}
```

### Read a category

- **Method**: GET
- **URL**: `/v3/company/{realmID}/item/<itemId >?minorversion=4`

Retrieves the details of a item object that has been previously created.

**Response**:
```json
{
  "Item": {
    "FullyQualifiedName": "Trees",
    "domain": "QBO",
    "Name": "Trees",
    "SyncToken": "0",
    "sparse": false,
    "Active": true,
    "Type": "Category",
    "Id": "29",
    "MetaData": {
      "CreateTime": "2015-10-06T08:50:34-07:00",
      "LastUpdatedTime": "2015-10-06T08:50:34-07:00"
    }
  },
  "time": "2015-10-06T08:50:34.863-07:00"
}
```

### Read an item

- **Method**: GET
- **URL**: `/v3/company/{realmID}/item/{itemId}`

Retrieves the details of a item object that has been previously created.

**Response**:
```json
{
  "Item": {
    "FullyQualifiedName": "Office Supplies",
    "domain": "QBO",
    "Id": "37",
    "Name": "Office Supplies",
    "TrackQtyOnHand": true,
    "Type": "Inventory",
    "PurchaseCost": 35,
    "QtyOnHand": 10,
    "IncomeAccountRef": {
      "name": "Sales of Product Income",
      "value": "79"
    },
    "AssetAccountRef": {
      "name": "Inventory Asset",
      "value": "81"
    },
    "Taxable": true,
    "MetaData": {
      "CreateTime": "2015-04-22T11:03:23-07:00",
      "LastUpdatedTime": "2015-04-22T11:03:24-07:00"
    },
    "sparse": false,
    "Active": true,
    "SyncToken": "0",
    "InvStartDate": "2013-02-19",
    "UnitPrice": 25,
    "ExpenseAccountRef": {
      "name": "Cost of Goods Sold",
      "value": "80"
    },
    "PurchaseDesc": "This is the purchasing description.",
    "Description": "This is the sales description."
  },
  "time": "2015-04-22T11:01:37.346-07:00"
}
```

### Full update an item

- **Method**: POST
- **URL**: `/v3/company/{realmID}/item`

Use this operation to update any of the writable fields of an existing item object. The request body must include all writable fields of the existing object as returned in a read response. Writable fields omitted from the request body are set to NULL. The ID of the object to update is specified in t...

**Request Body**:
```json
{
  "FullyQualifiedName": "Rock Fountain",
  "domain": "QBO",
  "Id": "5",
  "Name": "Rock Fountain",
  "TrackQtyOnHand": true,
  "Type": "Inventory",
  "PurchaseCost": 125,
  "QtyOnHand": 2,
  "IncomeAccountRef": {
    "name": "Sales of Product Income",
    "value": "79"
  },
  "AssetAccountRef": {
    "name": "Inventory Asset",
    "value": "81"
  },
  "Taxable": true,
  "MetaData": {
    "CreateTime": "2014-09-16T10:42:19-07:00",
    "LastUpdatedTime": "2014-09-19T13:16:17-07:00"
  },
  "sparse": false,
  "Active": true,
  "SyncToken": "2",
  "InvStartDate": "2014-09-19",
  "UnitPrice": 275,
  "ExpenseAccountRef": {
    "name": "Cost of Goods Sold",
    "value": "80"
  },
  "PurchaseDesc": "Rock Fountain",
  "Description": "New, updated description for Rock Fountain"
}
```

**Response**:
```json
{
  "Item": {
    "FullyQualifiedName": "Rock Fountain",
    "domain": "QBO",
    "Id": "5",
    "Name": "Rock Fountain",
    "TrackQtyOnHand": true,
    "Type": "Inventory",
    "PurchaseCost": 125,
    "QtyOnHand": 2,
    "IncomeAccountRef": {
      "name": "Sales of Product Income",
      "value": "79"
    },
    "AssetAccountRef": {
      "name": "Inventory Asset",
      "value": "81"
    },
    "Taxable": true,
    "MetaData": {
      "CreateTime": "2014-09-16T10:42:19-07:00",
      "LastUpdatedTime": "2015-04-22T11:10:18-07:00"
    },
    "sparse": false,
    "Active": true,
    "SyncToken": "3",
    "InvStartDate": "2014-09-19",
    "UnitPrice": 275,
    "ExpenseAccountRef": {
      "name": "Cost of Goods Sold",
      "value": "80"
    },
    "PurchaseDesc": "Rock Fountain",
    "Description": "New, updated description for Rock Fountain"
  },
  "time": "2015-04-22T11:08:31.596-07:00"
}
```

### Update a category

- **Method**: POST
- **URL**: `/v3/company/{realmID}/item?minorversion=4`

Use this operation to update any of the writable fields of an existing category object. The ID of the object to update is specified in the request body.

**Request Body**:
```json
{
  "SyncToken": "1",
  "domain": "QBO",
  "Name": "Trees",
  "sparse": false,
  "Type": "Category",
  "Id": "29"
}
```

**Response**:
```json
{
  "Item": {
    "FullyQualifiedName": "Organic Trees",
    "domain": "QBO",
    "Name": "Organic Trees",
    "SyncToken": "2",
    "sparse": false,
    "Active": true,
    "Type": "Category",
    "Id": "29",
    "MetaData": {
      "CreateTime": "2015-10-06T08:50:34-07:00",
      "LastUpdatedTime": "2015-10-07T12:38:03-07:00"
    }
  },
  "time": "2015-10-07T12:40:29.199-07:00"
}
```
