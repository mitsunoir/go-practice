# Golang の練習

## ミニ図書館 API (WIP)

- book の GET, POST だけ実装して満足気味
- Echo を使いつつクリーンアーキテクチャっぽくしたつもり
- ぬくもりてぃ溢れるフルハンドライティング

構成は以下

```
api
├── handler
│   └── book_handler.go
└── router
    └── router.go
cmd
└── server
    └── main.go
domain
├── author
├── book
│   ├── entity.go
│   ├── repo.go
│   └── service.go
├── borrow
└── user
infra
└── db
    └── book_repo.go
```

handler, infra も domain みたいにディレクトリを切った方が命名が簡素になるのでそうしたい。
