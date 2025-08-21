# Sample API Documentation

A comprehensive sample API demonstrating all OpenAPI features supported by the swagger-markdown generator

API Support

https://example.com/support

support@example.com

## Paths

| Path | Operations |
| --- | --- |
| [/products](#path/products) | GET |
| [/users](#path/users) | GET, POST |
| [/users/{userId}](#path/users/{userid}) | DELETE, GET, PUT |


## <span id="path/products">/products</span>

### GET

Retrieve a list of products with advanced filtering

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| category | false | string | Product category | electronics |
| price_min | false | number | Minimum price filter | 10.99 |
| price_max | false | number | Maximum price filter | 999.99 |
| in_stock | false | boolean | Filter by stock availability | true |

**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [Products retrieved successfully](#/definitions/ProductListResponse) |


---

## <span id="path/users">/users</span>

### GET

Retrieve a list of all users with optional filtering

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| page | false | integer | Page number for pagination | 1 |
| limit | false | integer | Number of items per page | 20 |
| status | false | string | Filter users by status | active |
| tags | false | array | Filter by tags (comma-separated) | [premium verified] |

**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [Successfully retrieved users](#/definitions/UserListResponse) |
| 400 | [Invalid request parameters](#/definitions/ErrorResponse) |
| 500 | [Internal server error](#/definitions/ErrorResponse) |


---

### POST

Create a new user account

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| user | true |  | User data | <nil> |

**Responses:**

| Status Code | Description |
| --- | --- |
| 201 | [User created successfully](#/definitions/User) |
| 400 | [Invalid user data](#/definitions/ErrorResponse) |
| 409 | [User already exists](#/definitions/ErrorResponse) |


---

## <span id="path/users/{userId}">/users/{userId}</span>

### DELETE

Delete a user account

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| userId | true | string | User ID | <nil> |

**Responses:**

| Status Code | Description |
| --- | --- |
| 204 | [User deleted successfully]() |
| 404 | [User not found](#/definitions/ErrorResponse) |


---

### GET

Retrieve a specific user by their ID

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| userId | true | string | Unique identifier for the user | 123e4567-e89b-12d3-a456-426614174000 |
| include | false | array | Additional fields to include | [profile preferences] |

**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [User found](#/definitions/User) |
| 404 | [User not found](#/definitions/ErrorResponse) |


---

### PUT

Update an existing user's information

**Parameters:**

| Name | Required | Type | Description | Example |
| --- | --- | --- | --- | --- |
| userId | true | string | User ID | <nil> |
| user | true |  | Updated user data | <nil> |

**Responses:**

| Status Code | Description |
| --- | --- |
| 200 | [User updated successfully](#/definitions/User) |
| 400 | [Invalid user data](#/definitions/ErrorResponse) |
| 404 | [User not found](#/definitions/ErrorResponse) |


---

## Definitions

### <span id="/definitions/ComplexNestedExample">ComplexNestedExample</span>

<a id="/definitions/ComplexNestedExample"></a>

Complex Nested Example

Demonstrates complex nested structures with arrays and maps

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| mixed_type_map | object | Map with different value types |  |
| multi_level_array | [][]string | Array of arrays of strings |  |
| object_array | []**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| name | string |  |  |
| nested_map | object |  |  |
 | Array of objects with nested properties |  |
| reference_array | [][User](#/definitions/User) | Array of user references |  |




---

### <span id="/definitions/CreateUserRequest">CreateUserRequest</span>

<a id="/definitions/CreateUserRequest"></a>

Create User Request

Request payload for creating a new user

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| email | string | User's email address | newuser@example.com |
| full_name | string | User's full name | New User |
| password | string | User's password | securePassword123 |
| username | string | Desired username | new_user |




---

### <span id="/definitions/ErrorResponse">ErrorResponse</span>

<a id="/definitions/ErrorResponse"></a>

Error Response

Standard error response format

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| details | []**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| code | string | Specific error code for this field | INVALID_FORMAT |
| field | string | Field that caused the error | email |
| message | string | Error message for this field | Email format is invalid |
 | Detailed error information |  |
| error | string | Error code | INVALID_REQUEST |
| message | string | Human-readable error message | The request parameters are invalid |
| request_id | string | Unique request identifier for debugging | req-abc123xyz |
| timestamp | string | When the error occurred | 2023-01-15T10:30:00Z |




---

### <span id="/definitions/NestedObjectMap">NestedObjectMap</span>

<a id="/definitions/NestedObjectMap"></a>

Nested Object Map

A map containing nested object structures

**Type:** map[*]->**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| name | string |  |  |
| nested_array | []string |  |  |
| value | string |  |  |




---

### <span id="/definitions/NotificationSettings">NotificationSettings</span>

<a id="/definitions/NotificationSettings"></a>

Notification Settings

User's notification preferences

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| email_enabled | boolean | Enable email notifications |  |
| frequency | string ([enums](#/enums/frequency)) | Notification frequency | daily |
| push_enabled | boolean | Enable push notifications |  |
| types | []string | Types of notifications to receive |  |


## Enums

**<span id="/enums/frequency"></span>frequency:**

| frequency |
| --- |
|immediate, daily, weekly, never|





---

### <span id="/definitions/PaginationInfo">PaginationInfo</span>

<a id="/definitions/PaginationInfo"></a>

Pagination Info

Pagination metadata

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| has_next | boolean | Whether there is a next page |  |
| has_previous | boolean | Whether there is a previous page |  |
| limit | integer | Items per page |  |
| page | integer | Current page number |  |
| total_pages | integer | Total number of pages |  |




---

### <span id="/definitions/PrivacySettings">PrivacySettings</span>

<a id="/definitions/PrivacySettings"></a>

Privacy Settings

User's privacy configuration

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| allow_indexing | boolean | Allow search engines to index profile |  |
| profile_public | boolean | Make profile publicly visible |  |
| show_email | boolean | Show email in public profile |  |




---

### <span id="/definitions/Product">Product</span>

<a id="/definitions/Product"></a>

Product

A product in the catalog

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| category | string | Product category | electronics |
| currency | string ([enums](#/enums/currency)) | Price currency | USD |
| description | string | Product description | High-quality wireless headphones with noise cancellation |
| id | string | Product ID | prod-12345 |
| in_stock | boolean | Whether the product is in stock |  |
| name | string | Product name | Wireless Headphones |
| price | number | Product price |  |
| specifications | [StringMap](#/definitions/StringMap) |  |  |
| stock_quantity | integer | Number of items in stock |  |
| variants | [][ProductVariant](#/definitions/ProductVariant) | Product variants |  |


## Enums

**<span id="/enums/currency"></span>currency:**

| currency |
| --- |
|USD, EUR, GBP, JPY|





---

### <span id="/definitions/ProductListResponse">ProductListResponse</span>

<a id="/definitions/ProductListResponse"></a>

Product List Response

Response containing a list of products

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| filters_applied | [StringMap](#/definitions/StringMap) |  |  |
| products | [][Product](#/definitions/Product) | Array of products |  |
| total_count | integer | Total number of products matching filters |  |




---

### <span id="/definitions/ProductVariant">ProductVariant</span>

<a id="/definitions/ProductVariant"></a>

Product Variant

A variant of a product

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| attributes | object | Variant attributes |  |
| id | string | Variant ID | var-001 |
| name | string | Variant name | Black - Large |
| price_adjustment | number | Price adjustment for this variant |  |




---

### <span id="/definitions/SocialLinks">SocialLinks</span>

<a id="/definitions/SocialLinks"></a>

Social Links

Map of social platform names to URLs

**Type:** map[*]->string



---

### <span id="/definitions/StringMap">StringMap</span>

<a id="/definitions/StringMap"></a>

String Map

A map of string keys to string values

**Type:** map[*]->string



---

### <span id="/definitions/UpdateUserRequest">UpdateUserRequest</span>

<a id="/definitions/UpdateUserRequest"></a>

Update User Request

Request payload for updating an existing user

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| email | string | Updated email address | updated@example.com |
| full_name | string | Updated full name | Updated Name |
| status | string ([enums](#/enums/status)) | Updated user status | active |


## Enums

**<span id="/enums/status"></span>status:**

| status |
| --- |
|active, inactive, pending, suspended|





---

### <span id="/definitions/User">User</span>

<a id="/definitions/User"></a>

User

A user in the system

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| age | integer | User's age |  |
| created_at | string | When the user was created | 2023-01-15T10:30:00Z |
| email | string | User's email address | user@example.com |
| full_name | string | User's full name | John Doe |
| id | string | Unique identifier for the user | 123e4567-e89b-12d3-a456-426614174000 |
| metadata | [StringMap](#/definitions/StringMap) |  |  |
| preferences | [UserPreferences](#/definitions/UserPreferences) |  |  |
| profile | [UserProfile](#/definitions/UserProfile) |  |  |
| status | string ([enums](#/enums/status)) | Current status of the user account | active |
| tags | []string | User tags for categorization |  |
| updated_at | string | When the user was last updated | 2023-01-20T15:45:00Z |
| username | string | User's unique username | john_doe |


## Enums

**<span id="/enums/status"></span>status:**

| status |
| --- |
|active, inactive, pending, suspended|





---

### <span id="/definitions/UserArrayMap">UserArrayMap</span>

<a id="/definitions/UserArrayMap"></a>

User Array Map

A map where values are arrays of users

**Type:** map[*]->[][User](#/definitions/User)



---

### <span id="/definitions/UserListResponse">UserListResponse</span>

<a id="/definitions/UserListResponse"></a>

User List Response

Response containing a list of users with pagination

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| pagination | [PaginationInfo](#/definitions/PaginationInfo) |  |  |
| total_count | integer | Total number of users |  |
| users | [][User](#/definitions/User) | Array of users |  |




---

### <span id="/definitions/UserPreferences">UserPreferences</span>

<a id="/definitions/UserPreferences"></a>

User Preferences

User's application preferences and settings

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| language | string ([enums](#/enums/language)) | Preferred language | en |
| notifications | [NotificationSettings](#/definitions/NotificationSettings) |  |  |
| privacy | [PrivacySettings](#/definitions/PrivacySettings) |  |  |
| theme | string ([enums](#/enums/theme)) | UI theme preference | dark |
| timezone | string | User's timezone | America/New_York |


## Enums

**<span id="/enums/language"></span>language:**

| language |
| --- |
|en, es, fr, de, ja, zh|

**<span id="/enums/theme"></span>theme:**

| theme |
| --- |
|light, dark, auto|





---

### <span id="/definitions/UserProfile">UserProfile</span>

<a id="/definitions/UserProfile"></a>

User Profile

Extended user profile information

**Type:** object

**Properties:**

| Name | Type | Description | Example |
| --- | --- | --- | --- |
| avatar_url | string | URL to user's avatar image | https://example.com/avatars/user123.jpg |
| bio | string | User's biography | Software developer passionate about open source |
| location | string | User's location | San Francisco, CA |
| skills | []string | User's skills |  |
| social_links | [SocialLinks](#/definitions/SocialLinks) |  |  |
| website | string | User's personal website | https://johndoe.com |




---

