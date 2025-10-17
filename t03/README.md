# Go Naming Conventions

## 1. Package Names

- Hepsi küçük harf
- Çoğul değil **tekil** form
- `_` (underscore) kullanma
- camelCase kullanma

**Örnek:**
```go
package tempconv
```

## 2. File Names

- Büyük harfle başlama, hepsi küçük harf
- camelCase kullanma
- `-` (tire) kullanma
- `_` (underscore) kullan

**Örnekler:**
- ✅ `user.go`, `user_test.go`, `http_client.go`
- ❌ `User.go`, `userManager.go`, `http-client.go`

## 3. Variables, Constants, Functions

- **Exported** (public): Büyük harfle başlar
- **Private** (unexported): Küçük harfle başlar
- camelCase kullanır, `_` kullanmaz
- **Getter** **function**'larda `Get` prefix'i yok (Java gibi değil)

**Örnekler:**
- ✅ `var userCount int`, `func GetUserByID()`, `func (u *User) Name()`
- ❌ `var user_count int`, `func (u *User) GetName()`

## 4. Acronyms / Kısaltmalar

Hepsi aynı case olmalı (tamamı büyük ya da tamamı küçük).

**Örnekler:**
- ✅ `userID`, `apiURL`, `httpServer`, `jsonData`
- ❌ `userId`, `apiUrl`, `httpserver`, `jsondata`

---

## Claude Code Komutu

`.claude/commands/nameco`

```md
---
description: Check Go naming conventions (package names, file names, variables, acronyms)
---

# Go Naming Convention Checker

Please analyze all Go files that I changed in this branch in the current project and check for naming convention violations according to these rules:

## 1. Package Names
- Must be lowercase only (no uppercase letters)
- Must be singular form (not plural, except standard library packages like strings, bytes, errors)
- Must NOT contain underscores (_)
- Must NOT use camelCase
- Should be short and descriptive

Examples:
- ✅ Good: `package user`, `package http`, `package tempconv`
- ❌ Bad: `package User` (uppercase), `package userManager` (camelCase), `package user_manager` (underscore)

## 2. File Names
- Must start with lowercase
- Must use lowercase throughout
- Must NOT use camelCase
- Must NOT use dashes (-)
- Should use underscores (_) for word separation

Examples:
- ✅ Good: `user.go`, `user_test.go`, `http_client.go`
- ❌ Bad: `User.go` (uppercase), `userManager.go` (camelCase), `http-client.go` (dash)

## 3. Variables, Constants, Functions
- Exported (public): Must start with uppercase letter
- Unexported (private): Must start with lowercase letter
- Must use camelCase (not snake_case)
- Getter methods should NOT have "Get" prefix (unlike Java)

Examples:
- ✅ Good: `var userCount int`, `func GetUserByID()`, `func (u *User) Name()` (no Get prefix)
- ❌ Bad: `var user_count int` (snake_case), `func (u *User) GetName()` (unnecessary Get)

## 4. Acronyms / Initialisms
- Must be all uppercase or all lowercase (consistent case)
- Common acronyms: ID, URL, HTTP, JSON, XML, API, SQL, HTML, CPU, DB

Examples:
- ✅ Good: `userID`, `apiURL`, `httpServer`, `jsonData`
- ❌ Bad: `userId`, `apiUrl`, `httpserver`, `jsondata`

## Analysis Instructions

For each Go file found, check:

1. **Package declaration**: Extract package name and verify it follows package naming rules
2. **File name**: Check if the filename follows file naming conventions
3. **Variable/Function declarations**: Look for snake_case usage (should be camelCase)
4. **Getter methods**: Look for methods with "Get" prefix on receiver methods
5. **Acronyms**: Find common wrong acronyms like `Id`, `Url`, `Http`, `Json`, `Api`, `Sql`, `Html`, `Xml`

## Output Format

Please provide a report in this format:

🔍 Go Naming Convention Check Results
=====================================

📦 Package Names:
[List any violations found]

📁 File Names:
[List any violations found]

🔤 Variables/Constants/Functions:
[List any violations found]

🔡 Acronyms:
[List any violations found]

=====================================
Summary: X violations found (or ✅ All checks passed!)

For each violation, show:
- File path and line number (if applicable)
- What was found
- What it should be
- Brief explanation

Exclude:
- Files in `vendor/` directory
- Generated files
- Test files when checking certain rules (if appropriate)

Start the analysis now for all `.go` files in the project that I changed in this branch.
```