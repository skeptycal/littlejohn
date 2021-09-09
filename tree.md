# littlejohn

> This is the initial directory tree for littlejohn. Use the make_tree_md.sh script ([GNU-tree required][get_tree]) to update it if you wish. It is safe to delete this file.

### Directory Structure

```sh
.
├── .VERSION
├── .editorconfig
├── .github
│   ├── FUNDING.yml
│   ├── ISSUE_TEMPLATE
│   │   ├── bug_report.md
│   │   └── feature_request.md
│   ├── dependabot.yml
│   └── workflows
│       ├── codeql-analysis.yml
│       └── go.yml
├── .gitignore
├── .pre-commit-config.yaml
├── CODE_OF_CONDUCT.md
├── LICENSE
├── README.md
├── SECURITY.md
├── bot
│   ├── bot.go
│   └── graph.go
├── cmd
│   └── examples
│       ├── gather
│       │   ├── pricedatatype
│       │   │   └── main.go
│       │   ├── randomint
│       │   │   └── main.go
│       │   └── randomstring
│       │       └── main.go
│       ├── littlejohn
│       │   └── main.go
│       └── logger
│           └── main.go
├── config
│   ├── ansicolormap.go
│   ├── config.go
│   └── util.go
├── contributing.md
├── coverage.txt
├── data_sources.json
├── docs
│   ├── _config.yml
│   ├── _config.yml~
│   ├── docs.md
│   ├── docs.md~
│   ├── index.html
│   ├── index.html~
│   ├── template.md
│   └── template.md~
├── example.go
├── examples
│   └── examples.go
├── gather
│   ├── data_sources.folder
│   ├── getdata.go
│   ├── stockdata.go
│   └── util.go
├── go.doc
├── go.mod
├── go.sum
├── go.test.sh
├── idea.md
├── instruments
│   ├── current.go
│   ├── historical.go
│   ├── instruments.go
│   └── session.go
├── littlejohn.go
├── logging
│   ├── enverror.go
│   ├── errors.go
│   └── logging.go
├── make_tree_md.sh
├── trades
│   ├── confirmation.go
│   ├── order.go
│   ├── orderTypes.go
│   ├── refType.go
│   └── triggerTypes.go
└── tree.md

20 directories, 60 files
```

[get_tree]: (http://mama.indstate.edu/users/ice/tree/)
