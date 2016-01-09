go-json-example
===

orgmode file to generate sample json files and sample go files

pre-requisites
===
1. recent emacs with orgmode; babel languages "sh" enabled
2. optional: go-mode.el installed for src block editing via  C-c '
3. go installed; go env vars set; go tools in path

usage
===
1. edit go-json.org
2. expand all sections: tab tab
3. run setup to create directories for tangled (generated) files: "C-c C-c y" in the block named create-dirs-for-tangled-code
4. tangle (generate) .json and .go files: "C-c C-v t"
5. build and run the go programs: "C-c C-c y" in the sh block named build-run-go

exporting html
===
1. edit go-json.org
2. export to go-json.html: "C-c C-E h o y"

generated html
===
Browse [sample output](http://wrightmikea.github.io/go-json-example.html "exported html")
