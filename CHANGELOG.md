# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added
- **Functional Options Pattern**: `New` function now accepts functional options like `WithSandbox()` for flexible client initialization.
- **Nested Structs**: `PaymentRequest` now uses nested structs (`Customer`, `Shipping`, `Product`, `Extra`) for cleaner initialization.
- **Constants**: Added constants for `ProductCategory`, `ShippingMethod`, `ProductProfile`, `PaymentStatus`, and `Currency` in the `models` package.
- **Cart Builder**: Added `SetCart` method to `PaymentRequest` to easily construct the `cart` parameter from a slice of products.
- **IPN Parser**: Added `ParseIPN` helper function to parse incoming IPN requests into an `IpnResponse` struct.
- **Models Package**: Created a dedicated `models` package to organize request and response structures.

### Changed
- **Module Name**: Renamed module to `github.com/sagar290/sslcommerz-go`.
- **Package Name**: Shortened package name from `ssl_wireless_pgw_golang_sdk` to `ssl`.
- **Client Initialization**: Replaced `Init` with `New`.
- **Request Construction**: Removed `Set...` methods in favor of direct struct literal initialization.
- **URL Encoding**: Switched to `github.com/google/go-querystring` for more robust URL parameter encoding, using `url` tags.
- **Address Fields**: Corrected `url` tags for address fields (e.g., `cus_add1` instead of `cus_add_1`) to match API requirements.

### Removed
- **Structs Package**: Removed the `structs` package; contents moved to `models`.
- **Helpers Package**: Removed `helpers.StructToMap` as it was replaced by `go-querystring`.
