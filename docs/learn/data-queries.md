# Query Operations and Syntax

You can use the query operation to get details about a specific API entity.

## How query operations work

You can query most of our API entities. [Visit the API Explorer](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/account) to see which operations each entity supports.

The query operation is similar to a pared down SQL query select statement, but with a few limitations. These limitations ensure requests don't overload server-side resources:

- Server responses return all attributes for each API entity
- Server responses only return attributes with values
- Wildcard character support for LIKE clauses is limited to "%" (wildcard that substitutes for 0 or more characters)

Query operations don't support:

- Projections
- OR operations in WHERE clauses
- GROUP BY clauses
- JOIN clauses
- Special characters

## Use the query operation

Use a **GET** operation and include the select statement as a URI parameter:

### Select statement URI

```
https://quickbooks.api.intuit.com/v3/company/<realmId>/query?query=<select_statement>
```

### Select statement syntax

Use the select statement to specify the selection criteria, entity attributes, sort order, and pagination.

```sql
Select Statement = SELECT * | count(*)
  FROM IntuitEntity
  [WHERE WhereClause]
  [ORDERBY OrderByClause]
  [STARTPOSITION Number]
  [MAXRESULTS Number]
```

Use a **POST** operation to execute a query by sending the query details in the request body. It's useful for complex queries that may not fit well in a URI parameter.

### Query Request Body

The request body should be the query statement and content-type header should be set as `application/text`. The query statement specifies the selection criteria, entity attributes, sort order, and pagination.

```
URI: https://quickbooks.api.intuit.com/v3/company/<realmId>/query
Header: content-type: application/text
Body: SELECT * | count(*) FROM IntuitEntity [WHERE WhereClause] [ORDERBY OrderByClause] [STARTPOSITION Number] [MAXRESULTS Number]
```

| Parameter | Value | Description |
|-----------|-------|-------------|
| `Select *` | | All fields are returned. |
| `count(*)` | | The number of records that satisfy the query criteria. Returns a single value. |
| `Entity` | Customer, Vendor, Invoice, etc | The name of the queried entity. For example: Customer, Vendor, Invoice, etc. Case sensitive. **Tip**: You can only query one entity at a time. Learn more [about API entities](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/most-commonly-used/account). |
| `WhereClause` | PropertyName, Operator Value [AND WhereClause] | The WhereClause filters the returned data according to the value of the PropertyName. Multiple clauses (filters) use AND. **Note**: The OR operation isn't supported. |
| `Operator` | IN, =, <, >, <=, >=, LIKE | Wildcard character support for LIKE clauses is limited to "%" (wildcard that substitutes for 0 or more characters). **Note:** Filtering on the ID field does not support operators >, <, >=, <=, !=, and LIKE. Only = and IN are supported for the ID field. |
| `Value` | (Value [,Value]), 'value_in_quote', value_without_quote, true, false, ' ' | This is only supported for IN clauses. **Note**: Null is represented as the string ' ' (i.e., a space within single quotes). Expression evaluation isn't supported. |
| `OrderByClause` | PropertyName [ASC, DESC], [OrderByClause] | The OrderByClause sorts the result. |
| `Number` | | Must be a positive integer. |

**Example query select statement:**

```
GET https://quickbooks.api.intuit.com/v3/company/1234/query?query=SELECT%20FROM%20Customer%20WHERE%20Metadata.LastUpdatedTime%20%3E+%272011-08-10T10%3A20%3A30-0700%27
```

**Example query request body statement:**

```
POST https://quickbooks.api.intuit.com/v3/company/1234/query
-H 'content-type:application/text'
-d 'SELECT * FROM Customer WHERE Metadata.LastUpdatedTime > '2011-08-10T10:20:30''
```

## Server responses for queries

Server responses contain a **\<QueryResponse\>** element. This corresponds with the original request's select statement and contains the API entities that matched its criteria.

**Example server response:**

```xml
<IntuitResponse xmlns="http://schema.intuit.com/finance/v3" time="2013-04-03T10:22:55.766Z">
  <QueryResponse startPosition="10" maxResults="2">
    <Customer>
      <Id>2123</Id>
      <SyncToken>0</SyncToken>
      ...
      <GivenName>Srini</GivenName>
    </Customer>
    <Customer>
      <Id>2124</Id>
      <SyncToken>0</SyncToken>
      ...
      <GivenName>Peter</GivenName>
    </Customer>
  </QueryResponse>
</IntuitResponse>
```

| Attribute | Description |
|-----------|-------------|
| `startposition` | The starting point of the response for pagination. |
| `maxresults` | The number of transactions or records in the **\<QueryResponse\>** element. The maximum number of entities that can be returned in a response is 1,000. If the result size isn't specified, the default number is 100. **Tip**: If a query returns many entities, use pagination to fetch the entities in chunks. To determine the number of entities a query returns, use the COUNT keyword. |
| `sparse` | Set to **true** if the response doesn't return all attributes for the entity. |

If the query contains an error, the \<QueryResponse\> element will contain **\<Fault\>**:

```xml
<IntuitResponse xmlns="http://schema.intuit.com/finance/v3" time="2013-04-03T10:22:55.766Z">
  <QueryResponse>
    <Fault type="Validation">
      <Error code="100" element="query">
        <Message>OBJECT name not available</Message>
      </Error>
      <Error code="300" element="query">
        <Message>Time value is incorrect</Message>
      </Error>
    </Fault>
  </QueryResponse>
</IntuitResponse>
```

