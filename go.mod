// A Go repository typically contains only one module, located at the root of
// the repository. A file named go.mod there declares the module path (The
// import path prefix for all packages within the module)
//
// The module contains the packages in the directory containing its go.mod
// file as well as subdirectories of that directory, up to the next
// subdirectory containing another go.mod file (if any).
module github.com/cloether/go-module-template

go 1.14

// -------------------------------
// Managing Go Module Dependencies
// -------------------------------
//
// Adding Module Dependencies
// --------------------------
// Manual:
//          // Add Multiple
//          require (
//            github.com/ardanlabs/conf v1.2.0
//            github.com/pkg/errors v0.9.1
//          )
//
//          // Add One
//          require github.com/ardanlabs/conf v1.2.0
//
// Substitute Imported Modules
// ---------------------------
// Manual:
//          replace github.com/go-chi/chi => ./packages/chi
//
// CLI:
//          $ go mod edit -dropreplace github.com/go-chi/chi
