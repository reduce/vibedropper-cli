# Changelog

## 0.1.0 (2026-05-02)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/reduce/vibedropper-cli/compare/v0.0.1...v0.1.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([34c6ce5](https://github.com/reduce/vibedropper-cli/commit/34c6ce5b106314edf8b0c6bb38231bbe9e17bbc2))
* add default description for enum CLI flags without an explicit description ([6c69bc2](https://github.com/reduce/vibedropper-cli/commit/6c69bc237da1fdc7cbe2a3f6a20fc875781a6c77))
* add support for file downloads from binary response endpoints ([eb061ff](https://github.com/reduce/vibedropper-cli/commit/eb061ff61f4e0862218a2056ed62c5251718f29d))
* allow `-` as value representing stdin to binary-only file parameters in CLIs ([1b7b36f](https://github.com/reduce/vibedropper-cli/commit/1b7b36f7cc376fb0e9c4b5d8c1fa3a25d680d564))
* **api:** api update ([b82ba33](https://github.com/reduce/vibedropper-cli/commit/b82ba334f6e9b27b946585a985ca2e63722938bd))
* **api:** api update ([b5aa6c0](https://github.com/reduce/vibedropper-cli/commit/b5aa6c06de93ed64a537fd58e8c3fa15142879be))
* **api:** manual updates ([7e2fb63](https://github.com/reduce/vibedropper-cli/commit/7e2fb6336e33506b32549e566dcf6d228032b73a))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([9a22a14](https://github.com/reduce/vibedropper-cli/commit/9a22a14534e73a8e54d6c4982012910731d815ea))
* binary-only parameters become CLI flags that take filenames only ([103c078](https://github.com/reduce/vibedropper-cli/commit/103c0781e982747584c2c3fca68a9f853211cbb7))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([a2f50ac](https://github.com/reduce/vibedropper-cli/commit/a2f50aca857e71149016593b31bebfa4ec971a13))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([336d8b6](https://github.com/reduce/vibedropper-cli/commit/336d8b6e14881d378ca0744ea23597ec0a10ed20))
* **cli:** send filename and content type when reading input from files ([172a216](https://github.com/reduce/vibedropper-cli/commit/172a2164552f7df78461a648be2ffcd145999ac5))
* improved documentation and flags for client options ([bb401ea](https://github.com/reduce/vibedropper-cli/commit/bb401ea7d82eeba9c4b7f5ea8d56182533b82038))
* set CLI flag constant values automatically where `x-stainless-const` is set ([3824c15](https://github.com/reduce/vibedropper-cli/commit/3824c15463c741e4cfbcb69f0a6dc27d04e6f1b0))
* support passing path and query params over stdin ([699d039](https://github.com/reduce/vibedropper-cli/commit/699d0390e9b5de91e1273e960c7a9793bc2b7cb5))
* support passing required body params through pipes ([f7ff615](https://github.com/reduce/vibedropper-cli/commit/f7ff615f36cafec0e4eb2749e33121742c490d3b))


### Bug Fixes

* avoid printing usage errors twice ([6a36803](https://github.com/reduce/vibedropper-cli/commit/6a36803168f123e1d8311484c6e49c7aa981e8cc))
* avoid reading from stdin unless request body is form encoded or json ([08ca3eb](https://github.com/reduce/vibedropper-cli/commit/08ca3eb786b48c60383e1f6661dfb53cf3322e55))
* better support passing client args in any position ([a1af9a2](https://github.com/reduce/vibedropper-cli/commit/a1af9a23c3660a293f65f4988c4440c37c935a09))
* cli no longer hangs when stdin is attached to a pipe with empty input ([c5c1d6d](https://github.com/reduce/vibedropper-cli/commit/c5c1d6dd43c7944bbcc4b87764f572e1b5c6633e))
* **cli:** correctly load zsh autocompletion ([f3da599](https://github.com/reduce/vibedropper-cli/commit/f3da5992a54cb1dfb6105d8adb755835bbe61e2c))
* fall back to main branch if linking fails in CI ([f9884e9](https://github.com/reduce/vibedropper-cli/commit/f9884e9943fc371c4f4956ee63d7327f461aaf93))
* fix for encoding arrays with `any` type items ([36ba377](https://github.com/reduce/vibedropper-cli/commit/36ba377ce2e709804ba4424357a2696743d319b8))
* fix for failing to drop invalid module replace in link script ([05b7b56](https://github.com/reduce/vibedropper-cli/commit/05b7b5659ea5de177ff4de13b911a3c0f6a26044))
* fix for off-by-one error in pagination logic ([7b4087f](https://github.com/reduce/vibedropper-cli/commit/7b4087f595593cbd690c0e625109989c4fd85da6))
* fix for test cases with newlines in YAML and better error reporting ([e35fbf9](https://github.com/reduce/vibedropper-cli/commit/e35fbf9a2c664b7868073ce531dc7949257ddeb4))
* fix quoting typo ([a4c8e23](https://github.com/reduce/vibedropper-cli/commit/a4c8e2355e8ecad98690221d4b031f891601b8e8))
* flags for nullable body scalar fields are strictly typed ([caaa396](https://github.com/reduce/vibedropper-cli/commit/caaa39680d4852343b21749a62e67d9d1789f610))
* handle empty data set using `--format explore` ([da31d29](https://github.com/reduce/vibedropper-cli/commit/da31d2915cae10cbe8092f913c07deff96d97e89))
* improve linking behavior when developing on a branch not in the Go SDK ([619ea0a](https://github.com/reduce/vibedropper-cli/commit/619ea0a2e1059217904b361a4f67a685a4468e09))
* improved workflow for developing on branches ([f13805d](https://github.com/reduce/vibedropper-cli/commit/f13805d9ba8648d344bcc354f3b92c03976ea335))
* no longer require an API key when building on production repos ([4a70906](https://github.com/reduce/vibedropper-cli/commit/4a70906227517b29372d62e99db7e33417a06ed9))
* only set client options when the corresponding CLI flag or env var is explicitly set ([94719ba](https://github.com/reduce/vibedropper-cli/commit/94719ba9d866ee4bb17bca4e0fe349599ce85099))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([9969fc6](https://github.com/reduce/vibedropper-cli/commit/9969fc6d4cb91cfbb2b220efcd21b5f4df928583))


### Chores

* add documentation for ./scripts/link ([f69d230](https://github.com/reduce/vibedropper-cli/commit/f69d23046bff75dd632557abde6f889cd64073de))
* **ci:** skip lint on metadata-only changes ([71190d9](https://github.com/reduce/vibedropper-cli/commit/71190d9f684073c5b900c0ddb1ce8ef72422ecb7))
* **ci:** skip uploading artifacts on stainless-internal branches ([26819be](https://github.com/reduce/vibedropper-cli/commit/26819bed6d43173ebe205b858b53cbd90e4fffdf))
* **ci:** support manually triggering release workflow ([82746c7](https://github.com/reduce/vibedropper-cli/commit/82746c75b9613ee357c717f8ba15a216552d6764))
* **cli:** additional test cases for `ShowJSONIterator` ([b1031fd](https://github.com/reduce/vibedropper-cli/commit/b1031fd0407fe8bdb788ca2d9732b2db7a4500a5))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([9bafc9d](https://github.com/reduce/vibedropper-cli/commit/9bafc9d902155ed0ee2a048a3f1260211f88532e))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([5e29e08](https://github.com/reduce/vibedropper-cli/commit/5e29e08acbb99413461aba3ab2caef86fe5864b7))
* **cli:** switch long lists of positional args over to param structs ([190aff4](https://github.com/reduce/vibedropper-cli/commit/190aff4fabdb6a1f0894889102e1f16c839b1f0d))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([888884c](https://github.com/reduce/vibedropper-cli/commit/888884ca68c93309d430a1f13ea239cdfa179d0f))
* configure new SDK language ([5df30bb](https://github.com/reduce/vibedropper-cli/commit/5df30bb2e1cce81717bc1aa6adb62af3118f8d1d))
* **internal:** codegen related update ([28b8a04](https://github.com/reduce/vibedropper-cli/commit/28b8a04a4de6bab5fdec296ed7d5e5d874ecad3d))
* **internal:** more robust bootstrap script ([974bf75](https://github.com/reduce/vibedropper-cli/commit/974bf7528e6e65aa9ffa65f98eabb866071dfe05))
* **internal:** tweak CI branches ([a9dfc4c](https://github.com/reduce/vibedropper-cli/commit/a9dfc4c9d1021e578bc5dd6439cf73d728001c4f))
* **internal:** update gitignore ([307d7d1](https://github.com/reduce/vibedropper-cli/commit/307d7d141299c832ae33cb765748177e85da8047))
* mark all CLI-related tests in Go with `t.Parallel()` ([7faec3d](https://github.com/reduce/vibedropper-cli/commit/7faec3db8f079610969a584ce2ada37e4217c4f1))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([1a504f8](https://github.com/reduce/vibedropper-cli/commit/1a504f88c1414bd7802557d0d3bfab1f239263ef))
* omit full usage information when missing required CLI parameters ([9aed547](https://github.com/reduce/vibedropper-cli/commit/9aed547ec9b7a7fec19078f8e0e777d8820aada6))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([3cb81e4](https://github.com/reduce/vibedropper-cli/commit/3cb81e4fee8151a0a89f226ccae9f5b494b44298))
* sync repo ([1bb0554](https://github.com/reduce/vibedropper-cli/commit/1bb0554f603cbd29a9871f018a2d5e4ac2fbb64d))
* update SDK settings ([6525394](https://github.com/reduce/vibedropper-cli/commit/652539440f04390fb28a652cdf1615b9cb3b3b89))
* update SDK settings ([96e04f1](https://github.com/reduce/vibedropper-cli/commit/96e04f1f59116dbf0906bf468dd6aeab21a4f560))
* update SDK settings ([9849e39](https://github.com/reduce/vibedropper-cli/commit/9849e392d8f961250c49ed2b51a0fac5b3ad9265))

## 0.0.2 (2026-02-28)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/reduce/vibedropper-cli/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([5df30bb](https://github.com/reduce/vibedropper-cli/commit/5df30bb2e1cce81717bc1aa6adb62af3118f8d1d))
* sync repo ([f763b13](https://github.com/reduce/vibedropper-cli/commit/f763b139caff52beea4b9074baa0319dbf840e12))
* update SDK settings ([b419b80](https://github.com/reduce/vibedropper-cli/commit/b419b808ef4846dfeb1f7b6c0942e6ee9b89f8f2))
* update SDK settings ([3a8cf1d](https://github.com/reduce/vibedropper-cli/commit/3a8cf1df8a76dae2602ec374cd65a586c69b5c2f))
* update SDK settings ([9dd7e5c](https://github.com/reduce/vibedropper-cli/commit/9dd7e5c66cbf9e615eb8979124c39ab1e0ee8155))