## Query syntax

### Case syntax

Reserved words, entity names, and attribute names are not case sensitive. Attribute values are also not case sensitive.

For example, while the following two statements use different cases, they're equivalent:

```sql
SELECT * FROM Customer WHERE GivenName = 'greg' STARTPOSITION 10
SELECT * from cUstomer where givenname = 'Greg' Startposition 10
```

### Escape character

Use backslash (\\) to escape special characters like apostrophes (').

For example, to find a **Customer** entity for a QuickBooks Online company named "Adam's Candy Shop," use the following query:

```sql
SELECT * FROM Customer WHERE CompanyName = 'Adam\'s Candy Shop'
```

## Query language operations

### Filters

The WHERE clause determines which entities are returned by a query. Filter criteria is based on the value of the attribute.

For example, the following query returns all **Invoice** entities with a TotalAmt attribute greater than 1000.0:

```sql
SELECT * FROM Invoice WHERE TotalAmt > '1000.0'
```

**Tip**: Most attributes support filters. [Visit the API Explorer](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/most-commonly-used/account) and the specific entity reference to see which can be filtered.

**Note**: Filtering on the `ID` field does not support operators `>`, `<`, `>=`, `<=`, `!=`, and `LIKE`. Only `=` and `IN` are supported for the `ID` field.

### Multiple filters

To use multiple filters together, use the AND keyword. The OR operation isn't supported.

For example, the following query returns all **Invoice** entities created during the specified time range:

```sql
SELECT * FROM Invoice WHERE MetaData.CreateTime >= '2009-10-14T04:05:05-07:00' AND MetaData.CreateTime <= '2012-10-14T04:05:05-07:00'
```

The following statement has two AND operators:

```sql
SELECT * FROM Invoice WHERE id in ('64523', '18761', '35767') AND MetaData.CreateTime >= '1990-12-12T12:50:30Z' AND MetaData.LastUpdatedTime <='1990-12-12T12:50:30Z'
```

If the WHERE clause is omitted, all entities up to the maximum number are returned.

If a query returns many entities, use pagination to fetch the entities in chunks.

### More filter examples

To get all invoices for a given customer, filter using the **CustomerRef** attribute:

```sql
SELECT * FROM Invoice WHERE CustomerRef = '123'
```

To get all active and inactive customers (only [list entities](https://developer.intuit.com/app/developer/qbo/docs/learn/explore-the-quickbooks-online-api) like **Customer**, **Vendor**, **Item**, etc. can be "active" or "inactive"), filter using the **Active** attribute:

```sql
SELECT * FROM Customer WHERE Active IN (true, false)
```

To get all active **Customer** entities:

```sql
SELECT * FROM Customer WHERE active = true
```

To get the **Customer** entity for the specified identifier, filter using the **Id** attribute:

```sql
SELECT * FROM Customer WHERE Id = '123456'
```

To get all attributes for the **Customer** entities whose first name starts with "K," ends with "h", and whose last name starts with "Palm":

```sql
SELECT * FROM Customer WHERE GivenName LIKE 'K%h' AND FamilyName LIKE 'Palm%'
```

To get the header properties for the **Invoice** entities whose **CustomerId** attributes match the values of the IN clause:

```sql
SELECT * FROM Invoice WHERE CustomerId IN ('12')
```

To get all attributes for the **Invoice** entities with transaction dates between Jan. 1, 2011 and the current date:

```sql
SELECT * FROM Invoice WHERE TxnDate > '2011-01-01' AND TxnDate <= CURRENT_DATE
```

### Sorting

To sort query results, include the ORDERBY clause and specify an attribute as the sort field.

For example, the following query sorts **Customer** entities by the FamilyName attribute:

```sql
SELECT * FROM Customer ORDERBY FamilyName
```

The default sort order is ascending. To indicate descending sort order, include DESC:

```sql
SELECT * FROM Customer ORDERBY FamilyName DESC
```

**Tip**: Many attributes can be used to sort. [Visit the API Explorer](https://developer.intuit.com/app/developer/qbo/docs/api/accounting/all-entities/most-commonly-used/account) and the specific entity reference to see which can sort results.

### Pagination

To page through the results, specify **STARTPOSITION** (position of the entity in the query results) and **MAXRESULTS** (maximum number of entities in the result).

For example, suppose you have 25 invoices. The following query gets invoices 1 - 10:

```sql
SELECT * FROM Invoice STARTPOSITION 1 MAXRESULTS 10
```

The next statement gets invoices 11 - 20:

```sql
SELECT * FROM Invoice STARTPOSITION 11 MAXRESULTS 10
```

The next gets invoices 21 - 25:

```sql
SELECT * FROM Invoice STARTPOSITION 21 MAXRESULTS 10
```

### Count

To find out how many entities will be returned by a query, specify the COUNT keyword.

The number of entities is returned as the value of the `totalCount` field of the **\<QueryResponse\>** element. QuickBooks Online data objects, such as Customers, aren't in the response.

For example, the following query returns the total number of customers for a single QuickBooks Online company:

```sql
SELECT COUNT(*) FROM Customer
```
