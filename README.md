# go-vue-check
Check, which files you should migrate to the Composition API from Options API

## Usage
`go get` - Install dependencies

`go build` - Build package

`./go-vue-check -p <project folder> [-s src <src folder>] -generate -open`

## Flags
`-p` - Project folder, default is current directory where program is executed

`-s` - Source folder, default is `src`

`-generate` - To generate report

`-open` - To open in browser (working only with active flag -generate)
